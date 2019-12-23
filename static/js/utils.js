

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