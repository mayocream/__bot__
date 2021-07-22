package alicloud

import (
	"github.com/pkg/errors"

	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// GetInstances ...
func (client *ECSClient) GetInstances() (*ecs.DescribeInstancesResponse, error) {
	request := ecs.CreateDescribeInstancesRequest()
	request.Scheme = "https"

	response, err := client.DescribeInstances(request)
	if err != nil {
		return nil, err
	}
	if !response.IsSuccess() {
		return nil, errors.New(response.String())
	}
	return response, nil
}

// StopCharging ...
func (client *ECSClient) StopCharging() error {
	request := ecs.CreateStopInstanceRequest()
	request.Scheme = "https"

	request.ConfirmStop = requests.NewBoolean(false)
	request.StoppedMode = "StopCharging"

	response, err := client.StopInstance(request)
	if err != nil {
		return err
	}
	if !response.IsSuccess() {
		return errors.New(response.String())
	}
	return nil
}

