package service

import (
	"github.com/astaxie/beego/orm"
	"yinji/models/bean"
	"io"
	"yinji/ffmpeg"
	"yinji/utils"
	"strconv"
	"yinji/service/db"
	"os"
)

type ResourceImageService struct {
	resourceService *ResourceService
}

/**
	根据对应的 insert 来生成对应的信息
 */
func ( self *ResourceImageService ) Insert( o orm.Ormer , resourceImage *bean.ResourceImage) error {
	var _  , insertErr = o.Insert(resourceImage)
	return insertErr
}

const (
	RESOURCES_IMAGES_FATHER_PATH = "resources/image/r"
)

/**

 */
func ( self *ResourceImageService) UploadImage( o orm.Ormer , filename string  , reader io.Reader ) (*bean.ResourceImage,error) {
	var hostResourceService = GetHostResourceServiceInstance()

	//获取对应的fatherPath
	var fatherPath = RESOURCES_IMAGES_FATHER_PATH

	var resource , hostresource  , uploadErr = hostResourceService.Upload(o , fatherPath ,filename , reader )
	var resourceImage = bean.ResourceImage{}
	resourceImage.New()
	resourceImage.OriginResourceId = resource.Id
	var insertErr = self.Insert(o , &resourceImage)
	if insertErr != nil {
		return nil , insertErr
	}

	self.thumb(hostresource.Path , filename , &resourceImage)

	return &resourceImage , uploadErr
}

func ( self *ResourceImageService) UpdateImage(o orm.Ormer, resourceImage *bean.ResourceImage,filename string  , reader io.Reader ) (*bean.ResourceImage,error) {
	var hostResourceService = GetHostResourceServiceInstance()

	//获取对应的fatherPath
	var fatherPath = RESOURCES_IMAGES_FATHER_PATH

	//根据resourceImage 获取 resource 的信息
	var hostResource , getErr = GetHostResource(o,resourceImage.OriginResourceId);

	if getErr != nil {
		return nil , getErr
	}

	var updateErr = hostResourceService.Update(o , hostResource, fatherPath ,filename , reader )

	return resourceImage , updateErr
}

func ( self *ResourceImageService ) thumb( origin_name string , filename string,resourceImage *bean.ResourceImage){
	var fileService = GetFileServiceInstance()
	var httpFileService = GetHttpFileServiceInstance()
	var hostResourceService = GetHostResourceServiceInstance()
	var ormService = db.GetOrmServiceInstance()
	go func() {

		var suffix  = fileService.Ext(filename)
		var new_name = RESOURCES_IMAGES_FATHER_PATH +"/" + strconv.FormatInt(utils.Rann(6) ,10) + suffix
		origin_name = httpFileService.BuildServerPath( origin_name )
		new_name = httpFileService.BuildServerPath(new_name)

		var thumbErr = ffmpeg.Thumb( origin_name ,  new_name)
		if thumbErr!= nil {
			return
		}
		defer os.Remove(new_name)
		var file ,openErr = os.Open(new_name)
		if openErr != nil {
			return
		}
		ormService.Transaction(func(o orm.Ormer) (interface{}, error) {
			var resource , _ , uploadErr = hostResourceService.Upload(o,RESOURCES_IMAGES_FATHER_PATH,new_name,file)

			if uploadErr != nil {
				return nil , uploadErr
			}

			resourceImage.ThumbResourceId = resource.Id
			return o.Update(resourceImage)
		})


	}()
}
var RESOURCE_IMAGE_SERVICE_INSTANCE = &ResourceImageService{}

func GetResourceImageServiceInstance() *ResourceImageService{
	return RESOURCE_IMAGE_SERVICE_INSTANCE
}

