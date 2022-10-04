//go:build ydb && gocdk && tikv
// +build ydb,gocdk,tikv

package command

// set true if gtags are set
func init() {
	isFullVersion = true
}
