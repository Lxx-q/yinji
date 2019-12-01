package service

import (
	"bytes"
)

type AppService struct {
	Context string
}

func (self AppService) GetUrl( url string) string {
	var buffer = bytes.Buffer{};

	// 可以为 相对应的 项目 设置 项目头部信息
	buffer.WriteString( self.Context );
	buffer.WriteString(url);

	return buffer.String();
}