new Vue({
    el:"#app",
    data:{
        //当前页面情况 
        /**
            对应的程序逻辑说明
        */
        page:["index","charts","forms","register","tables" , "collection"]
        ,currentPage:"index"
        ,navMain:[
            { name:"Index" , iClass:"icon-home",page:0},
            { name:"Tables" ,iClass:"icon-home" ,actived:true , page:4},
            { name:"Charts" ,iClass:"fa fa-bar-chart" ,page:1},
            { name:"Forms"  ,iClass:"icon-padnote" , page:2},
            { name:"Example dropdown" ,iClass:"icon-windows" , hasChildren:true , connection:"exampledropdownDropdown" , childrens:[
                { name :"Page" , page:1 },
                { name: "Page" , page:2 },
                { name: "Page" , page:3 }
            ]},{
                name:"收藏",iClass:"icon-windows" , hasChildren:true , connection:"exampledropdownDropdown_collection" ,childrens:[
                    { name:" 哈哈" ,page:5 , clickListener:function( item ){
                        alert("this is item :" + item.name );
                        alert("this is this :" + this.name  );
                    }},
                    { name:"哈哈哈哈" ,page:5 }
                ]}
        ],navExtras:[
            { name:"Demo" ,iClass:"icon-settings" , page:1 },
            { name:"Demo" ,iClass:"icon-writing-whiteboard",page:0},
            { name:"Demo" ,iClass:"icon-chart",page:0}
        ],currentCollection:[
            //输出收藏的信息 ，输出对应的 collection 的信息
            { id:123,userId:123,audioId:"" , createTime:"xxx",createTimeStruct:"",audio:{ id:123,name:"ss",introduction:"sss" }},
            { id:123,userId:123,audioId:"" , createTime:"xxx",createTimeStruct:"",audio:{ id:123,name:"ss",introduction:"sss" }},
            { id:123,userId:123,audioId:"" , createTime:"xxx",createTimeStruct:"",audio:{ id:123,name:"ss",introduction:"sss" }},
            { id:123,userId:123,audioId:"" , createTime:"xxx",createTimeStruct:"",audio:{ id:123,name:"ss",introduction:"sss" }},
            { id:123,userId:123,audioId:"" , createTime:"xxx",createTimeStruct:"",audio:{ id:123,name:"ss",introduction:"sss" }}
        ],currentFolder:{
            //表示当前收藏夹的信息
            id:0,name:"",introduction:"xxx"
        }
    },methods:{
        selectPage:function( event , item ){
            //获取对应的 信息
            var target = $( event.target );
            var currentPage = target.find(".currentPage");

            var pageName = currentPage.text().trim();

            this.currentPage = pageName;

            if( item.clickListener != undefined ){
                item.clickListener( item );
            }

        },initIndex:function(){
            //初始化 相对应的 时
            
        },initCollectionFolder:function(){
            //初始化收藏文件夹的参数
            var folderObj =  {name:"收藏",iClass:"icon-windows" , hasChildren:true , connection:"exampledropdownDropdown_collection_1" }

            var childrens = [];

            //下面开始初始化信息

            var userId = 2
            //开始进行请求

            var func = this.searchCollectionAudio
            window.AJAX_ENGINE.ajax({
                url:"/yinji/api/collection/folder/all",
                data:{
                    userId:userId
                },async:false,
                dataType:"json",
                success:function( result , status , xhr ){
                    //将对应的result转化为信息，然后插入 childrens 之中

                    for( var index = 0 ;  index < result.length ; index ++ ){
                        var item = result[index];

                        if(item.name == ""){
                            item.name="未命名";
                        }

                        var _item = item;
                        _item.page = 5;
                        _item.clickListener = func
                        childrens.push(_item);
                    }
                }
            })

            //将对应的结果保存进入对应的信息
            folderObj.childrens = childrens;

            this.navMain.push(folderObj);
        },searchCollectionAudio:function( item ){
            //获取对应的信息 ， 然后我们输出对应的信息
            var id = item.id



        },readyAddFolder:function(){
            //准备进行对应的新建收藏夹的输出化工作
            alert("helloworld");
        }
    },created:function(){
        this.initCollectionFolder();
    }
});