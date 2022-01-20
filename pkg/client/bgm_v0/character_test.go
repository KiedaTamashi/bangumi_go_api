package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestClient_GetCharacterById(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCharacterById(ctx, "", "77")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
