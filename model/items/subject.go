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

type CollectionType int

const (
	CollectionTypeWantWatch  = CollectionType(1) //想看
	CollectionTypeHasWatched = CollectionType(2) //看过
	CollectionTypeWatching   = CollectionType(3) //在看
	CollectionTypePutAside   = CollectionType(4) //搁置
	CollectionTypeForgive    = CollectionType(5) //抛弃
)

func (ct CollectionType) IsSupported() bool {
	return ct == CollectionTypeWantWatch || ct == CollectionTypeHasWatched || ct == CollectionTypeWatching || ct == CollectionTypePutAside || ct == CollectionTypeForgive
}

func (ct CollectionType) ToString() string {
	switch ct {
	case CollectionTypeWantWatch:
		return "1"
	case CollectionTypeHasWatched:
		return "2"
	case CollectionTypeWatching:
		return "3"
	case CollectionTypePutAside:
		return "4"
	case CollectionTypeForgive:
		return "5"
	default:
		return ""
	}
}

//SubjectUserCollection 查询用户收藏时返回的subject 结构
type SubjectUserCollection struct {
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
