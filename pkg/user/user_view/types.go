package user_view

import (
	"data_monitor/pkg/user/user_model"
	"fmt"
	"strings"
)

type CreateUserRequest struct {
	Name    *string
	IsAdmin *int
}

type DeleteUserRequest struct {
	Id   *int
	Name *string
}

type UpdateUserRequest struct {
	Id      *int64
	Name    *string
	IsAdmin *int
}

func (r CreateUserRequest) CopyToModel(m *user_model.User) {
	if r.Name != nil {
		m.Name = *r.Name
	}
	if r.IsAdmin != nil {
		m.IsAdmin = *r.IsAdmin
	}
}

func (r CreateUserRequest) CheckParams() error {
	var errMsgs []string
	if r.Name == nil {
		errMsgs = append(errMsgs, fmt.Sprint("username cannot be empty string"))
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "; "))
	}
	return nil
}

func (r UpdateUserRequest) CopyToModel(m *user_model.User) {
	if r.Id != nil {
		m.Id = *r.Id
	}
	if r.Name != nil {
		m.Name = *r.Name
	}
	if r.IsAdmin != nil {
		m.IsAdmin = *r.IsAdmin
	}
}

func (r UpdateUserRequest) CheckParams() error {
	var errMsgs []string
	if r.Name == nil {
		errMsgs = append(errMsgs, fmt.Sprint("username cannot be empty string"))
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "; "))
	}
	return nil
}
