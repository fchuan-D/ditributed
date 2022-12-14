package registry

type Registration struct {
	ServiceName      ServiceName
	ServiceURL       string
	RequireServices  []ServiceName
	ServiceUpdateURL string
}

type ServiceName string

const (
	LogService     = ServiceName("LogService")
	GradingService = ServiceName("GradingService")
	PortalService  = ServiceName("PortalService")
)

type patchEntry struct {
	Name ServiceName
	URL  string
}

type patch struct {
	Added   []patchEntry
	Removed []patchEntry
}
