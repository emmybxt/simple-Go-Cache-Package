package Cache

import (
	"fmt"
	"testing"
	// "time"
)

func TestSetAndGet(t *testing.T){
    c := New()
    
    c.Set("$$$eMMMY_DEV", "THIS IS CUTE", 0)
    
    fmt.Print(c.Get("$$$eMMMY_DEV"))
    if _, found := c.Get("$$$eMMMY_DEV"); !found {
        t.Error("Item not found")
    }
}

func TestCacheDelete(t *testing.T) {
	c := New()

	c.Set("key1", "value1", 0)
	c.Delete("key1")
	_, found := c.Get("key1")

	if found {
		t.Errorf("Expected 'key1' to be deleted")
	}
}
