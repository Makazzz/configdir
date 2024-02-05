package configdir

import "os"

var (
	systemConfig []string
	localConfig  string
	localCache   string
)

// For a portable configuration with https://github.com/dweymouth/supersonic
// Part of code inspired by https://siongui.github.io/2020/01/29/go-check-if-environment-variable-exist/
func findPaths() {
	isEnvExist1("SUPERSONIC_PROGRAMDATA_DIR")
	isEnvExist2("SUPERSONIC_APPDATA_DIR")
	isEnvExist3("SUPERSONIC_LOCALAPPDATA_DIR")
}

func isEnvExist1(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		systemConfig = []string{os.Getenv("SUPERSONIC_PROGRAMDATA_DIR")}
		return true
	}
	systemConfig = []string{os.Getenv("PROGRAMDATA")}
	return false
}
func isEnvExist2(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		localConfig = os.Getenv("SUPERSONIC_APPDATA_DIR")
		return true
	}
	localConfig = os.Getenv("APPDATA")
	return false
}
func isEnvExist3(key string) bool {
	if _, ok := os.LookupEnv(key); ok {
		localCache = os.Getenv("SUPERSONIC_LOCALAPPDATA_DIR")
		return true
	}
	localCache = os.Getenv("LOCALAPPDATA")
	return false
}
