/**

对应的方法目录
1.新建收藏夹 [readyAddFolder]
2.修改收藏夹信息 [ readyUpdateFolder]
3.删除收藏夹 [ readyDeleteFolder ]
4.准备复制收藏关系 readyCopyCollection
*/
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
        },currentFolderIndex:-1,
        collectionFolders:[

        ],folderObj:{

        }
    },methods:{
        selectPage:function( event , item ){
            //获取对应的 信息
            var target = $( event.target );
            var currentPage = target.find(".currentPage");

            var pageName = currentPage.text().trim();

            this.gotoPage(item);

        },gotoPage:function( item ){
            this.currentPage = this.page[item.page];

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

            var vue = this;
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

                        var _item = vue.collectionFolderToItem(item);
                        childrens.push( _item );
                    }
                }
            })

            //将对应的结果保存进入对应的信息
            folderObj.childrens = childrens;

            this.collectionFolders = childrens;

            this.folderObj = folderObj;

            this.navMain.push(this.folderObj);
        },collectionFolderToItem:function(item){
            var _item = item;
            _item.page = 5;
            _item.clickListener = this.searchCollectionAudio
            return _item;
        },searchCollectionAudio:function( item ){
            //根据目标的 id ( 收藏夹 [collection_folder]  , 来搜索 ，该id下面的所有的信息)

            var id = item.id;

            var url = "/yinji"  + "/" + "api/collection/all/and";

            var vue = this;
            
            //首先清空对应的信息
            vue.currentCollection = [];

            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    folderId:id
                },async:true
                ,dataType:"json"
                ,success:function( result ,status , xhr ){
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


        },readyAddFolder:function(){
            //准备添加修改文件夹
            var vue = this;
            //准备进行对应的新建收藏夹的输出化工作
            window.CONFIRM.confirm({
                theme:"black",
                title:"输入对应的信息",
                content:'<div class="form-group"><label>收藏夹名称</label><input autofocus type="text" id="collectionFolder_name" placeholder="Lets greet Ourselves." class="form-control"><label>简介</label><textarea id="collectionFolder_introduction" class="form-control" rows="4" placeholder="请写的稍微详细一点哦^_^..."></textarea></div>'   
                ,confirm:function(){
                    var name = this.$b.find("#collectionFolder_name");
                    var introduction = this.$b.find("#collectionFolder_introduction");
                    
                    var name_val = name.val();
                    var introduction_val = introduction.val();

                    var url = "/yinji/" + "api/collection/folder/insert"
                    window.AJAX_ENGINE.ajax({
                        url:url,
                        data:{
                            userId:2,
                            name:name_val,
                            introduction:introduction_val
                        },async:true,
                        dataType:"json",
                        success:function( result , status , xhr ){
                            //新建陈宫之后 ， 将对应的 result 加入 队列中的第一名
                            //理论性上的第一名 ， 因为第一默认认为是无情况
                            //将新建的信息
                            var item = vue.collectionFolderToItem(result);
                            vue.collectionFolders.unshift( item );
                        }
                    })


                }
            })

        },readyUpdateFolder:function( item , index  ){
            //根据用户的意愿去修改对应的收藏夹的对应的信息 

            var old_item_name = item.name;
            var old_item_introduction = item.introduction;

            var content = '<div class="form-group"><label>收藏夹名称</label><input autofocus type="text" id="collectionFolder_name" placeholder="Lets greet Ourselves." class="form-control" value="' + old_item_name + '"><label>简介</label><textarea id="collectionFolder_introduction" class="form-control" rows="4" placeholder="请写的稍微详细一点哦^_^...">' + old_item_introduction + '</textarea></div>'   
                
            var vue = this;

            //获取对应的收藏夹的信息
            var id = item.id;

            //准备进行对应的新建收藏夹的输出化工作
            window.CONFIRM.confirm({
                theme:"black",
                title:"输入对应的信息",
                content:content
                ,confirm:function(){
                    var name = this.$b.find("#collectionFolder_name");
                    var introduction = this.$b.find("#collectionFolder_introduction");
                    
                    var name_val = name.val();
                    var introduction_val = introduction.val();

                    var url = "/yinji/" + "api/collection/folder/update"
                    window.AJAX_ENGINE.ajax({
                        url:url,
                        data:{
                            id:id,
                            name:name_val,
                            introduction:introduction_val
                        },async:true,
                        dataType:"json",
                        success:function( result , status , xhr ){
                            //新建陈宫之后 ， 将对应的 result 加入 队列中的第一名
                            //理论性上的第一名 ， 因为第一默认认为是无情况
                            //将新建的信息
                            var new_item = vue.collectionFolderToItem(result)
                            vue.currentFolder.id = new_item.id;
                            vue.currentFolder.name = new_item.name;
                            vue.currentFolder.introduction = new_item.introduction;

                        }
                    })


                }
            })

        },readyDeleteFolder:function( collectionFolder ){
            //删除对应的收藏夹关系
            var vue = this;
            var id = collectionFolder.id;

            var url = "/yinji/" + "api/collection/folder/delete";
            window.CONFIRM.confirm({
                title:"警告",
                content:"请三思而后行哦(*￣︶￣)",
                confirm:function(){
                    window.AJAX_ENGINE.ajax({
                        url:url,
                        data:{
                            id:id
                        },async:true,
                        dataType:"json",
                        success:function( result , status , xhr ){
                            vue.deleteCollectionFolder( collectionFolder )
                        }
                    })
                }
            })
        },readyCancelCollection:function( item , index ){
            var vue = this;
            //准备开始取消收藏的方法
            window.CONFIRM.confirm({
                theme:"black",
                title:"警告",
                content:"是否取消收藏",
                confirm:function(){
                    //设置确定的信息
                    vue.cancelCollection( item , index );
                    //下面设置对应的信息
                }
            })
        },cancelCollection:function( item , index ){
            //根据相对应传递进来的 index 来进行删除对应的信息
            var currentCollection = this.currentCollection;

            var url = "/yinji/" + "api/collection/delete"
            window.AJAX_ENGINE.ajax({
                url:url,
                data:{
                    id:item.id
                },async:true,
                dataType:"json",
                success:function( result , status , xhr ){
                    $.alert("再见o(╥﹏╥)o");
                }
            })

            currentCollection.splice( index , 1);

        },buildFolderContentHtml:function(){
            var content = new String();
            content = '<ul class="list-group" >';
            for( var index = 0 ; index < this.collectionFolders.length ; index ++ ){
                var folderItem = this.collectionFolders[index];
                var html = '<li class="list-group-item black-color">'
                html+= '<div class="row">'
                //对应的 col-md-4 信息
                html += '<div class="col-md-10">';
                html += folderItem.name;
                html += '</div>';
                //对应的 col-md-2信息
                html += '<div class="col-md-2">';
                html += '<input name="folderId" type="radio" value="' + folderItem.id +'">'
                html += '</div>';

                html+="</div>"
                html += "</li>";
                content = content + html;
            }
            content += "</ul>";

            return content;
        },readyRemoveCollection:function( item , index ){
            //将收藏关系转移到其他的收藏夹中
            var content = this.buildFolderContentHtml();

            var vue = this;

            //准备进行转移文件夹
            window.CONFIRM.confirm({
                title:"列表",
                content:content,
                confirm:function(){
                    var check = this.$b.find("input[name='folderId']:checked");
                    var check_val = check.val();
                    var check_int = parseInt( check_val );

                    var url = "/yinji/" + "api/collection/update"
                    window.AJAX_ENGINE.ajax({
                        url:url,
                        data:{
                            id:item.id,
                            folderId:check_int
                        },dataType:"json",
                        async:true,
                        success:function( result , status , xhr ){
                            //若成功 ， 则直接删除
                            //对应的信息
                            console.log(vue.currentCollection);
                            vue.currentCollection.splice( index , 1 );
                            window.CONFIRM.alert({
                                title:"转移成功",
                                content:"已经转移过去喽！(。-ω-)zzz"
                            });
                        }
                    })
                }
            })
        },readyCopyCollection:function( item  , index ){
           //将对应的collection 复制到 对应的文件夹下面
           var content = this.buildFolderContentHtml();

            var vue = this;

            var url = "/yinji/" + "api/collection/insert";

            //准备进行转移文件夹
            window.CONFIRM.confirm({
                title:"列表",
                content:content,
                confirm:function(){

                    var check = this.$b.find("input[name='folderId']:checked");
                    var check_val = check.val();
                    var check_int = parseInt( check_val );

                    window.AJAX_ENGINE.ajax({
                        url:url,
                        data:{
                            userId:2,
                            audioId:item.audioId,
                            folderId:check_int
                        },async:true,
                        dataType:"json",
                        success:function( result , status , xhr ){
                            window.CONFIRM.alert({
                                title:"插入成功",
                                content:"复制成功哦o(*￣︶￣*)o"
                            })
                        }
                    })
                }
            });

        },deleteCollectionFolder:function( item ){
            var collectionFolders = this.collectionFolders;
            for( var index = 0 ; index < collectionFolders.length ; index++ ){
                var _item = collectionFolders[ index ];
                if( item.id == _item.id ){
                    collectionFolders.splice( index , 1 );
                    break;
                }
            }

            this.folderObj.childrens = collectionFolders;
            this.collectionFolders = collectionFolders;

            if( this.navMain.length > 0 ){
                this.gotoPage( this.navMain[0] )
            }
        }
    },created:function(){
        this.initCollectionFolder();
    }
});