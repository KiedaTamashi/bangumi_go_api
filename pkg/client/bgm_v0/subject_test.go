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
	resp, err := cli.GetSubject(ctx, "", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetSubjectPersons(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetSubjectPersons(ctx, "", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetSubjectCharacters(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetSubjectCharacters(ctx, "", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetSubjectRelations(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetSubjectRelations(ctx, "", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}

func TestClient_GetCalendar(t *testing.T) {
	ctx := context.Background()
	cli := NewBgmClient("test")
	resp, err := cli.GetCalendar(ctx)
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
