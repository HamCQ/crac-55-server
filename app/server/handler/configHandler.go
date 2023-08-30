package handler

import (
	"crac55/app/server/code"
	"crac55/app/server/entities"
	"crac55/common/conf"
	"crac55/common/tools"
	"net/http"
)

// Config 获取对应配置信息
func (c *CracHandler) Config(w http.ResponseWriter, r *http.Request) {
	var (
		res  entities.Config
		list []int
	)
	conf := conf.AppConf.SiteConfig
	values := r.URL.Query()
	year := values.Get("year")
	res.CurrentActiveYear, list = tools.MaxYear(conf)
	info, ok := conf[year]
	if !ok {
		code.SuccussJSON(w, res)
		return
	}
	res.Cover = info.Cover
	res.Title = info.Title
	res.SubTitle = info.SubTitle
	res.TitleEn = info.TitleEn
	res.SubTitleEn = info.SubTitleEn
	res.CracDesc = info.CracDesc
	res.YearList = list
	code.SuccussJSON(w, res)
}
