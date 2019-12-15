package main

import (
	"CloudStorage/store/ceph"
	"os"
)

func main() {
	bucket := ceph.GetCephBucket("userfile")
	d, _ := bucket.Get("/ceph/21248974b7c16c75db561b07eb607b1d402ef7c3")
	tmpFile, _ := os.Create("/tmp/test_file123")
	tmpFile.Write(d)
	return
}
