package ssl

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/v2/rainyun/common"
)

// 获取SSL证书列表
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetSSLCertificateList(options string) (*SslCertificateList, error) {
	path := "/product/sslcenter/?options=" + options

	var resp SslCertificateList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 上传SSL证书
//
// cert: 证书
//
// key： 私钥
func (c *Client) UploadSSLCertificate(cert string, key string) (*common.BasicOperationResponse, error) {
	path := "/product/sslcenter/"

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, SslCertificate{
		Cert: cert,
		Key:  key,
	}, &resp)

	return &resp, err
}

// 获取SSL证书详情
//
// id: SSL证书ID
func (c *Client) GetSslDetail(id int) (*SslDetail, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp SslDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 替换SSL证书
//
// id: SSL证书ID
//
// cert: 证书
//
// key： 私钥
func (c *Client) ReplaceSsl(id int, cert, key string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PUT", path, SslCertificate{
		Cert: cert,
		Key:  key,
	}, &resp)

	return &resp, err
}

// 删除SSL证书
//
// id: SSL证书ID
func (c *Client) DeleteSsl(id int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/sslcenter/%d", id)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}
