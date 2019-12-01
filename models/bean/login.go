package bean

type Login struct {

	Id int64  `orm:"column(id);pk"`
	Password string
}

func (login *Login) TableName() string {
	return getLoginTableName();
}


func getLoginTableName() string{
	return "login";
}

