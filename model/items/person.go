package items

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

//CharacterRelatedPerson [PersonCharacter] character相关的，person信息
type CharacterRelatedPerson struct {
	Id            int           `json:"id"`
	Name          string        `json:"name"`
	Type          CharacterType `json:"type"`
	Images        *ImageBgm     `json:"images,omitempty"`
	SubjectId     int           `json:"subject_id,omitempty"`
	SubjectName   string        `json:"subject_name,omitempty"`
	SubjectNameCn string        `json:"subject_name_cn,omitempty"`
}
