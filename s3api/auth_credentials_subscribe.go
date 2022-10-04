package s3api

import (
	"github.com/golang/glog"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/pb"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/s3api/s3_constants"
	"github.com/bookstairs/bookworm/util"
)

func (s3a *S3ApiServer) subscribeMetaEvents(clientName string, lastTsNs int64, prefix string, directoriesToWatch []string) {

	processEventFn := func(resp *filer_pb.SubscribeMetadataResponse) error {

		message := resp.EventNotification
		if message.NewEntry == nil {
			return nil
		}

		dir := resp.Directory

		if message.NewParentPath != "" {
			dir = message.NewParentPath
		}
		fileName := message.NewEntry.Name
		content := message.NewEntry.Content

		_ = s3a.onIamConfigUpdate(dir, fileName, content)
		_ = s3a.onCircuitBreakerConfigUpdate(dir, fileName, content)
		_ = s3a.onBucketMetadataChange(dir, message.OldEntry, message.NewEntry)

		return nil
	}

	var clientEpoch int32
	util.RetryForever("followIamChanges", func() error {
		clientEpoch++
		return pb.WithFilerClientFollowMetadata(s3a, clientName, s3a.randomClientId, clientEpoch, prefix, directoriesToWatch, &lastTsNs, 0, 0, processEventFn, pb.FatalOnError)
	}, func(err error) bool {
		glog.V(0).Infof("iam follow metadata changes: %v", err)
		return true
	})
}

// reload iam config
func (s3a *S3ApiServer) onIamConfigUpdate(dir, filename string, content []byte) error {
	if dir == filer.IamConfigDirectory && filename == filer.IamIdentityFile {
		if err := s3a.iam.LoadS3ApiConfigurationFromBytes(content); err != nil {
			return err
		}
		glog.V(0).Infof("updated %s/%s", dir, filename)
	}
	return nil
}

// reload circuit breaker config
func (s3a *S3ApiServer) onCircuitBreakerConfigUpdate(dir, filename string, content []byte) error {
	if dir == s3_constants.CircuitBreakerConfigDir && filename == s3_constants.CircuitBreakerConfigFile {
		if err := s3a.cb.LoadS3ApiConfigurationFromBytes(content); err != nil {
			return err
		}
		glog.V(0).Infof("updated %s/%s", dir, filename)
	}
	return nil
}

// reload bucket metadata
func (s3a *S3ApiServer) onBucketMetadataChange(dir string, oldEntry *filer_pb.Entry, newEntry *filer_pb.Entry) error {
	if dir == s3a.option.BucketsPath {
		if newEntry != nil {
			s3a.bucketRegistry.LoadBucketMetadata(newEntry)
			glog.V(0).Infof("updated bucketMetadata %s/%s", dir, newEntry)
		} else {
			s3a.bucketRegistry.RemoveBucketMetadata(oldEntry)
			glog.V(0).Infof("remove bucketMetadata  %s/%s", dir, newEntry)
		}
	}
	return nil
}