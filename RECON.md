# Recon — Donkey Kong Country: Tropical Freeze (Wii U)

Source: USA `rs10_production.rpx` (decompressed with `recon/unrpx.py`, zlib
sections) cross-referenced with kinnay's `nexwiiu.json`.

- NEX is **statically linked** (no `nn_nex` RPL). ~550 `nn::nex` symbols;
  protocol clients present: TicketGranting, SecureConnection, **DataStore**
  (heaviest), **Ranking**, AccountManagement, Health, Monitoring,
  Notification, NintendoNotificationEvent, RemoteLogDevice. **No MatchMaking /
  MatchmakeExtension client.** NAT Traversal engine linked (presence).
- Auth: `AcquireNexServiceToken(ACTNexAuthenticationResult*, u32 GameID)` with
  log string `GameID:%08x Success:%d AcquireNexServiceTokenResult:%08x`.
- Game server ID: code at file offset `0x710be8` builds `0x10144800`
  (`lis r6,0x1014; addi r6,r6,0x4800`) as a literal — matches kinnay
  `id 269764608`. Not derived from the title ID (USA = `0005000010137f00`).
- Game features hinting at DataStore/Ranking: Time Attack upload
  (`BeginTimeAttackUpload`, `isInLeaderboardReplay`), Miiverse posting
  (separate stack, out of scope).
- Kinnay: build `3_4_13_3_0`, branch `release/ngs/3.4.x.3`, key `7fcf384a`.

`recon/dkc.bin` / `strings.txt` are derived from the proprietary executable
and are git-ignored; regenerate with `python recon/unrpx.py <rpx> dkc.bin`.
