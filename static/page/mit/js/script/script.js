/**

对应的方法目录

对应的

1.新建收藏夹 [readyAddFolder]
2.修改收藏夹信息 [ readyUpdateFolder]
3.删除收藏夹 [ readyDeleteFolder ]
4.准备复制收藏关系 readyCopyCollection
5.重新修改密码 [ rewritePwd ]
*/
new Vue({
    el: "#app",
    data: {
        string: "helloworld",
        url:"",//目前的url
        //当前页面情况 
        /**
            对应的程序逻辑说明
        */
        page: ["index", "password", "forms", "register", "tables", "collection", "audio"],
        currentPage: "index",
        navMain: [{
            name: "数据中心",
            iClass: "icon-chart",
            page: 0
        },
        {
            name: "用户信息",
            iClass: "icon-padnote",
            page: 2,
            clickListener:function( item , vue ){
                /**
                    下面所要进行的操作
                    搜索用户的一些详细信息
                */
                var userId = vue.userId;

                var url = getServerUrl("api/user/details/find");
                window.AJAX_ENGINE.ajax({
                    url:url,
                    data:{
                        id:userId
                    },dataType:"json",
                    async:true,
                    success:function( result , status , xhr ){
                        vue.userDetails = result;
                    }
                })
            }
        },
        {
            name: "修改密码",
            iClass: "icon-settings",
            page: 1
        },
        {
            name: "发布作品",
            iClass: "icon-padnote",
            page: 6,
            clickListener:function( item , vue ){

                //初始化发布作品的
                var userId = vue.userId;

                //获取播放量最高的五个信息
                var limit = 3 ;
                //获取url
                var url = getServerUrl("api/audio/search/dashboard");

                window.AJAX_ENGINE.ajax({
                    url:url,
                    data:{
                        userId:userId,
                        limit:limit
                    },async:true,
                    dataType:"json",
                    success:function( result , status , xhr ){
                        vue.mostPlayAudio = result;

                        var labels = [];
                        var datas = [];
                        //循环的播放对应的界面
                        for( var index = 0 ; index < result.length ; index++){
                            var audio = result[index];
                            //添加对应的属性
                            labels.push(audio.name);
                            datas.push(audio.browseCount);
                        }

                        var data = {
                            labels: labels,
                            datasets: [
                            {
                            data: datas,
                            borderWidth: 0,
                            backgroundColor: [
                                '#723ac3',
                                "#864DD9",
                                "#9762e6",
                                "#a678eb"
                            ],
                            hoverBackgroundColor: [
                                '#723ac3',
                                "#864DD9",
                                "#9762e6",
                                "#a678eb"
                            ]
                        }]
                        };

                        //输入属性
                        var jq = $("#salesBarChart1")
                        var pie = window.CHART.pie( jq , data , function( event , legendItem ){
                            console.log( event );
                            console.log( legendItem );
                        } );
                    }
                });
                
                //初始化名字
                vue.searchAudio( vue.currentAudioIndex , vue.currentAudioCount);

                var count_url = getServerUrl("api/audio/count");
                window.AJAX_ENGINE.ajax({
                    url:count_url,
                    data:{
                        userId:userId
                    },async:true,
                    dataType:"json",
                    success:function( result , status , xhr ){
                        var len = result;
                        vue.currentAudioPageLen = len / vue.currentAudioCount;
                    }
                });
            }
        }],
        label:{
            details:{
                //用户详细信息
                title:"用户详细信息"
            },user:{
                title:"用户信息",
                name:"名字",
                introduction:"介绍信息",
                image:"用户头像",
                sex:"性别",
                age:"年龄",
                birthday:"出生年月",
                address:"家庭地址",
                cancelButton:"取消",
                saveButton:"保存",
            },login:{
                password:"登录密码",
                newPassword:"请输入新密码",
                pwdAgain:"请重新输入密码"
            }
        },
        navExtras: [{
            name: "Demo",
            iClass: "icon-settings",
            page: 1
        },
        {
            name: "Demo",
            iClass: "icon-writing-whiteboard",
            page: 0
        },
        {
            name: "Demo",
            iClass: "icon-chart",
            page: 0
        }]
        //用户与其有关的信息
        ,
        userId: null,
        user: {
            id: "",
            name: ""
        },userDetails:{
            //详细信息
        },
        dateArray: {
            //日期列表排期
        },
        userDashboard: {
            //根据目标的信息来生成对应的xinxi
            browseCount: "--",
            forwardCount: "",
            loveCount: "--"
        },
        userDateDashboard: {
            //用户近十日的数据
        }
        //收藏视频有关的数据信息
        ,currentCollection: [
        //输出收藏的信息 ，输出对应的 collection 的信息
        {
            id: 123,
            userId: 123,
            audioId: "",
            createTime: "xxx",
            createTimeStruct: "",
            audio: {
                id: 123,
                name: "ss",
                introduction: "sss"
            }
        },
        {
            id: 123,
            userId: 123,
            audioId: "",
            createTime: "xxx",
            createTimeStruct: "",
            audio: {
                id: 123,
                name: "ss",
                introduction: "sss"
            }
        },
        {
            id: 123,
            userId: 123,
            audioId: "",
            createTime: "xxx",
            createTimeStruct: "",
            audio: {
                id: 123,
                name: "ss",
                introduction: "sss"
            }
        },
        {
            id: 123,
            userId: 123,
            audioId: "",
            createTime: "xxx",
            createTimeStruct: "",
            audio: {
                id: 123,
                name: "ss",
                introduction: "sss"
            }
        },
        {
            id: 123,
            userId: 123,
            audioId: "",
            createTime: "xxx",
            createTimeStruct: "",
            audio: {
                id: 123,
                name: "ss",
                introduction: "sss"
            }
        }],
        currentFolder: {
            //表示当前收藏夹的信息
            id: 0,
            name: "",
            introduction: "xxx"
        },
        currentFolderIndex: -1,
        collectionFolders: [

        ],
        folderObj: {

        },mostPlayAudio:{
            //最多播放量的节目
        },currentAudio:{
            //当前音频信息

        },currentAudioIndex:0, // 当前audio 的 index 标签,
        currentAudioCount:7, //当前一个页面 的数量,
        currentAudioPageLen:0
    },
    methods: {
        selectPage: function(event, item) {
            //获取对应的 信息
            var target = $(event.target);
            var currentPage = target.find(".currentPage");

            var pageName = currentPage.text().trim();

            this.gotoPage(item);

        },
        gotoPage: function(item) {

            var toPage = this.page[item.page];
            /*if( this.currentPage == toPage ){
                return;
            }
            */
            this.currentPage = toPage;

            if (item.clickListener != undefined) {
                item.clickListener(item , this );
            }

        },
        init: function() {
            //初始化 相对应的 时
            this.initUser();
            this.initCollectionFolder();
            this.initIndexPage();
            //初始化绘图方法
            this.initCharts();
        },
        initUser: function() {
            //获取对应程序背后的 userId 信息 INIT_USER_INFORMATION
            var userId = GetQueryString("userId");
            this.userId = userId;
            var vue = this;
            var url = getServerUrl("api/user/find/id");
            window.AJAX_ENGINE.ajax({
                url: url,
                data: {
                    id: userId
                },
                async: true,
                dataType: "json",
                success: function(result, status, xhr) {
                    vue.user = result
                }
            })
        },
        initCollectionFolder: function() {
            //初始化收藏文件夹的参数
            var folderObj = {
                name: "收藏",
                iClass: "icon-windows",
                hasChildren: true,
                connection: "exampledropdownDropdown_collection_1"
            }

            var childrens = [];

            //下面开始初始化信息
            var userId = this.userId;
            //开始进行请求
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url: "/yinji/api/collection/folder/all",
                data: {
                    userId: userId
                },
                async: false,
                dataType: "json",
                success: function(result, status, xhr) {
                    //将对应的result转化为信息，然后插入 childrens 之中
                    for (var index = 0; index < result.length; index++) {
                        var item = result[index];

                        if (item.name == "") {
                            item.name = "未命名";
                        }

                        var _item = vue.collectionFolderToItem(item);
                        childrens.push(_item);
                    }
                }
            })

            //将对应的结果保存进入对应的信息
            folderObj.childrens = childrens;

            this.collectionFolders = childrens;

            this.folderObj = folderObj;

            this.navMain.push(this.folderObj);
        },
        initIndexPage: function() {
            //初始化index page的信息
            var vue = this;
            var userId = this.userId;

            /**
                
                if( userId == null ){
                    return ;
                }

            */

            var url = getServerUrl("api/dashboard/user/find");
            //获取对应的
            window.AJAX_ENGINE.ajax({
                url: url,
                data: {
                    id: userId
                },
                dataType: "json",
                async: true,
                success: function(result, status, xhr) {
                    vue.userDashboard = result;
                }
            });

        },
        initCharts: function() {
            //初始化 画图函数
            var init_func = window.CHART_FUNC;
            init_func(this);
            //获取今日的时间 ，以及前十天的准确日期
            var today = Today();

            var end = today.getTime();
            var start = end - 10 * DAY_MILLI_SEC;

            /*
            var times_array = [];

            
            var labels = [];
            //获取时间数组
            
            
            for( var index = 0 ; index < 10 ; index++ ){
                var times = start + index * DAY_MILLI_SEC
                var times_struct = parseTimes( times );
                //输出 名称 月 ， 日期
                var label = times_struct.month + "-" + times_struct.day;
                labels.push( label);
                //将转化的数据转化到tims_array之中
                times_array.push( times_struct );
            }

            this.dateArray = times_array;
            */

            var url = getServerUrl("api/dashboard/user/date/search");
            var userId = this.userId;

            var vue = this;
            window.AJAX_ENGINE.ajax({
                url: url,
                data: {
                    userId: userId,
                    start: start,
                    end: end
                },
                async: true,
                dataType: "json",
                success: function(result, status, xhr) {
                    vue.userDateDashboard = result;

                    var dateArray = vue.dateArray;

                    var min = 0;
                    var max = 0;

                    //播放量与点击量的数组
                    var lovesArray = [];
                    var browseArray = [];
                    var times_array = [];
                    var labels = [];

                    for (var index = 0; index < 10; index++) {
                        var times = start + index * DAY_MILLI_SEC
                        var times_struct = parseTimes(times);
                        //输出 名称 月 ， 日期
                        var label = times_struct.month + "-" + times_struct.day;
                        labels.push(label);
                        //将转化的数据转化到tims_array之中
                        times_array.push(times_struct);

                        var value = result[times];

                        //得到对应的点赞数量与播放数量的值
                        var love = 0;
                        var browse = 0;

                        if (value != undefined) {
                            love = value.loveCount;
                            browse = value.browseCount;
                        }

                        if( love > browse ){
                            //若果love 值 大于 browse 值
                            max = max > love ? max :love;
                            min = min < browse ? min : browse;
                        }else{
                            max = max > browse ? max : love ;
                            min = min < love ? min : love;
                        }

                        lovesArray.push(love);
                        browseArray.push(browse);
                    }


                    

                    var LINECHART = $('#lineCahrt');
                    window.CHART.line(LINECHART, {
                        labels: labels,
                        datasets: [{
                            label: "播放量",
                            fill: true,
                            lineTension: 0.2,
                            backgroundColor: "transparent",
                            borderColor: '#864DD9',
                            pointBorderColor: '#864DD9',
                            pointHoverBackgroundColor: '#864DD9',
                            borderCapStyle: 'butt',
                            borderDash: [],
                            borderDashOffset: 0.0,
                            borderJoinStyle: 'miter',
                            borderWidth: 2,
                            pointBackgroundColor: "#fff",
                            pointBorderWidth: 5,
                            pointHoverRadius: 5,
                            pointHoverBorderColor: "#fff",
                            pointHoverBorderWidth: 2,
                            pointRadius: 1,
                            pointHitRadius: 0,
                            //data: [20, 27, 20, 35, 30, 40, 33, 25, 39, 31],
                            data:browseArray,
                            spanGaps: false
                        },
                        {
                            label: "点赞量",
                            fill: true,
                            lineTension: 0.2,
                            backgroundColor: "transparent",
                            borderColor: "#EF8C99",
                            pointBorderColor: '#EF8C99',
                            pointHoverBackgroundColor: "#EF8C99",
                            borderCapStyle: 'butt',
                            borderDash: [],
                            borderDashOffset: 0.0,
                            borderJoinStyle: 'miter',
                            borderWidth: 2,
                            pointBackgroundColor: "#fff",
                            pointBorderWidth: 5,
                            pointHoverRadius: 5,
                            pointHoverBorderColor: "#fff",
                            pointHoverBorderWidth: 2,
                            pointRadius: 1,
                            pointHitRadius: 10,
                            //data: [25, 17, 28, 25, 33, 27, 30, 33, 27, 31],
                            data:lovesArray,
                            spanGaps: false
                        }]
                    },
                    min, max)

                    

                    vue.dateArray = times_array;
                }
            })
        },
        collectionFolderToItem: function(item) {
            var _item = item;
            _item.page = 5;
            _item.clickListener = this.searchCollectionAudio
            return _item;
        },
        searchCollectionAudio: function(item , _vue ) {
            //根据目标的 id ( 收藏夹 [collection_folder]  , 来搜索 ，该id下面的所有的信息)
            var id = item.id;

            var url = "/yinji" + "/" + "api/collection/all/and";

            var vue = this;

            //首先清空对应的信息
            vue.currentCollection = [];

            window.AJAX_ENGINE.ajax({
                url: url,
                data: {
                    folderId: id
                },
                async: true,
                dataType: "json",
                success: function(result, status, xhr) {
                    //将对应的结果输出
                    /*
                    for( var index = 0 ; index < result.length ; index ++ ){
                        var item = result[index];
                        //得到某行对应的信息
                    }
                    */
                    vue.currentCollection = result;

                }
            });

            this.currentFolder = item;

        },
        readyAddFolder: function() {
            //准备添加修改文件夹
            var vue = this;
            var userId = this.userId;
            //准备进行对应的新建收藏夹的输出化工作
            window.CONFIRM.confirm({
                theme: "black",
                title: "输入对应的信息",
                content: '<div class="form-group"><label>收藏夹名称</label><input autofocus type="text" id="collectionFolder_name" placeholder="Lets greet Ourselves." class="form-control"><label>简介</label><textarea id="collectionFolder_introduction" class="form-control" rows="4" placeholder="请写的稍微详细一点哦^_^..."></textarea></div>',
                confirm: function() {
                    var name = this.$b.find("#collectionFolder_name");
                    var introduction = this.$b.find("#collectionFolder_introduction");

                    var name_val = name.val();
                    var introduction_val = introduction.val();

                    var url = "/yinji/" + "api/collection/folder/insert";
                    window.AJAX_ENGINE.ajax({
                        url: url,
                        data: {
                            userId: userId,
                            name: name_val,
                            introduction: introduction_val
                        },
                        async: true,
                        dataType: "json",
                        success: function(result, status, xhr) {
                            //新建陈宫之后 ， 将对应的 result 加入 队列中的第一名
                            //理论性上的第一名 ， 因为第一默认认为是无情况
                            //将新建的信息
                            var item = vue.collectionFolderToItem(result);
                            vue.collectionFolders.unshift(item);
                        }
                    })

                }
            })

        },
        readyUpdateFolder: function(item, index) {
            //根据用户的意愿去修改对应的收藏夹的对应的信息 
            var old_item_name = item.name;
            var old_item_introduction = item.introduction;

            var content = '<div class="form-group"><label>收藏夹名称</label><input autofocus type="text" id="collectionFolder_name" placeholder="Lets greet Ourselves." class="form-control" value="' + old_item_name + '"><label>简介</label><textarea id="collectionFolder_introduction" class="form-control" rows="4" placeholder="请写的稍微详细一点哦^_^...">' + old_item_introduction + '</textarea></div>'

            var vue = this;

            //获取对应的收藏夹的信息
            var id = item.id;

            //准备进行对应的新建收藏夹的输出化工作
            window.CONFIRM.confirm({
                theme: "black",
                title: "输入对应的信息",
                content: content,
                confirm: function() {
                    var name = this.$b.find("#collectionFolder_name");
                    var introduction = this.$b.find("#collectionFolder_introduction");

                    var name_val = name.val();
                    var introduction_val = introduction.val();

                    var url = "/yinji/" + "api/collection/folder/update";
                    window.AJAX_ENGINE.ajax({
                        url: url,
                        data: {
                            id: id,
                            name: name_val,
                            introduction: introduction_val
                        },
                        async: true,
                        dataType: "json",
                        success: function(result, status, xhr) {
                            //新建陈宫之后 ， 将对应的 result 加入 队列中的第一名
                            //理论性上的第一名 ， 因为第一默认认为是无情况
                            //将新建的信息
                            var new_item = vue.collectionFolderToItem(result) ;
                            vue.currentFolder.id = new_item.id;
                            vue.currentFolder.name = new_item.name;
                            vue.currentFolder.introduction = new_item.introduction;

                        }
                    })

                }
            })

        },
        readyDeleteFolder: function(collectionFolder) {
            //删除对应的收藏夹关系
            var vue = this;
            var id = collectionFolder.id;

            var url = "/yinji/" + "api/collection/folder/delete";
            window.CONFIRM.confirm({
                title: "警告",
                content: "请三思而后行哦(*￣︶￣)",
                confirm: function() {
                    window.AJAX_ENGINE.ajax({
                        url: url,
                        data: {
                            id: id
                        },
                        async: true,
                        dataType: "json",
                        success: function(result, status, xhr) {
                            vue.deleteCollectionFolder(collectionFolder)
                        }
                    })
                }
            })
        },
        readyCancelCollection: function(item, index) {
            var vue = this;
            //准备开始取消收藏的方法
            window.CONFIRM.confirm({
                theme: "black",
                title: "警告",
                content: "是否取消收藏",
                confirm: function() {
                    //设置确定的信息
                    vue.cancelCollection(item, index);
                    //下面设置对应的信息
                }
            })
        },
        cancelCollection: function(item, index) {
            //根据相对应传递进来的 index 来进行删除对应的信息
            var currentCollection = this.currentCollection;

            var url = "/yinji/" + "api/collection/delete";
            window.AJAX_ENGINE.ajax({
                url: url,
                data: {
                    id: item.id
                },
                async: true,
                dataType: "json",
                success: function(result, status, xhr) {
                    $.alert("再见o(╥﹏╥)o");
                }
            })

            currentCollection.splice(index, 1);

        },
        buildFolderContentHtml: function() {
            var content = new String();
            content = '<ul class="list-group" >';
            for (var index = 0; index < this.collectionFolders.length; index++) {
                var folderItem = this.collectionFolders[index];
                var html = '<li class="list-group-item black-color">';
                html += '<div class="row">'
                //对应的 col-md-4 信息
                html += '<div class="col-md-10">';
                html += folderItem.name;
                html += '</div>';
                //对应的 col-md-2信息
                html += '<div class="col-md-2">';
                html += '<input name="folderId" type="radio" value="' + folderItem.id + '">';
                html += '</div>';

                html += "</div>";
                html += "</li>";
                content = content + html;
            }
            content += "</ul>";

            return content;
        },
        readyRemoveCollection: function(item, index) {
            //将收藏关系转移到其他的收藏夹中
            var content = this.buildFolderContentHtml();

            var vue = this;

            //准备进行转移文件夹
            window.CONFIRM.confirm({
                title: "列表",
                content: content,
                confirm: function() {
                    var check = this.$b.find("input[name='folderId']:checked");
                    var check_val = check.val();
                    var check_int = parseInt(check_val);

                    var url = "/yinji/" + "api/collection/update";
                    window.AJAX_ENGINE.ajax({
                        url: url,
                        data: {
                            id: item.id,
                            folderId: check_int
                        },
                        dataType: "json",
                        async: true,
                        success: function(result, status, xhr) {
                            //若成功 ， 则直接删除
                            //对应的信息
                            console.log(vue.currentCollection);
                            vue.currentCollection.splice(index, 1);
                            window.CONFIRM.alert({
                                title: "转移成功",
                                content: "已经转移过去喽！(。-ω-)zzz"
                            });
                        }
                    })
                }
            })
        },
        readyCopyCollection: function(item, index) {
            //将对应的collection 复制到 对应的文件夹下面
            var content = this.buildFolderContentHtml();

            var vue = this;

            var url = "/yinji/" + "api/collection/insert";

            var userId = this.userId;

            //准备进行转移文件夹
            window.CONFIRM.confirm({
                title: "列表",
                content: content,
                confirm: function() {

                    var check = this.$b.find("input[name='folderId']:checked");
                    var check_val = check.val();
                    var check_int = parseInt(check_val);

                    window.AJAX_ENGINE.ajax({
                        url: url,
                        data: {
                            userId: userId,
                            audioId: item.audioId,
                            folderId: check_int
                        },
                        async: true,
                        dataType: "json",
                        success: function(result, status, xhr) {
                            window.CONFIRM.alert({
                                title: "插入成功",
                                content: "复制成功哦o(*￣︶￣*)o"
                            })
                        }
                    })
                }
            });

        },
        deleteCollectionFolder: function(item) {
            var collectionFolders = this.collectionFolders;
            for (var index = 0; index < collectionFolders.length; index++) {
                var _item = collectionFolders[index];
                if (item.id == _item.id) {
                    collectionFolders.splice(index, 1);
                    break;
                }
            }

            this.folderObj.childrens = collectionFolders;
            this.collectionFolders = collectionFolders;

            if (this.navMain.length > 0) {
                this.gotoPage(this.navMain[0])
            }
        },rewritePwd:function(){
            var newPassword = $("#rewrite_newPassword");
            var pwdAgain = $("#rewrited_pwdAgain");

            //获取对应的信息
            var newPassword_val = newPassword.val();
            var pwdAgain_val = pwdAgain.val();

            if ( newPassword_val != pwdAgain_val ){
                window.CONFIRM.alert("两次密码不一致");
                return ;
            }

            if( newPassword_val.length <= 6 ){
                window.CONFIRM.alert("密码格式不正确，长度小于6")
                return;
            }

            var userId = this.userId;
            var url = getServerUrl("api/login/update/pwd")
            var vue = this;
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    id:userId,
                    password:newPassword_val
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    window.CONFIRM.alert("密码修改成功")
                    vue.login = result;
                    newPassword.val("");
                    pwdAgain.val("");
                }
            })
            //修改密码
            console.log("准备开始重写新密码");
        },searchAudio:function( page , count ){
            //根据对应的 起始页面 数据 来获取对应的信息
            //先获取对应的 
            var userId = this.userId;
            var vue = this;
            var url = getServerUrl("api/audio/user");

            if( page < 0 ){
                window.CONFIRM.alert("现在已经是第一页喽^_^");
                return ;
            }

            if( page > this.currentAudioPageLen ){
                window.CONFIRM.alert("超过页面限度o(>﹏<)o");
                return;
            }

            this.currentAudioIndex = page;
            
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    userId:userId,
                    page:page,
                    count:count
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    //设置对应的信息
                    vue.currentAudio = result;
                }
            });


        },updateUser:function(){
            //更新user信息
            var userId = this.userId;

            var name = $("#form_name");

            var name_val = name.val();
            var vue = this;
            var url = getServerUrl("api/user/update")

            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    id:userId,
                    name:name_val
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    vue.user = result;
                    window.CONFIRM.alert("修改成功喽^_^");
                }
            })
        },updateUserDetails:function(){
            var userId = this.userId;

            var introduction = $("#details_introduction");
            var introduction_val = introduction.val();

            var sex = $("#details_sex option:selected");
            var sex_val = sex.val();

            //出生年月
            var birthday = $("#details_birthday ");
            var birthday_val = birthday.val();

            //家庭住址
            var address = $("#details_address");
            var address_val = address.val();

            var vue = this;
            var url = getServerUrl("api/user/details/update");

            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    id:userId,
                    introduction:introduction_val,
                    sex:sex_val,
                    birthday:birthday_val,
                    address:address_val
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    vue.userDetails = result;
                    window.CONFIRM.alert("修改成功哦");
                }
            });


        }
    },
    created: function() {
        var vue = this;
        vue.init();
        this.url = window.location.href;
    }
});