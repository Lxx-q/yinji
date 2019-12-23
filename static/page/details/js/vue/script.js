new Vue({
    el: "#app",
    data: {
        string: "hello , world",
        //评论区信息
        comments: [{
            id: 1,
            createTime: "2019-10-17 10:15:16",
            content: "我觉得，你说的很对， 但是我不听你的",
            userId: 1
        },{
        	id:2,
        	createTime: "2019-10-17 10:15:16",
        	content:"今天是个好日子？？",
        	userId:1
        }],
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
        	2:[
        		{id:"xx",name:"",createTime: "2019-10-17 10:15:16",content:"加油， 奥利给",userId:2}
        	]
     
        },
        currentUserId:2,
        audioId:1574954699699,
        $commentTextarea:null,

        //用来设定回复的参数
        currentTargetId:null,
        currentAudioCommentId:null
    },
    methods: {
    	//写明伦之后的方法
    	writeComment:function(){
            var $commentTextarea = this.$commentTextarea;
    		//获取对应的信息内容然后将其输出
    		var content = $commentTextarea.val();
    		//之后进行一次 ajax ， 将对应的 this.currentUserId ， 发送出去
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/comment/insert",
                data:{
                    audioId:vue.audioId,
                    userId:vue.currentUserId,
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
            var userId = vue.currentUserId;
            var audioCommentId = vue.currentAudioCommentId;

            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/comment/reply/insert",
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
        },parseComment: function(data) {
            //对应的 数据库信息 ， 转化为 当前内部的 comment
            var comment = {};
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
                async:true,
                dataType:"json",
                success:function( data , status , xhr ){
                    
                }
            })
        }
    },created:function(){
    	var vue = this;
    	$(document).ready(function(){
    		//将对应的评论框对象进行输入
    		vue.$commentTextarea = $("#content");
    	})
    	
    }
})