package bgm_v0

import (
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"testing"
)

func TestClient_GetPersonById(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetPersonById(ctx, "", "3914")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetPersonRelatedSubject(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetPersonRelatedSubject(ctx, "", "3914")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetPersonRelatedCharacter(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetPersonRelatedCharacter(ctx, "", "3914")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
