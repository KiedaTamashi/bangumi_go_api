package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
	"strconv"
)

type ManageCollectionParam struct {
	SubjectId string                     `json:"subject_id"`
	Status    items.CollectionStatusType `json:"status"`
	Comment   *string                    `json:"comment,omitempty"`
	Tags      []string                   `json:"tags,omitempty"`
	Rating    *int64                     `json:"rating,omitempty"`
	Privacy   *int64                     `json:"privacy,omitempty"`
}

//GetCollection 根据subject id 获得某个条目的收藏信息，需要auth
func (cli *Client) GetCollection(ctx context.Context, authToken string, subjectId string) (*items.CollectionBgm, error) {
	var resp = &items.CollectionBgm{}
	param := map[string]string{}
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	if authToken == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "auth token is required!")
	}
	err := cli.GET(ctx, "/collection/"+subjectId, authToken, 0, param, nil, resp)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	resp.Id, _ = strconv.ParseInt(subjectId, 10, 64)
	return resp, nil
}

//ManageCollection 根据subject id 修改某个条目的收藏信息，需要auth，是post。但还是用param的形式（颇为迷惑
func (cli *Client) ManageCollection(ctx context.Context, authToken string, subjectId string, status items.CollectionStatusType,
	comment string, tags []string, rating int64, privacy int64) (*items.CollectionBgm, error) {
	var resp = &items.CollectionBgm{}
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	if authToken == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "auth token is required!")
	}
	if status.String() == "" || !status.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "status is invalid!")
	}

	tagStr := ""
	for i, tag := range tags {
		if i == 0 {
			tagStr = tag
		} else {
			tagStr = fmt.Sprintf("%s,%s", tagStr, tag)
		}
	}
	ratingStr := ""
	if 1 <= rating && rating <= 10 {
		ratingStr = fmt.Sprintf("%d", rating)
	}
	privacyStr := ""
	if privacy == 1 || privacy == 0 {
		privacyStr = fmt.Sprintf("%d", privacy)
	}
	param := map[string]string{}
	body := map[string]string{
		"status":  status.String(),
		"comment": comment,
		"tags":    tagStr,
		"rating":  ratingStr,
		"privacy": privacyStr,
	}
	// 统一使用update了，系统会自动判断是新建还是更新
	err := cli.POST(ctx, "/collection/"+subjectId+"/update", authToken, 0, param, body, resp, true)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	resp.Id, _ = strconv.ParseInt(subjectId, 10, 64)
	return resp, nil
}

//ManageCollectionRaw 根据subject id 修改某个条目的收藏信息，需要auth，是post。但还是用param的形式（颇为迷惑
// todo [refine] 可选参数，你确实可以可选，但它会设成默认值。rating 是0， private是false
func (cli *Client) ManageCollectionRaw(ctx context.Context, authToken string, params ManageCollectionParam) (*items.CollectionBgm, error) {
	var resp = &items.CollectionBgm{}
	subjectId := params.SubjectId
	status := params.Status
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	if authToken == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "auth token is required!")
	}
	if status.String() == "" || !status.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "status is invalid!")
	}
	param := map[string]string{}
	body := map[string]string{
		"status": status.String(),
	}
	if params.Comment != nil {
		body["comment"] = *params.Comment
	}

	if params.Tags != nil {
		tagStr := ""
		for i, tag := range params.Tags {
			if i == 0 {
				tagStr = tag
			} else {
				tagStr = fmt.Sprintf("%s,%s", tagStr, tag)
			}
		}
	}
	if params.Rating != nil {
		ratingStr := ""
		if 1 <= *params.Rating && *params.Rating <= 10 {
			ratingStr = fmt.Sprintf("%d", *params.Rating)
			body["rating"] = ratingStr
		}
	}
	if params.Privacy != nil {
		privacyStr := ""
		if *params.Privacy == 1 || *params.Privacy == 0 {
			privacyStr = fmt.Sprintf("%d", *params.Privacy)
			body["privacy"] = privacyStr
		}
	}
	// 统一使用update了，系统会自动判断是新建还是更新
	err := cli.POST(ctx, "/collection/"+subjectId+"/update", authToken, 0, param, body, resp, true)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	resp.Id, _ = strconv.ParseInt(subjectId, 10, 64)
	return resp, nil
}
