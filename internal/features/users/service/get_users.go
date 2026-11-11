package users_service

import (
	"context"
	"fmt"

	"github.com/XZSDFVB213/golang-todoapp/internal/core/domain"
	core_errors "github.com/XZSDFVB213/golang-todoapp/internal/core/errors"
)

func (s *UsersService) GetUsers(ctx context.Context, limit *int, offset *int) ([]domain.User, error) {
	if limit != nil && *limit < 0 {
		return nil, fmt.Errorf("Limit must be non-negtive: %w", core_errors.ErrInvalidArgument)
	}
	if offset != nil && *offset < 0 {
		return nil, fmt.Errorf("offset must be non-negtive: %w", core_errors.ErrInvalidArgument)
	}
	users, err := s.usersRepository.GetUsers(ctx,limit,offset)
	if err != nil {
		return nil, fmt.Errorf("Get users from repo: %w", err)
	}
	return users, nil
}