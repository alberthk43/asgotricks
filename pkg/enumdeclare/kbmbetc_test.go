package enumdeclare

import (
	"fmt"
	"testing"
)

const (
	_  = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
)

func ToHumanReadable(size int64) string {
	switch {
	case size >= TB:
		return fmt.Sprintf("%.2f TB", float64(size)/TB)
	case size >= GB:
		return fmt.Sprintf("%.2f GB", float64(size)/GB)
	case size >= MB:
		return fmt.Sprintf("%.2f MB", float64(size)/MB)
	case size >= KB:
		return fmt.Sprintf("%.2f KB", float64(size)/KB)
	default:
		return fmt.Sprintf("%d B", size)
	}
}

func TestToHummanReadable(t *testing.T) {
	tests := []struct {
		size     int64
		expected string
	}{
		{size: 100, expected: "100 B"},
		{size: 1024, expected: "1.00 KB"},
		{size: 1024 * 1024, expected: "1.00 MB"},
		{size: 500 * 1024 * 1024, expected: "500.00 MB"},
		{size: 1024 * 1024 * 1024, expected: "1.00 GB"},
		{size: 1024 * 1024 * 1024 * 1024, expected: "1.00 TB"},
	}
	for _, tt := range tests {
		if got := ToHumanReadable(tt.size); got != tt.expected {
			t.Errorf("ToHumanReadable(%d) got %s, want %s", tt.size, got, tt.expected)
		}
	}
}
