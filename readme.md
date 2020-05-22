<h1 align="center"><a href="https://github.com/nyaa8/gemshards">💎 gemshards</a></h1>
<p align="center">Snowflake-like Unique ID generator</p>

## Usage

Arguments:

- **EpochOffset** _5730518316208800 (example)_ Epoch offset in microseconds (should remain consistent in all instances)
- **GemID** _0-15 (default config)_ Must be unique across generator instances; otherwise, IDs can duplicate!

[Example code](example/main.go):

```go
package main

import (
	"fmt"
	"nyaa.science/gemshards"
)

func main() {
	generator := gemshards.Generator{EpochOffset: 5730518316208800, GeneratorID: 0}
	shard := generator.Generate()

	fmt.Println("Unique ID:", shard.ID)
}
```

# 🙌 Contributing

Please write commit messages according to [Chris Beams' guide](chris.beams.io/posts/git-commit) 😊

Thank you very much for reading this 🙇🏼‍♀️
