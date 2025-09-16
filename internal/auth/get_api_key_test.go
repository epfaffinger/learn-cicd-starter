// go
package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := []struct {
		name      string
		header    string
		wantKey   string
		wantErr   bool
		wantNoHdr bool
	}{
		{"no header", "", "", true, true},
		{"wrong scheme", "Bearer abc123", "", true, false},
		{"malformed one-part", "ApiKey", "", true, false},
		{"empty key", "ApiKey ", "", true, false},
		{"valid", "ApiKey abc123", "abc123", false, false},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			req, _ := http.NewRequest(http.MethodGet, "/", nil)
			if tc.header != "" {
				req.Header.Set("Authorization", tc.header)
			}
			key, err := GetAPIKey(req.Header)

			if tc.wantErr {
				if err == nil {
					t.Fatalf("expected error, got nil")
				}
				if tc.wantNoHdr && err != ErrNoAuthHeaderIncluded {
					t.Fatalf("expected ErrNoAuthHeaderIncluded, got %v", err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			if key != tc.wantKey {
				t.Fatalf("got key %q, want %q", key, tc.wantKey)
			}
		})
	}
}
