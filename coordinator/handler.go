package coordinator

import (
	"net"

	"github.com/Wikia/influxdb/protocol"
)

type Handler interface {
	HandleRequest(*protocol.Request, net.Conn) error
}
