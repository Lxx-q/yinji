

/*
	{
    "Name": "来不及勇敢",
    "Url": "1",
    "Image": "",
    "Code": 1,
    "Id": "m_1",
    "CreateTime": "2019-11-04T20:44:38+08:00",
    "ModifyTime": "2019-11-04T20:44:45+08:00"
  	}
*/

//目前专门属于 index.html 页面的 信息 转化的 页面

window.AJAX_HANDLER = {
	//接受某种数据 ， 我们将其转化为某一种数据
	receive:function( data  , index ){
		var new_data = {};

		var name = data.name;

		if(name == ""){
			name = "未命名";
		}
		new_data = data;
		new_data.title = name;
		new_data.artist = name;
		new_data.mp3 = '/yinji/api/resource/audio/t?type=1&id=' + data.resourceAudioId;
		new_data.poster = '/yinji/api/resource/image?type=1&id=' + data.resourceImageId;
		new_data.img = new_data.poster;
		new_data.time = this.formatTimeLength( data.timeLength)
		
		//最后返回对应的数据
		return new_data;
	},receiveArray:function( data_array ){
		//倘若转化的是一个数组数据
		var length = data_array.length;
		var new_data_array = [];
		for( var index = 0 ; index < length ; index++ ){
			var data = data_array[ index ];

			var new_data = this.receive( data  , index );

			new_data_array[new_data_array.length] = new_data;
		}

		return new_data_array;
	},formatTimeLength:function( timeLength ){
		//根据对应的 信息 来进行 获取
		var minutes = Math.floor(timeLength / 60);
		var seconds = timeLength % 60 ;

		var minutesString = minutes;

		if( minutes < 10 ){
			minutesString = "0" + minutesString;
		}

		var secondsString = seconds;

		if( seconds < 10 ){
			secondsString = "0" + secondsString;
		}

		return minutesString + ":" +  secondsString;

	}
}