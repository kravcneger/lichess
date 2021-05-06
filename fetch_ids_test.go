package lichess

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFetchIds(t *testing.T) {
	assert.EqualValues(t, "id1,id2", fetchIds("id1,id2"))
	assert.EqualValues(t, "id1,id2", fetchIds([]string{"id1", "id2"}))
	assert.EqualValues(t, "id1,id2", fetchIds([]User{{ID: "id1"}, {ID: "id2"}}))
	assert.EqualValues(t, "id1", fetchIds(User{ID: "id1"}))
	assert.Panics(t, func() { fetchIds(4) }, "The code did not panic")
}
