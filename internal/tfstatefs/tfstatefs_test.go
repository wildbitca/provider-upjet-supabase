package tfstatefs

import (
	"bytes"
	"errors"
	"io/fs"
	"testing"

	"github.com/spf13/afero"
)

func state(id string) []byte {
	return []byte(`{"version":4,"resources":[{"mode":"managed","type":"supabase_project","name":"x","instances":[{"attributes":{"name":"a","id":"` + id + `"}}]}]}`)
}

func TestWriteState(t *testing.T) {
	cases := map[string]struct {
		content    []byte
		wantExists bool
	}{
		"EmptyIDIsDropped":      {content: state(""), wantExists: false},
		"MissingIDIsDropped":    {content: []byte(`{"version":4,"resources":[{"instances":[{"attributes":{"name":"a"}}]}]}`), wantExists: false},
		"NonEmptyIDIsPersisted": {content: state("abc"), wantExists: true},
		"NoResourcesPersisted":  {content: []byte(`{"version":4,"resources":[]}`), wantExists: true},
		"InvalidJSONPersisted":  {content: []byte(`not json`), wantExists: true},
	}
	for name, tc := range cases {
		t.Run(name, func(t *testing.T) {
			mem := afero.NewMemMapFs()
			a := afero.Afero{Fs: New(mem)}
			if err := a.WriteFile("/ws/terraform.tfstate", tc.content, 0o600); err != nil {
				t.Fatalf("WriteFile: %v", err)
			}
			got, err := afero.ReadFile(mem, "/ws/terraform.tfstate")
			if tc.wantExists {
				if err != nil || !bytes.Equal(got, tc.content) {
					t.Fatalf("state not persisted intact: %v %q", err, got)
				}
				return
			}
			if !errors.Is(err, fs.ErrNotExist) {
				t.Fatalf("expected no state file, got err=%v content=%q", err, got)
			}
		})
	}
}

func TestOtherFilesUntouched(t *testing.T) {
	mem := afero.NewMemMapFs()
	a := afero.Afero{Fs: New(mem)}
	body := []byte(`{"resources":[{"instances":[{"attributes":{"id":""}}]}]}`)
	if err := a.WriteFile("/ws/main.tf.json", body, 0o600); err != nil {
		t.Fatal(err)
	}
	if got, err := afero.ReadFile(mem, "/ws/main.tf.json"); err != nil || !bytes.Equal(got, body) {
		t.Fatalf("other file altered: %v %q", err, got)
	}
}
