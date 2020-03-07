package base

type DashboardBase struct {
	LoveCount int64 `orm:"column(love_count)" json:"loveCount"`
	ForwardCount int64 `orm:"column(forward_count)" json:"forwardCount"`
	CollectionCount int64 `orm:"column(collection_count)" json:"collectionCount"`
	BrowseCount int64 `orm:"column(browse_count)" json:"browseCount"`
	CommentCount int64 `orm:"column(comment_count)" json:"commentCount"`
}

func ( self *DashboardBase ) Add( base *DashboardBase ){
	self.LoveCount += base.LoveCount
	self.ForwardCount += base.ForwardCount
	self.CollectionCount += base.CollectionCount
	self.BrowseCount += base.BrowseCount
	self.CommentCount += base.CommentCount
}

func ( self *DashboardBase ) New() {
	self.LoveCount = 0
	self.ForwardCount = 0
	self.CollectionCount = 0
	self.BrowseCount = 0
	self.CommentCount = 0
}

func NewDashboardBase() *DashboardBase{
	var dashboard = &DashboardBase{}
	dashboard.New()
	return dashboard
}
