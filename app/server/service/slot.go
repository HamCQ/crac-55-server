package service

import (
	"crac55/app/server/entities"
	"crac55/app/server/model"
	"crac55/common/clog"
	"crac55/common/tools"
	"strconv"
	"time"
)

// Slot slot信息
func (s *Service) Slot() ([]entities.SlotStation, error) {
	var (
		res []entities.SlotStation
	)
	slot := entities.SlotMap()
	info, err := s.Dao.CracSlotAll(strconv.Itoa(time.Now().Year()))
	if err != nil {
		clog.Log().Errorln(err)
		return res, err
	}
	for _, v := range info {
		s.dealSlot(v, &slot)
	}

	return res, nil
}

// dealSlot 归类处理slot
func (s *Service) dealSlot(info model.CracSlot, t *[]entities.SlotStation) {
	mode := tools.CateMode(info.Mode)
	for i := range *t {
		if info.BnCra == (*t)[i].CallsignStation {
			if mode == entities.ModeCW {
				s.dealSlotCw(info, &(*t)[i])
			}
			if mode == entities.ModePhone {
				s.dealSlotPhone(info, &(*t)[i])
			}
			if mode == entities.ModeDigi {
				s.dealSlotDigi(info, &(*t)[i])
			}
		}
	}
}

// dealSlotCw 归类处理slot cw
func (s *Service) dealSlotCw(info model.CracSlot, v *entities.SlotStation) {
	detail := entities.SlotDetail{
		Callsign: info.Callsign,
		Time:     tools.RencentTime(info.CreatedTime),
	}
	switch info.Band {
	case entities.BandM160:
		v.CW.M160 = detail
		s.slotScore(v)
	case entities.BandM80:
		v.CW.M80 = detail
		s.slotScore(v)
	case entities.BandM40:
		v.CW.M40 = detail
		s.slotScore(v)
	case entities.BandM30:
		v.CW.M30 = detail
		s.slotScore(v)
	case entities.BandM20:
		v.CW.M20 = detail
		s.slotScore(v)
	case entities.BandM17:
		v.CW.M17 = detail
		s.slotScore(v)
	case entities.BandM15:
		v.CW.M15 = detail
		s.slotScore(v)
	case entities.BandM12:
		v.CW.M12 = detail
		s.slotScore(v)
	case entities.BandM10:
		v.CW.M10 = detail
		s.slotScore(v)
	case entities.BandM6:
		v.CW.M6 = detail
		s.slotScore(v)
	}
}

// dealSlotPhone 归类处理slot phone
func (s *Service) dealSlotPhone(info model.CracSlot, v *entities.SlotStation) {
	detail := entities.SlotDetail{
		Callsign: info.Callsign,
		Time:     tools.RencentTime(info.CreatedTime),
	}
	switch info.Band {
	case entities.BandM160:
		v.Phone.M160 = detail
		s.slotScore(v)
	case entities.BandM80:
		v.Phone.M80 = detail
		s.slotScore(v)
	case entities.BandM40:
		v.Phone.M40 = detail
		s.slotScore(v)
	case entities.BandM30:
		v.Phone.M30 = detail
		s.slotScore(v)
	case entities.BandM20:
		v.Phone.M20 = detail
		s.slotScore(v)
	case entities.BandM17:
		v.Phone.M17 = detail
		s.slotScore(v)
	case entities.BandM15:
		v.Phone.M15 = detail
		s.slotScore(v)
	case entities.BandM12:
		v.Phone.M12 = detail
		s.slotScore(v)
	case entities.BandM10:
		v.Phone.M10 = detail
		s.slotScore(v)
	case entities.BandM6:
		v.Phone.M6 = detail
		s.slotScore(v)
	}
}

// dealSlotDigi 归类处理digi
func (s *Service) dealSlotDigi(info model.CracSlot, v *entities.SlotStation) {
	detail := entities.SlotDetail{
		Callsign: info.Callsign,
		Time:     tools.RencentTime(info.CreatedTime),
	}
	switch info.Band {
	case entities.BandM160:
		v.Digi.M160 = detail
		s.slotScore(v)
	case entities.BandM80:
		v.Digi.M80 = detail
		s.slotScore(v)
	case entities.BandM40:
		v.Digi.M40 = detail
		s.slotScore(v)
	case entities.BandM30:
		v.Digi.M30 = detail
		s.slotScore(v)
	case entities.BandM20:
		v.Digi.M20 = detail
		s.slotScore(v)
	case entities.BandM17:
		v.Digi.M17 = detail
		s.slotScore(v)
	case entities.BandM15:
		v.Digi.M15 = detail
		s.slotScore(v)
	case entities.BandM12:
		v.Digi.M12 = detail
		s.slotScore(v)
	case entities.BandM10:
		v.Digi.M10 = detail
		s.slotScore(v)
	case entities.BandM6:
		v.Digi.M6 = detail
		s.slotScore(v)
	}
}

// slotScore 累计slot单元格
func (s *Service) slotScore(v *entities.SlotStation) {
	v.Score += 1
}
