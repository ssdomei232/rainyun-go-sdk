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

	if resp.StatusCode() == 200 {
		return err
	} else if resp.StatusCode() == 403 {
		var baseResp BaseResponse
		json.Unmarshal(resp.Body(), &baseResp)

		if baseResp.Code == CodeApikeyError {
			return fmt.Errorf("apikey error")
		}

		return fmt.Errorf("unknown error")
	} else if resp.StatusCode() == 400 {
		var baseResp BaseResponse
		json.Unmarshal(resp.Body(), &baseResp)

		if baseResp.Code == CodeInvalidInputParameter {
			return fmt.Errorf("invalid input parameter")
		}

		return fmt.Errorf("unknown error")
	} else {
		return fmt.Errorf("unknown error")
	}
}
