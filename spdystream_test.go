package peerstream_spdystream

import (
	"testing"

	test "github.com/libp2p/go-stream-muxer/test"
)

func TestSpdyStreamTransport(t *testing.T) {
	test.SubtestAll(t, Transport)
}
