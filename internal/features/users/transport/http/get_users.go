package users_transport_http

import (
	"fmt"
	"net/http"

	core_logger "github.com/XZSDFVB213/golang-todoapp/internal/core/logger"
	core_http_utils "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/utils"
	core_http_response "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/response"
)
type GetUsersResponse []UserDTOResponse
func (h *UsersHTTPHandler) GetUsers(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	limit, offset, err := getLimitOffsetQueryParams(r)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get 'limit/offset' limit param")
		return
	}
	userDomains, err := h.usersService.GetUsers(ctx, limit, offset)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to Get Users")
		return
	}
	response := GetUsersResponse(UsersDTOFromDomain(userDomains))
	responseHandler.JSONResponse(response, http.StatusOK)
}
func getLimitOffsetQueryParams(r *http.Request) (*int, *int, error){
	limit, err := core_http_utils.GetIntQueryParam(r, "limit")
	if err != nil{
		return nil, nil, fmt.Errorf("get 'limit' query param: %w", err)
	}
	offset, err := core_http_utils.GetIntQueryParam(r, "offset")
	if err != nil{
		return nil, nil, fmt.Errorf("get 'offset' query param: %w", err)
	}
	return limit, offset, nil
}