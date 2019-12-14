package url

import "bytes"

func BuildApiUrl( url string ) string{
	var buffer = bytes.Buffer{}
	buffer.WriteString("http://localhost:8080/yinji/")
	buffer.WriteString( url );
	return buffer.String()
}