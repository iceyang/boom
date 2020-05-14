package boom

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoom(t *testing.T) {
	code := 500
	boom := Boom(code)
	assert.Equal(t, http.StatusText(code), boom.Message())
	assert.Equal(t, code, boom.Status())
	assert.Error(t, boom)
}
