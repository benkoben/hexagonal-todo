package http

import (
	"fmt"
    "net/http"
    "os/signal"
    "os"
    "context"
    "syscall"
    "time"
	"github.com/gin-gonic/gin"
    service "github.com/benkoben/hexagonal-todo/internal/core/service"
)

const (
    defaultAddr = "0.0.0.0"
    defaultPort = "8080"
)

type logger interface {
    Printf(format string, v...any)
    Println(v...any)
    Fatalf(format string, v...any)
    Fatalln(v...any)
}

type ServerOptions struct {
    Address string
    Port string
    Router *gin.Engine
    Log logger
}

// server is a wrapper for gin.Engine and logging and the todo service
type server struct {
    srv *http.Server
    // Logger implements function for writing logs 
    log logger
}

func NewServer(o ServerOptions, ls *service.ListService) (*server, error) {
    if o.Address == "" {
        o.Address = defaultAddr
    }

    if o.Port == "" {
        o.Port = defaultPort
    }

    if o.Router == nil {
        return nil, fmt.Errorf("router field cannot be nil in http server options")
    }

    if o.Log == nil {
        return nil, fmt.Errorf("log field cannot be nil in http server options")
    }

    bindSocket := fmt.Sprintf("%s/%s", o.Address, o.Port)
    router, err := NewRouter(service.ListService{})
    if err != nil {
        return nil, fmt.Errorf("could not create new router: %v", err)
    }
    
    return &server{
        srv: &http.Server{
            Addr: bindSocket,
            Handler: router.Engine,
        },
    }, nil
}

func (s server)Start() {
    go func(){
        if err := s.srv.ListenAndServe(); err != nil {
             s.log.Fatalf("could not start http server: %v", err)
        }
        s.log.Printf("Server stopped.")
    }()
    s.log.Printf("Server is listening on %s", s.srv.Addr)
    s.stop()
}

func (s server)stop() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	sig := <-stop

	s.log.Printf("Shutting down server. Reason: %s.\n", sig.String())

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	s.srv.SetKeepAlivesEnabled(false)
	if err := s.srv.Shutdown(ctx); err != nil {
		s.log.Printf("Server shutdown: %v.\n", err)
	}
}

// write sends a response to the client
func write (w http.ResponseWriter, response response) {
	w.Header().Set("Content-Type", "application/json;charset=UTF-8")
	w.WriteHeader(response.Code())
	w.Write(response.JSON())
}
