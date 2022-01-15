
## 这是一个对Bangumi开源API的go封装轮子库  An Basic Go Client of Bangumi OpenAPI - Simple kits

![](https://img.shields.io/github/go-mod/go-version/gohugoio/hugo)

[Bangumi OpenAPI](https://bangumi.github.io/api) \
**Go Client API in */pkg/client/bgm_v0* \
A function encapsulation based on *Bangumi-openAPI-v0***

---
###简单使用 - Simple Usage

```go
package main

import (
	"github.com/XiaoSanGit/bangumi_go_api/pkg/client/bgm_v0"
	"context"
	"fmt"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
)

func main() {
	cli := bgm_v0.NewBgmClient("test")
	ctx := context.Background()
	resp, err := cli.GetSubject(ctx, "", "311310")
	if err != nil {
		return
	}
	fmt.Printf("%s", common.JsonFormat(resp))
}
```

-- 更多测试和封装方法在pkg/client/bgm_v0，每个都有对应的单测 \
-- More Test and Functions in pkg/client/bgm_v0, with unit testing for each func.
