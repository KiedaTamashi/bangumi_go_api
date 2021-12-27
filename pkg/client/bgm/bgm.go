package bgm

import "net/http"

type BgmClient struct {
	Host     string
	httpCli  *http.Client
	SkipPost bool
}

func NewBgmClient() *BgmClient {
	res := &BgmClient{
		Host:     "https://api.bgm.tv/v0/",
		SkipPost: false,
		httpCli:  &http.Client{},
	}
	return res
}

type Client struct {
	Caller string

	host            string
	proxyUrl        string
	cli             *http.Client
	concurrencyChan chan bool
	skip            bool
}

func NewClient(caller string) *Client {
	return NewClientWithMaxConcurrency(0, caller)
}

func NewClinetConcurrency(chanLength uint, caller string) *Client {
	return NewClientWithMaxConcurrency(chanLength, caller)
}

func NewClientWithMaxConcurrency(c uint, caller string) *Client {
	cli := &Client{
		Caller: caller,
		host:   "https://api.bgm.tv/v0/",
		//proxyUrl: "https://api.bgm.tv/v0/",
		cli: &http.Client{},
	}
	if c > 0 {
		cli.concurrencyChan = make(chan bool, int(c))
	}
	return cli
}
