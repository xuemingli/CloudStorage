package config

import (
	cmn "CloudStorage/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/data/fileserver/"
	// TempPartRootDir : 分块文件在本地临时存储地址的路径
	TempPartRootDir = "/data/fileserver_part/"
	// TempCephRootDir : Ceph存储地址的路径
	TempCephRootDir = "/ceph/"
	// TempOssRootDir : Oss存储地址的路径
	TempOssRootDir = "oss/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreCeph
)
