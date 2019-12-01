package main

import (
	"fmt"
	"yinji/utils"
)

func main() {
	/*

	var fileFile = &service.FileService{}

	//获取当前文件的路径
	var path = "D://audio/1.mp3"

	var newPath = "D://audio/new/1.mp3"

	var file, err = fileFile.OpenFile(path)

	if (err != nil) {
		fmt.Println("path is error :" + err.Error())
	}
	var newFile, newErr = fileFile.GetFile(newPath)

	defer file.Close()

	if (newErr != nil) {
		fmt.Println("newErr is error : " + newErr.Error())
	}

	var fileAssembly = service.FileAssembly{newFile }
	var stringAssembly = service.StringAssembly{&bytes.Buffer{}}
	var assemblys = make([]service.DownloadAssembly , 0 )

	assemblys = append( assemblys , fileAssembly )
	assemblys =append( assemblys , stringAssembly )
	var ArrayAssembly = service.ArrayAssembly{ assemblys}

	fileFile.Write( ArrayAssembly , file)
	*/

	fmt.Println( utils.Now())

}
