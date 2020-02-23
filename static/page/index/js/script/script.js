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
        {
            title: "The wheel of fortune",
            artist: "Yang Bingyin",
            mp3: "http://music.163.com/song/media/outer/url?id=476515941.mp3",
            poster: "http://p1.music.126.net/4xHOkSVWH-n6p5pB3Jf0yQ==/109951162922204274.jpg?param=130y130"
        },
        {
            title: "you are my king",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=22731459.mp3",
            poster: "http://p1.music.126.net/Xy8iaLLLVT21Mr9wDDJBNQ==/792747883650465.jpg?param=130y130"
        },
        {
            title: "初恋的小美好",
            artist: "Sunny是个小太阳",
            mp3: "http://music.163.com/song/media/outer/url?id=529823229.mp3",
            poster: "http://p1.music.126.net/7DvTdaOADAw7KiKVkcMcag==/109951163105691071.jpg?param=130y130"
        },
        {
            title: "Lilac",
            artist: "MANYO",
            mp3: "http://music.163.com/song/media/outer/url?id=28466087.mp3",
            poster: "http://p1.music.126.net/scAe3f-GkQvo2h91Gpb0Bg==/1729531800492601.jpg?param=130y130"
        },
        {
            title: "一人静",
            artist: "姫神",
            mp3: "http://music.163.com/song/media/outer/url?id=25917069.mp3",
            poster: "http://p1.music.126.net/FnElEjsh00iuHWtHXrcY9g==/5638295627307829.jpg?param=130y130"
        },
        {
            title: "Through My Blood",
            artist: "Aimer",
            mp3: "http://music.163.com/song/media/outer/url?id=409872505.mp3",
            poster: "http://p1.music.126.net/45Qck88DR58FGgdHdDqfBw==/18181524277717113.jpg?param=130y130"
        },
        {
            title: "Too Late To Say",
            artist: "Sayulee",
            mp3: "http://music.163.com/song/media/outer/url?id=486194219.mp3",
            poster: "http://p1.music.126.net/Oa_1_y0a8NCmxB1zYKAFpw==/19187577416516096.jpg?param=130y130"
        },
        {
            title: "琥珀色の海へ",
            artist: "40㍍P",
            mp3: "http://music.163.com/song/media/outer/url?id=836411.mp3",
            poster: "http://p1.music.126.net/80iknoJfJHHLUSjz_EiSDw==/18736777650906624.jpg?param=130y130"
        },
        {
            title: "からくりピエロ (instrumental)",
            artist: "MUSIRISCA",
            mp3: "http://music.163.com/song/media/outer/url?id=33004714.mp3",
            poster: "http://p1.music.126.net/9kAx7AtyLbLop_XhJy3y5w==/3415083117091737.jpg?param=130y130"
        },
        {
            title: "M04",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=591901.mp3",
            poster: "http://p1.music.126.net/Rm8d72Gom9BZcDOmPBPPkA==/2923601420858492.jpg?param=130y130"
        },
        {
            title: "M35",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=28267707.mp3",
            poster: "http://p1.music.126.net/a-fgC2MhmKzNNOrmrQ_CMA==/5972547162126209.jpg?param=130y130"
        },
        {
            title: "哈尔的移动城堡",
            artist: "久石譲",
            mp3: "http://music.163.com/song/media/outer/url?id=481390254.mp3",
            poster: "http://p1.music.126.net/BmJ2bUsQwinDU2KiDsKkEQ==/5998935441331080.jpg?param=130y130"
        },
        {
            title: "流れ星 ",
            artist: "久石譲",
            mp3: "http://music.163.com/song/media/outer/url?id=26902975.mp3",
            poster: "http://p1.music.126.net/cJrHNkktHNG62uKdYvGahg==/4453022092508799.jpg?param=130y130"
        },
        {
            title: "Memories",
            artist: "Within Temptation",
            mp3: "http://music.163.com/song/media/outer/url?id=407002778.mp3",
            poster: "http://p1.music.126.net/cqNHhKOTcfkwIM8_YIaT1w==/3275445150564990.jpg?param=130y130"
        },
        {
            title: "それがあなたの幸せとしても",
            artist: "rairu",
            mp3: "http://music.163.com/song/media/outer/url?id=41654821.mp3",
            poster: "http://p1.music.126.net/SpovasHBud2A1qXXADXsBg==/109951163167455610.jpg?param=130y130"
        },
        {
            title: "Angel",
            artist: "阿桑",
            mp3: "http://music.163.com/song/media/outer/url?id=205276.mp3",
            poster: "http://p1.music.126.net/8cSVJulJa2tiLydRxyXuTg==/109951162938339077.jpg?param=130y130"
        },
        {
            title: "兰若词",
            artist: "刘亦菲",
            mp3: "http://music.163.com/song/media/outer/url?id=255739.mp3",
            poster: "http://p1.music.126.net/L4Sah2hA5QYBPUnpjjUQ0Q==/26388279081790.jpg?param=130y130"
        },
        {
            title: "Pieces Of My Words-言の花-",
            artist: "刘亦菲",
            mp3: "http://music.163.com/song/media/outer/url?id=255805.mp3",
            poster: "http://p1.music.126.net/L4Sah2hA5QYBPUnpjjUQ0Q==/26388279081790.jpg?param=130y130"
        }],
        //另一方面的 信息显示 输出表
        spirit: [{
            title: "M19+20",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=591753.mp3 ",
            // 上传者的 头像 图标
            poster: "http://p1.music.126.net/lW4YKD6cMgm32nI66CzWVg==/5702067301704441.jpg?param=130y130"
        },
        {
            title: "Sis puella magica!",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=496902072.mp3",
            poster: "http://p1.music.126.net/tFTRt1H87rReNTyO1K9IDQ==/18498183627713149.jpg?param=130y130"
        },
        {
            title: "月は优しく (月)",
            artist: "梶浦由記",
            mp3: "http://music.163.com/song/media/outer/url?id=590623.mp3",
            poster: "http://p1.music.126.net/7xaV2qB-T1d9m8b1XZC6tQ==/725677674344222.jpg?param=130y130",
        },
        {
            title: "Euterpe エウテルペ ",
            artist: "染音若蔡",
            mp3: "http://music.163.com/song/media/outer/url?id=452654214.mp3",
            poster: "http://p1.music.126.net/Iqckrd2sOB1ztqrSOw4XzA==/109951162841140691.jpg?param=130y130"
        },
        {
            title: "幻光",
            artist: "杨秉音",
            mp3: "http://music.163.com/song/media/outer/url?id=526989692.mp3",
            poster: "http://p1.music.126.net/4xHOkSVWH-n6p5pB3Jf0yQ==/109951162922204274.jpg"
        },
        {
            title: "世界の約束",
            artist: "神罗Shinra",
            mp3: "http://music.163.com/song/media/outer/url?id=429460870.mp3",
            poster: "http://p1.music.126.net/rkJSVKRZkfLXOoVwXtiB4w==/18283778858733705.jpg?param=130y130"
        },
        {
            title: "Cave OF Mind",
            artist: "久石譲",
            mp3: "http://music.163.com/song/media/outer/url?id=28457572.mp3",
            poster: "http://p1.music.126.net/HdmtedPRZEEBduHcmUnk3w==/853221023209311.jpg?param=130y130"
        },
        {
            title: "幽灵公主",
            artist: "K. Williams",
            mp3: "http://music.163.com/song/media/outer/url?id=22812274.mp3",
            poster: "http://p1.music.126.net/GNKtRK8w7edPw3jAsavL2A==/5980243743832365.jpg?param=130y130"
        },
        {
            title: "心之逆鳞",
            artist: "魏小涵",
            mp3: "http://music.163.com/song/media/outer/url?id=591753.mp3",
            poster: "http://p1.music.126.net/k_WRxDY1qQ4ztB5uFFrvoA==/17907745881679448.jpg?param=130y130"
        },
        {
            title: "愛を教えてくれた君へ",
            artist: "Qaijff",
            mp3: "http://music.163.com/song/media/outer/url?id=521416051.mp3",
            poster: "http://p1.music.126.net/-c3qURPNRNLe-YJMbiZoKA==/109951163072509863.jpg?param=130y130"
        },
        {
            title: "群雄疾走",
            artist: "川井憲次",
            mp3: "http://music.163.com/song/media/outer/url?id=448153.mp3",
            poster: "http://p1.music.126.net/r4TK33y6f8cwlntVidXZbQ==/931286348726555.jpg?param=130y130"
        },
        {
            title: "Ghost of a smile",
            artist: "EGOIST",
            mp3: "http://music.163.com/song/media/outer/url?id=35955908.mp3",
            poster: "http://p1.music.126.net/ivONokvElv9ZCzyrZp84FQ==/3297435373557125.jpg?param=130y130"
        },
        {
            title: "樱子小姐的脚下埋着尸体",
            artist: "大竹佑季",
            mp3: "http://music.163.com/song/media/outer/url?id=36271375.mp3",
            poster: "http://p1.music.126.net/Q4Dg5QXwft213TBKMv26_A==/3276544653004159.jpg?param=130y130"
        },
        {
            title: "非科学的表裏一体",
            artist: "豚乙女",
            mp3: "http://music.163.com/song/media/outer/url?id=30870899.mp3",
            poster: "http://p1.music.126.net/84dpde0vkfsDAVsNNjulXg==/7906588115750467.jpg?param=130y130"
        },
        {
            title: "You're the Shine",
            artist: "：FELT",
            mp3: "http://music.163.com/song/media/outer/url?id=26260757.mp3",
            poster: "http://p1.music.126.net/b04i7LFbHLJkmkzwhwRLMA==/2343059278838229.jpg?param=130y130"
        },
        {
            title: "旅の途中",
            artist: "清浦夏実",
            mp3: "http://music.163.com/song/media/outer/url?id=26220167.mp3",
            poster: "http://p1.music.126.net/4BgAnUbCDFex3m4z-hWULA==/2509085534622060.jpg?param=130y130"
        },
        {
            title: "夏祭り",
            artist: "東山奈央",
            mp3: "http://music.163.com/song/media/outer/url?id=488388729.mp3",
            poster: "http://p1.music.126.net/3eyBH8RjxjXG-EqWShU1wg==/18887410742154555.jpg?param=130y130"
        },
        {
            title: "Sway",
            artist: "Nevve",
            mp3: "http://music.163.com/song/media/outer/url?id=475073464.mp3",
            poster: "http://p1.music.126.net/KmPcFcxxg61d15R8yu5x_A==/18681802069425034.jpg?param=130y130"
        },
        {
            title: "Vanish",
            artist: " Breathe Carolina",
            mp3: "http://music.163.com/song/media/outer/url?id=427542077.mp3",
            poster: "http://p1.music.126.net/xaX_RkkW0cT4f38k62N8yg==/3413983630702236.jpg?param=130y130"
        },
        {
            title: "It's Over",
            artist: "MEIDEN",
            mp3: "http://music.163.com/song/media/outer/url?id=477933011.mp3",
            poster: "http://p1.music.126.net/foJM2P9nq8pXHnCZjcf75w==/19047939439716625.jpg?param=130y130"
        }],
        searchInput: null,
        searchInputId: "searchInput",
        test: "helloworld",
        //尾部的link信息的 信息
        links:[
            { title:"作品集",text:"个人作品集合",image:"images/314601.jpg",link:"page/pbl/main"},
            { title:"数据中心",text:"可爱的基友在哪里？？",image:"images/316718.jpg",link:"page/mit/index"},
            { title:"上传作品",text:"哦,你有什么宝贝 , 者行孙！！",image:"images/317839.jpg" , link:"page/upload/audio"}
        ],
        //当前userId
        userId:null
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
        },
        init: function() {
            for (var index = 0; index < this.latest.length; index++) {
                var item = this.latest[index];
                this.format(item);
            }

            for (var index = 0; index < this.spirit.length; index++) {
                var item = this.spirit[index];
                this.format(item);
            }

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
        },clickPlayMe: function(e) {
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
    }
});