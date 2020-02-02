
var server = "/" + window.SERVER + "/" + "js" + "/" + "confirm" ;

//输入对应的信息
window.import.js(server + "/" +  "js/jquery.sticky.min.js")
		.js(server + "/" + "js/jquery-confirm.js")
		.css(server + "/" + "css/jquery-confirm.css");


window.CONFIRM = {
	//基础配置
	config:function( settings ){
		settings.theme = "black";
		settings.keyboardEnabled = true;
		if(settings.columnClass == null ){
			settings.columnClass = "col-md-6";
		}
		settings.confirmButton = "是";
		settings.cancelButton = "否";
		settings.keyboardEnabled = true;
	},confirm:function( settings ){
		this.config( settings );
		$.confirm( settings )
	},alert:function(settings){
		this.config( settings );
		$.alert( settings );
	}
}