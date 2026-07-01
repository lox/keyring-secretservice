package secretservice

import "github.com/lox/keyring/v2"

type Item = keyring.Item
type Metadata = keyring.Metadata

const SecretServiceBackend = keyring.SecretServiceBackend

var (
	ErrKeyNotFound              = keyring.ErrKeyNotFound
	ErrMetadataNeedsCredentials = keyring.ErrMetadataNeedsCredentials
	ErrUnavailable              = keyring.ErrUnavailable
)
