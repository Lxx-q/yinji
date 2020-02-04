window.AJAX_ENGINE = {
	buildAjaxClass:function( data , success = function( result , status , xhr ){}, fail = function( type , content ){} ){
		var successFunc = function( result , status , xhr ){
			//2 便是操作的方法
			if( result.type == 2 ){
				fail( result.type , result.content);
			}else{
				//调用成功的方法
				success(result.content ,status , xhr);
			}
		}

		//输入对应的错误的函数
		data.success = successFunc;

		data.fail = fail;


		return data;

	},ajax:function(data){
		var newData = this.buildAjaxClass( data , data.success , data.error );
		$.ajax(newData);
	}
}

