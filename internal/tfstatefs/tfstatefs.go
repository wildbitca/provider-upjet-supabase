// Package tfstatefs wraps the Terraform workspace file system so that a
// synthetic Terraform state carrying an empty resource ID is never persisted.
//
// Upjet reconstructs the Terraform state of a managed resource from its
// external name before every refresh. When the external name is empty (the
// resource was never created) the reconstructed state holds the resource
// attributes with id="". SDKv2 providers treat an empty ID as "resource gone"
// and drop the instance during refresh, so Upjet sees no state and calls
// Create. The Plugin Framework based Supabase provider instead keeps the
// instance, so Upjet reports the resource as existing (a phantom Ready=True)
// and Create never runs. Not persisting the state restores the SDKv2
// semantics: the refresh starts from an empty state, Upjet observes that the
// resource does not exist, and Terraform apply creates it.
package tfstatefs

import (
	"bytes"
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/spf13/afero"
)

const stateFileName = "terraform.tfstate"

// New returns an afero.Fs that behaves like base, except that writes of a
// terraform.tfstate whose resource instances all have an empty "id" attribute
// are silently discarded.
func New(base afero.Fs) afero.Fs {
	return &guardFs{Fs: base}
}

type guardFs struct {
	afero.Fs
}

// OpenFile intercepts writable opens of terraform.tfstate.
func (f *guardFs) OpenFile(name string, flag int, perm os.FileMode) (afero.File, error) {
	file, err := f.Fs.OpenFile(name, flag, perm)
	if err != nil || filepath.Base(name) != stateFileName || flag&(os.O_WRONLY|os.O_RDWR) == 0 {
		return file, err
	}
	return &stateFile{File: file, fs: f.Fs, name: name}, nil
}

// stateFile buffers the state and drops it on Close if it has an empty ID.
type stateFile struct {
	afero.File
	fs   afero.Fs
	name string
	buf  bytes.Buffer
}

func (s *stateFile) Write(p []byte) (int, error) {
	// Buffer only; the underlying file is written on Close, once the full
	// document is known.
	return s.buf.Write(p)
}

func (s *stateFile) WriteString(str string) (int, error) {
	return s.buf.WriteString(str)
}

func (s *stateFile) Close() error {
	if hasEmptyID(s.buf.Bytes()) {
		cerr := s.File.Close()
		// The file was created by the open; remove the empty shell so that
		// Upjet sees a missing state.
		if err := s.fs.Remove(s.name); err != nil && !os.IsNotExist(err) {
			return err
		}
		return cerr
	}
	if _, err := s.File.Write(s.buf.Bytes()); err != nil {
		_ = s.File.Close()
		return err
	}
	return s.File.Close()
}

// hasEmptyID reports whether the Terraform v4 state document holds at least
// one instance and every instance has an empty (or missing) string "id".
func hasEmptyID(raw []byte) bool {
	var st struct {
		Resources []struct {
			Instances []struct {
				Attributes json.RawMessage `json:"attributes"`
			} `json:"instances"`
		} `json:"resources"`
	}
	if err := json.Unmarshal(raw, &st); err != nil {
		return false
	}
	seen := false
	for _, r := range st.Resources {
		for _, i := range r.Instances {
			var attrs map[string]any
			if err := json.Unmarshal(i.Attributes, &attrs); err != nil {
				return false
			}
			if id, _ := attrs["id"].(string); id != "" {
				return false
			}
			seen = true
		}
	}
	return seen
}
