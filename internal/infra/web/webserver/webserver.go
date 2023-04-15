package webserver

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type EchoHandlerCompose struct {
	method  string
	path    string
	handler echo.HandlerFunc
}

type WebServer struct {
	Router        *echo.Echo
	Handlers      []EchoHandlerCompose
	WebServerPort string
}

func NewWebServer(serverPort string) *WebServer {
	return &WebServer{
		Router:        echo.New(),
		Handlers:      make([]EchoHandlerCompose, 0),
		WebServerPort: serverPort,
	}
}

func (s *WebServer) AddHandler(method, path string, handler echo.HandlerFunc) {
	s.Handlers = append(s.Handlers, EchoHandlerCompose{
		method: method, path: path, handler: handler,
	})
}

func (s *WebServer) Start() {
	s.Router.Use(middleware.Logger())
	for _, handlerCompose := range s.Handlers {
		s.Router.Add(handlerCompose.method, handlerCompose.path, handlerCompose.handler)
	}
	http.ListenAndServe(s.WebServerPort, s.Router)
}
