package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestClient_SearchMediumSubjectByKeywords(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")

	resp, err := cli.SearchMediumSubjectByKeywords(ctx, "", "人形电脑天使心", 0, "", 0, 5)
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
