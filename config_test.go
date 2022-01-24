package ablib

import (
	"strings"
	"testing"
)

func TestMakeConfig_error(t *testing.T) {
	tests := []struct {
		name     string
		filename string
		want     string
	}{
		{
			name:     "file not found",
			filename: "testdata/file_not_found.yml",
			want:     "no such file",
		},
		{
			name:     "malformed yaml",
			filename: "testdata/malformed.yml",
			want:     "cannot unmarshal",
		},
		{
			name:     "wrong file type",
			filename: "testdata/wrong_file_type.json",
			want:     "cannot unmarshal",
		},
	}

	for _, dat := range tests {
		_, err := MakeConfig(dat.filename)
		if err == nil {
			t.Errorf("expected MakeConfig to return error")
		}
		if !strings.Contains(err.Error(), dat.want) {
			t.Errorf("%s: expected substr %v, got %v", dat.name, dat.want, err.Error())
		}
	}
}
