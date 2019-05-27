package quic

import (
	"sync"
	"time"
	"os/exec"
	"github.com/lucas-clemente/quic-go/internal/utils"
	"github.com/lucas-clemente/quic-go/internal/protocol"
)

var bufferPool sync.Pool

func getPacketBuffer() []byte {
	cmd := exec.Command("cat", "/proc/net/udp")

    dtTime := time.Now().UnixNano()

	out, err := cmd.CombinedOutput()

	if err != nil {
		 utils.Errorf("cmd.Run() failed with %s\n", err)
	}
	utils.Infof("%d - UDP QUEUE:\n %s", dtTime, string(out))
	return bufferPool.Get().([]byte)
}

func putPacketBuffer(buf []byte) {
	if cap(buf) != int(protocol.MaxReceivePacketSize) {
		panic("putPacketBuffer called with packet of wrong size!")
	}
	bufferPool.Put(buf[:0])
}

func init() {
	bufferPool.New = func() interface{} {
		return make([]byte, 0, protocol.MaxReceivePacketSize)
	}
}
