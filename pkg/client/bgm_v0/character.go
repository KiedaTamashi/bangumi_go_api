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
		return nil, errno.Errorf(errno.ErrBadRequest, "character id is required!")
	}
	err := cli.GET(ctx, "/v0/characters/"+characterId, authToken, 0, param, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//GetCharacterRelatedSubject 获得角色相关条目
func (cli *Client) GetCharacterRelatedSubject(ctx context.Context, authToken string, characterId string) ([]*items.CharacterRelatedSubject, error) {

	var resp = make([]*items.CharacterRelatedSubject, 0)
	param := map[string]string{}
	if characterId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "character id is required!")
	}
	err := cli.GET(ctx, "/v0/characters/"+characterId+"/subjects", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//GetCharacterRelatedPerson 获得角色相关人物
func (cli *Client) GetCharacterRelatedPerson(ctx context.Context, authToken string, characterId string) ([]*items.PersonCharacter, error) {

	var resp = make([]*items.PersonCharacter, 0)
	param := map[string]string{}
	if characterId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "character id is required!")
	}
	err := cli.GET(ctx, "/v0/characters/"+characterId+"/persons", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
