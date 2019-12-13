package main

import (
	"CloudStorage/store/ceph"
	"os"
)

func main() {
	bucket := ceph.GetCephBucket("userfile")
	d, _ := bucket.Get("/ceph/702b37b7b78e96c08ebe902e39f714a012e7d197")
	tmpFile, _ := os.Create("/tmp/test_file456")
	tmpFile.Write(d)
	return
}
