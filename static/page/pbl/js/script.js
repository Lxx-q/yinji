new Vue({
	el:"#app",
	data:{
		string:"helloworld",
		audios:[
			{name:"大闹天宫",image:"http://localhost:8080/yinji/image/1574951701701.jpg"}
		]
	},created:function(){

		var vue = this;
		$.ajax({
			url:"/yinji/api/audio/user",
			data:{
				userId:2
			},
			async:false,
			dataType:"json",
			success:function(  result , status , xhr ){
				var newResult = [];

				for( var index = 0 ; index < result.length ; index ++ ){
					var item = result[index];
					console.log( result[index]);
					var newItem = {};
					newItem.name = item.Name;
					newItem.image = "http://localhost:8080/yinji" + "/" + item.Image;
					newResult[newResult.length] = newItem;
				}


				vue.audios = newResult;

			},error:function( xhr,status,error ){
				alert("hello , world");
			}
		})
	}
})