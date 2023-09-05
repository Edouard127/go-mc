package basic

// Settings of client
type Settings struct {
	Locale             string
	ViewDistance       int
	ChatMode           int
	ChatColors         bool
	DisplayedSkinParts uint8
	MainHand           int

	// Enables filtering of text on signs and written book titles.
	// Currently, always false (i.e. the filtering is disabled)
	EnableTextFiltering bool
	AllowListing        bool

	// The brand string presented to the server.
	Brand string
}

// Used by Settings.DisplayedSkinParts.
// For each bit set if shows match part.
const (
	_ = 1 << iota
	Jacket
	LeftSleeve
	RightSleeve
	LeftPantsLeg
	RightPantsLeg
	Hat
)

// DefaultSettings are the default settings of client
var DefaultSettings = Settings{
	Locale:             "en_us",
	ViewDistance:       32,
	ChatMode:           0,
	DisplayedSkinParts: Jacket | LeftSleeve | RightSleeve | LeftPantsLeg | RightPantsLeg | Hat,
	MainHand:           1,

	EnableTextFiltering: false,
	AllowListing:        true,

	Brand: "vanilla",
}
