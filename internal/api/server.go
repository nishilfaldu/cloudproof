package api

import (
	"encoding/json"
	"net/http"

	"github.com/nishilfaldu/cloudproof/internal/checks"
)

func FindingsHandler(w http.ResponseWriter, r *http.Request) {
	findings := checks.RunAll(r.Context(), []checks.Check{checks.S3Encryption, checks.IAMMFA})
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(findings)
}
