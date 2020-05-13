# nicolive

Fetch nicolive chats.

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
		// handle error
	}

	n.Listen("LIVE_ID", func(c *nicolive.Chat) {
		fmt.Println(c.Text)
	})
}
```

## CLI

```bash
$ go get github.com/dqn/nicolive/cmd/nicolive
$ nicolive <mail> <password> <live-id>
```
