package service

import (
	"os"
	"io"
	"path/filepath"
	"fmt"
	"bytes"
	"path"
)

/**
 	用于计算机自己操作自己本地的文件
 */
type FileService struct{

}

//只能打开已存在的文件 , 不支持删除文件
func ( self *FileService) OpenFile(path string) (*os.File, error) {
	return os.Open(path)
}

func ( self *FileService) GetFile(path string) (*os.File, error) {

	//首先获取目标文件的上级文件夹 ， 查看其是否存在
	var dirPath = filepath.Dir(path)

	os.MkdirAll(dirPath, 0666)

	//之后获取对应的 文件信息
	return os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0666)
}

func ( self *FileService) Delete(path string) error {
	var err = os.Remove(path)
	return err
}

/**
	不知道怎么的 ， 反正 ， 我们 便使用这个方法 来进行操作
 */
func ( self *FileService) Write( assembly DownloadAssembly, reader io.Reader ) error {

	var BUFFER_SIZE = 1024
	var bytes = make([]byte, BUFFER_SIZE)
	defer assembly.Release()

	for {

		var length, err = reader.Read(bytes)

		//情况不太好说 ， 因为 可能我们离开循环需要 err 的帮助
		if (err != nil || length == 0) {
			//下载循环的出口
			break
		}

		assembly.Download(bytes, length)

	}

	return nil

}

//获取文件后缀名
func ( self *FileService) Ext( fileName string  ) string {
	return path.Ext( fileName )
}

//这个东西 ， 掌握了如何进行 下载的目标
type DownloadAssembly interface {
	Download(bytes []byte, length int)
	Release()
}

type FileAssembly struct {
	File *os.File
}

func (self FileAssembly) Download(bytes []byte, length int) {
	self.File.Write(bytes[:length])
}

func (self FileAssembly) Release() {
	self.File.Close()
}



type StringAssembly struct {
	Buffer *bytes.Buffer
}

func (self StringAssembly) Download(bytes []byte, length int) {
	self.Buffer.Write(bytes[:length])
}

func (self StringAssembly) Release() {
	fmt.Println( self.Buffer.String() )
}

type ArrayAssembly struct {
	Assemblys []DownloadAssembly
}

func (self ArrayAssembly) Download(bytes []byte, length int) {
	var assemblyLength = len( self.Assemblys )
	for index := 0 ; index < assemblyLength ; index ++  {
		var assembly = self.Assemblys[index]
		assembly.Download(bytes , length)
	}
}

func (self ArrayAssembly) Release() {
	var length = len( self.Assemblys )
	for index := 0 ; index < length ; index ++  {
		var assembly = self.Assemblys[index]
		assembly.Release()
	}
}

var FILE_SERVICE_INSTANCE = &FileService{}

func GetFileServiceInstance() *FileService {
	return FILE_SERVICE_INSTANCE
}