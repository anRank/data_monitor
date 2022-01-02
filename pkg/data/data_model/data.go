package data_model

type Data struct {
	Id         int64
	Data1      float64
	Data2      float64
	Data3      float64
	Data4      float64
	Data5      float64
	Data6      float64
	Data7      float64
	Data8      float64
	Data9      float64
	DeviceId   int64
	AreaId     string
	CreateTime string
}

func (Data) TableName() string {
	return "data"
}