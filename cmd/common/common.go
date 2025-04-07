package common

import (
	"net/http"
	"sync"
	"time"

	"github.com/Nathene/SyntheticMonitor/cmd/httpclient"
	"github.com/spf13/viper"
)

const (
	FlagAuthBasic = "auth.basic.password"

	FlagAuthDisabled    = "auth.disabled"
	DefaultAuthDisabled = false

	FlagDebug    = "debug"
	DefaultDebug = false

	DefaultServer = "http://localhost:9000/api/graphql/v1"

	DefaultTimeout = 10 * time.Second

	FlagServer = "server"
)

var (
	httpClientOnce sync.Once
	httpClient     *http.Client
)

var nickNames = map[string]string{
	"local": DefaultServer,
}

func DefaultHTTPClient() *http.Client {
	httpClientOnce.Do(func() {
		httpClient = httpclient.New().Client
	})
	return httpClient
}

func ServerURL() string {
	server := viper.GetString(FlagServer)

	if n, ok := nickNames[server]; ok {
		return n
	}
	return server
}
