package Cache

import (
    "time"
)

type Item struct {
    Value     interface{}
    Expiration   int64
}

func (item *Item) Expired() bool {
    if item.Expiration == 0 {
        return false
    }
    return time.Now().UnixNano() > item.Expiration
}