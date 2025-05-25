package constants

import "time"

const (
	RedisCampaignDetailKeyPrefix = "campaign:"

	RedisCampaignDetailKeyExpiry = 24 * time.Hour
)
