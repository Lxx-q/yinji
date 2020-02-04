/**
    1.点赞方法 [ addLove ]
*/
new Vue({
    el: "#app",
    data: {
        //是否已经点赞
        isLove:false,
        loveCount:"-",
        //评论区信息
        comments: [],
        maxComments:10,
        maxReplies:5,
        //用户信息
        userContainer: {
            //当前评论区出现的所有的user用户
            1 : {
                id: 1,
                name: "lx",
                image: "./images/img.jpg"
            }
            //当前用户的 id账号
        },replyContainer:{
        	//根据不同的用户信息来获取不同的用户信息内容
        	//其中每个人的内容根据对应的回复时间而定
        },
        userId:2,
        audioId:1574954699699,
        $commentTextarea:null,

        //用来设定回复的参数
        currentTargetId:null,
        currentAudioCommentId:null
    },
    methods: {
        //点赞方法
        addLove:function(){
            var url = getServerUrl("api/love/insert");
            var userId = this.userId;
            var audioId = this.audioId;
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    userId:userId,
                    audioId:audioId
                },dataType:"json",
                async:true,
                success:function( result , status , xhr ){
                    window.CONFIRM.alert("射射兄弟(●￣(ｴ)￣●)");
                    vue.isLove = true
                },fail:function( type , result ){
                    //等等 ， 你好像已经关注了哦
                }
            })
        },cancalLove:function(){
            var url = getServerUrl("api/love/delete");
            var userId = this.userId;
            var audioId = this.audioId;
            var vue = this;

            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    userId:userId,
                    audioId:audioId
                },dataType:"json",
                async:true,
                success:function( result , status , xhr ){
                    window.CONFIRM.alert("辣鸡，快给老子滚o(╥﹏╥)o");
                    vue.isLove = false;
                }
            })
        },
    	//写明伦之后的方法
    	writeComment:function(){
            var $commentTextarea = this.$commentTextarea;
    		//获取对应的信息内容然后将其输出
    		var content = $commentTextarea.val();
    		//之后进行一次 ajax ， 将对应的 this.userId ， 发送出去
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/comment/insert",
                data:{
                    audioId:vue.audioId,
                    userId:vue.userId,
                    content:content
                },success:function( result , status , xhr){
                    alert("插入成功");
                },error:function( xhr , status , err ){
                    //输入对应的信息
                    alert("插入失败");
                }
            })

    		//之后清空content
    		$commentTextarea.val("");
    	},writeCommentReply:function( content ){
            var vue = this;

            var targetId = vue.currentTargetId;
            var userId = vue.userId;
            var audioCommentId = vue.currentAudioCommentId;

            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/reply/comment/insert",
                data:{
                    userId:userId,
                    targetId:targetId,
                    commentId:audioCommentId,
                    content:content
                },success:function( result , status , xhr ){
                    alert("插入成功");
                },error:function( xhr , status , err ){
                    alert("插入失败")
                }
            })
        },replyClick:function( event , targetId ,commentId){

            //先将目标target存储
            this.currentTargetId = targetId;
            this.currentAudioCommentId = commentId;

        	//点击回复按钮之后出现的信息
        	var target = event.target;
        	var $this = $(target);

        	if($this.parent().parent().find(".replybox").length > 0){
				$(".replybox").remove();
			}else{
				$(".replybox").remove();
				this.showReply($this);
			}
        },
        //显示对应的信息
        showReply: function( el ) {

            var vue = this;
        	
            el.parent().parent().append("<div class='replybox'><textarea cols='80' rows='50' placeholder='来说几句吧......' class='mytextarea' ></textarea><span class='send'>发送</span></div>")
            .find(".send").click(function() {

                var $this = $(this);
                var content = $this.prev().val();
                //获取对应的信息
                if (content != "") {
                    //下面是点击放松之后的 方法
                    /*
                	//得到对应的信息内容之后，开始输出
                    var parentEl = $(this).parent().parent().parent().parent();
                    var obj = new Object();
                    obj.replyName = "匿名";
                    if (el.parent().parent().hasClass("reply")) {
                        console.log("1111");
                        obj.beReplyName = el.parent().parent().find("a:first").text();
                    } else {
                        obj.beReplyName = parentEl.find("h3").text();
                    }
                    obj.content = content;
                    obj.time = getNowDateFormat();
                    var replyString = createReplyComment(obj);
                    $(".replybox").remove();
                    parentEl.find(".reply-list").append(replyString).find(".reply-list-btn:last").click(function() {
                        alert("不能回复自己");
                    });
                    */
                    vue.writeCommentReply(content);
                } else {
                    alert("空内容");
                }

                $(this).prev().val("");
                $(".replybox").remove();
            });
        },loadAudioComments:function(){

            /**
                在加载信息的时候 ， 这里有 三个元素 
                1. 用户信息（id ， 名字， 图片）
                2.评论信息
                3.回复信息

                我们现在有 2 中方案 ， 
                第一种 ， 最简单的 方法 ， 用上用户外键 ，指利用 用户外键 来获取用户信息 。
                然后将相对应的 评论也加载出来 同样附带着用户信息
                第二种
                在加载完用户,回复的信息之后，收集其对应的信息之后 ， 再输出页面
            */

            var vue = this;
            //加载对应的 评论信息
            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/comment/find/audio",
                data:{
                    audioId:vue.audioId
                },
                dataType:"json",
                success:function( data , status , xhr ){
                    var length = data.length;
                    for( var index = 0 ; index < length ; index ++ ){
                        var item = data[index];
                        var audioComment = window.AJAX_FORMAT.toAudioComment(item);
                        var user = window.AJAX_FORMAT.toUser(item);
                        //输入对应的信息
                        vue.comments.push( audioComment );
                        vue.userContainer[user.id] = user;
                        //加载对应的信息
                        vue.loadReplies( audioComment.id );
                    }
                },error:function( xhr , status , data ){
                    //程序出现错误

                }
            })

        },loadReplies:function( commentId ){
            //通过对方的 commentid 来加载对应的 audio 的数据
            var vue = this;

            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/reply/comment/find/comment",
                data:{
                    commentId:commentId
                },
                async:false 
                ,success:function( data , status , xhr ){
                    var length = data.length;
                    var replys = [];
                    for( var index = 0 ; index < length ;  index++ ){
                        var item = data[index];
                        var reply = window.AJAX_FORMAT.toAudioCommentReply(item);
                        
                        var target = window.AJAX_FORMAT.toReplyUser( item.target );
                        var user = window.AJAX_FORMAT.toReplyUser( item.user );

                        vue.$set(vue.userContainer[item.userId], user);
                        vue.$set(vue.userContainer[item.targetId] , target);

                        //插入对应的数据
                        replys.push(reply);
                    }

                    //vue.replyContainer[commentId] = replys;
                    //vue.$set(vue.replyContainer[commentId],replys);
                    vue.replyContainer[commentId] = replys;

                },error:function( xhr , status , data ){

                }
            })
        },initLove:function(){
            var url = getServerUrl("api/love/find")
            var userId = this.userId;
            var audioId = this.audioId;
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    userId:userId,
                    audioId:audioId
                },dataType:"json",
                async:true,
                success:function( result , status , xhr ){
                    vue.isLove = true
                },fail:function( type , result ){
                    vue.isLove = false;
                }
            })
        }
    },created:function(){
    	var vue = this;

        //获取对应的开始的信息
        var audioId = GetQueryString("audioId");
        var userId = GetQueryString("userId");
        
        this.audioId = audioId;
        this.userId = userId;
        
    	$(document).ready(function(){
    		//将对应的评论框对象进行输入
    		vue.$commentTextarea = $("#content");
    	})
        vue.loadAudioComments();
        this.initLove();
    	
    }
})