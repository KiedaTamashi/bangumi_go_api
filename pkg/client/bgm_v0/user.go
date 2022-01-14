package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
	"strconv"
)

func (cli *Client) GetMe(ctx context.Context, authToken string) (*items.UserBgm, error) {
	//type respStruct struct {
	//	//Code int `json:"code"`
	//	//Message string `json:"message"`
	//	Data *items.UserBgm  `json:"data"`
	//}
	userBaseInfo := &items.UserBgm{}
	if authToken == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "auth token is required!")
	}
	//var resp = &respStruct{}
	param := map[string]string{
		//"pageIndex": 1,
		//"pageSize":  100,
	}
	err := cli.GET(ctx, "/v0/me", authToken, 0, param, nil, userBaseInfo)
	if err != nil {
		return nil, err
	}
	return userBaseInfo, nil
}

//GetUser query can be user id/username
func (cli *Client) GetUser(ctx context.Context, authToken string, query string) (*items.UserBgm, error) {
	var resp = &items.UserBgm{}
	param := map[string]string{
		//"pageIndex": 1,
		//"pageSize":  100,
	}
	if query == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	err := cli.GET(ctx, "/user/"+query, authToken, 0, param, nil, resp)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	return resp, nil
}

//GetUserCollection 根据用户username 和筛选条件 搜索用户收藏. 获取对应用户的收藏，查看私有收藏需要access token。
//todo [refine] golang optional input 改造
func (cli *Client) GetUserCollection(ctx context.Context, authToken string, username string, subjectType items.SubjectType,
	collectionStatusId items.CollectionStatusId, limit int64, offset int64) ([]*items.UserCollection, error) {
	//var resp = struct {
	//	data []*items.UserCollection
	//}{}
	var resp = make([]*items.UserCollection, 0)
	if !subjectType.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject type %v is illegal!", subjectType)
	}
	collectionStatusType := collectionStatusId.Type()
	if !collectionStatusType.IsSupported() {
		return nil, errno.Errorf(errno.ErrBadRequest, "collection type %v is illegal!", collectionStatusType)
	}
	param := map[string]string{
		"subject_type": subjectType.ToString(),
		"type":         collectionStatusType.String(),
		"limit":        strconv.FormatInt(limit, 10),
		"offset":       strconv.FormatInt(offset, 10),
	}
	if username == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	err := cli.GET(ctx, "/v0/users/"+username+"/collections", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
