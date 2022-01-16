
## 这是一个对Bangumi开源API的go封装轮子库 An Basic Go Client of Bangumi OpenAPI - Simple kits

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
Return
```
{
"id": 311310,
"type": 2,
"name": "海賊王女",
"name_cn": "海贼王女",
"summary": "これは、ある少女の記憶から始まる物語。\r\n18世紀、大西洋。\r\n\r\n父と船旅へ出ていたフェナ・ハウトマンは海賊に襲われ、\r\nたった一人小型ボートで漂流し、命をつなぎ留める。\r\nフェナが漂着したのは、国家が黙認する娼婦・男娼の島\u003cシャングリラ\u003eだった。\r\n\r\n10年後──。\r\n雪のような肌と白銀に光る髪を持ち、\r\n美しく成長したフェナは、初めての“仕事”を目前に控えていた。\r\nだが、それを受け入れることはできず、\r\n幾度となく想像してきた島からの脱出を決心する。\r\n\r\n迫りくる追っ手に絶体絶命のフェナは、\r\n真っ赤な鎧に鹿の角の兜をまとった青年・雪丸に救われる。\r\n雪丸は、フェナを「見つけ出す」と約束した少年だった。\r\nそして2 人の“再会”は、フェナ自身に眠っていた言葉\u003cエデン\u003eを呼び起こす。\r\n\r\n炎に包まれ、沈みゆく船。\r\n「必ず俺が見つけ出す！」と約束した少年。\r\nそして最愛の父が叫んだ、あの言葉──。\r\n\r\n「\u003cエデン\u003eに向かえ！」\r\n\r\nフェナはその真意を知るべく、\r\n雪丸たちと共に\u003cエデン\u003eの謎を解く船旅へ出る。",
"date": "2021-10-02T00:00:00+08:00",
"platform": "TV",
    "images": {
    "large": "https://lain.bgm.tv/pic/cover/l/2a/8e/311310_1NXnu.jpg",
    "common": "https://lain.bgm.tv/pic/cover/c/2a/8e/311310_1NXnu.jpg",
    "medium": "https://lain.bgm.tv/pic/cover/m/2a/8e/311310_1NXnu.jpg",
    "small": "https://lain.bgm.tv/pic/cover/s/2a/8e/311310_1NXnu.jpg",
    "grid": "https://lain.bgm.tv/pic/cover/g/2a/8e/311310_1NXnu.jpg"
    }
......
```

-- 更多测试和封装方法在pkg/client/bgm_v0，每个都有对应的单测 \
-- More Test and Functions in pkg/client/bgm_v0, with unit testing for each func.
---
##TODO LIST
- [x] BaseClient/Common Func
- [x] Users
- [x] Subject
- [x] Search
- [x] Collection
- [ ] Chapter
- [ ] Character
- [ ] Persons
- [ ] EditHistory
- [ ] Schema Definition According to Bangumi OpenAPI
- [ ] Refine