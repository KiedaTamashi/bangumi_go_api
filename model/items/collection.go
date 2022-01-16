package items

import (
	"encoding/json"
	"fmt"
	"time"
)

type CollectionStatusType string

const (
	CollectionStatusTypeWish    = CollectionStatusType("wish")    //想看
	CollectionStatusTypeCollect = CollectionStatusType("collect") //看过
	CollectionStatusTypeDo      = CollectionStatusType("do")      //在看
	CollectionStatusTypeOnHold  = CollectionStatusType("on_hold") //搁置
	CollectionStatusTypeDropped = CollectionStatusType("dropped") //抛弃
)

func (ct CollectionStatusType) IsSupported() bool {
	return ct == CollectionStatusTypeWish || ct == CollectionStatusTypeCollect || ct == CollectionStatusTypeDo || ct == CollectionStatusTypeOnHold || ct == CollectionStatusTypeDropped
}

//ToString 其实就是int转string
func (ct CollectionStatusType) String() string {
	switch ct {
	case CollectionStatusTypeWish:
		return "wish"
	case CollectionStatusTypeCollect:
		return "collect"
	case CollectionStatusTypeDo:
		return "do"
	case CollectionStatusTypeOnHold:
		return "on_hold"
	case CollectionStatusTypeDropped:
		return "dropped"
	default:
		return ""
	}
}

//Id 产出id
func (ct CollectionStatusType) ToId() CollectionStatusId {
	switch ct {
	case CollectionStatusTypeWish:
		return CollectionStatusIdWish
	case CollectionStatusTypeCollect:
		return CollectionStatusIdCollect
	case CollectionStatusTypeDo:
		return CollectionStatusIdDo
	case CollectionStatusTypeOnHold:
		return CollectionStatusIdOnHold
	case CollectionStatusTypeDropped:
		return CollectionStatusIdDropped
	default:
		return 0
	}
}

//ToName 对应Bangumi中的CollectionStatusName
func (ct CollectionStatusType) ToName() string {
	switch ct {
	case CollectionStatusTypeWish:
		return "想做"
	case CollectionStatusTypeCollect:
		return "做过"
	case CollectionStatusTypeDo:
		return "在做"
	case CollectionStatusTypeOnHold:
		return "搁置"
	case CollectionStatusTypeDropped:
		return "抛弃"
	default:
		return ""
	}
}

type CollectionStatusId int64

const (
	CollectionStatusIdWish    = CollectionStatusId(1) //想看
	CollectionStatusIdCollect = CollectionStatusId(2) //看过
	CollectionStatusIdDo      = CollectionStatusId(3) //在看
	CollectionStatusIdOnHold  = CollectionStatusId(4) //搁置
	CollectionStatusIdDropped = CollectionStatusId(5) //抛弃
)

//ToName 对应Bangumi中的CollectionStatusName
func (ct CollectionStatusId) Type() CollectionStatusType {
	switch ct {
	case CollectionStatusIdWish:
		return CollectionStatusTypeWish
	case CollectionStatusIdCollect:
		return CollectionStatusTypeCollect
	case CollectionStatusIdDo:
		return CollectionStatusTypeDo
	case CollectionStatusIdOnHold:
		return CollectionStatusTypeOnHold
	case CollectionStatusIdDropped:
		return CollectionStatusTypeDropped
	default:
		return ""
	}
}

func (ct CollectionStatusId) String() string {
	switch ct {
	case CollectionStatusIdWish:
		return "1"
	case CollectionStatusIdCollect:
		return "2"
	case CollectionStatusIdDo:
		return "3"
	case CollectionStatusIdOnHold:
		return "4"
	case CollectionStatusIdDropped:
		return "5"
	default:
		return ""
	}
}

type CollectionStatusName string

const (
	CollectionStatusNameWish    = CollectionStatusName("想看") //想看
	CollectionStatusNameCollect = CollectionStatusName("看过") //看过
	CollectionStatusNameDo      = CollectionStatusName("在看") //在看
	CollectionStatusNameOnHold  = CollectionStatusName("搁置") //搁置
	CollectionStatusNameDropped = CollectionStatusName("抛弃") //抛弃
)

func (ct CollectionStatusName) String() string {
	switch ct {
	case CollectionStatusNameWish:
		return "想看"
	case CollectionStatusNameCollect:
		return "看过"
	case CollectionStatusNameDo:
		return "在看"
	case CollectionStatusNameOnHold:
		return "搁置"
	case CollectionStatusNameDropped:
		return "抛弃"
	default:
		return ""
	}
}

type CollectionStatus struct {
	Id   CollectionStatusId   `json:"id,required"`
	Type CollectionStatusType `json:"type,omitempty"`
	Name CollectionStatusName `json:"name,omitempty"`
}

type CollectionBgm struct {
	Id        int64            `json:"id"` //这个collection 的 subject id
	Status    CollectionStatus `json:"status"`
	Rating    int              `json:"rating"`
	Comment   string           `json:"comment"`
	Private   bool             `json:"private"`
	Tag       []string         `json:"tag"`
	EpStatus  int              `json:"ep_status"`
	VolStatus int              `json:"vol_status"`
	Lasttouch time.Time        `json:"lasttouch"`
	User      *UserBgm         `json:"user"`
}

func (c *CollectionBgm) UnmarshalJSON(data []byte) error {
	tmp := struct {
		Id        int              `json:"id"` //这个collection 的 subject id
		Status    CollectionStatus `json:"status"`
		Rating    int              `json:"rating,omitempty"`
		Comment   string           `json:"comment,omitempty"`
		Private   int              `json:"private,omitempty"`
		Tag       []string         `json:"tag,omitempty"`
		EpStatus  int              `json:"ep_status,omitempty"`
		VolStatus int              `json:"vol_status,omitempty"`
		Lasttouch int64            `json:"lasttouch"`
		User      *UserBgm         `json:"user,omitempty"`
	}{}
	if err := json.Unmarshal(data, &tmp); err != nil {
		fmt.Printf("failed to parse object: %v", err)
		return err
	}
	(c).Status = tmp.Status
	(c).Rating = tmp.Rating
	(c).Comment = tmp.Comment
	if tmp.Private == 0 {
		(c).Private = false
	} else {
		(c).Private = true
	}
	(c).Tag = tmp.Tag
	(c).Lasttouch = time.Unix(tmp.Lasttouch, 0)
	(c).EpStatus = tmp.EpStatus
	(c).VolStatus = tmp.VolStatus
	(c).User = tmp.User
	return nil
}
