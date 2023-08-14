package service

import (
	"crac55/app/server/entities"
	"crac55/app/server/model"
	"crac55/common/clog"
	"crac55/common/tools"
)

// Search 首页搜索
func (s *Service) Search(callsign, year string,
) ([]entities.SearchStationDetail, entities.SearchStationDetail) {
	var bxcra = entities.SearchBxCraDetailMap()
	var bncra = entities.SearchBnCraMap()
	info, err := s.Dao.CracLogSearch(callsign, year)
	if err != nil {
		clog.Log().Errorln(err)
		return bxcra, bncra
	}
	for _, v := range info {
		s.dealBxCra(v, &bxcra)
		s.dealBncra(v, &bncra)
	}
	return bxcra, bncra
}

// dealBxCra 归类数据 b0-9cra
func (s *Service) dealBxCra(qso model.CracLog, t *[]entities.SearchStationDetail) {
	mode := tools.CateMode(qso.Mode)
	for i := range *t {
		if qso.Station == (*t)[i].CallsignStation {
			if mode == entities.ModeCW {
				s.dealCWDetail(qso, &(*t)[i])
			}
			if mode == entities.ModePhone {
				s.dealPhoneDetail(qso, &(*t)[i])
			}
			if mode == entities.ModeDigi {
				s.dealDigiDetail(qso, &(*t)[i])
			}
		}
	}
}

// dealBncra 归类数据 bncra
func (s *Service) dealBncra(qso model.CracLog, t *entities.SearchStationDetail) {
	mode := tools.CateMode(qso.Mode)
	if mode == entities.ModeCW {
		s.dealCWDetail(qso, t)
	}
	if mode == entities.ModePhone {
		s.dealPhoneDetail(qso, t)
	}
	if mode == entities.ModeDigi {
		s.dealDigiDetail(qso, t)
	}
}

// dealCWDetail 处理预览数据-cw
func (s *Service) dealCWDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.CW.M160 = append(v.CW.M160, detail)
		s.bxcraScore(v.CW.M160, v)
	case entities.BandM80:
		v.CW.M80 = append(v.CW.M80, detail)
		s.bxcraScore(v.CW.M80, v)
	case entities.BandM40:
		v.CW.M40 = append(v.CW.M40, detail)
		s.bxcraScore(v.CW.M40, v)
	case entities.BandM30:
		v.CW.M30 = append(v.CW.M30, detail)
		s.bxcraScore(v.CW.M30, v)
	case entities.BandM20:
		v.CW.M20 = append(v.CW.M20, detail)
		s.bxcraScore(v.CW.M20, v)
	case entities.BandM17:
		v.CW.M17 = append(v.CW.M17, detail)
		s.bxcraScore(v.CW.M17, v)
	case entities.BandM15:
		v.CW.M15 = append(v.CW.M15, detail)
		s.bxcraScore(v.CW.M15, v)
	case entities.BandM12:
		v.CW.M12 = append(v.CW.M12, detail)
		s.bxcraScore(v.CW.M12, v)
	case entities.BandM10:
		v.CW.M10 = append(v.CW.M10, detail)
		s.bxcraScore(v.CW.M10, v)
	case entities.BandM6:
		v.CW.M6 = append(v.CW.M6, detail)
		s.bxcraScore(v.CW.M6, v)
	}
}

// dealPhoneDetail 处理预览数据-phone
func (s *Service) dealPhoneDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.Phone.M160 = append(v.Phone.M160, detail)
		s.bxcraScore(v.Phone.M160, v)
	case entities.BandM80:
		v.Phone.M80 = append(v.Phone.M80, detail)
		s.bxcraScore(v.Phone.M80, v)
	case entities.BandM40:
		v.Phone.M40 = append(v.Phone.M40, detail)
		s.bxcraScore(v.Phone.M40, v)
	case entities.BandM30:
		v.Phone.M30 = append(v.Phone.M30, detail)
		s.bxcraScore(v.Phone.M30, v)
	case entities.BandM20:
		v.Phone.M20 = append(v.Phone.M20, detail)
		s.bxcraScore(v.Phone.M20, v)
	case entities.BandM17:
		v.Phone.M17 = append(v.Phone.M17, detail)
		s.bxcraScore(v.Phone.M17, v)
	case entities.BandM15:
		v.Phone.M15 = append(v.Phone.M15, detail)
		s.bxcraScore(v.Phone.M15, v)
	case entities.BandM12:
		v.Phone.M12 = append(v.Phone.M12, detail)
		s.bxcraScore(v.Phone.M12, v)
	case entities.BandM10:
		v.Phone.M10 = append(v.Phone.M10, detail)
		s.bxcraScore(v.Phone.M10, v)
	case entities.BandM6:
		v.Phone.M6 = append(v.Phone.M6, detail)
		s.bxcraScore(v.Phone.M6, v)
	}
}

// dealDigiDetail 处理预览数据-digi
func (s *Service) dealDigiDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.Digi.M160 = append(v.Digi.M160, detail)
		s.bxcraScore(v.Digi.M160, v)
	case entities.BandM80:
		v.Digi.M80 = append(v.Digi.M80, detail)
		s.bxcraScore(v.Digi.M80, v)
	case entities.BandM40:
		v.Digi.M40 = append(v.Digi.M40, detail)
		s.bxcraScore(v.Digi.M40, v)
	case entities.BandM30:
		v.Digi.M30 = append(v.Digi.M30, detail)
		s.bxcraScore(v.Digi.M30, v)
	case entities.BandM20:
		v.Digi.M20 = append(v.Digi.M20, detail)
		s.bxcraScore(v.Digi.M20, v)
	case entities.BandM17:
		v.Digi.M17 = append(v.Digi.M17, detail)
		s.bxcraScore(v.Digi.M17, v)
	case entities.BandM15:
		v.Digi.M15 = append(v.Digi.M15, detail)
		s.bxcraScore(v.Digi.M15, v)
	case entities.BandM12:
		v.Digi.M12 = append(v.Digi.M12, detail)
		s.bxcraScore(v.Digi.M12, v)
	case entities.BandM10:
		v.Digi.M10 = append(v.Digi.M10, detail)
		s.bxcraScore(v.Digi.M10, v)
	case entities.BandM6:
		v.Digi.M6 = append(v.Digi.M6, detail)
		s.bxcraScore(v.Digi.M6, v)
	}
}

// bxcraScore b0-9计分
func (s *Service) bxcraScore(d []entities.SearchQsoDetail, v *entities.SearchStationDetail) {
	if len(d) == 1 {
		v.Score += 1
	}
}
