package main

import (
	"github.com/XiaoSanGit/bangumi_go_api/pkg/client/bgm_v0"
)

func main() {
	cli := bgm_v0.NewBgmClient("test")
	print(cli)
}
