package weed_server

import (
	"fmt"

	"github.com/bookstairs/bookworm/pb/volume_server_pb"
	"github.com/bookstairs/bookworm/storage"
	"github.com/bookstairs/bookworm/storage/needle"
)

func (vs *VolumeServer) ReadAllNeedles(req *volume_server_pb.ReadAllNeedlesRequest, stream volume_server_pb.VolumeServer_ReadAllNeedlesServer) (err error) {

	for _, vid := range req.VolumeIds {
		if err := vs.streamReadOneVolume(needle.VolumeId(vid), stream, err); err != nil {
			return err
		}
	}
	return nil
}

func (vs *VolumeServer) streamReadOneVolume(vid needle.VolumeId, stream volume_server_pb.VolumeServer_ReadAllNeedlesServer, err error) error {
	v := vs.store.GetVolume(vid)
	if v == nil {
		return fmt.Errorf("not found volume id %d", vid)
	}

	scanner := &storage.VolumeFileScanner4ReadAll{
		Stream: stream,
		V:      v,
	}

	offset := int64(v.SuperBlock.BlockSize())

	err = storage.ScanVolumeFileFrom(v.Version(), v.DataBackend, offset, scanner)

	return err
}
