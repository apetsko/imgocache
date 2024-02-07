# Imgocache
Imgocache - Simple In-memory key-value cache with sync.RWMutex in Go.

1) The Imgocache type is a structure with an embedded sync.RWMutex and a map db to store key-value pairs, likely used for an image cache.
2) The type SM is a wrapper for the sync.Map type, ensuring safe concurrent read-write operations.

See it in action:

## Example 1

```go
package main

import (
	"fmt"
	"github.com/apetsko/imgocache"
)

func main(){
	var icache imgocache.Cacher
	icache = imgocache.New() // Creating a variable of type interface and assigning it a value of type Imgocache
	icache.Set("userId", 42)
	userID, exist := icache.Get("userId")

	fmt.Println(userID, exist)

	icache.Delete("userId")
	userID, exist = icache.Get("userId")

	fmt.Println(userID, exist)
}
```
#### output
```bash
$ go run .
42 true
<nil> false
```

## Example 2

```go
package main

import (
	"fmt"
	"github.com/apetsko/imgocache"
)

func main() {
	var icache imgocache.Cacher
	icache = imgocache.NewSM() // Creating a variable of type interface and assigning it a value of type Imgocache
	icache.Set("userId", 42)
	userID, exist := icache.Get("userId")

	fmt.Println(userID, exist)

	icache.Delete("userId")
	userID, exist = icache.Get("userId")

	fmt.Println(userID, exist)
}
```
#### output
```bash
$ go run .
42 true
<nil> false
```