package secretservice

import (
	"context"

	"github.com/lox/keyring/v2"
)

const Backend = keyring.SecretServiceBackend

type Option func(*Config)

type Config struct {
	ServiceName             string
	LibSecretCollectionName string
}

func Collection(name string) Option {
	return func(cfg *Config) { cfg.LibSecretCollectionName = name }
}

func Provider(opts ...Option) keyring.Provider {
	cfg := Config{}
	for _, opt := range opts {
		if opt != nil {
			opt(&cfg)
		}
	}
	return keyring.Provider{
		Backend: Backend,
		Open: func(ctx context.Context, open keyring.OpenOptions) (keyring.Keyring, error) {
			if err := ctx.Err(); err != nil {
				return nil, err
			}
			openCfg := cfg
			openCfg.ServiceName = open.ServiceName
			opener, ok := supportedBackends[Backend]
			if !ok {
				return nil, keyring.ErrUnavailable
			}
			ring, err := opener(openCfg)
			if err != nil {
				return nil, err
			}
			return adapter{ring: ring}, nil
		},
	}
}

type opener func(Config) (backendKeyring, error)

var supportedBackends = map[keyring.Backend]opener{}

type backendKeyring interface {
	Get(string) (keyring.Item, error)
	GetMetadata(string) (keyring.Metadata, error)
	Set(keyring.Item) error
	Remove(string) error
	Keys() ([]string, error)
}

type adapter struct {
	ring backendKeyring
}

func (a adapter) Get(ctx context.Context, key string) (keyring.Item, error) {
	if err := ctx.Err(); err != nil {
		return keyring.Item{}, err
	}
	return a.ring.Get(key)
}

func (a adapter) Set(ctx context.Context, item keyring.Item) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return a.ring.Set(item)
}

func (a adapter) Remove(ctx context.Context, key string) error {
	if err := ctx.Err(); err != nil {
		return err
	}
	return a.ring.Remove(key)
}

func (a adapter) Keys(ctx context.Context) ([]string, error) {
	if err := ctx.Err(); err != nil {
		return nil, err
	}
	return a.ring.Keys()
}

func (a adapter) Metadata(ctx context.Context, key string) (keyring.Metadata, error) {
	if err := ctx.Err(); err != nil {
		return keyring.Metadata{}, err
	}
	return a.ring.GetMetadata(key)
}
