package domain

import (
	"fmt"
	"regexp"

	core_errors "github.com/XZSDFVB213/golang-todoapp/internal/core/errors"
)

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber *string
}

func NewUser(id int, version int, fullname string, phoneNumber *string) User {
	return User{
		ID:          id,
		Version:     version,
		FullName:    fullname,
		PhoneNumber: phoneNumber,
	}
}
func NewUserUnitialized(fullname string, phoneNumber *string) User {
	return NewUser(UninitializedId, UninitializedVersion, fullname, phoneNumber)
}
func (u *User) Validate() error {
	fullNameLength := len([]rune(u.FullName))
	if fullNameLength < 3 || fullNameLength > 100 {
		return fmt.Errorf("Invalid `full name` len: %d: %w", fullNameLength, core_errors.ErrInvalidArgument)
	}
	if u.PhoneNumber != nil {
		phoneNumberLen := len([]rune(*u.PhoneNumber))
		if phoneNumberLen < 10 || phoneNumberLen > 15 {
			return fmt.Errorf("Invalid `phone number` `len: %d: %w", phoneNumberLen, core_errors.ErrInvalidArgument)
		}
		re := regexp.MustCompile(`^\+[0-9]+$`)
		if !re.MatchString(*u.PhoneNumber) {
			return fmt.Errorf("Invalid `PhoneNumber`: %w", core_errors.ErrInvalidArgument)
		}
	}
	return nil
}

type UserPatch struct {
	Fullname    Nullable[string]
	PhoneNumber Nullable[string]
}

func (p *UserPatch) Validate() error {
	if p.Fullname.Set && p.Fullname.Value == nil {
		return fmt.Errorf("`FullName` can't be patched to NULL:%w", core_errors.ErrInvalidArgument)
	}
	return nil
}
func (u *User) ApplyPatch(patch UserPatch) error {
	if err := patch.Validate(); err != nil {
		return fmt.Errorf("Validate user patch:%w", err)
	}
	tmp := *u
	if patch.Fullname.Set {
		tmp.FullName = *patch.Fullname.Value
	}
	if patch.PhoneNumber.Set {
		tmp.PhoneNumber = patch.PhoneNumber.Value
	}
	if err := tmp.Validate(); err != nil {
		return fmt.Errorf("Validate patched user: %w", err)
	}
	*u = tmp
	return nil
}
