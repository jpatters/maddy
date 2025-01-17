package imapsql

import (
	"io"

	imapsql "github.com/foxcpp/go-imap-sql"
	"github.com/foxcpp/maddy/framework/module"
)

type ExtBlob struct {
	io.ReadCloser
}

func (e ExtBlob) Sync() error {
	panic("not implemented")
}

func (e ExtBlob) Write(p []byte) (n int, err error) {
	panic("not implemented")
}

type WriteExtBlob struct {
	module.Blob
}

func (w WriteExtBlob) Read(p []byte) (n int, err error) {
	panic("not implemented")
}

type ExtBlobStore struct {
	Base module.BlobStore
}

func (e ExtBlobStore) Create(key string) (imapsql.ExtStoreObj, error) {
	blob, err := e.Base.Create(key)
	if err != nil {
		return nil, imapsql.ExternalError{
			NonExistent: err == module.ErrNoSuchBlob,
			Key:         key,
			Err:         err,
		}
	}
	return WriteExtBlob{Blob: blob}, nil
}

func (e ExtBlobStore) Open(key string) (imapsql.ExtStoreObj, error) {
	blob, err := e.Base.Open(key)
	if err != nil {
		return nil, imapsql.ExternalError{
			NonExistent: err == module.ErrNoSuchBlob,
			Key:         key,
			Err:         err,
		}
	}
	return ExtBlob{ReadCloser: blob}, nil
}

func (e ExtBlobStore) Delete(keys []string) error {
	err := e.Base.Delete(keys)
	if err != nil {
		return imapsql.ExternalError{
			Key: "",
			Err: err,
		}
	}
	return nil
}
