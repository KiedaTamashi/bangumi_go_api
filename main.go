package main

import (
	"github.com/XiaoSanGit/bangumi_go_api/pkg/client/bgm"
)

func main() {
	cli := bgm.NewClient("test")
	print(cli)
}
