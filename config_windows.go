package configdir

import "os"

var (
	systemConfig []string
	localConfig  string
	localCache   string
)

// For PortableApps.com configuration
func findPaths() {
	systemConfig = []string{os.Getenv("PORTABLE_PROGRAMDATA")}
	localConfig = os.Getenv("PORTABLE_APPDATA")
	localCache = os.Getenv("PORTABLE_LOCALAPPDATA")
}
