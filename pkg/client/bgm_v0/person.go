package bgm_v0

import (
	"context"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
)

//GetPersonById 根据id获得单个人物的详细信息 cache with 60s; authToken 似乎没用
func (cli *Client) GetPersonById(ctx context.Context, authToken string, personId string) (*items.PersonDetail, error) {

	var resp = &items.PersonDetail{}
	param := map[string]string{}
	if personId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "person id is required!")
	}
	err := cli.GET(ctx, "/v0/persons/"+personId, authToken, 0, param, nil, resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//GetPersonRelatedSubject 获得人物相关条目， 复用了CharacterRelatedSubject的结构体。。。
func (cli *Client) GetPersonRelatedSubject(ctx context.Context, authToken string, personId string) ([]*items.CharacterRelatedSubject, error) {

	var resp = make([]*items.CharacterRelatedSubject, 0)
	param := map[string]string{}
	if personId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "person id is required!")
	}
	err := cli.GET(ctx, "/v0/persons/"+personId+"/subjects", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

//GetPersonRelatedCharacter 获得角色相关人物
func (cli *Client) GetPersonRelatedCharacter(ctx context.Context, authToken string, personId string) ([]*items.PersonCharacter, error) {

	var resp = make([]*items.PersonCharacter, 0)
	param := map[string]string{}
	if personId == "" {
		return nil, errno.Errorf(errno.ErrBadRequest, "person id is required!")
	}
	err := cli.GET(ctx, "/v0/persons/"+personId+"/characters", authToken, 0, param, nil, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
