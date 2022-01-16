package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/model/items"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestClient_GetCollection(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCollection(ctx, "Bearer xxx", "265")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_ManageCollection(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	var tags []string
	resp, err := cli.ManageCollection(ctx, "Bearer xxx", "265", items.CollectionStatusTypeCollect,
		"7", tags, 8, 0)
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
