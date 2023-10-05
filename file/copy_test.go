package file

import (
	"path/filepath"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func testdataPath(t *testing.T) string {
	t.Helper()
	_, current, _, _ := runtime.Caller(0)
	return filepath.Join(filepath.Dir(current), "testdata")
}

func Test_checkValidCopyPaths(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		args    string
		pass    bool
		wantErr error
	}{
		{
			name:    "normal",
			args:    filepath.Join(testdataPath(t)),
			pass:    true,
			wantErr: nil,
		},
		{
			name:    "not find pass",
			args:    filepath.Join(testdataPath(t), "no-exist"),
			pass:    false,
			wantErr: ErrNotFindPath,
		},
		{
			name:    "not directory",
			args:    filepath.Join(testdataPath(t), "option.txt"),
			pass:    false,
			wantErr: ErrNotDirectory,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := checkValidCopyPaths(tt.args)

			if tt.pass {
				assert.NoError(t, err)
			} else {
				assert.ErrorIs(t, err, tt.wantErr)
			}
		})
	}
}

func Test_convertToAbsForOption(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "normal",
			args: filepath.Join(testdataPath(t)),
			want: filepath.Join(testdataPath(t), "option.txt"),
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			if got := convertToAbsForOption(tt.args); got != tt.want {
				t.Errorf("convertToAbsForOption() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_copyToNewer(t *testing.T) {
	t.Parallel()

	t.Run("normal", func(t *testing.T) {

		src := filepath.Join(testdataPath(t), "option.txt")
		target := filepath.Join(t.TempDir(), "option.txt")
		err := copyToNewer(src, target)

		assert.NoError(t, err)
	})

	t.Run("try to open no existed file", func(t *testing.T) {

		src := filepath.Join(testdataPath(t), "no-existed.txt")
		target := filepath.Join(t.TempDir(), "option.txt")
		err := copyToNewer(src, target)

		assert.ErrorIs(t, err, ErrNotExistOptionFile)
	})

	t.Run("try to create no creatable file path", func(t *testing.T) {

		src := filepath.Join(testdataPath(t), "option.txt")
		target := ""
		err := copyToNewer(src, target)

		assert.ErrorIs(t, err, ErrFileCreationFailed)
	})
}
