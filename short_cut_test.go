package boom

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShortCutFunc(t *testing.T) {
	boom := BadRequest()
	assert.Equal(t, http.StatusText(http.StatusBadRequest), boom.Message())
	assert.Equal(t, 400, boom.Status())
	assert.Error(t, boom)
}
