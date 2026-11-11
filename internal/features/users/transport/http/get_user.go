package users_transport_http

import (
	"net/http"

	core_logger "github.com/XZSDFVB213/golang-todoapp/internal/core/logger"
	core_http_utils "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/http/utils"
	core_http_response "github.com/XZSDFVB213/golang-todoapp/internal/core/transport/response"
)

type GetUserResponse UserDTOResponse

func (h *UsersHTTPHandler) GetUser(rw http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	logger := core_logger.FromContext(ctx)
	responseHandler := core_http_response.NewHTTPResponseHandler(logger, rw)
	userId, err := core_http_utils.GetIntPathValue(r, "id")
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get userID path value")
		return
	}
	user, err :=  h.usersService.GetUser(ctx, userId)
	if err != nil {
		responseHandler.ErrorResponse(err, "Failed to get user")
		return
	}
	response := GetUserResponse(UserDTOFromDomain(user))
	responseHandler.JSONResponse(response, http.StatusOK)
}
