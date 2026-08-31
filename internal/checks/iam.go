package checks

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/nishilfaldu/cloudproof/internal/model"
)

func IAMMFA(ctx context.Context) model.Finding {
	cfg, err := sharedConfig(ctx)
	if err != nil {
		return model.NewFinding("iam-mfa", "iam-users", "error", "cannot load AWS config: "+err.Error())
	}

	client := iam.NewFromConfig(cfg)

	users, err := client.ListUsers(ctx, &iam.ListUsersInput{})
	if err != nil {
		return model.NewFinding("iam-mfa", "iam-users", "error", "ListUsers failed: "+err.Error())
	}

	noMFA := 0
	for _, u := range users.Users {
		devs, err := client.ListMFADevices(ctx, &iam.ListMFADevicesInput{UserName: u.UserName})
		if err != nil {
			return model.NewFinding("iam-mfa", "iam-users", "error", "ListMFADevices failed: "+err.Error())
		}
		if len(devs.MFADevices) == 0 {
			noMFA++
		}
	}

	total := len(users.Users)
	if noMFA == 0 {
		return model.NewFinding("iam-mfa", "iam-users", "pass", fmt.Sprintf("all %d IAM users have MFA", total))
	}

	return model.NewFinding("iam-mfa", "iam-users", "fail", fmt.Sprintf("%d of %d IAM users have no MFA device", noMFA, total))

	// select {
	// case <-time.After(150 * time.Millisecond):
	// 	return model.NewFinding("iam-mfa", "iam-users", "fail", "2 of 7 users have no MFA device")
	// case <-ctx.Done():
	// 	return model.NewFinding("iam-mfa", "iam-users", "error", "check timed out")
	// }
}
