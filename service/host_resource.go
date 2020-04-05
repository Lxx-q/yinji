package service

import (
	"github.com/astaxie/beego/orm"
	"io"
	"yinji/models/bean"
	"time"
	"strconv"
	"os"
)


type HostResourceService struct {
	httpFileService *HttpFileService
}

/**
	用来上传信息的接口 ， 我们必须这样的设置一个信息的名称
 */
func ( self *HostResourceService ) Upload( o orm.Ormer , fatherPath string , filename string  , reader io.Reader ) (*bean.Resource , *bean.HostResource , error ){
	var t = time.Now()

	// 下面准备开始主键对应的信息
	var hostResource = bean.HostResource{}
	hostResource.NewEntity( t )

	var fileService = GetFileServiceInstance()
	//通过对应的 path来生成绝对路径
	var suffix = fileService.Ext(filename) // 获取后缀名
	var name  = strconv.FormatInt(hostResource.Id, 10)
	var relativePath = fatherPath +"/" + name  + suffix //获取相对应的路径
	var abs = self.httpFileService.BuildServerPath(relativePath) //再用相对路径获取绝对路径
	var file , getFileErr  = fileService.GetFile( abs )

	if getFileErr != nil { //若查找文件出现错误，则会直接返回 ,但是不明确当，是否文件不存在，会不会出现对应的错误。
		return nil ,  nil , getFileErr
	}

	var fileAssembly = FileAssembly{ file }
	var writeErr = fileService.Write(fileAssembly,reader)

	if writeErr != nil {
		return nil , nil  , writeErr
	}


	hostResource.Path = relativePath

	var _ , insertHostResourceErr = o.Insert( &hostResource )

	if insertHostResourceErr != nil {
		return nil ,  nil  , insertHostResourceErr
	}

	//最后插入 resource

	var resourceService = GetResourceServiceInstance()

	var resource = bean.Resource{}
	resource.NewEntity( t )
	resource.HostResourceId = hostResource.Id
	var insertResourceErr = resourceService.Insert(o , &resource)

	if insertResourceErr != nil {
		return nil , nil  , insertResourceErr
	}

	return &resource , &hostResource  , nil

}

func ( self *HostResourceService ) Update( o orm.Ormer , hostResource *bean.HostResource  , fatherPath string , filename string , reader io.Reader) error {
	//先删除原本的文件
	self.RemoveFile(hostResource)
	var fileService = GetFileServiceInstance()
	//通过对应的 path来生成绝对路径
	var suffix = fileService.Ext(filename) // 获取后缀名
	var name  = strconv.FormatInt(hostResource.Id, 10)
	var relativePath = fatherPath +"/" + name  + suffix //获取相对应的路径
	var abs = self.httpFileService.BuildServerPath(relativePath) //再用相对路径获取绝对路径
	var file , getFileErr  = fileService.GetFile( abs )

	if getFileErr != nil { //若查找文件出现错误，则会直接返回 ,但是不明确当，是否文件不存在，会不会出现对应的错误。
		return  getFileErr
	}

	var fileAssembly = FileAssembly{ file }
	var writeErr = fileService.Write(fileAssembly,reader)

	if writeErr != nil {
		return  writeErr
	}


	hostResource.Path = relativePath

	hostResource.Refresh()

	var _ , error = o.Update(hostResource)

	return error
}

func ( self *HostResourceService )RemoveFile(  resource *bean.HostResource) error{
	var abs = self.httpFileService.BuildServerPath(resource.Path) //再用相对路径获取绝对路径
	return os.Remove(abs)
}

var HOST_RESOURCE_SERVICE_INSTANCE = &HostResourceService{httpFileService:GetHttpFileServiceInstance()}

func GetHostResourceServiceInstance() *HostResourceService{
	return HOST_RESOURCE_SERVICE_INSTANCE
}
