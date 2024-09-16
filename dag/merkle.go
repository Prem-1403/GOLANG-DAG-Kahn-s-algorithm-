package dag

import (
	"crypto/sha256"
	"encoding/hex"
	"sort"
	"strings"
)

func generateHash(data map[string]string) string {
	var fields []string
	for _, val := range data {
		fields = append(fields, val)
	}
	sort.Strings(fields)
	joinedData := strings.Join(fields, "")

	hash := sha256.Sum256([]byte(joinedData))
	return hex.EncodeToString(hash[:])
}

func appendHash(existingHash string) string {
	hash := sha256.Sum256([]byte(existingHash))
	return hex.EncodeToString(hash[:])
}

func generateGraphSignature(leaves []map[string]string) string {
	var hashes []string
	for _, leaf := range leaves {
		hashes = append(hashes, leaf["data_hash"])
	}
	sort.Strings(hashes)
	joinedHashes := strings.Join(hashes, "")

	finalHash := sha256.Sum256([]byte(joinedHashes))
	return hex.EncodeToString(finalHash[:])
}
