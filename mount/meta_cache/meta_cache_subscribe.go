package meta_cache

import (
	"context"

	"github.com/golang/glog"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/pb"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/util"
)

func SubscribeMetaEvents(mc *MetaCache, selfSignature int32, client filer_pb.FilerClient, dir string, lastTsNs int64) error {

	processEventFn := func(resp *filer_pb.SubscribeMetadataResponse) error {
		message := resp.EventNotification

		for _, sig := range message.Signatures {
			if sig == selfSignature && selfSignature != 0 {
				return nil
			}
		}

		dir := resp.Directory
		var oldPath util.FullPath
		var newEntry *filer.Entry
		if message.OldEntry != nil {
			oldPath = util.NewFullPath(dir, message.OldEntry.Name)
			glog.V(4).Infof("deleting %v", oldPath)
		}

		if message.NewEntry != nil {
			if message.NewParentPath != "" {
				dir = message.NewParentPath
			}
			key := util.NewFullPath(dir, message.NewEntry.Name)
			glog.V(4).Infof("creating %v", key)
			newEntry = filer.FromPbEntry(dir, message.NewEntry)
		}
		err := mc.AtomicUpdateEntryFromFiler(context.Background(), oldPath, newEntry)
		if err == nil {
			if message.OldEntry != nil && message.NewEntry != nil {
				oldKey := util.NewFullPath(resp.Directory, message.OldEntry.Name)
				mc.invalidateFunc(oldKey, message.OldEntry)
				if message.OldEntry.Name != message.NewEntry.Name {
					newKey := util.NewFullPath(dir, message.NewEntry.Name)
					mc.invalidateFunc(newKey, message.NewEntry)
				}
			} else if filer_pb.IsCreate(resp) {
				// no need to invalidate
			} else if filer_pb.IsDelete(resp) {
				oldKey := util.NewFullPath(resp.Directory, message.OldEntry.Name)
				mc.invalidateFunc(oldKey, message.OldEntry)
			}
		}

		return err

	}

	var clientEpoch int32
	util.RetryForever("followMetaUpdates", func() error {
		clientEpoch++
		return pb.WithFilerClientFollowMetadata(client, "mount", selfSignature, clientEpoch, dir, nil, &lastTsNs, 0, selfSignature, processEventFn, pb.FatalOnError)
	}, func(err error) bool {
		glog.Errorf("follow metadata updates: %v", err)
		return true
	})

	return nil
}
