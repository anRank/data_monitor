package user_model

type User struct {
	Id 			int64
	Name 		string
	IsAdmin 	int
	IsDeleted 	int
}

func (User)TableName() string{
	return "user"
}
