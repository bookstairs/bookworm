package s3api

import (
	"net/http"

	"github.com/bookstairs/bookworm/s3api/s3_constants"
	"github.com/bookstairs/bookworm/s3api/s3err"
)

func getAccountId(r *http.Request) string {
	id := r.Header.Get(s3_constants.AmzAccountId)
	if len(id) == 0 {
		return AccountAnonymous.Id
	} else {
		return id
	}
}

func (s3a *S3ApiServer) checkAccessByOwnership(r *http.Request, bucket string) s3err.ErrorCode {
	metadata, errCode := s3a.bucketRegistry.GetBucketMetadata(bucket)
	if errCode != s3err.ErrNone {
		return errCode
	}
	accountId := getAccountId(r)
	if accountId == AccountAdmin.Id || accountId == *metadata.Owner.ID {
		return s3err.ErrNone
	}
	return s3err.ErrAccessDenied
}
