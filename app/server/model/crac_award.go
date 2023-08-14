package model

import (
	"errors"

	"gorm.io/gorm"
)

const CracAwardTableName = "crac_award"

type CracAward struct {
	ID          int    `gorm:"primary_key;column:id"`
	Call        string `gorm:"column:callsign"`
	Continent   string `gorm:"column:continent"`
	AwardType   int8   `gorm:"column:type"`
	Combination int64  `gorm:"column:combination"`
	BncraNum    int64  `gorm:"column:bncra_num"`
	Year        int64  `gorm:"column:year"`
	ImgUrl      string `gorm:"column:img_url"`
	CreatedTime int64  `gorm:"column:created_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
	Status      int8   `gorm:"column:status"`
	Remark      string `gorm:"column:remark"`
}

// CracAwardCountAll 返回总数信息
func (c *Dao) CracAwardCountAll(year string) (int64, error) {
	var (
		count int64
	)
	err := c.DB.Table(CracAwardTableName).Where("year = ? and status = ?", year, 1).Count(&count).Error
	return count, err
}

// CracAwardQueryByCall 根据呼号查询奖状信息
func (c *Dao) CracAwardQueryByCall(callsign, year string) (CracAward, error) {
	var t CracAward
	err := c.DB.Table(CracAwardTableName).
		Where("callsign = ? and year = ? and status = ?", callsign, year, 1).First(&t).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return t, err
	}
	return t, nil
}
