# Imgocache
Simple In-memory key-value cache with RWMutex in Go.

See it in action:

## Example

```go
func main() {

	cache := cache.New()

	cache.Set("userId", 42)
	userId := cache.Get("userId")

	fmt.Println(userId)

	cache.Delete("userId")
	userId, exist := cache.Get("userId")

	fmt.Println(userId)
}
```