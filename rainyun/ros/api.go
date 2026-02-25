package ros

import (
	"fmt"

	"github.com/ssdomei232/rainyun-go-sdk/rainyun/common"
)

// 创建对象存储实例
//
// duration: 购买时长（月）
//
// planID: 套餐ID
//
// withCouponID: 使用的优惠券ID
func (c *Client) CreateRosInstance(duration int, planID int, withCouponID int) (*CreateRosInstanceResponse, error) {
	path := "/product/ros/instance"

	var resp CreateRosInstanceResponse
	err := c.DoRequest("POST", path, CreateRosInstanceRequest{
		Duration:     duration,
		PlanID:       planID,
		WithCouponID: withCouponID,
	}, &resp)

	return &resp, err
}

// 获取对象存储实例详情
//
// instanceID: 实例ID
func (c *Client) GetRosInstanceDetail(instanceID int) (*RosInstanceDetail, error) {
	path := fmt.Sprintf("/product/ros/instance/%d", instanceID)

	var resp RosInstanceDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取对象存储桶列表
//
// ⚠️注意这个接口是拿不到AK和SK的，响应里面的AK和SK都是空的
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRosBucketList(options string) (*RosBucketList, error) {
	path := fmt.Sprintf("/product/ros/bucket/?options=%s", options)

	var resp RosBucketList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取对象存储实例列表
//
// ⚠️注意这个接口是拿不到AK和SK的，响应里面的AK和SK都是空的
//
// options: 查询参数 可以用 EncodingStandardQueryParameters 获取.
func (c *Client) GetRosInstanceList(options string) (*RosInstanceList, error) {
	path := fmt.Sprintf("/product/ros/instance/?options=%s", options)

	var resp RosInstanceList
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 创建对象存储桶
//
// bucketName 存储桶名（>= 3 字符 <= 63 字符）
//
// instanceID 实例ID
func (c *Client) CreateRosBucket(bucketName string, instanceID int) (*CreateRosBucketResponse, error) {
	path := "/product/ros/bucket"

	var resp CreateRosBucketResponse
	err := c.DoRequest("POST", path, CreateRosBucketRequest{
		BucketName: bucketName,
		InstanceID: instanceID,
	}, &resp)

	return &resp, err
}

// 获取对象存储桶详情
//
// bucketID 存储桶ID
func (c *Client) GetRosBucketDetail(bucketID int) (*RosBucketDetail, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d", bucketID)

	var resp RosBucketDetail
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 删除对象存储桶
//
// bucketID 存储桶ID
func (c *Client) DeleteRosBucket(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d", bucketID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("DELETE", path, nil, &resp)

	return &resp, err
}

// 获取对象存储桶监控数据
//
// bucketID： 存储桶ID
//
// startdate：开始时间（timestamp）
//
// enddate：结束时间（timestamp）
func (c *Client) GetRosBucketMonitorData(bucketID int, startdate int, enddate int) (*RosMonitorData, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/monitor?start_date=%d&end_date=%d", bucketID, startdate, enddate)

	var resp RosMonitorData
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 获取对象存储实例监控数据
//
// instanceID： 实例ID
//
// startdate：开始时间（timestamp）
//
// enddate：结束时间（timestamp）
func (c *Client) GetRosInstanceMonitorData(instanceID int, startdate int, enddate int) (*RosMonitorData, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/monitor?start_date=%d&end_date=%d", instanceID, startdate, enddate)

	var resp RosMonitorData
	err := c.DoRequest("GET", path, nil, &resp)

	return &resp, err
}

// 修改存储桶Proxy设置
//
// bucketID：存储桶ID
func (c *Client) ModifyRosBucketProxySettings(bucketID int, settings *ModifyRosBucketProxySettingsRequest) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/proxy", bucketID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, settings, &resp)

	return &resp, err
}

// 对象存储桶重新生成密钥
//
// bucketID：存储桶ID
func (c *Client) ReGenerateRosBucketKeys(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/regenerate-keys", bucketID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 开关对象存储桶匿名访问
//
// bucketID：存储桶ID
func (c *Client) SetRosBucketPublicAccess(bucketID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/bucket/%d/toggle-public-access", bucketID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 开关对象实例桶匿名访问
//
// instanceID：实例ID
func (c *Client) SetRosInstancePublicAccess(instanceID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/toggle-public-access", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 对象存储实例重新生成密钥
//
// instanceID：实例ID
func (c *Client) ReGenerateRosInstanceKeys(instanceID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/regenerate-keys", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, nil, &resp)

	return &resp, err
}

// 对象存储实例续费
//
// instanceID：实例ID
//
// duratoin： 续费时长(月)
//
// couponID: 优惠券ID
func (c *Client) RenewRosInstance(instanceID int, duration int, couponID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/renew", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, RenewRosInstanceRequest{Duration: duration, WithCouponID: couponID}, &resp)

	return &resp, err
}

// 对象存储实例自动续费选项
//
// instanceID: 实例ID
//
// isOpen: 是否开启自动续费
func (c *Client) SetRosInstanceAutoRenewOption(instanceID int, isOpen bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/renew/option", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, RosInstanceAutoRenewOption{
		AutoRenewOption: isOpen,
	}, &resp)

	return &resp, err
}

// ROS实例缩放
//
// instanceID: 实例ID
//
// desrPlan: 目标套餐ID
//
// couponID：优惠券ID
func (c *Client) ScaleRosInstance(instanceID int, desrPlan int, couponID int) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/scale", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("POST", path, ScaleRosInstanceRequest{
		DestPlan:     desrPlan,
		WithCouponID: couponID,
	}, &resp)

	return &resp, err
}

// 设置对象存储实例标签
//
// instanceID: 实例ID
//
// tag： 标签
func (c *Client) SetRosInstanceTags(instanceID int, tag string) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/tag", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, SetRosInstnceTagRequest{TagName: tag}, &resp)

	return &resp, err
}

// 开关对象存储实例的弹性计费选项
//
// instanceID： 实例ID
//
// isEnable： 是否开启
func (c *Client) ToggleRosInstanceExtraAccounting(instanceID int, isEnable bool) (*common.BasicOperationResponse, error) {
	path := fmt.Sprintf("/product/ros/instance/%d/toggle-extra-accounting", instanceID)

	var resp common.BasicOperationResponse
	err := c.DoRequest("PATCH", path, ToggleRosInstanceExtraAccountingRequest{IsEnable: isEnable}, &resp)

	return &resp, err
}
