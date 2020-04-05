package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"io"
)

type ResourceAudioService struct {

}

/**
	根据对应的 insert 来生成对应的信息
 */
func ( self *ResourceAudioService ) Insert( o orm.Ormer , resourceImage *bean.ResourceAudio) error {
	var _  , insertErr = o.Insert(resourceImage)
	return insertErr
}

/**
	resource
 */

const (
	RESOURCES_AUDIOS_FATHER_PATH = "resources/audio/n"
)

/**

 */
func ( self *ResourceAudioService) UploadAudio( o orm.Ormer , filename string  , reader io.Reader ) (*bean.ResourceAudio,error) {
	var hostResourceService = GetHostResourceServiceInstance()

	//获取对应的fatherPath
	var fatherPath = RESOURCES_AUDIOS_FATHER_PATH

	var resource , _  , uploadErr = hostResourceService.Upload(o , fatherPath ,filename , reader )
	var resourceAudio = bean.ResourceAudio{}
	resourceAudio.New()
	resourceAudio.OriginResourceId = resource.Id

	var insertErr = self.Insert(o , &resourceAudio)

	if insertErr != nil {
		return nil , insertErr
	}

	return &resourceAudio, uploadErr
}

func ( self *ResourceAudioService) UpdateAudio( o orm.Ormer , resourceAudio *bean.ResourceAudio,filename string  , reader io.Reader ) (*bean.ResourceAudio,error) {
	var hostResourceService = GetHostResourceServiceInstance()

	//获取对应的fatherPath
	var fatherPath = RESOURCES_AUDIOS_FATHER_PATH

	//根据resourceImage 获取 resource 的信息
	var hostResource , getErr = GetHostResource(o,resourceAudio.OriginResourceId);

	if getErr != nil {
		return nil , getErr
	}

	var updateErr = hostResourceService.Update(o , hostResource, fatherPath ,filename , reader )

	return resourceAudio , updateErr
}


var RESOURCE_Audio_SERVICE_INSTANCE = &ResourceAudioService{}

func GetResourceAudioServiceInstance() *ResourceAudioService{
	return RESOURCE_Audio_SERVICE_INSTANCE
}
