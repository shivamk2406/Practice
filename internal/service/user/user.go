package user

import (
	"time"

	"gorm.io/gorm"
)

type Metadata struct {
	CreateTime time.Time `gorm:"column:create_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"create_time"`
	UpdateTime time.Time `gorm:"column:update_time;type:datetime;default:CURRENT_TIMESTAMP;" json:"update_time"`
	Creator    string    `gorm:"column:creator;type:varchar;size:31;" json:"creator"`
	Updater    string    `gorm:"column:updater;type:varchar;size:31;" json:"updater"`
	Active     bool      `gorm:"column:active;type:tinyint;default:1;" json:"active"`
}

// Model struct is a row record of the customer table in the summarise database.
type Model struct {
	Metadata     Metadata `gorm:"embedded"`
	ID           string   `gorm:"primary_key;column:id;type:varchar;size:36;" json:"id"`
	Name         string   `gorm:"column:name;type:varchar;size:1000;" json:"name"`
	Subscription string   `gorm:"column:subscription;type:type:varchar;size:1000;" json:"subscription"`
}

func (a *Model) TableName() string {
	return "test_practice"
}

func (a *Model) BeforeCreate(tx *gorm.DB) (err error) {
	a.Metadata.Creator = "system"
	a.Metadata.Updater = "system"
	a.Metadata.CreateTime = time.Now().UTC()
	a.Metadata.UpdateTime = time.Now().UTC()
	return nil
}

func (a *Model) BeforeUpdate(tx *gorm.DB) (err error) {
	a.Metadata.UpdateTime = time.Now().UTC()
	return nil
}

func (a *Model) whereClause() []func(*gorm.DB) *gorm.DB {
	var clauses []func(*gorm.DB) *gorm.DB

	if a.ID != "" {
		clauses = append(clauses, func(db *gorm.DB) *gorm.DB {
			return db.Where("id = ?", a.ID)
		})
	}

	if a.Subscription != "" {
		clauses = append(clauses, func(db *gorm.DB) *gorm.DB {
			return db.Where("subscription = ?", a.Subscription)
		})
	}
	return clauses
}
