package matchEngineModel

import "time"

type TargetRule struct {
	ID         int64      `gorm:"column:id;primaryKey;autoIncrement"`
	Code       string     `gorm:"column:code;type:varchar(25);not null;uniqueIndex"`
	CampaignID int64      `gorm:"column:campaign_id;index;not null"`
	Dimension  string     `gorm:"column:dimension;type:varchar;check:dimension IN ('app','country','os')"`
	RuleType   string     `gorm:"column:rule_type;type:varchar;check:rule_type IN ('INCLUDE','EXCLUDE')"`
	Values     []string   `gorm:"column:values;type:text[]"`
	CreatedAt  time.Time  `gorm:"column:created_at;type:timestamptz;default:now()"`
	UpdatedAt  time.Time  `gorm:"column:updated_at;type:timestamptz;default:now()"`
	Deleted    bool       `gorm:"column:deleted"`
	DeletedAt  *time.Time `gorm:"column:deleted_at"`
}

func (TargetRule) TableName() string {
	return "target_rules"
}
