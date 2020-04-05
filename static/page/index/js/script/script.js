//对应的 js 脚本 语言

/**
   简单的说 ， 我们搜索两个信息

   1.收藏的信息
   2.
*/
new Vue({
    el: "#app",
    data: {
        //对应的头部循环栏
        headerSwiper:[1,2,3],
        jsPlayList: null,
        latestStartLimit:0,
        latestEndLimit:0,
        latestOffset:40,
        latestContent:null,
        latest: [

        /**
        * 对应的 数据 格式 

            {
                title:数据的 名称
                ,artist:"" 作者的姓名
                ,mp3:"" 对应的 音频的链接
                ，poster："对应上传者的头像"
                ,length:"对应时间的长度"
                ,img:" 目标视频的链接 "

            }
        */
        // 对应分辨显示页中的其中一个 信息
        ],
        //另一方面的 信息显示 输出表
        spirit: [],
        searchInput: null,
        searchInputId: "searchInput",
        test: "helloworld",
        //尾部的link信息的 信息
        links:[// 0 为 pc 与 手机 两个 ，  1 为 只为电脑 2 为 手机端 
            { title:"作品集",text:"个人作品集合",image:"images/314601.jpg",link:"page/pbl/main" , status:0},
            { title:"数据中心",text:"数据宝典",image:"images/316718.jpg",link:"page/mit/index" , status:0},
            { title:"上传作品",text:"上传本地音频",image:"images/317839.jpg" , link:"page/uploadnew/template.html" , status:1},
            { title:"上传作品",text:"上传本地音频",image:"images/317839.jpg" , link:"page/uploadnew/index.html" , status:2}

        ],
        //当前userId
        userId:null,
        isPc:true //倘若是 true 那就是 电脑 , 不是 那就是 手机
    },
    methods: {
        format: function(item) {
            // 通过信息来进行 转化
            item.img = item.poster;
            item.time = "03:30";
        },
        newjPlayerPlaylist: function() {
            this.jsPlayList = new jPlayerPlaylist({
                jPlayer: "#jplayer_N",
                cssSelectorAncestor: "#jp_container_N"
            },
            [
            //对应播放器 内 存储的 歌单的 信息
            ], {
                playlistOptions: {
                    enableRemoveControls: true,
                    autoPlay: false
                },
                swfPath: "js/jPlayer",
                supplied: "webmv, ogv, m4v, oga, mp3",
                useStateClassSkin: true,
                autoBlur: false,
                smoothPlayBar: true,
                keyEnabled: true,
                audioFullScreen: false
            });
            
        },initPcAndMobile:function(){ //判断是否是
            var userAgentInfo = navigator.userAgent;
            var Agents = ["Android", "iPhone",
                "SymbianOS", "Windows Phone",
                "iPad", "iPod"];
            var flag = true;
            for (var v = 0; v < Agents.length; v++) {
            if (userAgentInfo.indexOf(Agents[v]) > 0) {
                flag = false;
                break;
            }

            }
            this.isPc = flag; // 设定当前的状态
        },init: function() {

            //获取目标的 信息

            var vue = this;
            $(document).ready(function(){
                vue.searchInput = $("#" + vue.searchInputId);
            })
            
            this.initjPlayer();
            this.initLastest();
            this.initSpirit();

        },
        initjPlayer: function() {

            var vue = this;
            $(document).ready(function() {

                vue.newjPlayerPlaylist();
                $(document).on($.jPlayer.event.pause, vue.jsPlayList.cssSelector.jPlayer,
                function() { //$(document).on('click','要选择的元素',function(){})  on方法包含很多事件，点击，双击等等事件。
                    $('.musicbar').removeClass('animate'); //look this class and remove it
                    $('.jp-play-me').removeClass('active');
                    $('.jp-play-me').parent('li').removeClass('active');
                });

                $(document).on($.jPlayer.event.play, vue.jsPlayList.cssSelector.jPlayer,
                function() {
                    $('.musicbar').addClass('animate'); // when the player add a animate
                });
                $("#before,#after").addClass("set_imd");
            });

        },addjPlayList:function(item){ //将目标的音频文件添加到目标的播放列表之中
            this.jsPlayList.add(item);
        },writePlayDashboard:function( item ){ // 请求服务器,并且记录下播放事件
            var userId = this.userId;
            if( userId == null ){
                return;
            }
            var audioId = item.id; //获取对应的 audioId
            var url = getServerUrl("api/audio/history/add")
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    lastTime:0, // 目前只能暂时默认这个lastTime 为0
                    userId:userId,
                    audioId:audioId,
                    count:1
                },async:true,
                success:function( data , status , xhr ){
                    console.log( data );
                }
            })
        },clickPlayMe: function(e , item ) {
            e && e.preventDefault();
            var $this = $(e.target);
            if (!$this.is('a')) $this = $this.closest('a'); //closest() 方法获得匹配选择器的第一个祖先元素，从当前元素开始沿 DOM 树向上。
            $('.jp-play-me').not($this).removeClass('active');
            $('.jp-play-me').parent('li').not($this.parent('li')).removeClass('active');

            $this.toggleClass('active');
            $this.parent('li').toggleClass('active');
            if (!$this.hasClass('active')) {
                this.jsPlayList.pause();
            } else {
                var k = $this.parent("li").index();
                $(".poster-img").attr("src", this.latest[k].poster); //此方法返回一个函数改变src   $('a.cover1').html('<img src="' + latest[k].poster );
                $(".musicbar").addClass("animate").index();
                this.addjPlayList({
                    title: this.latest[k].title,
                    artist: this.latest[k].artist,
                    mp3: this.latest[k].mp3,
                    poster: this.latest[k].poster
                });
                this.jsPlayList.play( - 1);
            }
            this.writePlayDashboard( item );

        },
        clickPlayFun: function(e) {
            e && e.preventDefault();
            var $this = $(e.target);
            if (!$this.is('a')) $this = $this.closest('a'); //closest() 方法获得匹配选择器的第一个祖先元素，从当前元素开始沿 DOM 树向上。
            $('.jp-play-fun').not($this).removeClass('active');
            $('.jp-play-fun').parent('li').not($this.parent('li')).removeClass('active');

            $this.toggleClass('active');
            $this.parent('li').toggleClass('active');
            if (!$this.hasClass('active')) {
                this.jsPlayList.pause();
            } else {
                var k = $this.parent("li").index();

                $(".poster-img").attr("src", this.spirit[k].poster);
                //此方法返回一个函数改变src   $('a.cover1').html('<img src="' + spirit[k].poster );
                this.jsPlayList.add({
                    title: this.spirit[k].title,
                    artist: this.spirit[k].artist,
                    mp3: this.spirit[k].mp3,
                    poster: this.spirit[k].poster
                });
                this.jsPlayList.play( -1 );
            }

        },clickSearchButton: function() {
            this.latestContent = this.searchInput.val();
            this.latestStartLimit = 0 ;
            //为lasted , 加上对应的区间大小等等
            this.latestEndLimit = this.latestStartLimit + this.latestOffset;
            this.updateLatest();
        },updateLatest: function() {
            //利用对应的 信息 来进行 更换 相对应的 信息
            var url = window.URL_SERVICE.buildUrl(SEARCH_AUDIO_URL);
            var val = this.latestContent;
            //之后我们利用对应的 信息来进行搜索
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url: url,
                async: true,
                data: {
                    content: val,
                    startLimit:vue.latestStartLimit,
                    endLimit:vue.latestEndLimit
                },
                success: function(data, status) {

                    //之后我们将相对应的 数据 来进行 输出
                    var new_data = window.AJAX_HANDLER.receiveArray(data);
                    vue.latest = new_data;

                    //之后我们 来进行 得到对应的 信息
                },
                error: function(request, status, err) {
                    //对应错误之后的操作
                }
            })
        },clickLinks:function( link ){
            //点击对应的 链接页面
            var userId = this.userId;
            if( userId == null ){
                //那么 我们便开始进入 loca
                //这里需要插入对应的 是否愿意的参数
                this.locationLogin();
                return ;
            }

            if( link == null ){
                //倘若连接为空 ， 则返回自己
                return ;
            }
            var url = getServerUrl(link) + "?" + "userId" + "=" + this.userId;
            window.open( url );
        },clickHref:function( link ){
            var userId = this.userId;

            if( link == null ){
                //倘若连接为空 ， 则返回自己
                return ;
            }

            console.log(link)

            if( userId != null){
                link =  link + "&" + "userId" + "=" + userId;
            }
            var url = getServerUrl(link) ;
            window.open( url );
        },locationLogin:function(){
            var url = getServerUrl("page/mit/login");
            // 转移到对应的 登录页面
            window.open( url );
        },initLastest:function(){
            //开始初始化信息
            //初始化对应的播放时间
            var url=  getServerUrl("api/audio/search/browse/most");
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    page:0,
                    count:10,
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    var new_data = window.AJAX_HANDLER.receiveArray(result);
                    vue.latest = new_data;
                }
            })
        },initSpirit:function(){
            var url=  getServerUrl("api/audio/search/browse/date");

            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    page:0,
                    count:10,
                },async:false,
                dataType:"json",
                success:function( result , status , xhr ){
                    var new_data = window.AJAX_HANDLER.receiveArray(result);
                    vue.spirit = new_data;
                    //暂时的initheader 放在这里
                    vue.initHeaderSwiper();
                }
            });

        },initHeaderSwiper:function(){ //初始化对应的头部轮播栏
            var spirit = this.spirit;
            this.headerSwiper = [];
            var len = spirit.length;
            len = len > 5 ? 5 : len;
            for(var index = 0 ; index < len ; index++ ){
                var item = spirit[index];
                this.headerSwiper.push(item);
            }
        }
    },
    created: function() {
        this.userId = GetQueryString("userId");
        this.init();
        this.initPcAndMobile();
    }
});