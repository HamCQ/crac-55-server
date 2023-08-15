package service

import (
	"crac55/app/server/entities"
	"crac55/common/clog"
	"crac55/common/tools"
	"fmt"
	"strings"
)

// Award 奖状查询
func (s *Service) Award(callsign, year string) (entities.AwardDetail, error) {
	var (
		res entities.AwardDetail
	)
	info, err := s.Dao.CracAwardQueryByCall(callsign, year)
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	res = entities.AwardDetail{
		AwardType:      info.AwardType,
		AwardString:    tools.AwardTypeString(info.AwardType),
		Continent:      info.Continent,
		BncraNum:       info.BncraNum,
		Callsign:       strings.ToUpper(callsign),
		ImageUrl:       fmt.Sprintf("%s?x-oss-process=style/min_img", tools.UriFilterExcludeQueryString(info.ImgUrl)),
		ImageUrlOrigin: info.ImgUrl,
		Combination:    info.Combination,
	}
	return res, nil
}
