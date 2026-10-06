package database

import (
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

// PlayerName returns the PNID user_id the account/token service recorded for
// pid in PN_DKCTF_NAMES_DIR (default /names), or "" if unknown. The game's
// leaderboard "Identifiant" column shows this, NUL-terminated, from the
// ranking common data.
func PlayerName(pid uint64) []byte {
	dir := os.Getenv("PN_DKCTF_NAMES_DIR")
	if dir == "" {
		dir = "/names"
	}
	raw, err := os.ReadFile(filepath.Join(dir, strconv.FormatUint(pid, 10)))
	name := strings.TrimSpace(string(raw))
	if err != nil || name == "" {
		return nil
	}
	return append([]byte(name), 0)
}
