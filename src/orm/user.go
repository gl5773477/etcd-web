package orm

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

type TUser struct {
	ID   string
	Name string
}

func (c *TUser) TableName() string {
	return "user"
}

func (c *TUser) LoadAll(db *gorm.DB) ([]TUser, error) {
	var dts []TUser
	sql := fmt.Sprintf("select * from %s where trade_order = ? for update", c.TableName())
	if err := db.Raw(sql).Where("name = ?", c.Name).Scan(&dts).Error; err != nil {
		return nil, err
	}
	return dts, nil
}
