package models

import "time"

type Community struct {
	ID   int64  `json:"id" db:"community_id"`
	Name string `json:"name" db:"community_name"`
}

type CommunityDetail struct {
	ID           int64     `json:"id" db:"community_id"`
	Name         string    `json:"name" db:"community_name"`
	Introduction string    `json:"introduction" db:"introduction"`
	CreateTime   time.Time `json:"create_time" db:"create_time"` //time.Time 需要连接的时候增加parseTime=true
	UpdateTime   time.Time `json:"update_time" db:"update_time"`
}
