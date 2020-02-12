package bean

type Login struct {

	//Id int64  `orm:"column(id);pk"`
	Acount string `orm:"column(account)" json:"account"`
	Password string `orm:"column(password)" json:"password"`
	Tel *string `orm:"column(tel)" json:"tel"`
	Email *string `orm:"column(email)" json:"email"`
	BaseEntity
}

func (login *Login) TableName() string {
	return GetLoginTableName();
}


func GetLoginTableName() string{
	return "login";
}



