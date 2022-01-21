package items

import (
	"encoding/json"
	"fmt"
)

type CharacterType int

const (
	CharacterTypeCharacter = CharacterType(1) //角色
	CharacterTypeRobot     = CharacterType(2) //机体
	CharacterTypeVehicle   = CharacterType(3) //船舰
	CharacterTypeOrg       = CharacterType(4) //组织
)

func (ct CharacterType) Name() string {
	switch ct {
	case CharacterTypeCharacter:
		return "角色"
	case CharacterTypeRobot:
		return "机体"
	case CharacterTypeVehicle:
		return "船舰"
	case CharacterTypeOrg:
		return "组织"
	default:
		return ""
	}
}

type RelatedCharacter struct {
	Id       int           `json:"id,omitempty"`
	Name     string        `json:"name,omitempty"`
	Type     CharacterType `json:"type,omitempty"`
	Images   *ImageBgm     `json:"images,omitempty"`
	Relation string        `json:"relation,omitempty"`
}

type CharacterDetail struct {
	Id      int64         `json:"id"`
	Name    string        `json:"name"`
	Type    CharacterType `json:"type"`
	Images  *ImageBgm     `json:"images,omitempty"`
	Summary string        `json:"summary"`
	Locked  bool          `json:"locked"`
	//InfoBox []*detaiInfoBoxItemBgm
	Gender     string    `json:"gender,omitempty"`
	BloodType  BloodType `json:"blood_type,omitempty"`
	BirthYear  int       `json:"birth_year,omitempty"`
	BirthMonth int       `json:"birth_month,omitempty"`
	BirthDay   int       `json:"birth_day,omitempty"`
	Stat       Stat      `json:"stat"`

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
	Source          string   `json:"source,omitempty"` //url

}

func (cd *CharacterDetail) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id      int64         `json:"id"`
		Name    string        `json:"name"`
		Type    CharacterType `json:"type"`
		Images  *ImageBgm     `json:"images,omitempty"`
		Summary string        `json:"summary"`
		Locked  bool          `json:"locked"`
		//InfoBox []*detaiInfoBoxItemBgm
		Gender     string    `json:"gender,omitempty"`
		BloodType  BloodType `json:"blood_type,omitempty"`
		BirthYear  int       `json:"birth_year,omitempty"`
		BirthMonth int       `json:"birth_month,omitempty"`
		BirthDay   int       `json:"birth_day,omitempty"`
		Stat       Stat      `json:"stat"`
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
	(cd).Id = tmp.Id
	(cd).Type = tmp.Type
	(cd).Name = tmp.Name
	(cd).Summary = tmp.Summary
	(cd).Images = tmp.Images
	(cd).Locked = tmp.Locked
	(cd).Gender = tmp.Gender
	(cd).Images = tmp.Images
	(cd).BloodType = tmp.BloodType
	(cd).BirthYear = tmp.BirthYear
	(cd).BirthMonth = tmp.BirthMonth
	(cd).BirthDay = tmp.BirthDay
	(cd).Stat = tmp.Stat

	//info box赋值
	for _, box := range tmpNest.Boxes {
		switch box.Key {
		case "别名":
			for _, item := range box.Value.([]interface{}) {
				for k, v := range item.(map[string]interface{}) {
					if k == "v" {
						(cd).AliasName = append((cd).AliasName, v.(string))
					}
					//(cd).AliasName = append((cd).AliasName, v.(string))
				}
			}
		case "出生地":
			(cd).BirthPlace = box.Value.(string)
		case "简体中文名":
			(cd).NameCn = box.Value.(string)
		case "性别":
			(cd).Sex = box.Value.(string)
		case "生日":
			(cd).BirthDate = box.Value.(string)
		case "血型":
			(cd).BloodTypeString = box.Value.(string)
		case "身高":
			(cd).Height = box.Value.(string)
		case "年龄":
			(cd).Age = box.Value.(string)
		case "体重":
			(cd).Weight = box.Value.(string)
		case "引用来源":
			(cd).Source = box.Value.(string)
		}
	}
	return nil
}
