package device_model

type Device struct {
	Id 			int64
	Name 		string
	Owners 		string
	AreaId		string
}

type DeviceUser struct {
	Id 			int64
	UserId 		int64
	Username 	string
	DeviceId 	int64
	DeviceName  string
}

func (Device)TableName() string{
	return "device"
}

func (DeviceUser)TableName() string{
	return "device_user"
}