package users_transport_http

import "github.com/XZSDFVB213/golang-todoapp/internal/core/domain"

type UserDTOResponse struct {
	ID          int     `json:"id"`
	Version     int     `json:"version"`
	FullName    string  `json:"full_name"`
	PhoneNumber *string `json:"phone_number"`
}

func UserDTOFromDomain(user domain.User) UserDTOResponse {
	return UserDTOResponse{
		ID:          user.ID,
		Version:     user.Version,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
	}
}
func UsersDTOFromDomain(user []domain.User) []UserDTOResponse{
	usersDto := make([]UserDTOResponse, len(user))
	for i, user := range user{
		usersDto[i] = UserDTOFromDomain(user)
	}
	return usersDto
}