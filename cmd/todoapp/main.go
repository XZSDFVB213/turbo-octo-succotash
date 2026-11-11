package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_logger "github.com/XZSDFVB213/golang-todoapp/internal/core/logger"
	core_postgres_pool "github.com/XZSDFVB213/golang-todoapp/internal/core/repository/postgres/pool"
	core_http_middleware "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/middleware"
	core_http_server "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/server"
	users_postgres_repository "github.com/XZSDFVB213/golang-todoapp/internal/features/users/repository/postgres"
	users_service "github.com/XZSDFVB213/golang-todoapp/internal/features/users/service"
	users_transport_http "github.com/XZSDFVB213/golang-todoapp/internal/features/users/transport/http"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()
	log, err := core_logger.NewLogger(core_logger.NewConfigMust())
	if err != nil {
		fmt.Println("Failed to init application logger")
		os.Exit(1)
	}
	defer log.Close()
	log.Debug("Initializing postgres connection pool")
	pool, err := core_postgres_pool.NewConnectionPool(
		core_postgres_pool.NewConfigMust(),
		ctx,
	)
	if err != nil {
		log.Fatal("failed to init postgres connection pool", zap.Error(err))
	}
	defer pool.Close()
	log.Debug("Initializing feature", zap.String("feature", "users"))
	usersRepository := users_postgres_repository.NewUsersRepository(pool)
	usersService := users_service.NewUsersService(usersRepository)
	usersTransportHTTP := users_transport_http.NewUsersHTTPHandler(usersService)
	log.Debug("Initializing http server")
	ApiVersionRouter := core_http_server.NewApiVersion(core_http_server.ApiVersion1)
	ApiVersionRouter.RegisterRoutes(usersTransportHTTP.Routes()...)
	httpServer := core_http_server.NewHTTPServer(core_http_server.NewConfigMust(),
		log,
		core_http_middleware.RequestID(),
		core_http_middleware.Logger(log),
		core_http_middleware.Panic(),
		core_http_middleware.Trace(),
	)
	httpServer.RegisterAPiRouter(ApiVersionRouter)
	if err := httpServer.Run(ctx); err != nil {
		log.Error("Failed to start http server", zap.Error(err))
	}
}
