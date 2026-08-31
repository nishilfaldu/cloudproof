package checks

import (
	"context"
	"time"

	"github.com/nishilfaldu/cloudproof/internal/model"
)

func S3Encryption(ctx context.Context) model.Finding {
	select {
	case <-time.After(200 * time.Millisecond):
		return model.NewFinding("s3-encryption-at-rest", "all-buckets", "pass", "3 of 3 buckets have SSE-S3 enabled")
	case <-ctx.Done():
		return model.NewFinding("s3-encryption-at-rest", "all-buckets", "error", "check timed out")
	}
}

func S3PublicAccess(ctx context.Context) model.Finding {
	select {
	case <-time.After(200 * time.Millisecond):
		return model.NewFinding("s3-public-access", "all-buckets", "fail", "1 of 3 buckets allows public access")
	case <-ctx.Done():
		return model.NewFinding("s3-public-access", "all-buckets", "error", "check timed out")
	}
}
