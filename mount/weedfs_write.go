package mount

import (
	"fmt"
	"io"

	"github.com/golang/glog"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/operation"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/util"
)

func (wfs *WFS) saveDataAsChunk(fullPath util.FullPath) filer.SaveDataAsChunkFunctionType {

	return func(reader io.Reader, filename string, offset int64) (chunk *filer_pb.FileChunk, err error) {

		fileId, uploadResult, err, data := operation.UploadWithRetry(
			wfs,
			&filer_pb.AssignVolumeRequest{
				Count:       1,
				Replication: wfs.option.Replication,
				Collection:  wfs.option.Collection,
				TtlSec:      wfs.option.TtlSec,
				DiskType:    string(wfs.option.DiskType),
				DataCenter:  wfs.option.DataCenter,
				Path:        string(fullPath),
			},
			&operation.UploadOption{
				Filename:          filename,
				Cipher:            wfs.option.Cipher,
				IsInputCompressed: false,
				MimeType:          "",
				PairMap:           nil,
			},
			func(host, fileId string) string {
				fileUrl := fmt.Sprintf("http://%s/%s", host, fileId)
				if wfs.option.VolumeServerAccess == "filerProxy" {
					fileUrl = fmt.Sprintf("http://%s/?proxyChunkId=%s", wfs.getCurrentFiler(), fileId)
				}
				return fileUrl
			},
			reader,
		)

		if err != nil {
			glog.V(0).Infof("upload data %v: %v", filename, err)
			return nil, fmt.Errorf("upload data: %v", err)
		}
		if uploadResult.Error != "" {
			glog.V(0).Infof("upload failure %v: %v", filename, err)
			return nil, fmt.Errorf("upload result: %v", uploadResult.Error)
		}

		if offset == 0 {
			wfs.chunkCache.SetChunk(fileId, data)
		}

		chunk = uploadResult.ToPbFileChunk(fileId, offset)
		return chunk, nil
	}
}
