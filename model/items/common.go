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

type WeekNo string //星期几

const (
	Monday    = WeekNo("星期一")
	Tuesday   = WeekNo("星期二") // 内广，巨量引擎
	Wednesday = WeekNo("星期三") // 内广头条
	Thursday  = WeekNo("星期四") // 内广网盟(穿山甲流量)
	Friday    = WeekNo("星期五") //巨量千川
	Saturday  = WeekNo("星期六")
	Sunday    = WeekNo("星期日")
)

var Num2MonthMap = map[string]string{
	"01": "Jan",
	"02": "Feb",
	"03": "Mar",
	"04": "Apr",
	"05": "May",
	"06": "Jun",
	"07": "Jul",
	"08": "Aug",
	"09": "Sep",
	"10": "Oct",
	"11": "Nov",
	"12": "Dec",
}
