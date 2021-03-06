package blob

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Build a byte-array of one megabyte
func buildOneMeg() []byte {
	var buffer bytes.Buffer

	for i := 0; i < 100000; i++ {
		buffer.WriteString("HelloWorld")

	}
	return buffer.Bytes()
}

func buildOneMegHybridBlob() HybridRawGzipBlob {
	return NewRawBlob(buildOneMeg())
}

func TestNewRawBlob(t *testing.T) {
	assert.Equal(t, 1000000, len(buildOneMeg()))
	assert.Equal(t, buildOneMeg(), buildOneMegHybridBlob().Raw())
}

func TestNewGzippedBlob(t *testing.T) {
	gzippedBlob := NewGzippedBlob(buildOneMegHybridBlob().Gzipped())

	// once compressed the blob should be under 5k
	assert.Equal(t, true, len(gzippedBlob.Gzipped()) < 5000)

	assert.Equal(t, buildOneMeg(), gzippedBlob.Raw())
}
