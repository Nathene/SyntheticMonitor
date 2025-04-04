package types

import (
	"fmt"
	"slices"
)

type ProbeStatus string

const (
	Enabled  ProbeStatus = "enabled"
	Disabled ProbeStatus = "disabled"
)

func (ProbeStatus) Values() (kinds []string) {
	for _, s := range []ProbeStatus{Enabled, Disabled} {
		kinds = append(kinds, string(s))
	}
	return
}

func (ps ProbeStatus) Validate() error {
	if slices.Contains(ps.Values(), string(ps)) {
		return nil
	}
	return fmt.Errorf("invalid probe status: %s", ps)
}

type ProbeType string

const (
	HTTP ProbeType = "http"
	SQL  ProbeType = "sql"
	PING ProbeType = "ping"
)

func (ProbeType) Values() (kinds []string) {
	for _, s := range []ProbeType{HTTP, SQL, PING} {
		kinds = append(kinds, string(s))
	}
	return
}

func (pt ProbeType) Validate() error {
	if slices.Contains(pt.Values(), string(pt)) {
		return nil
	}
	return fmt.Errorf("invalid probe type: %s", pt)
}
