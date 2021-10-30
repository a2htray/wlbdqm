package wlbdqm

type AppMode uint

var appMode = AppModeDebug

const (
	AppModeDebug AppMode = iota
	AppModeProduction
)

func SetAppMode(mode AppMode) {
	appMode = mode
}
