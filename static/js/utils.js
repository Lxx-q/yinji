
const DEFAULT_IMAGE_PATH = "/yinji/resources/image/none.jpg";

//获取当前页面中的参数
function GetQueryString(name){
     var reg = new RegExp("(^|&)"+ name +"=([^&]*)(&|$)");
     var r = window.location.search.substr(1).match(reg);
     if(r!=null)return  unescape(r[2]); return null;
}

function ParseTime(date){
	
     var arr=date.split("T");
     var d=arr[0];
     var darr = d.split('-');
     var t=arr[1];
     var tarr = t.split('.000');
     var marr = tarr[0].split(':');
     var dd = parseInt(darr[0])+"/"+parseInt(darr[1])+"/"+parseInt(darr[2])+" "+parseInt(marr[0])+":"+parseInt(marr[1])+":"+parseInt(marr[2]);
     return dd;
}

function getServerUrl( url ){
	return "/yinji" + "/" + url;
}

function parseImage( path ){
	return path != null  && path != "" ? "/yinji" + "/" +  path : DEFAULT_IMAGE_PATH;  
}

/**
	获取对应事件的毫秒级别参数
*/

const HOUR_MILLI_SEC = 1000 * 60 * 60 ;
const DAY_MILLI_SEC = HOUR_MILLI_SEC * 24;

function Today(){
	var date = new Date();
	//获取当前的时间
	var year = date.getFullYear();
	var month = date.getMonth();
	var day = date.getDate();
	return new Date( year , month , day );
}

/**
	
	year:year,
	month:month,
	day:day,
	date:date,


*/
function parseTimeStruct( now ){

	var year = now.getFullYear(); //得到年份
	var month = now.getMonth() + 1 ; //得到月份
	var date = now.getDate(); //得到日期
	var day = now.getDay(); //得到周几
	var hour = now.getHours(); //得到小时
	var minu = now.getMinutes(); //得到分钟
	var sec = now.getSeconds(); //得到秒
	var times = now.getTime();

	var timeStruct = {
   		year: year,
    	month: month,
   	 	day: date,
    	weekDay: day,
    	hour: hour,
    	minute: minu,
    	milliTime: times
	}

	return timeStruct;
}

function parseTimes( times ){
	var date = new Date( times );
	return parseTimeStruct( date );
}


function splice( arr , start , end ){

	var new_arr = [];
	var len = arr.length;
	//倘若长度如果大于约定的长度 ， 那么我们就设定他的长度为固定的
	len = len > end ? end : len;

	for( var index = start ; index < len ; index ++ ){
		var item = arr[index];
		new_arr.push(item);
	}

	return new_arr;

}