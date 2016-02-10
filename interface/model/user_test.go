package model

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStatus(t *testing.T) {
	statusSlice := []struct {
		status UserStatus
		expected string
	}{
		{status: Member, expected: `"member"`},
		{status: Withdrawal, expected: `"withdrawal"`},
	}

	for _, status := range statusSlice {
		statusStr, err := json.Marshal(status.status)
		if err != nil {
			t.Fatal(err)
		}
		assert.Equal(t, string(statusStr), status.expected, "they should be equal")
	}
}
