package service

import (
	"crac55/app/server/entities"
	"crac55/app/server/model"
	"crac55/common/clog"
	"crac55/common/tools"
)

// Search 首页搜索
func (s *Service) Search(callsign, year string) []entities.SearchStationDetail {
	abstract := entities.SearchBxCraDetailMap()
	// point := entities.SearchBxCraPointMap()

	info, err := s.Dao.CracLogSearch(callsign, year)
	if err != nil {
		clog.Log().Errorln(err)
		return abstract
	}
	for _, v := range info {
		s.dealBxCra(v, &abstract)
	}
	return abstract
}

// dealBxCra 归类数据 b0-9cra
func (s *Service) dealBxCra(qso model.CracLog, t *[]entities.SearchStationDetail) {
	for _, v := range *t {
		mode := tools.CateMode(qso.Mode)
		if qso.Station == v.CallsignStation {
			if mode == entities.ModeCW {
				s.dealBxcraCWDetail(qso, &v)
			}
			if mode == entities.ModePhone {
				s.dealBxcraPhoneDetail(qso, &v)
			}
			if mode == entities.ModeDigi {
				s.dealBxcraDigiDetail(qso, &v)
			}
		}
	}
}

// dealBxcraCWDetail 处理预览数据-cw
func (s *Service) dealBxcraCWDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Call:     qso.Call,
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
		QsoDate:  qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.CW.M160 = append(v.CW.M160, detail)
	case entities.BandM80:
		v.CW.M80 = append(v.CW.M80, detail)
	case entities.BandM40:
		v.CW.M40 = append(v.CW.M40, detail)
	case entities.BandM30:
		v.CW.M30 = append(v.CW.M30, detail)
	case entities.BandM20:
		v.CW.M20 = append(v.CW.M20, detail)
	case entities.BandM17:
		v.CW.M17 = append(v.CW.M17, detail)
	case entities.BandM15:
		v.CW.M15 = append(v.CW.M15, detail)
	case entities.BandM12:
		v.CW.M12 = append(v.CW.M12, detail)
	case entities.BandM10:
		v.CW.M10 = append(v.CW.M10, detail)
	case entities.BandM6:
		v.CW.M6 = append(v.CW.M6, detail)
	}
}

// dealBxcraPhoneDetail 处理预览数据-phone
func (s *Service) dealBxcraPhoneDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Call:     qso.Call,
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
		QsoDate:  qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.Phone.M160 = append(v.Phone.M160, detail)
	case entities.BandM80:
		v.Phone.M80 = append(v.Phone.M80, detail)
	case entities.BandM40:
		v.Phone.M40 = append(v.Phone.M40, detail)
	case entities.BandM30:
		v.Phone.M30 = append(v.Phone.M30, detail)
	case entities.BandM20:
		v.Phone.M20 = append(v.Phone.M20, detail)
	case entities.BandM17:
		v.Phone.M17 = append(v.Phone.M17, detail)
	case entities.BandM15:
		v.Phone.M15 = append(v.Phone.M15, detail)
	case entities.BandM12:
		v.Phone.M12 = append(v.Phone.M12, detail)
	case entities.BandM10:
		v.Phone.M10 = append(v.Phone.M10, detail)
	case entities.BandM6:
		v.Phone.M6 = append(v.Phone.M6, detail)
	}
}

// dealBxcraDigiDetail 处理预览数据-digi
func (s *Service) dealBxcraDigiDetail(qso model.CracLog, v *entities.SearchStationDetail) {
	detail := entities.SearchQsoDetail{
		Call:     qso.Call,
		Oprator:  qso.Oprator,
		Frequecy: qso.Frequency,
		QsoDate:  qso.Frequency,
	}
	switch qso.Band {
	case entities.BandM160:
		v.Digi.M160 = append(v.Digi.M160, detail)
	case entities.BandM80:
		v.Digi.M80 = append(v.Digi.M80, detail)
	case entities.BandM40:
		v.Digi.M40 = append(v.Digi.M40, detail)
	case entities.BandM30:
		v.Digi.M30 = append(v.Digi.M30, detail)
	case entities.BandM20:
		v.Digi.M20 = append(v.Digi.M20, detail)
	case entities.BandM17:
		v.Digi.M17 = append(v.Digi.M17, detail)
	case entities.BandM15:
		v.Digi.M15 = append(v.Digi.M15, detail)
	case entities.BandM12:
		v.Digi.M12 = append(v.Digi.M12, detail)
	case entities.BandM10:
		v.Digi.M10 = append(v.Digi.M10, detail)
	case entities.BandM6:
		v.Digi.M6 = append(v.Digi.M6, detail)
	}
}
