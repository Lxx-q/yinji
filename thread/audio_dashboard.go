package thread

import (
	"yinji/models/bean"
)

/**
	算法依然不够成熟 ， 因此暂时只能先跳过这样的线程算法 ， 使用原始算法基础
*/

type AudioDashboardThread struct {
	dashboardMap map[int64] *bean.AudioDashboard
}

func ( self AudioDashboardThread) Start(){

}
