package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
)

func (cli *Client) GetMe(ctx context.Context, authToken string) (*items.UserBgm, error) {
	//type respStruct struct {
	//	//Code int `json:"code"`
	//	//Message string `json:"message"`
	//	Data *items.UserBgm  `json:"data"`
	//}
	userBaseInfo := &items.UserBgm{}

	//var resp = &respStruct{}
	param := map[string]string{
		//"pageIndex": 1,
		//"pageSize":  100,
	}
	err := cli.GET(ctx, "/me", authToken, 0, param, userBaseInfo)
	if err != nil {
		return nil, err
	}
	return userBaseInfo, nil
}
