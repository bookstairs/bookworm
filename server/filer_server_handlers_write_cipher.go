package weed_server

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/golang/glog"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/operation"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/storage/needle"
	"github.com/bookstairs/bookworm/util"
)

// handling single chunk POST or PUT upload
func (fs *FilerServer) encrypt(ctx context.Context, w http.ResponseWriter, r *http.Request, so *operation.StorageOption) (filerResult *FilerPostResult, err error) {

	fileId, urlLocation, auth, err := fs.assignNewFileInfo(so)

	if err != nil || fileId == "" || urlLocation == "" {
		return nil, fmt.Errorf("fail to allocate volume for %s, collection:%s, datacenter:%s", r.URL.Path, so.Collection, so.DataCenter)
	}

	glog.V(4).Infof("write %s to %v", r.URL.Path, urlLocation)

	// Note: encrypt(gzip(data)), encrypt data first, then gzip

	sizeLimit := int64(fs.option.MaxMB) * 1024 * 1024

	bytesBuffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(bytesBuffer)

	pu, err := needle.ParseUpload(r, sizeLimit, bytesBuffer)
	uncompressedData := pu.Data
	if pu.IsGzipped {
		uncompressedData = pu.UncompressedData
	}
	if pu.MimeType == "" {
		pu.MimeType = http.DetectContentType(uncompressedData)
		// println("detect2 mimetype to", pu.MimeType)
	}

	uploadOption := &operation.UploadOption{
		UploadUrl:         urlLocation,
		Filename:          pu.FileName,
		Cipher:            true,
		IsInputCompressed: false,
		MimeType:          pu.MimeType,
		PairMap:           pu.PairMap,
		Jwt:               auth,
	}
	uploadResult, uploadError := operation.UploadData(uncompressedData, uploadOption)
	if uploadError != nil {
		return nil, fmt.Errorf("upload to volume server: %v", uploadError)
	}

	// Save to chunk manifest structure
	fileChunks := []*filer_pb.FileChunk{uploadResult.ToPbFileChunk(fileId, 0)}

	// fmt.Printf("uploaded: %+v\n", uploadResult)

	path := r.URL.Path
	if strings.HasSuffix(path, "/") {
		if pu.FileName != "" {
			path += pu.FileName
		}
	}

	entry := &filer.Entry{
		FullPath: util.FullPath(path),
		Attr: filer.Attr{
			Mtime:  time.Now(),
			Crtime: time.Now(),
			Mode:   0660,
			Uid:    OS_UID,
			Gid:    OS_GID,
			TtlSec: so.TtlSeconds,
			Mime:   pu.MimeType,
			Md5:    util.Base64Md5ToBytes(pu.ContentMd5),
		},
		Chunks: fileChunks,
	}

	filerResult = &FilerPostResult{
		Name: pu.FileName,
		Size: int64(pu.OriginalDataSize),
	}

	if dbErr := fs.filer.CreateEntry(ctx, entry, false, false, nil, false); dbErr != nil {
		fs.filer.DeleteChunks(entry.Chunks)
		err = dbErr
		filerResult.Error = dbErr.Error()
		return
	}

	return
}