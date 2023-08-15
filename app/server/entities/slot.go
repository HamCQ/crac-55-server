package entities

import "fmt"

// SlotMap 组装初始数据
func SlotMap() []SlotStation {
	var t []SlotStation
	for i := 0; i <= 9; i++ {
		t = append(t, SlotStation{
			CallsignStation: fmt.Sprintf("B%dCRA", i),
		})
	}
	return t
}

// SlotStation slot 信息
type SlotStation struct {
	CallsignStation string `json:"callsign_station"`
	Score           int    `json:"score"`
	SlotMode
}

// SlotMode slot mode 归类
type SlotMode struct {
	Phone SlotBand `json:"phone"`
	CW    SlotBand `json:"cw"`
	Digi  SlotBand `json:"digi"`
}

// SlotBand slot band归类
type SlotBand struct {
	M160 SlotDetail `json:"160M"` //160M
	M80  SlotDetail `json:"80M"`
	M40  SlotDetail `json:"40M"`
	M30  SlotDetail `json:"30M"`
	M20  SlotDetail `json:"20M"`
	M17  SlotDetail `json:"17M"`
	M15  SlotDetail `json:"15M"`
	M12  SlotDetail `json:"12M"`
	M10  SlotDetail `json:"10M"`
	M6   SlotDetail `json:"6M"`
}

// SlotDetail slot单元格具体信息
type SlotDetail struct {
	Callsign string `json:"callsign"`
	Time     string `json:"time"`
}
