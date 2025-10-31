package rainyun_go_sdk

import (
	"fmt"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	apikey string
	client *resty.Client
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

func (c *Client) Get(endpoint string, params map[string]interface{}) (*resty.Response, error) {
	stringParams := make(map[string]string)
	for k, v := range params {
		stringParams[k] = fmt.Sprintf("%v", v)
	}

	return c.client.R().
		SetQueryParams(stringParams).
		Get(endpoint)
}

func (c *Client) Post(endpoint string, data interface{}) (*resty.Response, error) {
	return c.client.R().
		SetBody(data).
		Post(endpoint)
}

func (c *Client) Put(endpoint string, data interface{}) (*resty.Response, error) {
	return c.client.R().
		SetBody(data).
		Put(endpoint)
}

func (c *Client) Delete(endpoint string) (*resty.Response, error) {
	return c.client.R().
		Delete(endpoint)
}

func (c *Client) DoRequest(method, endpoint string, reqData interface{}, respData interface{}) error {
	req := c.client.R()

	if reqData != nil {
		req.SetBody(reqData)
	}

	if respData != nil {
		req.SetResult(respData)
	}

	_, err := req.Execute(method, endpoint)
	return err
}
