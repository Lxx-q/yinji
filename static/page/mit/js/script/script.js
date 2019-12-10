new Vue({
    el:"#app",
    data:{
        //当前页面情况 
        page:["index","charts","forms","register","tables"]
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
            ]}
        ],navExtras:[
            { name:"Demo" ,iClass:"icon-settings" , page:1 },
            { name:"Demo" ,iClass:"icon-writing-whiteboard",page:0},
            { name:"Demo" ,iClass:"icon-chart",page:0}
        ]
    },methods:{
        selectPage:function( event ){
            //获取对应的 信息
            var target = $( event.target );
            var currentPage = target.find(".currentPage");

            var pageName = currentPage.text().trim();

            this.currentPage = pageName;

        },initIndex:function(){
            //初始化 相对应的 时
            
        }
    }
});