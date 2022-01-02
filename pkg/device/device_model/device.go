package device_model

type Device struct {
	Id 			int64
	Name 		string
	Owners 		string
	AreaId		string
	IsDeleted   int
}

type DeviceUser struct {
	Id 			int64
	UserId 		int64
	Username 	string
	DeviceId 	int64
	DeviceName  string
	IsDeleted   int
}

func (Device)TableName() string{
	return "device"
}

func (DeviceUser)TableName() string{
	return "device_user"
}