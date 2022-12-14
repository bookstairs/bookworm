package command

import (
	_ "net/http/pprof"

	_ "github.com/bookstairs/bookworm/filer/leveldb"
	_ "github.com/bookstairs/bookworm/filer/leveldb2"
	_ "github.com/bookstairs/bookworm/filer/leveldb3"
	_ "github.com/bookstairs/bookworm/remote_storage/azure"
	_ "github.com/bookstairs/bookworm/remote_storage/gcs"
	_ "github.com/bookstairs/bookworm/remote_storage/s3"
	_ "github.com/bookstairs/bookworm/replication/sink/azuresink"
	_ "github.com/bookstairs/bookworm/replication/sink/b2sink"
	_ "github.com/bookstairs/bookworm/replication/sink/filersink"
	_ "github.com/bookstairs/bookworm/replication/sink/gcssink"
	_ "github.com/bookstairs/bookworm/replication/sink/localsink"
	_ "github.com/bookstairs/bookworm/replication/sink/s3sink"
)
