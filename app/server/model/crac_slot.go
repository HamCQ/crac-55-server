package model

import (
	"errors"

	"gorm.io/gorm"
)

const CracSlotTableName = "crac_slot"

type CracSlot struct {
	ID          int64  `gorm:"primary_key;column:id"`
	UserID      int64  `gorm:"column:user_id"`
	BnCra       string `gorm:"column:bncra"`
	Callsign    string `gorm:"column:callsign"`
	Year        int    `gorm:"column:year"`
	Mode        string `gorm:"column:mode"`
	Band        string `gorm:"column:band"`
	Status      int8   `gorm:"column:status"`
	CreatedTime int64  `gorm:"column:created_time"`
	UpdateTime  int64  `gorm:"column:update_time"`
}

// CracSlotAll 返回slot信息
func (c *Dao) CracSlotAll(year string) ([]CracSlot, error) {
	var (
		t []CracSlot
	)
	err := c.DB.Table(CracSlotTableName).Where("year = ? and status = ?", year, 1).Find(&t).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return t, err
	}
	return t, nil
}
