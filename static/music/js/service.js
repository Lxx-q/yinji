
//相对应的 信息
window.service = {
	//对应的 播放
	playerId:null,
	//数据 包名
	dataKey:"data_key",
	// 是否为 对应的 登录 主机
	hostKey:"host_key",
	playerListKey:"player_list_key",
	playerList:null,
	//playerMapKey:"player_map_key",
	//playerMap:null,
	isHost:false,
	playingKey:"playing",
	register:function(){
		//
		var time = new Date().getTime();
		//首先获取当前的 时间 ， 精确到 毫秒 ....

		// 获取playerid ，id 编号
		var playerIdTag = "pk";

		//生成 0 - 1000 的 一个 随机数
		var random = Math.floor( Math.random() * 1000 );
		var split = "_";
		//得到 自己的 相对应的 playerId
		this.playerId =  playerIdTag + split + time + split +random;
		this.isHost = false;

		//之后加载对应的 playerList
		
		var player_list = window.engine.load( this.playerListKey );

		if( player_list == null ){
			//倘若为空， 则进行初始化
			player_list = [];
		}
		/*
		if( player_map == null ){
			player_map = {};
		}
		*/
		//往数组内添加 对应的 数据
		var length = player_list.length;
		player_list[ length ] = this.playerId;
		//player_map[this.playerId] = length;

		window.engine.save( this.playerListKey ,  player_list );
		//window.engine.save( this.playerMapKey , player_map );

		//this.playerMap = player_map;
		this.playerList = player_list;

	},registerHost:function(){
		//目前的 信息 去 竞选  host
		//我们现在的 策略为 ， 只要 host 为空 ， 便允许去竞选
		var host_key = this.hostKey;

		var host_id = window.engine.load( host_key );
		//之后 我们 再 部署 进去

		if( host_id == null ){
			//倘若 没有 对应的 信息 ， 我们 便将 相对应的 信息 部署 进去
			window.engine.save( host_key , this.playerId );
			this.isHost = true;
		}

	},cancelHost:function(){
		this.isHost = false;
		window.engine.remove( this.hostKey );

	},cancel:function(){

		var service = this;
		// 获取对应的 信息 
		//var index = service.playerMap[ service.playerId ];

		//service.playerMap[ service.playerId ] = null;

		//为了防止其他的原属 ， 因此 ，我们采用这个方法
		var playerList = window.engine.load( service.playerListKey );

		for( var index = 0 ; index < playerList.length; index++ ){
			// 获取 本次 相对应的 playerId 的 信息 为
			var playerId = playerList[ index ];

			//倘若 删除本地值应该所在的 位置
			if( playerId == service.playerId ){
				playerList.splice( index , 1);
				break;
			}

		}

		service.playerList = playerList;

		// 将相对应的 属性 保存至 对应的 playerList
		window.engine.save( service.playerListKey ,playerList );

	},
	save:function( player ){
		//得到 相对应的 信息 属性
		var data = window.handler.send( player , {} );

		window.engine.save( this.dataKey , data);

	},update:function( player ){
		var data = window.engine.load( this.dataKey );
		window.handler.receive( player, data );
	}

}