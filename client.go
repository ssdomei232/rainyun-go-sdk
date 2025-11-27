// package rainyun_go_sdk provide a full-sdk for rainyun.com
package rainyun_go_sdk

import (
	"encoding/json"
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	apikey string
	client *resty.Client
}

type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewClient(apiKey string) *Client {
	restyClient := resty.New()
	restyClient.SetBaseURL("https://api.v2.rainyun.com")
	restyClient.SetHeader("x-api-key", apiKey)
	restyClient.SetHeader("User-Agent", "rainyun-go-sdk/1.0")
	restyClient.SetHeader("Accept", "application/json")

	return &Client{
		apikey: apiKey,
		client: restyClient,
	}
}

// DoRequest 发起一次http请求
func (c *Client) DoRequest(method, endpoint string, reqData any, respData any) error {
	req := c.client.R()

	if reqData != nil {
		req.SetBody(reqData)
	}

	if respData != nil {
		req.SetResult(respData)
	}

	resp, err := req.Execute(method, endpoint)

	// 错误处理
	if resp.StatusCode() == 200 {
		return err
	} else {
		var baseResp BaseResponse
		json.Unmarshal(resp.Body(), &baseResp)
		return fmt.Errorf(fmt.Sprint(baseResp.Code))
	}
}
