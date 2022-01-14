package items

type UserBgm struct {
	Id        int        `json:"id"`
	Url       string     `json:"url"`
	Username  string     `json:"username"`
	Nickname  string     `json:"nickname"`
	UserGroup UserGroup  `json:"user_group"`
	Avatar    *AvatarBgm `json:"avatar"`
	Sign      string     `json:"sign"` //签名
}

type UserGroup int

const (
	UserGroupAdmin        = UserGroup(1)  //管理员
	UserGroupAdminBangumi = UserGroup(2)  // bgm 管理员
	UserGroupAdminDoujin  = UserGroup(3)  //天窗管理员
	UserGroupBannedSpeak  = UserGroup(4)  //禁言用户
	UserGroupBannedEnter  = UserGroup(5)  //禁止访问用户
	UserGroupAdminPersons = UserGroup(8)  //人物管理员
	UserGroupAdminWiki    = UserGroup(9)  //维基条目管理员
	UserGroupUser         = UserGroup(10) //用户
	UserGroupWikier       = UserGroup(11) //维基人
)

func (ug UserGroup) IsSupported() bool {
	return ug == UserGroupAdmin || ug == UserGroupAdminBangumi || ug == UserGroupAdminDoujin || ug == UserGroupBannedSpeak || ug == UserGroupBannedEnter ||
		ug == UserGroupAdminPersons || ug == UserGroupAdminWiki || ug == UserGroupUser || ug == UserGroupWikier
}

func (ug UserGroup) ToString() string {
	switch ug {
	case UserGroupAdmin:
		return "1"
	case UserGroupAdminBangumi:
		return "2"
	case UserGroupAdminDoujin:
		return "3"
	case UserGroupBannedSpeak:
		return "4"
	case UserGroupBannedEnter:
		return "5"
	case UserGroupAdminPersons:
		return "8"
	case UserGroupAdminWiki:
		return "9"
	case UserGroupUser:
		return "10"
	case UserGroupWikier:
		return "11"
	default:
		return ""
	}
}

func (ug UserGroup) Name() string {
	switch ug {
	case UserGroupAdmin:
		return "管理员"
	case UserGroupAdminBangumi:
		return "Bangumi 管理猿"
	case UserGroupAdminDoujin:
		return "天窗管理猿"
	case UserGroupBannedSpeak:
		return "禁言用户"
	case UserGroupBannedEnter:
		return "禁止访问用户"
	case UserGroupAdminPersons:
		return "人物管理猿"
	case UserGroupAdminWiki:
		return "维基条目管理猿"
	case UserGroupUser:
		return "用户"
	case UserGroupWikier:
		return "维基人"
	default:
		return ""
	}
}
