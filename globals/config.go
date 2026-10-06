package globals

// NEX configuration for "Donkey Kong Country: Tropical Freeze" (Wii U,
// Retro Studios / Nintendo, 2014).
//
// GameServerID and AccessKey come from kinnay.github.io's public Wii U NEX
// database (id 269764608). The game server ID is a compile-time constant in
// the retail executable (rs10_production.rpx: lis/addi 0x1014,0x4800 feeding
// the NEX init), NOT derived from the per-region title ID (the USA title is
// 0005000010137f00), so every region uses the same ID.
const (
	GameServerID = "10144800" // 269764608
	AccessKey    = "7fcf384a"

	// kinnay build tag "build:3_4_13_3_0" (origin/release/ngs/3.4.x.3), the
	// same NEX build as Trine 2 / Sonic Lost World. Ranking is bumped to
	// 3.6.0 at runtime (see nex/secure.go), as on the Mario & Sonic servers.
	NEXMajor = 3
	NEXMinor = 4
	NEXPatch = 13
)
