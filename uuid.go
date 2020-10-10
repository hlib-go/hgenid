package hgenid

import (
	"github.com/google/uuid"
	"strings"
)

// UUID
func UUID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
