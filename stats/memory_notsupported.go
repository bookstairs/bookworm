//go:build !linux
// +build !linux

package stats

import "github.com/bookstairs/bookworm/pb/volume_server_pb"

func fillInMemStatus(status *volume_server_pb.MemStatus) {
	return
}
