
//首先我们 必须必须先存储初始化的 方法
var timeupdate = window.player.timeupdate;

window.player.timeupdate = function(){

    

    /*
    // 将我们 所需要获取到的 并将它 整合成 json
    var data = window.handler.send( this , {} );

    //然后 我们将 相对应的 item 转化为 相对应的 data

    //之后我们 将 相对应的 原子 信息 转化为 相对应的的 json 信息
    window.engine.save("as" , data );

    //之后我们 需要将 相对应的 属性 保存进 相对应的 data
    */

    if( window.service.isHost){
    
    	window.service.save( this );
    	window.player.currentAudioData();	
	}

	timeupdate();
}

var storage = window.onstorage;
window.onstorage = function( event ){
	if( storage  != null ){
		storage();
	}

	/*
		var data = window.engine.load("as");
		window.handler.receive( window.player , data );
	*/
	var key = event.key;

	var service = window.service;

	//输出 对应的 信息
	if( key == service.dataKey ){
		console.log("service");
		service.update( window.player );
		console.log( window.player);
	}else if( key == service.hostKey){
		//倘若是 hostKye 出现 的 修改 ， 那么我们便可以这么认为 是 相对应的 程序 出现的 问题

		var newValue = event.newValue;

		//是否 当前的 newValue 为null ， 那么 说明是时候到 需要换 host 的 时候了
		if( newValue == null){
			var playerId = service.playerList[1];

			if( service.playerId == playerId  ){
				window.engine.save( service.hostKey , service.playerId );
				service.isHost = true;
			}

		}	

	}else if( key == service.playerListKey ){
		service.playerList = window.engine.load( service.playerListKey );
	}else if( key == service.playingKey){
		//首先 ， 我们 获取 相对应的 细信息 
		window.player.isTimerPlaying = ! window.engine.load( service.playingKey );
		//之后 ，我们 开始运行整套的
		player_click_play();
	}

}

//首先 ， 我们 获取 原先的 初始化 时间

var load = window.onload;

window.onload =  function( event ){
	if( load != null ){
		load();
	}

	//首先初始化 页面的 所有的 信息 ， 
	window.service.register();
	//之后 我们 便将 进行 注册 ，查看是否能进行注册
	window.service.registerHost();

	//首先 ， 我们 获取 相对应的 细信息 
	window.player.isTimerPlaying = ! window.engine.load( service.playingKey );
	//之后 ，我们 开始运行整套的
	player_click_play();
}

var unload = window.onunload ;

window.onunload  = function(){
	if( unload != null ){
		unload();
	}

	var service = window.service;
	//倘若 你目前 是 相对应的 host
	if( service.isHost){
		service.cancelHost();
	}

	//关闭 对应的 信息
	service.cancel();
}

//修改其播放以及对应的修改方法
var player_play_audio = window.player.playAudio;

window.player.playAudio = function(){
	//倘若目标 并非是 主机 ， 那么 我们的 一般操作 便是 ， 禁止他进行播放
	var isHost = window.service.isHost;

	//倘若 该 isHost为false ，那么 杜绝其进行操作
	//console.log( isHost );
	
	
	player_play_audio();
	var number = window.player.currentTimeData;

	this.audio.currentTime = number; 
	if( !isHost){
		this.audio.volume = 0.0;

	}

}

var player_pause_audio = window.player.pauseAudio;

window.player.pauseAudio = function(){
	//倘若目标 并非是 主机 ， 那么 我们的 一般操作 便是 ， 禁止他进行播放
	var isHost = window.service.isHost;

	//倘若 该 isHost 为false ，那么 杜绝其进行操作
	//console.log( isHost );
	player_pause_audio();
	this.setVolume();
}

//整体点击播放状态 
var player_click_play = window.player.play;

window.player.play = function(){
	//我们将其 修改为 倘若 进行保存对应的状态
	player_click_play();

	window.engine.save( window.service.playingKey , window.player.isTimerPlaying);
}


var nextTrack = window.player.play;

window.player.nextTrack = function(){
	nextTrack();
	// 之后发送一些数据
	window.service.save( this );
}

