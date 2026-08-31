package checks

import (
	"context"
	"sync"
	"time"
	"github.com/nishilfaldu/cloudproof/internal/model"
)

type Check func(ctx context.Context) model.Finding

func RunAll(ctx context.Context, checks []Check) []model.Finding {
	out := make([]model.Finding, len(checks))
	var wg sync.WaitGroup
	for i, c := range checks {
		wg.Add(1)
		go func(i int, c Check) {
			defer wg.Done()
			cctx, cancel := context.WithTimeout(ctx, 3*time.Second)
			defer cancel()
			out[i] = c(cctx)
		}(i, c)
	}
	wg.Wait()
	return out
}