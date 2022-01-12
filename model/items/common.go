package items

type RatingBgm struct {
	Rank  int            `json:"rank"`
	Total int            `json:"total"` //打分的人
	Count map[string]int `json:"count"` //e.g. "1":3000,"2":2000
	Score float64        `json:"score"`
}

type CollectionBgm struct {
	Wish    int `json:"wish"`
	Collect int `json:"collect"`
	Doing   int `json:"doing"`
	OnHold  int `json:"on_hold"`
	Dropped int `json:"dropped"`
}

type AvatarBgm struct {
	Large  string `json:"large,omitempty"`
	Medium string `json:"medium,omitempty"`
	Small  string `json:"small,omitempty"`
}
