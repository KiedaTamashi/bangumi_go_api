package bgm_v0

import (
	"context"
	"encoding/json"
	"fmt"
	mError "github.com/XiaoSanGit/bangumi_go_api/model/error"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/client/httpcli"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/common"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/errno"
	"github.com/XiaoSanGit/bangumi_go_api/pkg/types"
	"net/http"
	"strings"
)

var logger = types.InitBasicLogger("bgmLog.log")

type Client struct {
	Caller string

	host            string
	proxyUrl        string
	cli             *http.Client
	concurrencyChan chan bool
	skip            bool
}

func NewBgmClient(caller string) *Client {
	return NewClinetConcurrency(0, caller)
}

type respType struct {
	Total  int64           `json:"total,omitempty"`
	Limit  int64           `json:"limit,omitempty"`
	Offset int64           `json:"offset,omitempty"`
	Data   json.RawMessage `json:"data"`
}

func NewClinetConcurrency(chanLength uint, caller string) *Client {
	return NewClientWithMaxConcurrency(chanLength, caller)
}

func NewClientWithMaxConcurrency(c uint, caller string) *Client {
	cli := &Client{
		Caller: caller,
		host:   "https://api.bgm.tv", //v0
		//proxyUrl: "https://api.bgm.tv/v0/",
		cli: &http.Client{},
	}
	if c > 0 {
		cli.concurrencyChan = make(chan bool, int(c))
	}
	return cli
}

func (cli *Client) call(ctx context.Context, method, url string,
	params map[string]string, headers map[string]string, body []byte, retryCount uint) (int, string, error) {
	if params == nil {
		params = make(map[string]string)
	}
	//params["query"] = url
	//params["caller"] = cli.Caller
	//return httpcli.HttpWithCli(ctx, cli.cli, method, cli.proxyUrl, params, headers, body, retryCount)
	return httpcli.HttpWithCli(ctx, cli.cli, method, url, params, headers, body, retryCount)
}

func (cli *Client) callEncodedUrl(ctx context.Context, method, url string,
	params map[string]string, headers map[string]string, body map[string]string, retryCount uint) (int, string, error) {
	if params == nil {
		params = make(map[string]string)
	}
	//params["query"] = url
	//params["caller"] = cli.Caller
	//return httpcli.HttpWithCli(ctx, cli.cli, method, cli.proxyUrl, params, headers, body, retryCount)
	return httpcli.HttpFromUrlEncode(ctx, cli.cli, method, url, params, headers, body, retryCount)
}

func (cli *Client) callFromUrlEncode(ctx context.Context, method, absolutePath, authToken string, retry uint, in interface{}, out interface{}) error {
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	if authToken != "" {
		headers["Authorization"] = authToken
	}

	if cli.concurrencyChan != nil {
		cli.concurrencyChan <- true
		defer func() {
			if cli.concurrencyChan != nil {
				<-cli.concurrencyChan
			}
		}()
	}
	if len(headers["Content-Type"]) == 0 {
		headers["Content-Type"] = "application/json"
	}
	for {
		body, ok := in.(map[string]string)
		if !ok {
			return errno.Errorf(errno.ErrBadRequest, "in is not map[string]string!")
		}
		statusCode, content, err := cli.callEncodedUrl(ctx, method, cli.host+absolutePath, //content ???body raw string
			nil, headers, in.(map[string]string), 1)
		if err != nil {
			return err
		}
		// ?????????????????????????????????
		if statusCode != 200 {
			var httpErrorReason string
			if statusCode == 404 {
				httpErrorReason = "Not Found"
			} else if statusCode == 422 {
				httpErrorReason = "Validation Error"
				httpErrorRsp := mError.ValidationError{}
				if err = json.Unmarshal([]byte(content), &httpErrorRsp); err == nil {
					//todo [refine] ??????????????????????????????????????????
				}
			}
			err = errno.Errorf(errno.ErrInternalServer, "http [%s] error, path: %s, request body: %s, status_code: %d, response body: %s",
				httpErrorReason, absolutePath, common.Json(body), statusCode, content)
			logger.Error("%v", err)
			return err
		}
		resp := &respType{}
		// todo [refine] ???????????????????????????????????????
		if err := json.Unmarshal([]byte(content), &resp); err != nil || resp.Data == nil {
			resp = &respType{
				Data: []byte(content),
			}
		}
		if out != nil {
			if err := json.Unmarshal(resp.Data, out); err != nil {
				err = errno.Errorf(errno.ErrInternalServer, "json.Unmarshal error, json string: %s", resp.Data)
				logger.Error("%v", err)
				return err
			}
		}
		return nil
	}
}

