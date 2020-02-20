package controllers

import (
	"yinji/service"
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models/bind"
	"yinji/config"
)

type AudioController struct {
	BeegoController
}

/**
* 输出 对应的 页面的 信息
*/
func (controller *AudioController) Player() {
	controller.TplName = "music/player.html"
}

func (controller *AudioController) IndexPage() {
	controller.Resource("index/index.html")
}

/**
	根据对应的 信息 ， 来进行对应的 信息 查询
 */

func (self *AudioController) SearchByString() {

	//获取对应的 xixni
	var content string = self.GetString("content");

	var startLimit , startLimitError = self.GetInt("startLimit")

	var endLimit , endLimitError = self.GetInt("endLimit")

	if startLimitError != nil {
		startLimit = 0
	}

	if endLimitError != nil {
		endLimit = startLimit + 10
	}

	//获取对应的 信息 ， 之后 来对对应的 信息 来进行 查询
	var audio_service = service.GetAudioServiceInstance();
	var ormService = db.GetOrmServiceInstance()

	//获取对应的方法
	var audioList = audio_service.SearchAudio(content , startLimit , endLimit );

	var audioAndUserList interface{} 

	//先暂时 用这个方法 来进行输出 信息
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		audio_service.SearchAudioAndUser( o , audioList)
		return nil, nil
	})
	
	self.Json(audioAndUserList)
}

func (self *AudioController) Favorites() {
	var music_service = service.GetAudioServiceInstance();
	var result = music_service.Favorites();
	self.Json(result);
}

func (self *AudioController) Delete() {
	//获取目标的 id
	var id , err = self.GetInt64("id" )

	if err != nil {
		self.FailJson(err)
		return
	}

	var httpFileService = service.GetHttpFileServiceInstance()
	var ormService = db.GetOrmServiceInstance()

	//查询对应的 code 的 信息
	var audio = bean.Audio{}
	audio.Id = id
	var _ , readErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var err = o.Read(&audio , "Id")
		return audio, err
	})

	//这里说明读取出错 ， 数据库中并不存在该属性
	if readErr != nil  {
		self.FailJson(readErr)
		return ;
	}

	//之后进行删除
	var _ , tranErr = ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		var _ , err = o.Delete(&audio , "Id")

		if err != nil {
			//上传错误
			return audio , err
		}

		var absPath = httpFileService.GetAudioFileName( &audio )

		var deleteAudioErr = httpFileService.Delete( absPath )

		//删除文件夹错误
		if deleteAudioErr != nil{
			return audio , deleteAudioErr
		}

		//如果audio.Image 为空，我们则不进行删除
		if audio.Image != "" {
			//会返回对应的信息 ， 删除是否出现问题 ， 还不如被注意
			var _ = httpFileService.Delete( audio.Image)

		}
		return nil , nil

	})

	if tranErr != nil {
		self.FailJson(tranErr)
		return
	}

	self.Json( audio )
}

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME = "name";

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_LENGTH = "length";

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_AUDIO = "audio[]"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_IMAGE = "image"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_INTRODUCTION ="introduction"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_USERID = "userId"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID = "id"



/**
	负责音频上传的 api 接口
 */
