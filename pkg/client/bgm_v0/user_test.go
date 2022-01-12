package bgm_v0

import (
	"context"
	"fmt"
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
