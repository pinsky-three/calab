package remote

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
)

// RemoteBinaryServer represents a remote server.
type RemoteBinaryServer struct {
	port int
	path string
	r    chi.Router
	data chan []byte
}

func (rs *RemoteBinaryServer) sessionHandler(w http.ResponseWriter, r *http.Request) {
	conn, _, _, err := ws.UpgradeHTTP(r, w)
	if err != nil {
		panic(err)
	}

	go func(data chan []byte) {
		defer conn.Close()

		for {
			if err = wsutil.WriteServerMessage(conn, ws.OpBinary, <-data); err != nil {
				log.Println(err)
				conn.Close()
				break
			}
		}
	}(rs.data)
}

func (rs *RemoteBinaryServer) init() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get(rs.path, rs.sessionHandler)

	rs.r = r
}

func (rs *RemoteBinaryServer) launch() {
	http.ListenAndServe(fmt.Sprintf(":%d", rs.port), rs.r)
}

// Run listen and server http chi router.
func (rs *RemoteBinaryServer) Run() {
	rs.launch()
}

// NewBinaryRemote create a new remote server.
func NewBinaryRemote(port int, route string, dataSource chan []byte) *RemoteBinaryServer {
	rs := &RemoteBinaryServer{
		port: port,
		path: route,
		data: dataSource,
	}

	rs.init()

	return rs
}
