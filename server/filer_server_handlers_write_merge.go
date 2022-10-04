package weed_server

import (
	"github.com/bookstairs/bookworm/operation"
	"github.com/bookstairs/bookworm/pb/filer_pb"
)

func (fs *FilerServer) maybeMergeChunks(so *operation.StorageOption, inputChunks []*filer_pb.FileChunk) (mergedChunks []*filer_pb.FileChunk, err error) {
	//TODO merge consecutive smaller chunks into a large chunk to reduce number of chunks
	return inputChunks, nil
}
