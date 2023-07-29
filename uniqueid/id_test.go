package main

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewUniqueIdSuccess(t *testing.T) {
	prevId := ""
	for i := 0; i < 10; i++ {
		id, err := NewUniqueId()
		assert.Nil(t, err)
		assert.Equal(t, 14, len(id))

		time.Sleep(500 * time.Millisecond)

		if prevId > id {
			t.Errorf("previous id '%s' is greater than current id '%s'", prevId, id)
		}
		prevId = id
	}
}

func TestNewUniqueIdConcurrentSuccess(t *testing.T) {
	// todo
	prevId := ""
	for i := 0; i < 10; i++ {
		id, err := NewUniqueId()
		assert.Nil(t, err)
		assert.Equal(t, 14, len(id))

		time.Sleep(500 * time.Millisecond)

		if prevId > id {
			t.Errorf("previous id '%s' is greater than current id '%s'", prevId, id)
		}
		prevId = id
	}
}
