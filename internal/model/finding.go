package model

import "time"

type Finding struct {
	Control     string    `json:"control"`
	Status      string    `json:"status"` // pass, fail, error
	Severity    string    `json:"severity"`
	Evidence    string    `json:"evidence"`
	Remediation string    `json:"remediation"`
	CheckedAt   time.Time `json:"checkedAt"`
	Resource    string    `json:"resource"`
}

var controlMeta = map[string]struct {
	Severity    string
	Remediation string
}{
	"s3-public-access": {
		"high", "Block public access at the bucket or account level",
	},
	"iam-mfa": {
		"critical", "Enforce MFA via IAM policy for all console users",
	},
	"s3-encryption-at-rest": {
		"high", "Enable default encryption (SSE-S3 or SSE-KMS) on all buckets",
	},
	// one entry per check
}

func NewFinding(control, resource, status, evidence string) Finding {
	meta, ok := controlMeta[control]
	if !ok {
		meta = struct {
			Severity    string
			Remediation string
		}{"unknown", ""}
	}
	return Finding{
		Control:     control,
		Status:      status,
		Severity:    meta.Severity,
		Evidence:    evidence,
		Remediation: meta.Remediation,
		Resource:    resource,
		CheckedAt:   time.Now(),
	}
}
