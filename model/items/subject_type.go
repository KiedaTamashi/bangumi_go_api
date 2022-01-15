package items

type SubjectType int

const (
	SubjectTypeBook  = SubjectType(1)
	SubjectTypeAnime = SubjectType(2)
	SubjectTypeMusic = SubjectType(3)
	SubjectTypeGame  = SubjectType(4)
	SubjectTypeReal  = SubjectType(5)
)

func (st SubjectType) IsSupported() bool {
	return st == SubjectTypeBook || st == SubjectTypeAnime || st == SubjectTypeMusic || st == SubjectTypeGame || st == SubjectTypeReal
}

func (st SubjectType) ToString() string {
	switch st {
	case SubjectTypeBook:
		return "1"
	case SubjectTypeAnime:
		return "2"
	case SubjectTypeMusic:
		return "3"
	case SubjectTypeGame:
		return "4"
	case SubjectTypeReal:
		return "5"
	default:
		return ""
	}
}

func (st SubjectType) Name() string {
	switch st {
	case SubjectTypeBook:
		return "书籍"
	case SubjectTypeAnime:
		return "动画"
	case SubjectTypeMusic:
		return "音乐"
	case SubjectTypeGame:
		return "游戏"
	case SubjectTypeReal:
		return "三次元"
	default:
		return ""
	}
}
