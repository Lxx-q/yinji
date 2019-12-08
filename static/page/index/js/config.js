
const URL_APPLICATION_HEADER = "yinji"

//  设定 相对应的搜索的 url
const SEARCH_AUDIO_URL = "api/audio/search";


window.URL_SERVICE = {

	/*
		根据对应的 信息来进行生成信息
	*/
	buildUrl:function( url ){
		return "/" + URL_APPLICATION_HEADER + "/" + url;
	}

}