package Cache

import (
    "time"
    "sync"
)

type Cache struct {
    items map[string]Item
    mutex sync.RWMutex
}

func New() *Cache {
    return &Cache{
        items: make(map[string]Item),
    }
}

func (c *Cache) Set(key string, value interface{}, duration time.Duration) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    var expiration int64
    if duration > 0 {
        expiration = time.Now().Add(duration).UnixNano()
    }

    if expiration > 0 {
        go func() {
            time.Sleep(duration)
            c.mutex.Lock()
            defer c.mutex.Unlock()
            delete(c.items, key)
        }()
    }
    c.items[key] = Item{
        Value: value,
        Expiration: expiration,
    }
}

func (c *Cache) Get(key string) (interface{}, bool) {
    c.mutex.RLock()
    defer c.mutex.RUnlock()

    item, found := c.items[key]
    if !found || item.Expired() {
        return nil, false
    }
    return item.Value, true
}

func (c *Cache) Delete(key string) {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    if _, found := c.items[key]; found {
        delete(c.items, key)
    }
}

func (c *Cache) DeleteExpired() {
    c.mutex.Lock()
    defer c.mutex.Unlock()

    for key, item := range c.items {
        if item.Expired() {
            delete(c.items, key)
        }
    }
}