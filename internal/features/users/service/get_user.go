package users_service

import (
	"context"
	"fmt"

	"github.com/XZSDFVB213/golang-todoapp/internal/core/domain"
)

func (s *UsersService) GetUser(ctx context.Context, id int) (domain.User, error){
	user, err := s.usersRepository.GetUser(ctx, id)
	if err != nil {
		return domain.User{}, fmt.Errorf("Get user from repo: %w", err)
	}
	return user, nil
}
