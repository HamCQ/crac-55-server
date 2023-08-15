package model

import (
	"errors"

	"gorm.io/gorm"
)

const CracLogTableName = "crac_log"

type CracLog struct {
	ID           int64  `gorm:"primary_key;column:id"`
	UserID       int64  `gorm:"column:user_id"`
	Year         int    `gorm:"column:year"`
	Continent    string `gorm:"column:continent"`
	Call         string `gorm:"column:call_obj"`
	CNRegionCode int    `gorm:"column:cn_region_code"`
	CNByCode     int8   `gorm:"column:cn_by_code"`
	Mode         string `gorm:"column:mode"`
	Band         string `gorm:"column:band"`
	QsoDate      string `gorm:"column:qso_date"`
	TimeOn       string `gorm:"column:time_on"`
	TimeOff      string `gorm:"column:time_off"`
	QsoTimestamp int    `gorm:"column:user_id"`
	Frequency    string `gorm:"column:frequency"`
	Station      string `gorm:"column:station_callsign"`
	Oprator      string `gorm:"column:operator"`
	LogFile      string `gorm:"column:log_file"`
	Status       int8   `gorm:"column:status"`
	CreatedTime  int64  `gorm:"column:created_time"`
	UpdateTime   int64  `gorm:"column:update_time"`
}

type CracLogProvince struct {
	CnRegionCode int `gorm:"column:cn_region_code"`
	Num          int `gorm:"column:num"`
}

// CracLogCountAll 返回总数信息
func (c *Dao) CracLogCountAll(year string) (int64, error) {
	var (
		count int64
	)
	err := c.DB.Table(CracLogTableName).Where("year = ? and status = ?", year, 1).Count(&count).Error
	return count, err
}

// CracLogCountGroupByCall 独立呼号个数
func (c *Dao) CracLogCountGroupByCall(year string) (int64, error) {
	var (
		count int64
	)
	err := c.DB.Table(CracLogTableName).Where("year = ? and status = ?", year, 1).Group("call_obj").Count(&count).Error
	return count, err
}

// CracLogSearch 首页搜索
func (c *Dao) CracLogSearch(callsign, year string) ([]CracLog, error) {
	var t []CracLog
	err := c.DB.Table(CracLogTableName).Select("call_obj", "mode", "band", "qso_date", "frequency", "station_callsign", "operator").
		Where("call_obj = ? and year = ? and status = ? ", callsign, year, 1).Find(&t).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return t, err
	}
	return t, nil
}

// CracLogBy09 BY电台通联数量统计(0-9区)
func (c *Dao) CracLogBy09(byCode int, mode []string, year string) (int64, error) {
	var (
		count int64
	)
	err := c.DB.Table(CracLogTableName).
		Where("cn_by_code = ? and mode in (?) and continent = ? and year = ? and status = ?",
			byCode, mode, "AS", year, 1).Count(&count).Error
	return count, err
}

// CracLogProvince 省份通联数量统计
func (c *Dao) CracLogProvince(year string) ([]CracLogProvince, error) {
	var res []CracLogProvince
	err := c.DB.Table(CracLogTableName).Select("cn_region_code,count(*) as num").
		Where("continent = ? and year = ? and status = ? and cn_region_code != ?", "AS", year, 1, 0).
		Group("cn_region_code").Find(&res).Error
	return res, err
}

// CracLogBnCra BnCRA 电台通联统计
func (c *Dao) CracLogBnCra(station, year string, mode []string) (int64, error) {
	var (
		count int64
	)
	err := c.DB.Table(CracLogTableName).
		Where("station_callsign = ? and mode in (?) and year = ? and status = ?",
			station, mode, year, 1).Count(&count).Error
	return count, err
}
