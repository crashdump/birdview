package sync

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_clients_GetHealthEvents(t *testing.T) {
	c := Init()

	got := c.GetHealthEvents()
	//assert.NoError(t, err)
	assert.NotEmpty(t, got)
}