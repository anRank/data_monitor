package device_view

import (
	"data_monitor/pkg/device/device_model"
	"data_monitor/pkg/device/device_service"
	"fmt"
	"strings"
)

type CreateDeviceRequest struct {
	Name     *string
	AreaId   *string
	AreaName *string
	Owners   *string
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

type Device2UserRequest struct {
	UserId     *int64  `json:"user_id"`
	UserName   *string `json:"user_name"`
	DeviceId   *int64  `json:"device_id"`
	DeviceName *string `json:"device_name"`
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
	if r.DeviceId == nil {
		errMsgs = append(errMsgs, fmt.Sprint("device id cannot be nil"))
	}
	if r.UserId == nil {
		errMsgs = append(errMsgs, fmt.Sprint("device id cannot be nil"))
	}

	e, err := device_service.IsRelationExisted(*r.UserId, *r.DeviceId)
	if err != nil {
		errMsgs = append(errMsgs, err.Error())
	}
	if e{
		errMsgs = append(errMsgs, fmt.Sprint("relation already existed"))
	}

	if len(errMsgs) > 0 {
		return fmt.Errorf(strings.Join(errMsgs, "; "))
	}
	return nil
}

type UpdateDeviceRequest struct {
	Id     *int64  `json:"id"`
	Name   *string `json:"name"`
	AreaId *string `json:"area_id"`
}

func (r UpdateDeviceRequest) CheckParams() error {
	if r.Id == nil {
		return fmt.Errorf("device id cannot be nil")
	}
	return nil
}

func (r UpdateDeviceRequest) CopyToModel(m *device_model.Device) {
	if r.Id != nil {
		m.Id = *r.Id
	}
	if r.Name != nil {
		m.Name = *r.Name
	}
	if r.AreaId != nil {
		m.AreaId = *r.AreaId
	}
}
