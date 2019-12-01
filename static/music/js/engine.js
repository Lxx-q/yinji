

window.engine = {

	// 相对应的 信息
	save:function( key , value ){
		//
		var data = {};
		//将信息输入
		data.content = value;

		//之后将 相对应的 信息转化为 对应的 json 格式的 信息

		var json = JSON.stringify( data );

		window.localStorage.setItem( key , json );

	},load:function( key ){
		//获取 相对应的 json的 信息 
		var json = window.localStorage.getItem( key );
		// 我们将其 转化为data
		var data = JSON.parse( json );

		//再之后 直接返回成 content
		if( data == null ){
			return null;
		}

		return data.content;
	},remove:function( key ){
		window.localStorage.removeItem( key );
	}

}