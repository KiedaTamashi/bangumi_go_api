package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestName(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetMe(ctx, "Bearer xxx")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetUser(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetUser(ctx, "Bearer xxx", "1")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetUserCollection(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetUserCollection(ctx, "", "sai", items.SubjectTypeAnime, items.CollectionStatusIdCollect, 40, 5)
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
