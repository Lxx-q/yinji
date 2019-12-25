window.AJAX_FORMAT = {
	toAudioComment:function( data ){
		var audioComment = {};

		audioComment.id = data.id;
		audioComment.userId = data.userId;
		audioComment.audioId = data.audioId;
		audioComment.content = data.content;
		audioComment.createTime = data.createTime;
		audioComment.code = data.code;
		audioComment.createTimeStruct = data.createTimeStruct;

		return audioComment;
	},toUser:function( data ){
		var user = {};

		user.id = data.userId;
		user.name = data.userName;
		user.image = parseImage(data.userImage);

		return user;
	},toAudioCommentReply:function( data ){
		var reply = {};

		reply.id = data.id;
		reply.userId = data.userId;
		reply.content = data.content;
		reply.targetId = data.targetId;
		reply.code = data.code;
		reply.createTime = data.createTime;
		reply.createTimeStruct = data.createTimeStruct;

		return reply;
	},toReplyUser:function( data ){
		var user = {};

		user.name = data.name;
		user.image = data.image

		return user;
	}	
}