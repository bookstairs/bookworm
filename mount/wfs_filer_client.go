package mount

import (
	"sync/atomic"

	"github.com/golang/glog"
	"google.golang.org/grpc"

	"github.com/bookstairs/bookworm/pb"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/util"
)

var _ = filer_pb.FilerClient(&WFS{})

func (wfs *WFS) WithFilerClient(streamingMode bool, fn func(filer_pb.SeaweedFilerClient) error) (err error) {

	return util.Retry("filer grpc", func() error {

		i := atomic.LoadInt32(&wfs.option.filerIndex)
		n := len(wfs.option.FilerAddresses)
		for x := 0; x < n; x++ {

			filerGrpcAddress := wfs.option.FilerAddresses[i].ToGrpcAddress()
			err = pb.WithGrpcClient(streamingMode, func(grpcConnection *grpc.ClientConn) error {
				client := filer_pb.NewSeaweedFilerClient(grpcConnection)
				return fn(client)
			}, filerGrpcAddress, false, wfs.option.GrpcDialOption)

			if err != nil {
				glog.V(0).Infof("WithFilerClient %d %v: %v", x, filerGrpcAddress, err)
			} else {
				atomic.StoreInt32(&wfs.option.filerIndex, i)
				return nil
			}

			i++
			if i >= int32(n) {
				i = 0
			}

		}
		return err
	})

}

func (wfs *WFS) AdjustedUrl(location *filer_pb.Location) string {
	if wfs.option.VolumeServerAccess == "publicUrl" {
		return location.PublicUrl
	}
	return location.Url
}

func (wfs *WFS) GetDataCenter() string {
	return wfs.option.DataCenter
}
