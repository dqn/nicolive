# nicolive

[![build status](https://github.com/dqn/nicolive/workflows/build/badge.svg)](https://github.com/dqn/nicolive/actions)

Nicolive comments fetcher.

## Installation

```bash
$ go get github.com/dqn/nicolive
```

## Usage

```go
package main

import (
  "fmt"

  "github.com/dqn/nicolive"
)

func main() {
  n, err := nicolive.New("MAIL", "PASSWORD")
  if err != nil {
    // Handle error.
  }

  err = n.Listen("LIVE_ID", func(c *nicolive.Chat) error {
    fmt.Println(c.Text)
    return nil
  })

  if err != nil {
    // Handle error.
  }
}
```

## CLI

```bash
$ go get github.com/dqn/nicolive/cmd/nicolive
$ nicolive <mail> <password> <live-id>
```

## License

MIT
