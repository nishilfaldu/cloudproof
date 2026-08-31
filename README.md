# CloudProof

A small, read-only AWS compliance scanner I built to learn how SOC 2 / continuous-compliance platforms (Vanta, Drata, Oneleet-style) work under the hood. This is a learning project, not a product.

It scans an AWS account against a few hygiene checks, normalizes the results into one finding model (control, status, severity, evidence, remediation), and renders them on a one-screen dashboard with a live re-scan button.

## Architecture

```
ui/ (Next.js)  --/api/findings-->  Go server :8080  -->  AWS (read-only)
```

- **Go server** (`cmd/server`, `internal/`): one endpoint, `GET /api/findings`. Each check runs in its own goroutine with a `sync.WaitGroup` and a 3-second `context.WithTimeout`, so a slow or stuck check can't hang the batch - it comes back as an `error` finding instead.
- **Finding model** (`internal/model`): every check returns the same shape - `control`, `status` (`pass` / `fail` / `error`), `severity`, `evidence`, `remediation`, `resource`, `checkedAt`. Severity and remediation text live in one `controlMeta` map.
- **UI** (`ui/`): Next.js dev server proxies `/api/*` to the Go server on `:8080` via `next.config.ts` rewrites. The page polls once on load, shows loading / error / populated states, and has a Re-scan button.

## Checks

| Control | What it looks for | Source |
| --- | --- | --- |
| `iam-mfa` | IAM console users with no MFA device (severity: critical) | Live AWS (`iam:ListUsers`, `iam:ListMFADevices`) |
| `s3-encryption-at-rest` | Buckets without default encryption (severity: high) | Fixture data (mocked while I build the real call) |
| `s3-public-access` | Buckets open to the public (severity: high) | Fixture data (mocked while I build the real call) |

## Running locally

Prerequisites:

- Go 1.26+
- Node.js + pnpm
- AWS credentials on the default provider chain (`~/.aws/credentials`, env vars, or SSO). Read-only is enough - the scanner only ever calls `List*` APIs.

```bash
# terminal 1 - API on :8080
go run ./cmd/server

# terminal 2 - UI on :3000
cd ui
pnpm install
pnpm dev
```

Open http://localhost:3000. The UI calls the Go API through the Next.js rewrite, so both processes need to be running.

You can also hit the API directly:

```bash
curl localhost:8080/api/findings | jq
```

## Layout

```
cmd/server/main.go      # entrypoint, serves /api/findings on :8080
internal/api/           # HTTP handler
internal/checks/        # one file per check + RunAll (concurrency, timeouts)
internal/model/         # Finding type + per-control severity/remediation
ui/                     # Next.js dashboard
```

## Notes

- Read-only by design: no mutating AWS calls, no auto-remediation.
- Error handling is part of the point: a check that fails or times out degrades to an `error` finding rather than crashing the scan.
