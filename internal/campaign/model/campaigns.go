package campaignModel

import "time"

type Campaign struct {
	ID        int64      `gorm:"primaryKey;autoIncrement"`
	Code      string     `gorm:"column:code;type:varchar(20);not null;uniqueIndex"`
	Name      string     `gorm:"column:name;type:varchar(255)"`
	Image     string     `gorm:"column:image;type:varchar(1024)"`
	CTA       string     `gorm:"column:cta;type:varchar(255)"`
	Status    string     `gorm:"column:status;type:varchar(10);check:status IN ('ACTIVE','INACTIVE')"`
	CreatedAt time.Time  `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt time.Time  `gorm:"column:updated_at;type:timestamptz;default:now()"`
	Deleted   bool       `gorm:"column:deleted"`
	DeletedAt *time.Time `gorm:"column:deleted_at"`
}

func (Campaign) TableName() string {
	return "campaigns"
}
