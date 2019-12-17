package main

import (
	"CloudStorage/store/ceph"
	"os"
)

func main() {
	bucket := ceph.GetCephBucket("userfile")
	d, _ := bucket.Get("/ceph/a3a8f72f07600789e43c320544e5878f7e378e06")
	tmpFile, _ := os.Create("/tmp/test_file2")
	tmpFile.Write(d)
	return
}