func (cli *Client) Call(ctx context.Context, method, absolutePath string,
	param map[string]string, headers map[string]string, retry uint, body []byte, out interface{}) error {
	if cli.concurrencyChan != nil {
		cli.concurrencyChan <- true
		defer func() {
			if cli.concurrencyChan != nil {
				<-cli.concurrencyChan
			}
		}()
	}

	if headers == nil {
		headers = make(map[string]string)
	}
	if len(headers["Content-Type"]) == 0 {
		headers["Content-Type"] = "application/json"
	}

	//retryN := uint(0)
	for {
		statusCode, content, err := cli.call(ctx, method, cli.host+absolutePath, //content ???body raw string
			param, headers, body, 1)
		if err != nil {
			return err
		}
		// ?????????????????????????????????
		if statusCode != 200 {
			var httpErrorReason string
			if statusCode == 404 {
				httpErrorReason = "Not Found"
			} else if statusCode == 422 {
				httpErrorReason = "Validation Error"
				httpErrorRsp := mError.ValidationError{}
				if err = json.Unmarshal([]byte(content), &httpErrorRsp); err == nil {
					//todo [refine] ??????????????????????????????????????????
				}
			}
			err = errno.Errorf(errno.ErrInternalServer, "http [%s] error, path: %s, request body: %s, status_code: %d, response body: %s",
				httpErrorReason, absolutePath, string(body), statusCode, content)
			logger.Error("%v", err)
			return err
		}
		// ????????????????????????????????????
		//var resp *struct {
		//	//Message string `json:"message"`
		//	//Code    int    `json:"code"`
		//	//Data    struct {
		//	//	StatusCode int    `json:"status_code"`
		//	//	RespData   string `json:"resp_data"`
		//	//} `json:"data"`
		//	Data 	string `json:"data"`
		//}
		resp := &respType{}
		// todo [refine] ???????????????????????????????????????
		if err := json.Unmarshal([]byte(content), &resp); err != nil || resp.Data == nil {
			//if err != nil {
			//	err = errno.Errorf(errno.ErrInternalServer, "json.Unmarshal error, json body: %s", content)
			//	logger.Error("%v", err)
			//}
			resp = &respType{
				Data: []byte(content),
			}
		}

		// todo [refine] ??????????????????????????????????????? Token??????????????????
		//if retryN < retry && resp.Code != 0 {
		//	logger.Debug("respond code = %d, message = %s, retry", resp.Code, resp.Message)
		//	retryN++
		//	continue
		//}
		if out != nil {
			if err := json.Unmarshal(resp.Data, out); err != nil {
				err = errno.Errorf(errno.ErrInternalServer, "json.Unmarshal error, json string: %s", resp.Data)
				logger.Error("%v", err)
				return err
			}
		}
		return nil
	}
}

func (cli *Client) callJson(ctx context.Context, method, absolutePath, authToken string, retry uint, in interface{}, out interface{}) error {
	var body []byte
	if in != nil {
		var err error
		body, err = json.Marshal(in)
		if err != nil {
			err = errno.Errorf(errno.ErrInternalServer, "json.Marshal error")
			logger.Error("%v", err)
			return err
		}
	}
	headers := make(map[string]string)
	headers["Content-Type"] = "application/x-www-form-urlencoded"
	if authToken != "" {
		headers["Authorization"] = authToken
	}
	return cli.Call(ctx, method, absolutePath, nil, headers, retry, body, out)
}

//POST fixme [auth] token???????????????????????????bgm???token?????????????????????????????????????????????????????????????????????????????????token??????????????????func
func (cli *Client) POST(ctx context.Context, absolutePath, authToken string, retry uint, param map[string]string, in map[string]string, out interface{}, isUrlEncoded bool) error {
	if cli.skip {
		return nil
	}
	if param != nil {
		absolutePath = fmt.Sprintf("%s?", absolutePath)
		for k, v := range param {
			if v == "" {
				continue
			}
			absolutePath = fmt.Sprintf("%s%s=%s&", absolutePath, k, v)
		}
		absolutePath = strings.TrimSuffix(absolutePath, "&")
	}
	if isUrlEncoded {
		return cli.callFromUrlEncode(ctx, "POST", absolutePath, authToken, retry, in, out)
	} else {
		return cli.callJson(ctx, "POST", absolutePath, authToken, retry, in, out)
	}
}

func (cli *Client) GET(ctx context.Context, absolutePath, authToken string, retry uint, param map[string]string, in interface{}, out interface{}) error {
	if cli.skip {
		return nil
	}
	if param != nil {
		absolutePath = fmt.Sprintf("%s?", absolutePath)
		for k, v := range param {
			absolutePath = fmt.Sprintf("%s%s=%s&", absolutePath, k, v)
		}
		absolutePath = strings.TrimSuffix(absolutePath, "&")
	}
	return cli.callJson(ctx, "GET", absolutePath, authToken, retry, []byte{}, out)
}
