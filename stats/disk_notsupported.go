//go:build openbsd || netbsd || plan9 || solaris
// +build openbsd netbsd plan9 solaris

package stats

import "github.com/bookstairs/bookworm/pb/volume_server_pb"

func fillInDiskStatus(status *volume_server_pb.DiskStatus) {
	return
}
