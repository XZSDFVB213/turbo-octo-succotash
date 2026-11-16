package core_http_server

import (
	"fmt"
	"net/http"

	core_http_middleware "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/middleware"
)

type ApiVersion string

var (
	ApiVersion1 = ApiVersion("v1")
	ApiVersion2 = ApiVersion("v2")
	ApiVersion3 = ApiVersion("v3")
)

type ApiVersionRouter struct {
	*http.ServeMux
	ApiVersion ApiVersion
	middleware []core_http_middleware.Middleware
}

func NewApiVersion(newAPIVersion ApiVersion, middleware ...core_http_middleware.Middleware) *ApiVersionRouter {
	return &ApiVersionRouter{
		ServeMux:   http.NewServeMux(),
		ApiVersion: newAPIVersion,
		middleware: middleware,
	}
}
func (r *ApiVersionRouter) RegisterRoutes(routes ...Route) {
	for _, route := range routes {
		pattern := fmt.Sprintf("%s %s", route.Method, route.Path)
		
		r.Handle(pattern, route.WithMiddleware())
	}
}
func (r *ApiVersionRouter) WithMiddleware() http.Handler{
	return core_http_middleware.ChainMiddleware(
			r,
			r.middleware...
		)
}
