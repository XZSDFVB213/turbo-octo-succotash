package core_http_server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	core_logger "github.com/XZSDFVB213/golang-todoapp/internal/core/logger"
	core_http_middleware "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/middleware"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux    *http.ServeMux
	config Config
	log    core_logger.Logger
	middleware []core_http_middleware.Middleware

}

func NewHTTPServer(
	config Config,
	log *core_logger.Logger,
	middleware ...core_http_middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:    http.NewServeMux(),
		config: config,
		log:    *log,
		middleware: middleware,
	}
}
func (s *HTTPServer) RegisterAPiRouter(routers ...*ApiVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.ApiVersion)
		
		s.mux.Handle(prefix+"/", http.StripPrefix(prefix, router.WithMiddleware()))
	}
}
func (s *HTTPServer) Run(ctx context.Context) error {
	mux := core_http_middleware.ChainMiddleware(s.mux, s.middleware...)
	server := &http.Server{
		Addr:    s.config.Addr,
		Handler: mux,
	}
	ch := make(chan error, 1)
	go func() {
		defer close(ch)
		s.log.Warn("Start http server", zap.String("Addr", s.config.Addr))
		if err := server.ListenAndServe(); err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				ch <- err
			}
		}
	}()
	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("listen and serve http : %w", err)
		}
	case <-ctx.Done():
		s.log.Warn("Shutdown http server")
		shutdownCtx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			_ = server.Close()
			return fmt.Errorf("Shutdown http server: %w", err)
		}
		s.log.Warn("Http server stopped")
	}
	return nil

}
