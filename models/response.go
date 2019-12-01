package models

//返回的成功的结果
const RESPONSE_TYPE_SUCCESS = 1
//返回错误的结果
const RESPONSE_TYPE_ERRPR = 2

type Response struct {
	Type int
	Content interface{}
}

func BuildResponse( Type int , Content interface{}) *Response {
	var response = &Response{ Type , Content }
	return response
}

func FailResponse( err error ) *Response{
	var errString = err.Error()
	return BuildResponse( RESPONSE_TYPE_ERRPR , errString )
}