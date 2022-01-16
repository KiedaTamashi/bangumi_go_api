package items

type ResponseGroup string

const (
	ResponseGroupSmall  = "small"
	ResponseGroupMedium = "medium"
)

func (rg ResponseGroup) IsMedium() bool {
	if rg == ResponseGroupMedium {
		return true
	} else {
		return false
	}
}

//SimpleSearchQuery Bangumi openAPI提供的简单搜索
type SimpleSearchQuery struct {
	Keywords      string        `json:"keywords,required"`
	Type          SubjectType   `json:"type"`
	ResponseGroup ResponseGroup `json:"responseGroup"`
	Start         int64         `json:"start"`
	MaxResults    int64         `json:"max_results"`
}

func (ss SimpleSearchQuery) Validate() bool {
	if ss.Keywords == "" {
		return false
	}
	if ss.Type != 0 && !ss.Type.IsSupported() {
		return false
	}
	if ss.Start < 0 {
		return false
	}
	if ss.MaxResults <= 0 {
		return false
	}
	return true
}
