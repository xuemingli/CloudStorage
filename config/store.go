package config

import (
	cmn "CloudStorage/common"
)

const (
	// TempLocalRootDir : 本地临时存储地址的路径
	TempLocalRootDir = "/data/fileserver/"
	// TempCephRootDir : Ceph存储地址的路径
	TempCephRootDir = "/ceph/"
	// CurrentStoreType : 设置当前文件的存储类型
	CurrentStoreType = cmn.StoreCeph
)
