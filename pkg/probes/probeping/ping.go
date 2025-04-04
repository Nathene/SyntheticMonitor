package probeping

type PingConfig struct {
	Host    string `json:"host"`              // Required
	Count   int    `json:"count,omitempty"`   // Default e.g. 3
	Timeout string `json:"timeout,omitempty"` // e.g., "1s"
}
