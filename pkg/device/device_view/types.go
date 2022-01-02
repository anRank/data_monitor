package device_view

import (
	"data_monitor/pkg/device/device_model"
	"fmt"
	"strings"
)

type CreateDeviceRequest struct {
	Name     *string
	AreaId   *string
	AreaName *string
	Owners   *string
}

type Device2UserRequest struct {
	UserId     *int64
	UserName   *string
	DeviceId   *int64
	DeviceName *string
}

func (r CreateDeviceRequest) CopyToModel(m *device_model.Device) {
	if r.Name != nil {
		m.Name = *r.Name
	}
	if r.AreaId != nil {
		m.AreaId = *r.AreaId
	}
	if r.Owners != nil {
		m.Owners = *r.Owners
	}
}

func (r CreateDeviceRequest) CheckParams() error {
	var errMsgs []string
	if r.Name == nil {
		errMsgs = append(errMsgs, fmt.Sprint("device name cannot be empty string"))
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "; "))
	}
	return nil
}

func (r Device2UserRequest) CopyToModel(m *device_model.DeviceUser) {
	if r.UserId != nil {
		m.UserId = *r.UserId
	}
	if r.UserName != nil {
		m.Username = *r.UserName
	}
	if r.DeviceId != nil {
		m.DeviceId = *r.DeviceId
	}
	if r.DeviceName != nil {
		m.DeviceName = *r.DeviceName
	}
}

func (r Device2UserRequest) CheckParams() error {
	var errMsgs []string
	if r.UserName == nil {
		errMsgs = append(errMsgs, fmt.Sprint("device name cannot be empty string"))
	}
	if r.DeviceName == nil {
		errMsgs = append(errMsgs, fmt.Sprint("device name cannot be empty string"))
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "; "))
	}
	return nil
}


type UpdateDataRequest struct {
	Id *int64
	Name *string
	AreaId *string
	Owners *string
}

func (r UpdateDataRequest) CheckParams() error {
	if r.Id == nil {
		return fmt.Errorf("device id cannot be nil")
	}
	return nil
}

func (r UpdateDataRequest) CopyToModel(m *device_model.Device) {
	if r.Id != nil {
		m.Id = *r.Id
	}
	if r.Name != nil {
		m.Name = *r.Name
	}
	if r.AreaId != nil {
		m.AreaId = *r.AreaId
	}
	if r.Owners != nil {
		m.Owners = *r.Owners
	}
}