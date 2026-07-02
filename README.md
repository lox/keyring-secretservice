keyring-secretservice
=====================
[![CI](https://github.com/lox/keyring-secretservice/actions/workflows/test.yml/badge.svg?branch=master)](https://github.com/lox/keyring-secretservice/actions/workflows/test.yml)
[![Go Reference](https://pkg.go.dev/badge/github.com/lox/keyring-secretservice.svg)](https://pkg.go.dev/github.com/lox/keyring-secretservice)

Secret Service provider for [`github.com/lox/keyring/v2`](https://github.com/lox/keyring).

## Usage

```bash
go get github.com/lox/keyring-secretservice
```

```go
import (
	"context"

	"github.com/lox/keyring/v2"
	secretservice "github.com/lox/keyring-secretservice"
)

ctx := context.Background()

ring, err := keyring.Open(ctx,
	keyring.WithServiceName("example"),
	keyring.WithProvider(secretservice.Provider()),
)
```

`secretservice.Provider` accepts the `Collection` option. On non-Linux
platforms, it returns `keyring.ErrUnavailable` during open.
