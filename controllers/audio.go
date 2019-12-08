package controllers

import (
	"yinji/service"
	"yinji/models/bean"
	"yinji/service/db"
	"github.com/astaxie/beego/orm"
	"yinji/models"
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

	//获取对应的方法
	var audioList = audio_service.SearchAudio(content , startLimit , endLimit );

	//先暂时 用这个方法 来进行输出 信息

	var audioAndUserList = audio_service.SearchAudioAndUser(audioList);
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
		var response = models.FailResponse( err )
		self.Json( response )
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
		var response = models.FailResponse(readErr)
		self.Json( response )
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
		var response = models.FailResponse( tranErr )
		self.Json( response )
		return
	}

	self.Json( audio )
}

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME = "name";

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_LENGTH = "length";

const AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_AUDIO = "audio[]"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_IMAGE = "image"

const AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID = "id"



/**
	负责音频上传的 api 接口
 */
func (self *AudioController) AudioUpload() {

	var httpFileService = service.GetHttpFileServiceInstance()

	var audio = new(bean.Audio)

	var name = self.GetString(AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME)

	var length , _ = self.GetInt( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_LENGTH )

	var instance = db.GetOrmServiceInstance()

	audio.New()
	audio.Name = name
	audio.Image = ""
	audio.UserId = 2
	audio.Url = ""
	audio.TimeLength = length

	//直接获取对应的 信息

	var file , fileHeader ,  err = self.GetFile( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_AUDIO )


	defer httpFileService.CloseMultipart( file )


	if( err != nil ){
		self.Fail( err )
		return
	}

	//var dir, _ = os.Getwd();
	//准备进行 相对应的 信息的 输入


	var fileName = httpFileService.BuildAudioFileName(audio , fileHeader)

	var audioUrl , uploadAudioErr = httpFileService.UploadAudio( fileName , file )

	if uploadAudioErr != nil {
		self.Fail( err )
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

	instance.Transaction(func(o orm.Ormer) (interface{}, error) {
		return o.Insert(audio)
	})


	self.Json(audio);
}


func ( self *AudioController ) AudioUpdate() {
	var httpFileService = service.GetHttpFileServiceInstance()
	var ormService = db.GetOrmServiceInstance();
	var audio = new(bean.Audio)
	//默认设置httpFileService ,修改的对象

	var id , getIdErr = self.GetInt64( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID )

	if getIdErr != nil {
		self.Fail( getIdErr )
		return
	}

	audio.Id = id;


	var _ ,  readErr = ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		 var err = o.Read(audio , "Id")
		 return nil , err
	});

	if readErr != nil {
		//程序错误 ， 那么 就进行退出程序
		self.Fail( readErr )
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

	audio.TimeLength = 30

	audio.Refresh()

	var name = self.GetString( AUDIO_CONTROL_AUDIOUPLOAD_PARMTER_NAME )

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

	var startLimit , startLimitError = self.GetInt("startLimit")
	var endLimit , endLimitError = self.GetInt("endLimit")

	if startLimitError != nil {
		startLimit = 0
	}

	if endLimitError != nil {
		endLimit = startLimit + 10
	}

	var offset = endLimit - startLimit

	var audioSlice []bean.Audio

	//获取对应的
	var ormService = db.GetOrmServiceInstance()

	ormService.Jdbc(func(o orm.Ormer) (interface{}, error) {
		var tableName = bean.GetAduioTableName()
		var qs = o.QueryTable( tableName )
		qs.Filter("user_id" , userId).Limit( offset , startLimit ).OrderBy("-create_time").All( &audioSlice )
		return nil, nil
	})

	self.Json( audioSlice )
}


func( self *AudioController) AudioPblPage(){
	self.Resource("pbl/main.html")
}

//
func( self *AudioController ) FindAudioById(){

	var id  , getInt64Err = self.GetInt64( AUDIO_CONTROL_AUDIOUPLOAD_PARAMTER_ID )

	if getInt64Err != nil {
		//那么返回错误的信息
		self.Fail( getInt64Err )
		return
	}

	var audioService = service.GetAudioServiceInstance();

	var audio , findAudioError = audioService.FindAudioById( id )

	if findAudioError != nil {
		self.Fail( findAudioError )
		return
	}

	self.Json( audio )
}