package users

import (
	"testing"

	"github.com/bmizerany/assert"
)

func TestCheckPasswordHash(t *testing.T) {

	tests := []struct {
		name     string
		password string
		hash     string
		want     bool
	}{
		{
			name:     "ok",
			password: "test",
			hash:     "$2a$14$pBSnpivvYO5L4HDmhGcC4e7Dm5zdC07jPftQ7QnG4prZeVnKLF/rK",
			want:     true,
		},
		{
			name:     "ok new hash",
			password: "test",
			hash:     "$2a$14$fm8S/a5svyKTFSAdbOFimudfCrSBkROq.JhSXvaoTECt6KlsV8mne",
			want:     true,
		},
		{
			name:     "blank",
			password: "",
			hash:     "$2a$14$odUzHzLa26poSFM8ZrhLFOA3ElzEEH89AAFCdACyX4mpulnlh2bcq",
			want:     true,
		},
		{
			name:     "100",
			password: "1234567890123456789012345678901234567890123456789012345678901234567890123456789012345678901234567890",
			hash:     "$2a$14$uPRTTd845aHmDV0srZ62a.YaL3aMyqONpzmmje5.EW/SryMUiE9JC",
			want:     true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := CheckPasswordHash(tt.password, tt.hash)
			assert.Equal(t, tt.want, got)
		})
	}
}
