package model

import "context"

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

// CracLogCountAll 返回总数信息
func (c *Dao) CracLogCountAll(ctx context.Context) (int64, error) {
	var (
		count int64
	)
	err := c.DB.WithContext(ctx).Table(CracLogTableName).Where("status = ?", 1).Count(&count).Error
	return count, err
}
