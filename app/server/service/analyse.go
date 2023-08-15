package service

import (
	"crac55/app/server/entities"
	"crac55/common/cache"
	"crac55/common/clog"
	"time"
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

// AnalyseRankTop 首页排名-排序前五
func (s *Service) AnalyseRankTop(year string) (entities.AnalyseRankTopRes, error) {
	var (
		res entities.AnalyseRankTopRes
		err error
	)
	res.CnDiffCra, err = cache.RankTop5(cache.KeyCnDiffCra(year))
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	res.CnCra, err = cache.RankTop5(cache.KeyCnCra(year))
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	res.GlobleDiffCra, err = cache.RankTop5(cache.KeyGlobleDiffCra(year))
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	res.GlobleCra, err = cache.RankTop5(cache.KeyGlobleCra(year))
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	updateTime, err := cache.RankUpdateTime(year)
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	res.UpdateTime = time.Unix(int64(updateTime), 0).Format("2006-01-02 15:04:05")
	if updateTime == 0 {
		res.UpdateTime = "Unknown"
	}
	return res, nil
}

// AnalyseRankAll 首页排名-所有
func (s *Service) AnalyseRankAll(year string, rankType, page int) ([]entities.AnalyseRankContent, error) {
	var (
		res []entities.AnalyseRankContent
		err error
	)
	switch rankType {
	case entities.CnDiffCra:
		res, err = cache.RankAll(cache.KeyCnDiffCra(year), page)
	case entities.CnCra:
		res, err = cache.RankAll(cache.KeyCnCra(year), page)
	case entities.GlobleDiffCra:
		res, err = cache.RankAll(cache.KeyGlobleDiffCra(year), page)
	case entities.GlobleCra:
		res, err = cache.RankAll(cache.KeyGlobleCra(year), page)
	}
	return res, err
}
