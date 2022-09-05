package types

// used as constants for asset platforms
type AssetPlatform int

const (
	Ethereum AssetPlatform = iota
	Solana
	Fantoms
	Polygon
	Avalanche
)

func (ap AssetPlatform) String() string {
	return []string{"ethereum", "solana", "fantom", "polygon-pos", "avalanche"}[ap]
}
