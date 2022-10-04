package shell

import (
	"fmt"
	"io"

	"google.golang.org/protobuf/proto"

	"github.com/bookstairs/bookworm/filer"
	"github.com/bookstairs/bookworm/pb/filer_pb"
	"github.com/bookstairs/bookworm/util"
)

func init() {
	Commands = append(Commands, &commandFsMetaCat{})
}

type commandFsMetaCat struct {
}

func (c *commandFsMetaCat) Name() string {
	return "fs.meta.cat"
}

func (c *commandFsMetaCat) Help() string {
	return `print out the meta data content for a file or directory

	fs.meta.cat /dir/
	fs.meta.cat /dir/file_name
`
}

func (c *commandFsMetaCat) Do(args []string, commandEnv *CommandEnv, writer io.Writer) (err error) {

	path, err := commandEnv.parseUrl(findInputDirectory(args))
	if err != nil {
		return err
	}

	dir, name := util.FullPath(path).DirAndName()

	return commandEnv.WithFilerClient(false, func(client filer_pb.SeaweedFilerClient) error {

		request := &filer_pb.LookupDirectoryEntryRequest{
			Name:      name,
			Directory: dir,
		}
		respLookupEntry, err := filer_pb.LookupEntry(client, request)
		if err != nil {
			return err
		}

		filer.ProtoToText(writer, respLookupEntry.Entry)

		bytes, _ := proto.Marshal(respLookupEntry.Entry)
		gzippedBytes, _ := util.GzipData(bytes)
		// zstdBytes, _ := util.ZstdData(bytes)
		// fmt.Fprintf(writer, "chunks %d meta size: %d gzip:%d zstd:%d\n", len(respLookupEntry.Entry.Chunks), len(bytes), len(gzippedBytes), len(zstdBytes))
		fmt.Fprintf(writer, "chunks %d meta size: %d gzip:%d\n", len(respLookupEntry.Entry.Chunks), len(bytes), len(gzippedBytes))

		return nil

	})

}
