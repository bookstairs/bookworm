package s3api

import (
	"testing"
	"time"

	"github.com/bookstairs/bookworm/s3api/s3err"
)

func TestCopyObjectResponse(t *testing.T) {

	// https://docs.aws.amazon.com/AmazonS3/latest/API/API_CopyObject.html

	response := CopyObjectResult{
		ETag:         "12345678",
		LastModified: time.Now(),
	}

	println(string(s3err.EncodeXMLResponse(response)))

}

func TestCopyPartResponse(t *testing.T) {

	// https://docs.aws.amazon.com/AmazonS3/latest/API/API_UploadPartCopy.html

	response := CopyPartResult{
		ETag:         "12345678",
		LastModified: time.Now(),
	}

	println(string(s3err.EncodeXMLResponse(response)))

}
