package items

import (
	"encoding/json"
	"fmt"
	"time"
)

// note 这里存放了小型的subject struct

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

//SubjectMedium Calendar 中出现的小型subject结构体. 对应schema里的subjectSmall，search里的subject medium
type SubjectMedium struct {
	Id         int                    `json:"id,omitempty"`
	Url        string                 `json:"url,omitempty"`
	Type       SubjectType            `json:"type,omitempty"`
	Name       string                 `json:"name,omitempty"`
	NameCn     string                 `json:"name_cn,omitempty"`
	Summary    string                 `json:"summary,omitempty"`
	AirDate    time.Time              `json:"air_date"`              //放送日期
	AirWeekday int                    `json:"air_weekday,omitempty"` //放送星期
	Images     *ImageBgm              `json:"images,omitempty"`
	Eps        int                    `json:"eps,omitempty"`
	EpsCount   int                    `json:"eps_count,omitempty"`
	Rating     *RatingBgm             `json:"rating,omitempty"`
	Rank       int                    `json:"rank"`
	Collection *CollectionOverviewBgm `json:"collection,omitempty"`
}

func (subs *SubjectMedium) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id      int         `json:"id,omitempty"`
		Url     string      `json:"url,omitempty"`
		Type    SubjectType `json:"type,omitempty"`
		Name    string      `json:"name,omitempty"`
		NameCn  string      `json:"name_cn,omitempty"`
		Summary string      `json:"summary,omitempty"`
		//AirDate             time.Time 			`json:"air_date"`              //放送日期
		AirWeekday int                    `json:"air_weekday,omitempty"` //放送星期
		Images     *ImageBgm              `json:"images,omitempty"`
		Eps        int                    `json:"eps,omitempty"`
		EpsCount   int                    `json:"eps_count,omitempty"`
		Rating     *RatingBgm             `json:"rating,omitempty"`
		Rank       int                    `json:"rank"`
		Collection *CollectionOverviewBgm `json:"collection,omitempty"`
	}{}
	tmpNest := struct {
		AirDate string `json:"air_date"` //放送日期
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	if err := json.Unmarshal(data, &tmpNest); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	(subs).Id = tmp.Id
	(subs).Type = tmp.Type
	(subs).Url = tmp.Url
	(subs).Name = tmp.Name
	(subs).NameCn = tmp.NameCn
	(subs).Summary = tmp.Summary
	(subs).AirDate, _ = time.ParseInLocation("2006-01-02 15:04:05", tmpNest.AirDate+" 00:00:00", time.Local) //2021-Jan-02
	(subs).AirWeekday = tmp.AirWeekday
	(subs).Images = tmp.Images
	(subs).Eps = tmp.Eps
	(subs).EpsCount = tmp.EpsCount
	(subs).Rating = tmp.Rating
	(subs).Rank = tmp.Rank
	(subs).Collection = tmp.Collection
	return nil
}

//SubjectSmall Search 中可选的small subject结构体. 对应schema里的subjectBase
type SubjectSmall struct {
	Id         int         `json:"id,omitempty"`
	Url        string      `json:"url,omitempty"`
	Type       SubjectType `json:"type,omitempty"`
	Name       string      `json:"name,omitempty"`
	NameCn     string      `json:"name_cn,omitempty"`
	Summary    string      `json:"summary,omitempty"`
	AirDate    time.Time   `json:"air_date"`              //放送日期
	AirWeekday int         `json:"air_weekday,omitempty"` //放送星期
	Images     *ImageBgm   `json:"images,omitempty"`
}

func (subs *SubjectSmall) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id      int         `json:"id,omitempty"`
		Url     string      `json:"url,omitempty"`
		Type    SubjectType `json:"type,omitempty"`
		Name    string      `json:"name,omitempty"`
		NameCn  string      `json:"name_cn,omitempty"`
		Summary string      `json:"summary,omitempty"`
		//AirDate             time.Time 			`json:"air_date"`              //放送日期
		AirWeekday int       `json:"air_weekday,omitempty"` //放送星期
		Images     *ImageBgm `json:"images,omitempty"`
	}{}
	tmpNest := struct {
		AirDate string `json:"air_date"` //放送日期
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	if err := json.Unmarshal(data, &tmpNest); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	(subs).Id = tmp.Id
	(subs).Type = tmp.Type
	(subs).Url = tmp.Url
	(subs).Name = tmp.Name
	(subs).NameCn = tmp.NameCn
	(subs).Summary = tmp.Summary
	(subs).AirDate, _ = time.ParseInLocation("2006-01-02 15:04:05", tmpNest.AirDate+" 00:00:00", time.Local) //2021-Jan-02
	(subs).AirWeekday = tmp.AirWeekday
	(subs).Images = tmp.Images
	return nil
}