func (self *AudioController) AudioUpload() {

	var httpFileService = service.GetHttpFileServiceInstance()

	var audio = new(bean.Audio)

	var name = self.GetString(AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME)

	var length , _ = self.GetInt( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_LENGTH )

	var introduction = self.GetString(AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_INTRODUCTION)

	//获取对应的 userId 的 信息
	var userId , getUserIdErr = self.GetInt64( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_USERID )

	if getUserIdErr != nil {
		//当出现 getUserId 的错误 ,
		self.FailJson( getUserIdErr )
		return
	}

	//获取对应相对应的 信息

	var instance = db.GetOrmServiceInstance()

	//下面输入对应的信息
	//audio.New()
	audio.Name = name
	audio.Image = ""
	audio.UserId = userId
	audio.Url = ""
	audio.TimeLength = length
	audio.Introduction = introduction

	//直接获取对应的 信息

	var file , fileHeader ,  err = self.GetFile( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_AUDIO )

	defer httpFileService.CloseMultipart( file )

	if( err != nil ){
		self.FailJson( err )
		return
	}
	
	var fileName = httpFileService.BuildAudioFileName(audio , fileHeader)

	var audioUrl , uploadAudioErr = httpFileService.UploadAudio( fileName , file )

	if uploadAudioErr != nil {
		self.FailJson( err )
		return
	}

	audio.Url = audioUrl

	//传输网音频之后 ， 便开始传输图片

	var imageFile , imageHeader , getImageErr = self.GetFile( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_IMAGE )

	//倘若 getImageErr 不为空 ， 才会顺便上传图片
	if getImageErr == nil {

		var newFileName = httpFileService.BuildAudioFileName(audio , imageHeader )

		var uploadFilePath , uploadFileErr = httpFileService.UploadImage( newFileName , imageFile )

		if uploadFileErr == nil{
			//设置对应的 Image 的 路径
			audio.Image = uploadFilePath
			//会不会错误那个另说
		}

	}

	var audioService = service.GetAudioServiceInstance()

	var _ , transacErr = instance.Transaction(func(o orm.Ormer) (interface{}, error) {
		var newErr = audioService.New( o , audio )
		return audio , newErr
	})

	if transacErr != nil {
		self.FailJson( transacErr )
		return
	}

	self.Json(audio);
}


func ( self *AudioController ) AudioUpdate() {
	var httpFileService = service.GetHttpFileServiceInstance()
	var ormService = db.GetOrmServiceInstance();
	var audio = new(bean.Audio)
	//默认设置httpFileService ,修改的对象

	var id , getIdErr = self.GetInt64( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID )

	if getIdErr != nil {
		self.FailJson( getIdErr )
		return
	}

	audio.Id = id;


	var _ ,  readErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		 var err = o.Read(audio , "Id")
		 return nil , err
	});

	if readErr != nil {
		//程序错误 ， 那么 就进行退出程序
		self.FailJson( readErr )
		return
	}

	//现根据对应的id 获取对应的 信息

	//获取文件
	var audioFile , audioFileHeader , audioErr = self.GetFile(AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_AUDIO)

	defer httpFileService.CloseMultipart( audioFile )

	if audioErr == nil {
		//倘若不存在 ， 那么 ， 我们就跳过

		var audioName = httpFileService.BuildAudioFileName( audio , audioFileHeader);

		//var audioUrl = "audio/" + audioName

		var audioUrl , _ = httpFileService.UploadAudio( audioName , audioFile )

		audio.Url = audioUrl

	}

	audio.Refresh()

	var name = self.GetString( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME )

	var introduction = self.GetString( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_INTRODUCTION )

	//下次输入对应的 信息

	audio.Introduction = introduction

	audio.Name = name

	var imageFile , imageHeader , imageErr = self.GetFile( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_IMAGE)

	defer httpFileService.CloseMultipart( imageFile )

	if imageErr == nil {
		var imageName = httpFileService.BuildAudioFileName( audio , imageHeader)
		var uploadImagePath , uploadImageErr = httpFileService.UploadImage( imageName , imageFile )

		if uploadImageErr == nil {
			audio.Image = uploadImagePath
		}
	}

	ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
		o.Update( audio )
		return nil , nil
	})
	self.Json( audio )
}


func (self *AudioController) AudioUploadPage() {
	self.Resource("upload/demo_1.html");
}

func( self * AudioController ) AudioUpdatePage(){
	self.Resource("upload/demo_2.html")
}

//获取目标用户的
func( self *AudioController) SearchAudioByUserId(){

	//获取目标的 id
	var userId , _ = self.GetInt64("userId")

	var page, getPageErr = self.GetInt("page")
	var count, getCountErr = self.GetInt("count")

	if getPageErr != nil {
		page = 0
	}

	if getCountErr != nil {
		count = 10
	}

	var offset = page * count

	var audioSlice []*bean.Audio

	//获取对应的
	var ormService = db.GetOrmServiceInstance()
	var audioService = service.GetAudioServiceInstance()
	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var tableName = bean.GetAudioTableName()
		var qs = o.QueryTable( tableName )
		qs.Filter("user_id" , userId).Limit( count , offset ).OrderBy("-create_time").All( &audioSlice )
		return nil, nil
	})

	audioService.ParseArr( &audioSlice )

	self.Json( audioSlice )
}


