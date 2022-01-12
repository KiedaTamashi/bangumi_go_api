package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestGetSubject(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetSubject(ctx, "Bearer db614f9a334c9d20cac4515015cc57f5aae1c30c", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
