cch-go
=======

Sipmle cache provider by unified interface

usage
------

### installation

```sh
go get -u github.com/0x75960/cch-go
```

### in code

1. define driver satisfy cch.CacheDriver

```go
// CacheDriver
type CacheDriver interface {
	HasItem(key interface{}) bool
	Add(item interface{})
	Remove(key interface{})
	Load()
	Dump()
}
```

2. initialze cch.Cache with cch.NewCache(driver CacheDriver)

```go
/*
    NOTE: `driver.Load` is called in NewCache
          and set callback that call `driver.Dump` when process has interrupted
*/
cache := cch.NewCache()
```

3. and set Dump to `defer` if you needed.

```go
defer cache.Dump() // Optional
```

4. use `Add` / `Remove` / `HasItem` in your process

```go
/* for example  */

if cache.HasItem(item) {
	// if already seen
	continue
}

// do something

cache.Add(item)
```
