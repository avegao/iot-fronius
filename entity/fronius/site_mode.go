package fronius

const (
	SiteModeProduceOnly   SiteMode = "produce-only"
	SiteModeMeter         SiteMode = "meter"
	SiteModeVagueMeter    SiteMode = "vague-meter"
	SiteModeBidirectional SiteMode = "bidirectional"
	SiteModeAcCoupled     SiteMode = "ac-coupled"
)

var (
	SiteModeMap = map[string]SiteMode{
		"produce-only":  SiteModeProduceOnly,
		"meter":         SiteModeMeter,
		"vague-meter":   SiteModeVagueMeter,
		"bidirectional": SiteModeBidirectional,
		"ac-coupled":    SiteModeAcCoupled,
	}
)

type SiteMode string

func (mode SiteMode) String() string {
	return string(mode)
}

func NewSiteModeFromString(modeString string) SiteMode {
	return SiteModeMap[modeString]
}
