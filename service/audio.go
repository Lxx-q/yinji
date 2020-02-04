package service

import (
	"bytes"
	"yinji/service/db"
	"yinji/models/bean"
	"github.com/astaxie/beego/orm"
	"yinji/models/bind"
	"fmt"
)

type AudioService struct {
	ormService *db.OrmService
}

//相对应的 地址 ， 去掉 了 ip ,去掉了 项目头等七七八八的东西 ，
const DEFAULT_AUDIO_IMAGE_URL = "image/default.jpg"

func ( self *AudioService ) BuildImage( audio bean.Audio ) string {
	var image = audio.Image

	if image == "" {
		image = DEFAULT_AUDIO_IMAGE_URL
	}

	return image
}
//获取目标的信息
func ( self *AudioService) BuildUrl( audio *bean.Audio) string{
	// 我们利用这个属性来获取 相对应的 信息

	var buffer = bytes.Buffer{}

	var head = "http://localhost:8080/none/"

	buffer.WriteString( head );

	//buffer.WriteString( strconv.FormatInt(music.Code , 10) )
	buffer.WriteString( audio.Url )

	return buffer.String();
}

/**
	注意 ， 这个接口 ， 并不是我们寻常所看见的搜搜功能
 */
func (self *AudioService) SearchAudio( content string , startLimit int , endLimit int ) [] *bean.Audio{
	var user_list [] *bean.Audio
	var offset = endLimit - startLimit
	var _ , error = self.ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qs = o.QueryTable("audio");
		return qs.Filter("name__icontains",content).OrderBy("-modify_time").Limit( offset , startLimit).All( &user_list );
	});

	if( error != nil){
		return nil;
	}

	return user_list;

}

func( self *AudioService) SearchAudioAndUser( o orm.Ormer , audioList []*bean.Audio)  []*bind.AudioAndUser{

	var user_service = GetUserServiceInstance();

	var length = len(audioList);

	var audioAndUserList = make([]*bind.AudioAndUser ,0)

	//虽然效果可以达到预期， 但是效率却令人深思
	for index := 0 ; index < length ; index++{
		//输出 相对应的 信息
		var audio = audioList[ index ];

		var user , _ = user_service.FindUserById( o , audio.UserId );

		var audioAndUser = new(bind.AudioAndUser);

		audioAndUser.Audio = *audio;
		audioAndUser.BindUser( *user );

		audioAndUserList = append( audioAndUserList , audioAndUser );
	}


	return audioAndUserList;

}

func (self *AudioService) Favorites() []*bean.Audio{
	var user_list []*bean.Audio = []*bean.Audio{}
	return user_list;
}

func ( self * AudioService ) FindAudioById( id int64) ( *bean.Audio , error ){
	var audio = bean.Audio{}
	audio.Id = id
	var _ , readErr = self.ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var ormErr = o.Read(&audio)
		return &audio , ormErr
	})

	return &audio , readErr
}

//根据对应的id 来进行获取对应的信息
func (self *AudioService ) FindById ( o orm.Ormer , id int64 ) ( *bean.Audio , error ){

	var audio = bean.Audio{}
	audio.Id = id
	var  readErr = o.Read(&audio)
	return &audio , readErr

}

func ( self *AudioService ) AllByCollectionFolder(o orm.Ormer , folderId int64  ){
	var sql = "SELECT * FROM audio a WHERE a.id IN ( SELECT auc.audio_id FROM audio_user_collection auc WHERE auc.collection_folder_id = ?);"

	//当 folderId 等于 0 的时候
	if folderId == 0 {
		sql = "SELECT * FROM audio a WHERE a.id IN ( )"
	}


	fmt.Println(sql)


}

func ( self *AudioService ) New( o orm.Ormer  , audio *bean.Audio ) error {

	var _ , insertErr = o.Insert( audio )

	if insertErr != nil {
		return  insertErr
	}

	//初始化 ， 阅读记录信息
	var audioBrowseHistoryService = GetAUdioBrowseHistoryServiceInstance()

	var _ ,newErr = audioBrowseHistoryService.NewByAudioId( o , audio.Id )

	//初始化 ， 点赞 ， 转化 ， 收藏初始化信息
	var dashboradService = GetDashboardServiceInstance()
	dashboradService.NewByAudioId( o , audio.Id )

	return newErr
}

//之后我们 便可以开始获取信息

var AUDIO_SERVICE_INSTANCE =&AudioService{db.GetOrmServiceInstance()}

func GetAudioServiceInstance() *AudioService{
	return AUDIO_SERVICE_INSTANCE
}