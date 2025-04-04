package probehttp

type HTTPConfig struct {
	URL          string            `json:"url"`                     // Required
	Method       string            `json:"method,omitempty"`        // Default GET
	ExpectedCode int               `json:"expected_code,omitempty"` // Default 200
	Headers      map[string]string `json:"headers,omitempty"`
	Body         string            `json:"body,omitempty"`
	Timeout      string            `json:"timeout,omitempty"` // e.g., "5s"
}
