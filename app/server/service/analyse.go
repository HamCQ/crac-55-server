package service

import (
	"crac55/app/server/entities"
	"crac55/common/clog"
)

// AnalyseNum 返回首页统计总数
func (s *Service) AnalyseNum(year string) (entities.AnalyseTotal, error) {
	var (
		total entities.AnalyseTotal
		err   error
	)
	//todo: 一次raw执行返回
	total.LogNum, err = s.Dao.CracLogCountAll(year)
	if err != nil {
		clog.Log().Errorln(err)
		return total, err
	}
	total.AwardNum, err = s.Dao.CracAwardCountAll(year)
	if err != nil {
		clog.Log().Errorln(err)
		return total, err
	}
	total.SingleCallNum, err = s.Dao.CracLogCountGroupByCall(year)
	if err != nil {
		clog.Log().Errorln(err)
		return total, err
	}
	return total, nil
}
