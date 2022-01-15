package items

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
