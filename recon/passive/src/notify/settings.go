package notify

var (
	s *settings
)

type settings struct {
	ErrorWHName   string
	LoggingWHName string
	ScanWHName    string
}

func Settings() *settings {
	if s == nil {
		s = &settings{
			ErrorWHName:   "RECON_DISCORD_WH_ERROR",
			LoggingWHName: "RECON_DISCORD_WH_LOGGING",
			ScanWHName:    "RECON_DISCORD_WH_SCAN",
		}
	}

	return s
}