func( self *AudioController) AudioPblPage(){
	self.Resource("pbl/main.html")
}

func( self *AudioController ) FindAudioById(){

	var id  , getInt64Err = self.GetInt64( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID )

	if getInt64Err != nil {
		//那么返回错误的信息
		self.FailJson( getInt64Err )
		return
	}

	var audioService = service.GetAudioServiceInstance();

	var audio , findAudioError = audioService.FindAudioById( id )

	if findAudioError != nil {
		self.FailJson( findAudioError )
		return
	}

	self.Json( audio )
}

/**
	接受的信息：
	userId : 目标用户的id
	输出的信息：
	直接输出数量
 */
func ( self *AudioController ) AudioLen( ){

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	var ormService = db.GetOrmServiceInstance()

	var result , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var qt = o.QueryTable( bean.GetAudioTableName() )

		var count , countErr = qt.Filter("userId",userId).Count()
		return count , countErr
	})


	if jdbcErr != nil {
		self.FailJson( jdbcErr )
	}

	self.Json( result )

}

/**
	根据对应的输入属性，来搜索根据某种排行的数据
	接收的信息：
		audioId:userId
		type:排序的类型
		page:页数,
		count:一张页面输出多少信息
	目前暂时只支持对应的 browse( 浏览 )

*/
func ( self *AudioController ) AudioByDashboard(  ) {

	var userId , getUserIdErr = self.GetInt64("userId")

	if getUserIdErr != nil {
		self.FailJson( getUserIdErr )
		return
	}

	//获取当前页数
	var page , getPageErr = self.GetInt("page")

	if getPageErr != nil {
		page =  0 ;
	}

	var count, getCountErr = self.GetInt("count")

	if getCountErr != nil {
		count = config.LIMIT_COUNT
	}

	var startLimit = page * count

	var endLimit = startLimit + count
	
	var ormService = db.GetOrmServiceInstance()

	var sql = "SELECT a.* , ad.browse_count , ad.love_count , ad.comment_count ,ad.collection_count FROM audio a LEFT JOIN audio_dashboard ad ON a.id = ad.id WHERE a.user_id = ? ORDER BY ad.browse_count DESC LIMIT ? , ? ;"
	var audioArr []*bind.AudioAndDashboard
	/**
		输出结果
	 */
	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var _ , queryErr =  o.Raw( sql , userId , startLimit , endLimit ).QueryRows(&audioArr)
		return audioArr , queryErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	for _ , dashboard := range audioArr {
		dashboard.Parse()
	}

	//输出结果
	self.Json( audioArr )

}

/**
	获取单曲播放量最多的信息
	获取的信息
	page:页数
	count:一个页面上有多少的数据
	type:各种的类型 // 目前也仅仅只是对应的

	PS. 这里的代码只是暂时的操作
*/

func ( self *AudioController ) SearchMostBrowseAudio(){
	//获取当前页数
	var page , getPageErr = self.GetInt("page")

	if getPageErr != nil {
		page =  0 ;
	}

	var count, getCountErr = self.GetInt("count")

	if getCountErr != nil {
		count = config.LIMIT_COUNT
	}

	var startLimit = page * count

	var endLimit = startLimit + count

	//暂时的输出对应的信息

	var ormService = db.GetOrmServiceInstance()
	var sql = "SELECT a.* , ad.browse_count , ad.love_count , ad.comment_count ,ad.collection_count FROM audio a LEFT JOIN audio_dashboard ad ON a.id = ad.id  ORDER BY ad.browse_count DESC LIMIT ? , ? ;"
	var audioArr []*bind.AudioAndDashboard

	var _ , jdbcErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var _ , queryErr =  o.Raw( sql , startLimit , endLimit ).QueryRows(&audioArr)
		return audioArr , queryErr
	})

	if jdbcErr != nil {
		self.FailJson( jdbcErr )
		return
	}

	for _ , dashboard := range audioArr {
		dashboard.Parse()
	}

	//输出结果
	self.Json( audioArr )

}
