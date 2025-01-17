package ui

// Config contains TUI-specific configuration.
type Config struct {
	ShowAllFiles    bool
	Gopath          string `env:"GOPATH"`
	HomeDir         string `env:"HOME"`
	GlamourMaxWidth uint
	GlamourStyle    string

	// Which directory should we start from?
	WorkingDirectory string

	// Which document types shall we show?
	DocumentTypes DocTypeSet

	// For debugging the UI
	Logfile              string `env:"GLOW_LOGFILE"`
	HighPerformancePager bool   `env:"GLOW_HIGH_PERFORMANCE_PAGER" default:"true"`
	GlamourEnabled       bool   `env:"GLOW_ENABLE_GLAMOUR" default:"true"`

	// Collected notes
	CNSiteID     string `env:"GLOW_CN_SITE_ID"`
	CNVisibility string `env:"GLOW_CN_VISIBLITY" default:"public"`
	CNAuthHeader string `env:"GLOW_CN_AUTH_HEADER"`
}

func (c Config) showLocalFiles() bool {
	return c.DocumentTypes.Contains(LocalDoc)
}

func (c Config) localOnly() bool {
	return c.DocumentTypes.Equals(NewDocTypeSet(LocalDoc))
}

func (c Config) stashedOnly() bool {
	return c.DocumentTypes.Contains(StashedDoc) && !c.DocumentTypes.Contains(LocalDoc)
}
