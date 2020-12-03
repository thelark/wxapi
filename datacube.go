package wxapi

import "reflect"

type dataCube struct {
	AccessToken string
}

func (t *dataCube) set(k, v string) {
	_value := reflect.ValueOf(t).Elem()
	_type := reflect.TypeOf(t).Elem()
	if _, ok := _type.FieldByName(k); ok {
		_field := _value.FieldByName(k)
		_field.SetString(v)
	}
}

// 方法 --------------------------------------------------------------------

// GetWeAnalysisAppIdDailyRetainInfo 获取用户访问小程序日留存
// POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyretaininfo?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdDailyRetainInfo() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdDailySummaryTrend 获取用户访问小程序数据概况
// POST https://api.weixin.qq.com/datacube/getweanalysisappiddailysummarytrend?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdDailySummaryTrend() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdDailyVisitTrend 获取用户访问小程序数据日趋势
// POST https://api.weixin.qq.com/datacube/getweanalysisappiddailyvisittrend?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdDailyVisitTrend() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdMonthlyRetainInfo 获取用户访问小程序月留存
// POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyretaininfo?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdMonthlyRetainInfo() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdMonthlyVisitTrend 获取用户访问小程序数据月趋势
// POST https://api.weixin.qq.com/datacube/getweanalysisappidmonthlyvisittrend?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdMonthlyVisitTrend() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdUserPortrait 获取小程序新增或活跃用户的画像分布数据。时间范围支持昨天、最近7天、最近30天。其中，新增用户数为时间范围内首次访问小程序的去重用户数，活跃用户数为时间范围内访问过小程序的去重用户数。
// POST https://api.weixin.qq.com/datacube/getweanalysisappiduserportrait?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdUserPortrait() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdVisitDistribution 获取用户小程序访问分布数据
// POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitdistribution?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdVisitDistribution() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdVisitPage 访问页面。目前只提供按 page_visit_pv 排序的 top200。
// POST https://api.weixin.qq.com/datacube/getweanalysisappidvisitpage?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdVisitPage() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdWeeklyRetainInfo 获取用户访问小程序周留存
// POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyretaininfo?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdWeeklyRetainInfo() {

}

//      --------------------------------------------------------------------

// GetWeAnalysisAppIdWeeklyVisitTrend 获取用户访问小程序数据周趋势
// POST https://api.weixin.qq.com/datacube/getweanalysisappidweeklyvisittrend?access_token=ACCESS_TOKEN
func (t *dataCube) GetWeAnalysisAppIdWeeklyVisitTrend() {

}
