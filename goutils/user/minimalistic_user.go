package user

import (
	"encoding/json"

	"github.com/bvv_utils/goutils/mistake"
)

const (
	maxUserLoginLen = 40
	maxUserPswLen   = 40
)

// MinimalisticUser is realization of MinimalisticUserI in minimalistic form
type MinimalisticUser struct {
	// nullable: false
	// example: mivanov
	// Extensions:
	// x-order: "0"
	Login string `json:"login"`
	// nullable:false
	// example: lUhqd8hR
	// Extensions:
	// x-order: "1"
	Password string `json:"password"`
}

func (m *MinimalisticUser) GetLogin() *string {
	if m == nil {
		return nil
	}
	if len(m.Login) == 0 {
		return nil
	}
	return &m.Login
}

func (m *MinimalisticUser) GetPassword() *string {
	if m == nil {
		return nil
	}
	if len(m.Password) == 0 {
		return nil
	}
	return &m.Password
}

// type to validate MinimalisticUser input:
type inputMinimalisticUserValidation MinimalisticUser

func (m *MinimalisticUser) UnmarshalJSON(b []byte) error {
	var av inputMinimalisticUserValidation
	if err := json.Unmarshal(b, &av); err != nil {
		return err
	}
	*m = MinimalisticUser(av)
	if s := m.GetLogin(); s == nil {
		return mistake.ErrUserLoginNil
	} else if len(*s) > maxUserLoginLen {
		return mistake.ErrUserLoginLen
	}
	if s := m.GetPassword(); s == nil {
		return mistake.ErrUserPswNil
	} else if len(*s) > maxUserPswLen {
		return mistake.ErrUserPswLen
	}
	return nil
}
