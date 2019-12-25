package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"

	cmn "CloudStorage/common"
	cfg "CloudStorage/config"
	mydb "CloudStorage/db"
	"CloudStorage/mq"
	dbcli "CloudStorage/service/dbproxy/client"
	"CloudStorage/service/dbproxy/orm"
	"CloudStorage/store/ceph"
	"CloudStorage/util"
)

// DoUploadHandler ： 处理文件上传
func DoUploadHandler(c *gin.Context) {
	errCode := 1
	defer func() {
		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
		if errCode < 0 {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传失败",
			})
		} else if errCode == 1 {
			c.JSON(http.StatusOK, gin.H{
				"code": 0,
				"msg":  "重复上传",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"code": errCode,
				"msg":  "上传成功",
			})
		}
	}()

	// 1. 从form表单中获得文件内容句柄
	file, head, err := c.Request.FormFile("file")
	if err != nil {
		log.Printf("Failed to get form data, err:%s\n", err.Error())
		errCode = -1
		return
	}
	defer file.Close()

	// 2. 把文件内容转为[]byte
	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		log.Printf("Failed to get file data, err:%s\n", err.Error())
		errCode = -2
		return
	}

	// 3. 构建文件元信息
	fileMeta := dbcli.FileMeta{
		FileName: head.Filename,
		FileSha1: util.Sha1(buf.Bytes()), //　计算文件sha1
		FileSize: int64(len(buf.Bytes())),
		UploadAt: time.Now().Format("2006-01-02 15:04:05"),
	}

	username := c.Request.FormValue("username")
	ufile, err := mydb.QueryUserFileMeta(username, fileMeta.FileSha1)
	if err != nil || ufile.FileName != fileMeta.FileName {

		// 4. 将文件写入临时存储位置
		fileMeta.Location = cfg.TempLocalRootDir + fileMeta.FileName // 临时存储地址
		newFile, err := os.Create(fileMeta.Location)
		if err != nil {
			log.Printf("Failed to create file, err:%s\n", err.Error())
			errCode = -3
			return
		}
		defer newFile.Close()

		nByte, err := newFile.Write(buf.Bytes())
		if int64(nByte) != fileMeta.FileSize || err != nil {
			log.Printf("Failed to save data into file, writtenSize:%d, err:%s\n", nByte, err.Error())
			errCode = -4
			return
		}

		newFile.Seek(0, 0)
		// 5. 同步或异步将文件转移到Ceph/OSS
		newFile.Seek(0, 0) // 游标重新回到文件头部
		if cfg.CurrentStoreType == cmn.StoreCeph {
			// 文件写入Ceph存储
			cephPath := cfg.TempCephRootDir + fileMeta.FileSha1

			// 判断写入ceph为同步还是异步
			if !cfg.AsyncTransferEnable {
				datafile, _ := ioutil.ReadAll(newFile)
				err = ceph.PutObject("userfile", cephPath, datafile)
				if err != nil {
					fmt.Println(err.Error())
					errCode = -5
					return
				}
				fileMeta.Location = cephPath
			} else {
				// 写入异步转移任务队列
				data := mq.TransferData{
					FileHash:      fileMeta.FileSha1,
					CurLocation:   fileMeta.Location,
					DestLocation:  cephPath,
					DestStoreType: cmn.StoreCeph,
				}
				pubData, _ := json.Marshal(data)
				pubSuc := mq.Publish(
					cfg.TransExchangeName,
					cfg.TransCephRoutingKey,
					pubData,
				)
				if !pubSuc {
					// TODO: 当前发送转移信息失败，稍后重试
				}
				fileMeta.Location = cephPath
			}
		}

		//6.  更新文件表记录
		_, err = dbcli.OnFileUploadFinished(fileMeta)
		if err != nil {
			errCode = -6
			return
		}

		// 7. 更新用户文件表
		username := c.Request.FormValue("username")
		upRes, err := dbcli.OnUserFileUploadFinished(username, fileMeta)
		if err == nil && upRes.Suc {
			errCode = 0
		} else {
			errCode = -7
		}
	} else {
		errCode = 1
	}
}

// TryFastUploadHandler : 尝试秒传接口
func TryFastUploadHandler(c *gin.Context) {

	// 1. 解析请求参数
	username := c.Request.FormValue("username")
	filehash := c.Request.FormValue("filehash")
	filename := c.Request.FormValue("filename")
	// filesize, _ := strconv.Atoi(c.Request.FormValue("filesize"))

	// 2. 从文件表中查询相同hash的文件记录
	fileMetaResp, err := dbcli.GetFileMeta(filehash)
	if err != nil {
		log.Println(err.Error())
		c.Status(http.StatusInternalServerError)
		return
	}

	// 3. 查不到记录则返回秒传失败
	if !fileMetaResp.Suc {
		resp := util.RespMsg{
			Code: -1,
			Msg:  "秒传失败，请访问普通上传接口",
		}
		c.Data(http.StatusOK, "application/json", resp.JSONBytes())
		return
	}

	// 4. 上传过则将文件信息写入用户文件表， 返回成功
	fmeta := dbcli.TableFileToFileMeta(fileMetaResp.Data.(orm.TableFile))
	fmeta.FileName = filename
	upRes, err := dbcli.OnUserFileUploadFinished(username, fmeta)
	if err == nil && upRes.Suc {
		resp := util.RespMsg{
			Code: 0,
			Msg:  "秒传成功",
		}
		c.Data(http.StatusOK, "application/json", resp.JSONBytes())
		return
	}
	resp := util.RespMsg{
		Code: -2,
		Msg:  "秒传失败，请稍后重试",
	}
	c.Data(http.StatusOK, "application/json", resp.JSONBytes())
	return
}
