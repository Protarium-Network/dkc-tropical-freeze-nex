# Protocol coverage

Console-observed calls (2026-10-06): Ranking `GetRanking`(0x09), `GetCommonData`(0x06), `UploadScore`(0x01). No DataStore call was seen yet. See [RECON.md](RECON.md).

| Endpoint | Protocol | Coverage |
|---|---|---|
| Auth | Ticket Granting | Login / LoginEx / RequestTicket, stock common handler |
| Secure | Secure Connection, Utility | stock |
| Secure | Ranking | `GetRankings`, `GetRankingsAndCount`, `UploadScore`, common-data get/upload; Postgres `dkctf_*` tables, generic (categories not tuned) |
| Secure | DataStore | `PostMetaBinary`, `PrepareGetObject`, `PreparePostObject` → S3 → `CompletePostObject`, meta updates, `GetMetasMultipleParam`, `GetRatings` (empty) |

Not registered: NAT Traversal, MatchMaking, MatchmakeExtension (no online
multiplayer in the game).

## Still to check

1. Replay upload/download over DataStore (needs S3) — untested.
2. EUR / JPN consoles.
