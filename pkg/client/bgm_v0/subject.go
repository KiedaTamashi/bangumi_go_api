package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
)

//GetSubject 根据id获得单个条目的详细信息
func (cli *Client) GetSubject(ctx context.Context, authToken string, subjectId string) (*items.AnimeBgm, error) {

	//var subjectList []*items.AnimeBgm
	//
	//for {
	var resp = &items.AnimeBgm{}
	param := map[string]string{
		//"pageIndex": 1,
		//"pageSize":  100,
	}
	if subjectId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	err := cli.GET(ctx, "/subjects/"+subjectId, authToken, 0, param, resp)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	return resp, nil
}
