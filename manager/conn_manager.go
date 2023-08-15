package manager

import (
	"github.com/kataras/iris/v12/websocket"
	"sync"
)

var ConnPool connPool

type connPool struct {
	sync.RWMutex
	conn map[string]*websocket.Conn
	once sync.Once
}

func (p *connPool) AddConn(id string, conn *websocket.Conn) {
	p.Lock()
	defer p.Unlock()
	p.conn[id] = conn
}

func (p *connPool) GetConn(id string) (*websocket.Conn, bool) {
	p.RLock()
	defer p.RUnlock()
	c, ok := p.conn[id]
	return c, ok
}

func (p *connPool) AllConn() map[string]*websocket.Conn {
	p.RLock()
	defer p.RUnlock()
	_connPool := make(map[string]*websocket.Conn)
	for _id, _conn := range p.conn {
		_connPool[_id] = _conn
	}
	return _connPool
}

func (p *connPool) DelConn(id string) {
	p.Lock()
	defer p.Unlock()
	delete(p.conn, id)
}

func init() {
	ConnPool = connPool{
		RWMutex: sync.RWMutex{},
		conn:    map[string]*websocket.Conn{},
		once:    sync.Once{},
	}
}
