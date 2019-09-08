package handlers

import "net/http"

// swagger:route GET /healthz healthz healthz
//
// Check service health
func LivenessHandler(w http.ResponseWriter, r *http.Request) {}
