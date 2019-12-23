new Vue({
	el:"#app",
	data:{
		string:"hello , world",
		items:[
			
		],
		startLimit:0,
		endLimit:0,
		space:10,
		userId:2,
		//对应的分组信息
		groups:[
			{ name:"默认分组" , "id":null }
		]
	},methods:{

		ajaxAudio:function(){

			var vue = this;

			if( vue.userId == null ){
				alert("并未查找到该用户");
				return ;
			}

			//进入下一页
			this.nextPage();
			//$.ajax({
			window.AJAX_ENGINE.ajax({
				url:"/yinji/api/audio/user",
				async:false,
				dataType:"json",
				data:{
					userId:vue.userId,
					startLimit:vue.startLimit,
					endLimit:vue.endLimit
				},success:function( result , status , xhr ){

					if( result.length != vue.space ){
						$(window).off("scroll");
					}

					var newItems = vue.loadAudio( result );

					vue.items = vue.items.concat(newItems);

				},error:function( xhr , status , error ){
					alert("hello , world")
				}
			});
		},loadAudio:function( items ){
			var newItems = [];

			for( var index = 0 ; index < items.length ;  index++ ){
				
				var item = items[ index ];
				var newItem = window.AUDIO_HANDLER.receive( item )
				newItems[newItems.length] = newItem;
			}

			return newItems;
		},deleteAudio:function( item , index ){

			var vue = this;
			
			$.ajax({
				url:"/yinji/api/audio/delete",
				dataType:"json",
				async:true,
				data:{
					id:item.id
				},success:function( result , status , xhr ){
					//倘若结果为true
					//那么， 我们便可以， 删除对应的 result
					vue.items.splice( index , 1);
				},error:function( xhr , status , result ){
					alert("this is error" + result);
				}
			})


		},intoPage:function( item ,  index ){
			var id = item.id;
			//暂时使用对应的方法 ， 来进行操作，主要是方便测试
			window.open("/yinji/page/upload/update" + "?" + "id=" + id + "&" + "userId" + "=" + this.userId); 
		},nextPage:function(){
			this.startLimit = this.endLimit;
			this.endLimit = this.startLimit + this.space;
		}
	},created:function(){
		var vue = this;

		vue.startLimit = 0;
		vue.endLimit = 0;

		//获取连接之中是否有对应的 userId
		vue.userId = GetQueryString("userId");

		vue.ajaxAudio();
		$(document).ready(function(){
			//开始绑定对应的滑动事件
			$(window).on("scroll",function(){
				var screenHeight = screen.availHeight;
				var currentHeight = $(document).scrollTop();
				var height = $(document).height();

				if( currentHeight > 0.8 * ( height - screenHeight) ){
					vue.ajaxAudio();
					window.pbl.init()
				}
				
				
			})
		})
	}
})