package sink

import (
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/replication/source"
	"github.com/bookstairs/bookworm/util"
)

type ReplicationSink interface {
	GetName() string
	Initialize(configuration util.Configuration, prefix string) error
	DeleteEntry(key string, isDirectory, deleteIncludeChunks bool, signatures []int32) error
	CreateEntry(key string, entry *filer_pb.Entry, signatures []int32) error
	UpdateEntry(key string, oldEntry *filer_pb.Entry, newParentPath string, newEntry *filer_pb.Entry, deleteIncludeChunks bool, signatures []int32) (foundExistingEntry bool, err error)
	GetSinkToDirectory() string
	SetSourceFiler(s *source.FilerSource)
	IsIncremental() bool
}

var (
	Sinks []ReplicationSink
)
