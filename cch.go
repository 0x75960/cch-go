package cch

import (
	"os"
	"os/signal"
)

// CacheDriver interface
type CacheDriver interface {
	HasItem(key interface{}) bool
	Add(item interface{})
	Remove(key interface{})
	Load()
	Dump()
}

// Cache provider
type Cache struct {
	driver CacheDriver
}

// NewCache with specified Driver
func NewCache(driver CacheDriver) Cache {

	// catch interrupt by user
	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	go func() {
		for range c {
			driver.Dump()
			os.Exit(0)
		}
	}()

	driver.Load()

	return Cache{
		driver: driver,
	}
}

// Dump cache if needed
func (c Cache) Dump() {
	c.driver.Dump()
}

// HasItem in cache
func (c Cache) HasItem(key interface{}) bool {
	return c.driver.HasItem(key)
}

// Add item to cache
func (c Cache) Add(item interface{}) {
	c.driver.Add(item)
}

// Remove item has key from cache
func (c Cache) Remove(key interface{}) {
	c.driver.Remove(key)
}
