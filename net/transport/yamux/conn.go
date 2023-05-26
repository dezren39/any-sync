package yamux

import (
	"context"
	"github.com/anyproto/any-sync/net/connutil"
	"github.com/hashicorp/yamux"
	"net"
	"time"
)

type yamuxConn struct {
	ctx    context.Context
	luConn *connutil.LastUsageConn
	addr   string
	*yamux.Session
}

func (y *yamuxConn) Open(ctx context.Context) (conn net.Conn, err error) {
	return y.Session.Open()
}

func (y *yamuxConn) LastUsage() time.Time {
	return y.luConn.LastUsage()
}

func (y *yamuxConn) Context() context.Context {
	return y.ctx
}

func (y *yamuxConn) Addr() string {
	return y.addr
}
