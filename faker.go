package faker

import (
	"bytes"
	"sync"
)

var BufferPool = &sync.Pool{
	New: func() interface{} {
		return bytes.NewBuffer(make([]byte, 256))
	},
}
