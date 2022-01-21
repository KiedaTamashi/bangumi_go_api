package items

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

//SubjectBgm bangumi anime条目的格式
type SubjectBgm struct {
	Id       int         `json:"id,omitempty"`
	Type     SubjectType `json:"type,omitempty"`
	Name     string      `json:"name,omitempty"`
	NameCn   string      `json:"name_cn,omitempty"`
	Summary  string      `json:"summary,omitempty"`
	Nsfw     bool        `json:"nsfw,omitempty"`
	Locked   bool        `json:"locked,omitempty"`
	Date     time.Time   `json:"date,omitempty"` //放送日期
	Platform string      `json:"platform,omitempty"`
	Images   *ImageBgm   `json:"images,omitempty"`
	//Infobox       []*detaiInfoBoxItemBgm 			  `json:"infobox,omitempty"` //[]*detaiInfoBoxItemBgm
	Volumes       int                    `json:"volumes,omitempty"`
	Eps           int                    `json:"eps,omitempty"`
	TotalEpisodes int                    `json:"total_episodes,omitempty"` //总集数
	Rating        *RatingBgm             `json:"rating,omitempty"`
	Collection    *CollectionOverviewBgm `json:"collection,omitempty"`
	Tags          []*TagBgm              `json:"tags,omitempty"`
	// infobox 具体内容 todo [refine] 合并infobox
	AliasName            []string  `json:"alias_name,omitempty"`
	AirDate              time.Time `json:"air_date"`              //放送日期
	AirWeekday           WeekNo    `json:"air_weekday,omitempty"` //放送星期
	AirTvSite            string    `json:"air_tv_site,omitempty"` //放送电视台
	OtherTvSite          string    `json:"other_tv_site,omitempty"`
	OfficialWebSite      string    `json:"official_web_site,omitempty"`
	Copyright            string    `json:"copyright,omitempty"`
	Other                string    `json:"other,omitempty"`    //也不知道干嘛的
	Original             string    `json:"original,omitempty"` //原作
	Director             string    `json:"director,omitempty"` //监督
	Scripter             string    `json:"scripter,omitempty"` //脚本
	Musician             string    `json:"musician,omitempty"`
	MusicProduce         string    `json:"music_produce,omitempty"`
	AnimeProduce         string    `json:"anime_produce,omitempty"`
	CharacterSetting     string    `json:"character_setting,omitempty"`
	ThemeSongLyrics      string    `json:"theme_song_lyrics,omitempty"`      //主题曲填词
	ThemeSongComposition string    `json:"theme_song_composition,omitempty"` //主题曲作曲
	ThemeSongArrange     string    `json:"theme_song_arrange,omitempty"`     //主题曲编曲
	ThemeSongDirecting   string    `json:"theme_song_directing,omitempty"`   //主题曲演出
	Directing            string    `json:"directing,omitempty"`              //演出
	AnimationDirector    string    `json:"animation_director,omitempty"`     //作监
}

type TagBgm struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

