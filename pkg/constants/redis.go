package constants

import "time"

const (
	RedisCampaignDetailKeyPrefix = "campaign:"
	RedisCampaignDetailKeyExpiry = 24 * time.Hour

	RedisQualifiedCampaignsKeyPrefix = "campaigns:qualified:"
	RedisQualifiedCampaignsKeyExpiry = time.Hour

	RedisCampaignsWithNoRulesKeyPrefix = "campaigns:no_rules:"
	RedisCampaignsWithNoRulesKeyExpiry = time.Hour
)
