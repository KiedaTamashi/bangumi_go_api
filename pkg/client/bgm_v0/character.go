package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
)

//GetCharacterById 根据id获得单个角色的详细信息 cache with 60s; authToken 似乎没用
func (cli *Client) GetCharacterById(ctx context.Context, authToken string, characterId string) (*items.CharacterDetail, error) {

	var resp = &items.CharacterDetail{}
	param := map[string]string{}
	if characterId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "subject id is required!")
	}
	err := cli.GET(ctx, "/v0/characters/"+characterId, authToken, 0, param, nil, resp)
	if err != nil {
		return nil, err
	}
	//subjectList = append(subjectList, resp)
	//}
	return resp, nil
}
