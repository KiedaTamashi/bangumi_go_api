package items

import (
	"encoding/json"
	"fmt"
	"time"
)

type PersonCareer string

const (
	PersonCareerProducer    = PersonCareer("producer")
	PersonCareerMangaka     = PersonCareer("mangaka") //漫画家
	PersonCareerArtist      = PersonCareer("artist")
	PersonCareerSeiyu       = PersonCareer("seiyu") //声优
	PersonCareerWriter      = PersonCareer("writer")
	PersonCareerIllustrator = PersonCareer("illustrator")
	PersonCareerActor       = PersonCareer("actor")
)

func (pc PersonCareer) String() string {
	switch pc {
	case PersonCareerProducer:
		return "producer"
	case PersonCareerMangaka:
		return "mangaka"
	case PersonCareerArtist:
		return "artist"
	case PersonCareerSeiyu:
		return "seiyu"
	case PersonCareerWriter:
		return "writer"
	case PersonCareerIllustrator:
		return "illustrator"
	case PersonCareerActor:
		return "actor"
	default:
		return ""
	}
}

type PersonType int

const (
	PeresonTypePersonal   = PersonType(1)
	PersonTypeBusiness    = PersonType(2)
	PersonTypeCombination = PersonType(3)
)

func (pt PersonType) Name() string {
	switch pt {
	case PeresonTypePersonal:
		return "个人"
	case PersonTypeBusiness:
		return "公司"
	case PersonTypeCombination:
		return "组合"
	default:
		return ""
	}
}

//RelatedPerson subject相关的person信息
type RelatedPerson struct {
	Id       int            `json:"id,omitempty"`
	Name     string         `json:"name,omitempty"`
	Type     PersonType     `json:"type,omitempty"`
	Career   []PersonCareer `json:"career,omitempty"`
	Images   *ImageBgm      `json:"images,omitempty"`
	Relation string         `json:"relation"` //比如oped设计，监督啥的
}

//PersonCharacter [PersonCharacter] 之所以叫person character。因为关联时person和character都用这个接口体= =
type PersonCharacter struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Type          CharacterType `json:"type"`
	Images        *ImageBgm     `json:"images,omitempty"`
	SubjectId     int           `json:"subject_id,omitempty"`
	SubjectName   string        `json:"subject_name,omitempty"`
	SubjectNameCn string        `json:"subject_name_cn,omitempty"`
}

type PersonDetail struct {
	Id           int64          `json:"id,omitempty"`
	Name         string         `json:"name,omitempty"`
	Type         PersonType     `json:"type,omitempty"`
	Career       []PersonCareer `json:"career,omitempty"`
	Images       *ImageBgm      `json:"images,omitempty"`
	Summary      string         `json:"summary,omitempty"`
	Locked       bool           `json:"locked,omitempty"`
	LastModified time.Time      `json:"last_modified,omitempty"`
	//InfoBox []...
	Gender     string    `json:"gender,omitempty"`
	BloodType  BloodType `json:"blood_type,omitempty"`
	BirthYear  int       `json:"birth_year,omitempty"`
	BirthMonth int       `json:"birth_month,omitempty"`
	BirthDay   int       `json:"birth_day,omitempty"`
	Stat       Stat      `json:"stat"`
	Img        string    `json:"img,omitempty"`
	// InfoBox Detail

	// infobox 具体内容 todo [refine] 合并infobox
	NameCn          string   `json:"name_cn,omitempty"`
	AliasName       []string `json:"alias_name,omitempty"`
	Sex             string   `json:"sex,omitempty"`
	BirthDate       string   `json:"birth_date,omitempty"` //x月x日。。。意义不明的字段
	BirthPlace      string   `json:"birth_place,omitempty"`
	BloodTypeString string   `json:"blood_type_string,omitempty"` //血型具体
	Height          string   `json:"height,omitempty"`
	Weight          string   `json:"weight,omitempty"`
	Age             string   `json:"age,omitempty"`
	Agency          string   `json:"agency,omitempty"`
	School          string   `json:"school,omitempty"`
	Spouse          []string `json:"spouse,omitempty"`
	Source          string   `json:"source,omitempty"` //url
	HomePageUrl     string   `json:"home_page_url,omitempty"`
}

func (pd *PersonDetail) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id           int64      `json:"id"`
		Name         string     `json:"name"`
		Type         PersonType `json:"type"`
		Images       *ImageBgm  `json:"images,omitempty"`
		Summary      string     `json:"summary"`
		Locked       bool       `json:"locked"`
		LastModified time.Time  `json:"last_modified,omitempty"`
		//InfoBox []*detaiInfoBoxItemBgm
		Gender     string    `json:"gender,omitempty"`
		BloodType  BloodType `json:"blood_type,omitempty"`
		BirthYear  int       `json:"birth_year,omitempty"`
		BirthMonth int       `json:"birth_month,omitempty"`
		BirthDay   int       `json:"birth_day,omitempty"`
		Stat       Stat      `json:"stat"`
		Img        string    `json:"img,omitempty"`
	}{}
	tmpNest := struct {
		Boxes []*DetailInfoBoxItemBgm `json:"infobox,omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	if err := json.Unmarshal(data, &tmpNest); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	//fmt.Printf("tmp object: %+v \n", tmp)
	(pd).Id = tmp.Id
	(pd).Type = tmp.Type
	(pd).Name = tmp.Name
	(pd).Summary = tmp.Summary
	(pd).Images = tmp.Images
	(pd).Locked = tmp.Locked
	(pd).LastModified = tmp.LastModified

	(pd).Gender = tmp.Gender
	(pd).Images = tmp.Images
	(pd).BloodType = tmp.BloodType
	(pd).BirthYear = tmp.BirthYear
	(pd).BirthMonth = tmp.BirthMonth
	(pd).BirthDay = tmp.BirthDay
	(pd).Stat = tmp.Stat
	(pd).Img = tmp.Img

	//info box赋值
	for _, box := range tmpNest.Boxes {
		switch box.Key {
		case "别名":
			for _, item := range box.Value.([]interface{}) {
				for k, v := range item.(map[string]interface{}) {
					if k == "v" {
						(pd).AliasName = append((pd).AliasName, v.(string))
					}
				}
			}
		case "出生地":
			(pd).BirthPlace = box.Value.(string)
		case "简体中文名":
			(pd).NameCn = box.Value.(string)
		case "性别":
			(pd).Sex = box.Value.(string)
		case "生日":
			(pd).BirthDate = box.Value.(string)
		case "血型":
			(pd).BloodTypeString = box.Value.(string)
		case "身高":
			(pd).Height = box.Value.(string)
		case "年龄":
			(pd).Age = box.Value.(string)
		case "体重":
			(pd).Weight = box.Value.(string)
		case "引用来源":
			(pd).Source = box.Value.(string)
		case "毕业院校":
			(pd).School = box.Value.(string)
		case "事务所":
			(pd).Agency = box.Value.(string)
		case "事务所个人页面":
			(pd).HomePageUrl = box.Value.(string)
		case "配偶":
			for _, item := range box.Value.([]interface{}) {
				for _, v := range item.(map[string]interface{}) {
					(pd).Spouse = append((pd).Spouse, v.(string))
				}
			}
		}
	}
	return nil
}
