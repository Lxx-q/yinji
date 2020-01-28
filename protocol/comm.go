package protocol

//下面我们设定一下需要返回的类型


const(
	//正常成功的响应;
 COMM_TYPE_SUCCESS  = 1
 //失败的响应
 COMM_TYPE_FAIL = 2
 //快速短小的文
 COMM_TYPE_FAST = 3
)

type CommStruct struct {
	Type int `json:"type"`
	Content interface{} `json:"content"`
	//Time time.Time `json:"time"`
}


