package items

import "time"

//UserCollection 查询用户收藏时返回的subject 结构
type UserCollection struct {
	SubjectId   int                `json:"subject_id"`
	SubjectType SubjectType        `json:"subject_type"`
	Rate        int                `json:"rate,omitempty"`
	Type        CollectionStatusId `json:"type,omitempty"`
	Tags        []string           `json:"tags,omitempty"`
	EpStatus    int                `json:"ep_status,omitempty"`
	VolStatus   int                `json:"vol_status,omitempty"`
	UpdatedAt   time.Time          `json:"updated_at"`
	Private     bool               `json:"private"` //是否为自由对象
}
