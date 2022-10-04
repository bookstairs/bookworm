package filer

import (
	"fmt"
	"strings"

	"github.com/bookstairs/bookworm/util"
)

func (f *Filer) CanRename(source, target util.FullPath) error {
	sourceBucket := f.DetectBucket(source)
	targetBucket := f.DetectBucket(target)
	if sourceBucket != targetBucket {
		return fmt.Errorf("can not move across collection %s => %s", sourceBucket, targetBucket)
	}
	return nil
}

func (f *Filer) DetectBucket(source util.FullPath) (bucket string) {
	if strings.HasPrefix(string(source), f.DirBucketsPath+"/") {
		bucketAndObjectKey := string(source)[len(f.DirBucketsPath)+1:]
		t := strings.Index(bucketAndObjectKey, "/")
		if t < 0 {
			bucket = bucketAndObjectKey
		}
		if t > 0 {
			bucket = bucketAndObjectKey[:t]
		}
	}
	return bucket
}
