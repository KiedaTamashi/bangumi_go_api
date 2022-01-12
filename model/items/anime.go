package items

//AnimeBgm bangumi anime条目的格式
type AnimeBgm struct {
	Id       int            `json:"id,omitempty"`
	Type     int            `json:"type,omitempty"`
	Name     string         `json:"name,omitempty"`
	NameCn   string         `json:"name_cn,omitempty"`
	Summary  string         `json:"summary,omitempty"`
	Nsfw     bool           `json:"nsfw,omitempty"`
	Locked   bool           `json:"locked,omitempty"`
	Date     string         `json:"date"`
	Platform string         `json:"platform,omitempty"`
	Images   *AnimeBgmImage `json:"images,omitempty"`
	//Infobox       string 			  `json:"infobox,omitempty"` //[]*detaiInfoBoxItemBgm
	Volumes       int            `json:"volumes,omitempty"`
	Eps           int            `json:"eps,omitempty"`
	TotalEpisodes int            `json:"total_episodes,omitempty"`
	Rating        *RatingBgm     `json:"rating,omitempty"`
	Collection    *CollectionBgm `json:"collection,omitempty"`
	Tags          []*tagBgm      `json:"tags,omitempty"`
}

type AnimeBgmImage struct {
	Large  string `json:"large,omitempty"` //均为url
	Common string `json:"common,omitempty"`
	Medium string `json:"medium,omitempty"`
	Small  string `json:"small,omitempty"`
	Grid   string `json:"grid,omitempty"`
}

type tagBgm struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type detaiInfoBoxItemBgm struct {
	Key   string `json:"key,omitempty"`
	Value string `json:"value,omitempty"` //别名有点问题，value可能是复杂结构
}
