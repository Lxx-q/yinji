/*
design by Voicu Apostol.
*/

const FROM_SERVER = "server";

const FROM_LOCATION = "location";


new Vue({
    el: "#app",
    data() {
        return {
            userId:2,
            audio: null,
            circleLeft: null,
            barWidth: null,
            duration: null,
            currentTime: null,
            //相对应的duration,以及currentTime的详细数据
            durationData:null,
            currentTimeData:null,
            isTimerPlaying: false,
            tracks: [
                /*
                  现在 ，我们 设定 会将 相对应的 track 信息 ， 加上 一个 index的 属性 ，
                   来表示 ， 自己在 原始 歌单中的 序列号 ， 并且 该，序列号自动生成,
                   我们再使用一个 相对应的 参数 ， 来表示 该歌曲 是否来自本地 ， 还是网络 ， 这个参数 ， 只能自己添加
                */
                {
                    name: "Mekanın Sahibi",
                    artist: "Norm Ender",
                    cover: "t2.png",
                    source: "mp3/2.amr",
                    url: "#",
                    favorited: false,
                    from: FROM_SERVER
                },{
                    name: "Everybody Knows",
                    artist: "Leonard Cohen",
                    cover: "t1.png",
                    source: "mp3/1.mp3",
                    url: "#",
                    favorited: true,
                    from: FROM_SERVER
                },
                 {
                    name: "Da yu",
                    artist: "Zhou 深",
                    cover: "t2.png",
                    source: "mp3/3.mp3",
                    url: "#",
                    favorited: true,
                    from: FROM_SERVER
                }, {
                    name: "来不及勇敢",
                    artist: "周深",
                    cover: "t2.png",
                    source: "mp3/4.mp3",
                    url: "#",
                    favorited: true,
                    from: FROM_SERVER
                }
            ],
            currentTrack: null,
            currentTrackIndex: 0,
            transitionName: null,

            //播放模式
            models: [
                {
                    name: "顺序播放", code: "0", svg: "#icon-next", settingPlayTrack(app) {
                        var tracks = app.tracks;
                        app.playTracks = tracks;

                        //我们 需要 给 所有 会返回 原本属性的 程序 都加上 相对应的 这一行
                        app.currentTrackIndex = app.currentTrack.index;

                    }, opened: function (app) {
                        app.nextTrack();
                    }
                },
                {
                    name: "单曲播放", code: "1", svg: "#icon-prev", settingPlayTrack(app) {
                        var tracks = app.tracks;
                        app.playTracks = tracks;
                        app.currentTrackIndex = app.currentTrack.index;

                    }, opened: function (app) {
                        app.prevTrack();
                        app.nextTrack();
                    }
                },
                {
                    name: "随机播放", code: "2", svg: "#icon-link", settingPlayTrack(app) {

                        var tracks = [];

                        for (var index = 0; index < app.tracks.length; index++) {

                            tracks[index] = app.tracks[index];
                        }


                        var length = tracks.length;
                        var currentTrackIndex = app.currentTrackIndex;

                        /*
                          我们 的 乱序算法 为如此 ， 让一个数组的 某一个数值 ，
                           从头到尾的开始计算 ， 倘若 ， 如果是 计算之后 允许 进行调换，那么 我们 便将他们的 位置 进行 调换
                           我们的 目的 ， 就是 让 数组不按预算的 被换组 ，
                           我们只需要注意 一件事 便可以 ， 便是 ， 倘若被调用的双方 中的 某个值 ， 等于当前的 播放的键值 ， 我们 便
                           不对其进行操作 ， 变能保证 现在播放的 歌曲 ， 还在之前的位置上
                        */
                        for (var index = 0; index < length; index++) {

                            for (var _index = 0; _index < length; _index++) {

                                if (index == _index) {
                                    //倘若 相对应的 index == _index , 那么 我们也可以跳过计算
                                    continue;
                                }

                                //倘若只要 交换双方只要有一个值等于相对应的 当前的 ，值 ， 我们就跳过相对应的交换
                                if (currentTrackIndex == index || currentTrackIndex == _index) {
                                    continue;
                                }


                                //获取 一个 随机数 0 -  99
                                var random = Math.random() * 100;

                                //倘若这个值 ， 大于 49 ， 我们便进行 相对应的 交换的工作
                                if (random > 49) {

                                    var temp_element = tracks[index];
                                    tracks[index] = tracks[_index];
                                    tracks[_index] = temp_element;
                                }

                            }
                        }

                        app.playTracks = tracks;

                    }, opened: function (app) {

                        //下面 我们 开始 设定 相对应的 属性

                        //倘若当前， 播放值最后一个 ， 我们还是将相对应的播放顺序打乱
                        var currentTrackIndex = app.currentTrackIndex;

                        var length = app.tracks.length;

                        if (currentTrackIndex == length - 1) {
                            this.settingPlayTrack(app);
                        }

                        app.nextTrack();

                    }
                }
            ],
            //相对应的 播放模式的 参数
            currentModel: null,
            currentModelIndex: 0,

            //相对应的 模仿 列表
            playTracks: null,
            appendInput: null,
            //下面 我们 自己设定的 reload 的 字段
            reloadString: "reload",
            volume:0.0
        };
    },
    methods: {
        play() {
            //if (this.audios.paused) {
            if( !this.isTimerPlaying){
                //this.audios.play();
                this.isTimerPlaying = true;
                this.playAudio();
            } else {
                //this.audios.pause();
                this.isTimerPlaying = false;
                this.pauseAudio();
            }
        },
        generateTime() {

            //获取 相对应的 两个参数 ， 一个为 duration
            //另一个为相对应的 currentTime
            var duration = this.durationData;

            var currentTime = this.currentTimeData;


            let width = (100 / duration) * currentTime;
            this.barWidth = width + "%";
            this.circleLeft = width + "%";
            let durmin = Math.floor(duration / 60);
            let dursec = Math.floor(duration - durmin * 60);
            let curmin = Math.floor(currentTime / 60);
            let cursec = Math.floor(currentTime - curmin * 60);
            if (durmin < 10) {
                durmin = "0" + durmin;
            }
            if (dursec < 10) {
                dursec = "0" + dursec;
            }
            if (curmin < 10) {
                curmin = "0" + curmin;
            }
            if (cursec < 10) {
                cursec = "0" + cursec;
            }
            this.duration = durmin + ":" + dursec;
            this.currentTime = curmin + ":" + cursec;
        },
        updateBar(x) {
            let progress = this.$refs.progress;
            let maxduration = this.durationData;
            let position = x - progress.offsetLeft;
            let percentage = (100 * position) / progress.offsetWidth;
            if (percentage > 100) {
                percentage = 100;
            }
            if (percentage < 0) {
                percentage = 0;
            }
            this.barWidth = percentage + "%";
            this.circleLeft = percentage + "%";
            this.audio.currentTime = (maxduration * percentage) / 100;
            
            this.currentAudioData();

            //下面 开始创建对应的 信息
            //this.audios.play();
            this.playAudio();
        },
        clickProgress(e) {
            this.isTimerPlaying = true;
            this.pauseAudio();

            this.updateBar(e.pageX);


        },
        initTracks: function () {
            var tracks = this.tracks;

            for (var index = 0; index < tracks.length; index++) {
                var track = tracks[index];
                track.index = index;
            }

        }
        ,

        //我们 在这边 设置 相对应的 播放列表的 函数
        settingPlayTrack: function () {
            //默认 我们 使用 当前播放的 方法
            //this.playTracks = this.tracks;

            var model = this.currentModel;

            if (model != null) {
                model.settingPlayTrack(this);
            }

        },
        settingCurrentTrack: function () {
            // 配置 相对应的 当前的 播放的 周期
            this.currentTrack = this.playTracks[this.currentTrackIndex];
            this.resetPlayer();
        },
        prevTrack() {

            this.transitionName = "scale-in";
            this.isShowCover = false;
            if (this.currentTrackIndex > 0) {
                this.currentTrackIndex--;
            } else {
                this.currentTrackIndex = this.playTracks.length - 1;
            }
            this.settingCurrentTrack();
        },
        nextTrack() {

            //倘若为 不允许 切换 ， 那么 就不允许 执行 该方法


            this.transitionName = "scale-out";
            this.isShowCover = false;
            if (this.currentTrackIndex < this.playTracks.length - 1) {
                this.currentTrackIndex++;
            } else {
                this.currentTrackIndex = 0;
            }

            this.settingCurrentTrack();
        },
        resetPlayer() {
            this.barWidth = 0;
            this.circleLeft = 0;
            this.audio.currentTime = 0;
            this.audio.src = this.currentTrack.source;

            setTimeout(() => {
                if (this.isTimerPlaying) {
                    //this.audios.play();
                    this.playAudio();
                } else {
                    //this.audios.pause();
                    this.pauseAudio();
                }
            }, 300);
            
        },
        favorite() {

            var userId = this.userId;
            var track = this.tracks[this.currentTrackIndex];

            //辨别当前的操作是进行点赞还是取消点赞
            track.favorited = !track.favorited;
        },
        helloworld: function () {
            alert("hello , world");
        },
        //设置 相对应的currentModel 的值
        settingCurrrentModel: function () {
            this.currentModel = this.models[this.currentModelIndex];
        },
        //下面 是 切换相对应的 播放模式的函数
        nextModel: function () {

            //获取 相对应的 model 的长度
            var length = this.models.length;

            //根据 相对应的 数字 来进行切换
            var currentModelIndex = this.currentModelIndex + 1;

            this.currentModelIndex = currentModelIndex % length;

            //最后 设置 相对应的 取值
            this.settingCurrrentModel();

            //之后 我们 我们 设置 相对应的
            this.settingPlayTrack();

        }, uploadFile: function () {

            //相对应的 上传 文件的 函数

            if (this.appendInput == null) {
                this.appendInput = document.getElementById("append_input");
            }

            var appendInput = this.appendInput;

            appendInput.click();
        }, appendFile: function () {
            // 将 相对应的 file 存进 表单

            var files = this.appendInput.files;

            for (var index = 0; index < files.length; index++) {
                //获取目标的 文件
                var file = files[index];

                this.fileToTrack(file);

            }
        }, fileToTrack: function (file) {

            var track = {};
            track.name = file.name;
            track.artist = "unknow";
            track.cover = "t1.png";
            track.source = window.URL.createObjectURL(file);
            track.url = "#";
            track.favorited = false;
            track.from = FROM_LOCATION;

            this.appendTrack(track);

        }, appendTrack(track) {
            //获取 相对应的 track 的 长度
            var track_length = this.tracks.length;
            var play_track_length = this.playTracks.length;

            track.index = track_length;

            //之后 将其 添加至 相对应的 地方
            this.$set(this.tracks, track_length, track);

            this.$set(this.playTracks, track_length, track);

            // 不知道 为什么 ， 总之 ， 我们 需要 下面的程序 刷新 下 页面的 数据
        }, download: function () {

            // 倘若目标来自于本地 ， 那么 我们便不允许 其下载
            //这个功能其实 可以放在 html 页面 ， 让 vue 来完成
            var track = this.currentTrack;

            if (track.from != FROM_SERVER) {
                alert("该文件为本地文件，不提供下载");
                return;
            }
            location.href = track.source;
        },
        //下面 我们  设置 相对应的 方法
        timeupdate: function () {
            this.generateTime();
        }, loadedmetadata: function () {
            //下面 我们 也设置 相对应的 事件
            this.generateTime();
        },//使用了对应的 信息 来进行覆盖方便对方覆盖对应的操作
        playAudio:function(){
            this.audio.play();
        },pauseAudio:function(){
            this.audio.pause();
        },currentAudioData:function(){
            this.currentTimeData = this.audio.currentTime;
            this.durationData = this.audio.duration;
        },setVolume:function(){
            this.audio.volume = this.volume;
        }
    },
    created() {

        let vm = this;

        window.player = vm;
        this.currentTrack = this.tracks[0];
        this.audio = new Audio();

        this.audio.src = this.currentTrack.source;
        this.audio.ontimeupdate = function () {
            vm.timeupdate();
        };
        this.audio.onloadedmetadata = function () {

            vm.currentAudioData();
            vm.loadedmetadata();
        };
        this.audio.onended = function () {

            var model = vm.currentModel;

            model.opened(vm);

            vm.isTimerPlaying = true;
        };


        // this is optional (for preload covers)
        for (let index = 0; index < this.tracks.length; index++) {
            const element = this.tracks[index];
            let link = document.createElement('link');
            link.rel = "prefetch";
            link.href = element.cover;
            link.as = "image"
            document.head.appendChild(link)
        }


        this.initTracks();
        //初始化 相对应的 Model
        this.settingCurrrentModel();
        this.settingPlayTrack();

        //添加对应的参数


        vm.volume =  this.audio.volume;

        console.log("hello , world");
    }
});