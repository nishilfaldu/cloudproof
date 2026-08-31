package checks

import (
	"context"
	"sync"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

var (
	cfgOnce sync.Once
	awsCfg  aws.Config
	cfgErr  error
)

func sharedConfig(ctx context.Context) (aws.Config, error) {
	cfgOnce.Do(func() {
		awsCfg, cfgErr = config.LoadDefaultConfig(ctx)
	})
	return awsCfg, cfgErr
}
