

/**
	相对应的 data 的 值 为 this.的 属性值
*/
window.handler  = {
	//第一个参数 为目标的 data 值 ， 第二个 为 存储值 ， 但是 实际 的 值的 属性值 为何 ， 为相对应的 
	send:function( player  , element ){

		var data = {};

		data.currentTime = player.currentTime;
		data.duration = player.duration;
		data.isTimerPlaying = player.isTimerPlaying;
		data.tracks = player.tracks;
		data.currentTrack = player.currentTrack;
		data.currentTrackIndex = player.currentTrackIndex;
		data.playTracks = player.playTracks;
		data.appendInput = player.appendInput;
		data.currentTimeData = player.currentTimeData;
		data.durationData = player.durationData;
		data.volume = player.volume;

		return data;

	},receive:function( player , data ){

		player.currentTime = data.currentTime;
		player.duration = data.duration;
		player.isTimerPlaying = data.isTimerPlaying;
		player.tracks = data.tracks;
		player.currentTrack = data.currentTrack;
		player.currentTrackIndex = data.currentTrackIndex;
		player.playTracks = data.playTracks;
		player.appendInput = data.appendInput;
		player.currentTimeData = data.currentTimeData;
		player.durationData = data.durationData;
		player.volume = data.volume;

		/*
		if( player.isTimerPlaying ){

			var audios = player.audios;

			if( !audios.isHost ){
				audios.volume = 0 ;
			}
			audios.play();
		}else{
			player.audios.pause();
		}
		*/

			
		return data;
	}

}