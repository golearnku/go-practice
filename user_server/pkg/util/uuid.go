package util

import (
	"strings"

	uuid "github.com/satori/go.uuid"
)

// NewUUID Create uuid
func NewUUID() string {
	gid := uuid.NewV4()
	return strings.ReplaceAll(gid.String(), "-", "")
}
