package items

type UserBgm struct {
	Id        int        `json:"id"`
	Url       string     `json:"url"`
	Username  string     `json:"username"`
	Nickname  string     `json:"nickname"`
	UserGroup int        `json:"user_group"`
	Avatar    *AvatarBgm `json:"avatar"`
	Sign      string     `json:"sign"`
}
