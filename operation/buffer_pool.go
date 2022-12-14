package operation

import (
	"sync/atomic"

	"github.com/valyala/bytebufferpool"
)

var bufferCounter int64

func GetBuffer() *bytebufferpool.ByteBuffer {
	defer func() {
		atomic.AddInt64(&bufferCounter, 1)
		// println("+", bufferCounter)
	}()
	return bytebufferpool.Get()
}

func PutBuffer(buf *bytebufferpool.ByteBuffer) {
	defer func() {
		atomic.AddInt64(&bufferCounter, -1)
		// println("-", bufferCounter)
	}()
	bytebufferpool.Put(buf)
}
