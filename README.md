# Donkey Kong Country: Tropical Freeze (Wii U) — NEX server

A preservation NEX server for **Donkey Kong Country: Tropical Freeze**
(Retro Studios / Nintendo, 2014), `game_server_id` `10144800`. It speaks the
game's PRUDP authentication + secure protocols and backs the online features
the retail executable links: **Ranking** (time-attack leaderboards) and
**DataStore** (ghost / replay data attached to scores).

Built on the [Pretendo Network](https://github.com/PretendoNetwork) NEX
libraries (`nex-go`, `nex-protocols-go`, `nex-protocols-common-go`), templated
from the Hyrule Warriors / Mario & Sonic servers in this org.
`internal/nex-protocols-common-go-patch/` is a vendored fork carrying the
DataStore/S3 changes the object-upload flow needs.

Evidence for each constant: [RECON.md](RECON.md). Per-protocol status:
[PROTOCOL_COVERAGE.md](PROTOCOL_COVERAGE.md).

## Recovered configuration

| Field | Value | Source |
|---|---|---|
| Game server ID | `10144800` (269764608) — **all regions** | kinnay [`nexwiiu.json`](https://kinnay.github.io/data/nexwiiu.json); also a hard-coded constant in the USA RPX |
| Access key | `7fcf384a` | kinnay `key` — **confirmed on hardware** |
| NEX version | `3.4.13` | kinnay `build:3_4_13_3_0` |
| `UseStructureHeader` | `false` | **confirmed on hardware** (LoginEx decodes) |
| `LegacyConnectionSignature` | `false` | **confirmed on hardware** — `true` gives `106-0502` on this title even though it is NEX 3.4.13 |
| Ranking library | bumped to `3.6.0` | as on the Mario & Sonic 3.4.x servers; leaderboard rows decode on hardware |

The game server ID is **not** the title ID (USA title is
`0005000010137f00`) and is not derived from it at runtime, so JPN / EUR / USA
consoles all request `10144800`. The account server's `nex_token` route must
therefore map that one ID, with **no title-ID allow-list**.

## Scope

Ticket Granting, Secure Connection, Utility, Ranking, DataStore. NAT
Traversal and MatchMaking are linked in the RPX only for the NEX presence
stack; the game has no online multiplayer, so they are not registered.

## Player names ("Identifiant" column)

The game never uploads a name: the leaderboard column is read from the
ranking *common data* (a NUL-terminated `char` string), and `GetCommonData`
is called for it. This server answers both from a per-PID name file
(`PN_DKCTF_NAMES_DIR`, default `/names`; one file per PID containing the
PNID). Those files are written by the account/NEX-token service —
[`account-server`](https://github.com/Protarium-Network/account-server) with
`PN_ACCOUNT_NAMES_DIR` set — from the `user_id` in the console's own profile
at every `nex_token` request. Without that service, names stay blank.
`docker compose` mounts the directory via `PN_DKCTF_NAMES_HOST_DIR`.

## Running

```bash
cp .env.example .env                     # PN_DKCTF_LOCAL_MODE=1 by default
cp settings.example.json settings.json   # your console's PID + NEX password
docker compose up --build                # UDP 28000 (auth) / 28001 (secure)
```

Local mode has no account server: NEX passwords come from `settings.json` and
the token is not validated — isolated networks only. Set
`PN_DKCTF_SECURE_HOST` to an address the console can reach (**short**, ~15
chars max: the retail binary truncates it).

**Shared mode:** `PN_DKCTF_LOCAL_MODE=0` plus `PN_DKCTF_NEX_TOKEN_AES_KEY`
(64 hex) and `PN_DKCTF_NEX_PASSWORD_SECRET` (≥32 bytes hex) matching your
account server; password = `HMAC-SHA256(secret, pid)`.

**Without Docker:** set the `PN_DKCTF_*` variables from `.env.example`, then
`go build -o dkctf-nex . && ./dkctf-nex`.

Optional `PN_S3_ENDPOINT` / `PN_S3_ACCESS_KEY` / `PN_S3_SECRET_KEY` /
`PN_S3_BUCKET` enable DataStore uploads (`PreparePostObject`); downloads work
without it.

## Status

Verified on a real Wii U (USA copy, 2026-10-06): token → `LoginEx` → secure
`RegisterEx` → Time Attack score upload (`UploadScore`) and "Top mondial"
leaderboard read (`GetRanking`) with the player's PNID shown. Not verified:
EUR/JPN consoles (same game server ID by construction, see above), DataStore
replay upload/download (needs `PN_S3_ENDPOINT`; the replay icon stays
crossed out without it), and multi-player ordering.

## License

AGPL-3.0. See [LICENSE](LICENSE) and [NOTICE](NOTICE). No proprietary
Nintendo or Retro Studios code or assets are included.
