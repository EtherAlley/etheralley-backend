package common

// These are intended to be generic settings that every app is expected to implement.
// packages in common can safely expect to have access to these settings
type IAppSettings interface {
	Appname() string
	Hostname() string
	Env() string
	IsDev() bool
	Port() string
}
