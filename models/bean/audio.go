package bean

import (
	"time"
	"yinji/models/base"
)

type Audio struct {

	//Id string `orm:"column(id);pk"`
	Name string `orm:"column(name)"  json:"name"`
	Url string `orm:"column(url)" json:"url"`
	Image string `orm:"column(image)" json:"image"`
	UserId int64 `orm:"column(user_id)" json:"userId"`
	TimeLength int `orm:"column(time_length)" json:"timeLength"`
	//商业介绍
	Introduction string `orm:"column(introduction)" json:"introduction"`
	AudioAlbumId int64 `orm:"column(audio_album_id)" json:"audioAlbumId"`

	//EntityBase
	base.IdAndCodeStruct
	base.CreateTimeAndModifyTimeStruct
}


func ( self *Audio) TableName() string {
	return GetAudioTableName()
}

func GetAudioTableName() string{
	return "audio";
}

func ( self *Audio ) NewEntity( t time.Time) {
	self.CreateTimeAndModifyTimeStruct.NewEntity( t )
	self.IdAndCodeStruct.NewEntity( t )

}

func ( self *Audio ) New(){
	var t = time.Now()
	self.NewEntity( t )
}

/*
func( self *Audio) NewToDbFunc() func(o orm.Ormer) (interface{}, error){

	return func(o orm.Ormer) (interface{}, error) {
		self.New();
		self.Code = "m_" +  strconv.FormatInt(self.Id,10)
		return o.Insert( self );
	};
}

func( self *Audio) NewToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction( self.NewToDbFunc());
}

func ( self *Audio) RefreshToDbFunc() func(o orm.Ormer) (interface{}, error){


	return func(o orm.Ormer) (interface{}, error) {
		self.Refresh();
		return o.Update( self );
	};

}



func (self *Audio) RefreshToDb(){
	var ormService = db.GetOrmServiceInstance();
	ormService.Transaction(self.RefreshToDbFunc());
}

*/

