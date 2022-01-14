package items

import "time"

type SubjectType int

const (
	SubjectTypeBook  = SubjectType(1)
	SubjectTypeAnime = SubjectType(2)
	SubjectTypeMusic = SubjectType(3)
	SubjectTypeGame  = SubjectType(4)
	SubjectTypeReal  = SubjectType(5)
)

func (st SubjectType) IsSupported() bool {
	return st == SubjectTypeBook || st == SubjectTypeAnime || st == SubjectTypeMusic || st == SubjectTypeGame || st == SubjectTypeReal
}

func (st SubjectType) ToString() string {
	switch st {
	case SubjectTypeBook:
		return "1"
	case SubjectTypeAnime:
		return "2"
	case SubjectTypeMusic:
		return "3"
	case SubjectTypeGame:
		return "4"
	case SubjectTypeReal:
		return "5"
	default:
		return ""
	}
}

//UserCollection 查询用户收藏时返回的subject 结构
type UserCollection struct {
	SubjectId   int         `json:"subject_id,omitempty"`
	SubjectType int         `json:"subject_type,omitempty"`
	Rate        int         `json:"rate,omitempty"`
	Type        SubjectType `json:"type,omitempty"`
	Tags        []string    `json:"tags,omitempty"`
	EpStatus    int         `json:"ep_status,omitempty"`
	VolStatus   int         `json:"vol_status,omitempty"`
	UpdatedAt   time.Time   `json:"updated_at"`
	Private     bool        `json:"private,omitempty"`
}
