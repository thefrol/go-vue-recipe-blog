package localstorage

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const testFolder = ".testfolder"

func TestFileStorage_tokensFile(t *testing.T) {

	tests := []struct {
		name   string
		folder string
		want   string
	}{
		{
			name:   "positive #1",
			folder: "fold",
			want:   "fold/tokens",
		},
		{
			name:   "positive #2",
			folder: "./fold",
			want:   "fold/tokens",
		},
		{
			name:   "windows #1",
			folder: "c:/fold",
			want:   "c:/fold/tokens",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := FileStorage{
				folder: tt.folder,
			}
			assert.Equal(t, tt.want, s.tokensFile())
		})
	}
}

func TestFileStorage_Token(t *testing.T) {
	tests := []struct {
		name      string
		hash      []byte
		wantFound bool
		wantErr   bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := os.Stat(testFolder); err != os.ErrNotExist {
				err := os.RemoveAll(testFolder)
				require.NoError(t, err, "Cant delete folder")
			}

			// TODO
			// сделать функцию withtestfolder(folde, func)
			//
			// а ещё можно воспользовать os.TempDir

			err := os.Mkdir(testFolder, os.FileMode(os.O_CREATE|os.O_RDWR))
			require.NoError(t, err, "Cant create test folder")

			s := FileStorage{
				folder: testFolder,
			}
			gotFound, err := s.Token(tt.hash)
			if (err != nil) != tt.wantErr {
				t.Errorf("FileStorage.Token() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if gotFound != tt.wantFound {
				t.Errorf("FileStorage.Token() = %v, want %v", gotFound, tt.wantFound)
			}

			err = os.RemoveAll(testFolder)
			assert.NoError(t, err, "Cant delete folder")
		})
	}
}