func (ani *SubjectBgm) UnmarshalJSON(data []byte) error {
	tmp := struct {
		//Nested
		Id       int         `json:"id,omitempty"`
		Type     SubjectType `json:"type,omitempty"`
		Name     string      `json:"name,omitempty"`
		NameCn   string      `json:"name_cn,omitempty"`
		Summary  string      `json:"summary,omitempty"`
		Nsfw     bool        `json:"nsfw,omitempty"`
		Locked   bool        `json:"locked,omitempty"`
		Date     string      `json:"date,omitempty"`
		Platform string      `json:"platform,omitempty"`
		Images   *ImageBgm   `json:"images,omitempty"`
		//Infobox       []*detaiInfoBoxItemBgm 			  `json:"infobox,omitempty"` //[]*detaiInfoBoxItemBgm
		Volumes       int                    `json:"volumes,omitempty"`
		Eps           int                    `json:"eps,omitempty"`
		TotalEpisodes int                    `json:"total_episodes,omitempty"`
		Rating        *RatingBgm             `json:"rating,omitempty"`
		Collection    *CollectionOverviewBgm `json:"collection,omitempty"`
		Tags          []*TagBgm              `json:"tags,omitempty"`
	}{}
	// unmarshal Nested alone
	tmpNest := struct {
		Boxes []*DetailInfoBoxItemBgm `json:"infobox,omitempty"`
	}{}
	//fmt.Printf("parsing object json %s \n", string(data))
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	// the Nested impl UnmarshalJSON, so it should be unmarshaled alone
	if err := json.Unmarshal(data, &tmpNest); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	//fmt.Printf("tmp object: %+v \n", tmp)
	(ani).Id = tmp.Id
	(ani).Type = tmp.Type
	(ani).Name = tmp.Name
	(ani).NameCn = tmp.NameCn
	(ani).Summary = tmp.Summary
	(ani).Nsfw = tmp.Nsfw
	(ani).Locked = tmp.Locked
	//date_sep := strings.Split(tmp.Date,"-")
	//date_sep[1] = Num2MonthMap[date_sep[1]]
	//date := strings.Join(date_sep,"-")
	(ani).Date, _ = time.ParseInLocation("2006-01-02 15:04:05", tmp.Date+" 00:00:00", time.Local) //2021-Jan-02
	(ani).Platform = tmp.Platform
	(ani).Images = tmp.Images
	(ani).Volumes = tmp.Volumes
	(ani).Eps = tmp.Eps
	(ani).TotalEpisodes = tmp.TotalEpisodes
	(ani).Rating = tmp.Rating
	(ani).Collection = tmp.Collection
	(ani).Tags = tmp.Tags

	//info box赋值
	for _, box := range tmpNest.Boxes {
		switch box.Key {
		case "别名":
			for _, item := range box.Value.([]interface{}) {
				for _, v := range item.(map[string]interface{}) {
					(ani).AliasName = append((ani).AliasName, v.(string))
				}
			}
		case "放送开始":
			if !ani.Date.IsZero() {
				(ani).AirDate = ani.Date
			} else {
				date := strings.Replace(strings.Replace(box.Value.(string), "年", "-", 1), "月", "-", 1)
				date = date[0 : len(date)-3]
				yearMonthDay := strings.Split(date, "-")
				(ani).AirDate, _ = time.Parse("2006-01-02 15:04:05", fmt.Sprintf("%4s-%2s-%2s", yearMonthDay[0], yearMonthDay[1], yearMonthDay[2])+" 00:00:00")
			}
		case "放送星期":
			(ani).AirWeekday = WeekNo(box.Value.(string))
		case "官方网站":
			(ani).OfficialWebSite = box.Value.(string)
		case "播放电视台":
			(ani).AirTvSite = box.Value.(string)
		case "其他电视台":
			(ani).OtherTvSite = box.Value.(string)
		case "其他":
			(ani).Other = box.Value.(string)
		case "Copyright":
			(ani).Copyright = box.Value.(string)
		case "原作":
			(ani).Original = box.Value.(string)
		case "导演":
			(ani).Director = box.Value.(string)
		case "脚本":
			(ani).Scripter = box.Value.(string)
		case "音乐":
			(ani).Musician = box.Value.(string)
		case "音乐制作":
			(ani).MusicProduce = box.Value.(string)
		case "动画制作":
			(ani).AnimeProduce = box.Value.(string)
		case "人物设定":
			(ani).CharacterSetting = box.Value.(string)
		case "主题歌作词":
			(ani).ThemeSongLyrics = box.Value.(string)
		case "主题歌作曲":
			(ani).ThemeSongComposition = box.Value.(string)
		case "主题歌编曲":
			(ani).ThemeSongArrange = box.Value.(string)
		case "主题歌演出":
			(ani).ThemeSongDirecting = box.Value.(string)
		case "演出":
			(ani).Directing = box.Value.(string)
		case "作画监督":
			(ani).AnimationDirector = box.Value.(string)
		}
	}
	return nil
}

//RelatedSubject 条目下所关联的条目
type RelatedSubject struct {
	Id       int         `json:"id"`
	Type     SubjectType `json:"type"`
	Name     string      `json:"name"`
	NameCn   string      `json:"name_cn"`
	Images   *ImageBgm   `json:"images,omitempty"`
	Relation string      `json:"relation"`
}

type CharacterRelatedSubject struct {
	Id     int    `json:"id"`
	Staff  string `json:"staff"`
	Name   string `json:"name"`
	NameCn string `json:"name_cn"`
	Image  string `json:"image,omitempty"`
}
