package alicloud

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/services/ecs"
)

// AccessKey ...
type AccessKey struct {
	RegionID string `mapstructure:"region-id"`
	AccessKeyID string `mapstructure:"access-key-id"`
	AccessKeySecret string `mapstructure:"access-key-secret"`
}

// ECSClient ...
type ECSClient struct {
	*ecs.Client
}

// NewECSClient ...
func NewECSClient(opt *AccessKey) (*ECSClient, error) {
	client, err := ecs.NewClientWithAccessKey(opt.RegionID, opt.AccessKeyID, opt.AccessKeySecret)
	if err != nil {
		return nil, err
	}
	return &ECSClient{Client: client}, nil
}
