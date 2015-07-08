package guid

import (
	"testing"
)

func TestGUID(t *testing.T) {
	count := 10000
	ids := make([]int64, 0, count)

	guid := NewGUID()

	for i := 0; i < count; i++ {
		id := guid.Next()

		ids = append(ids, id)

		for j := 0; j < i; j++ {
			if ids[j] == id {
				t.Error("The same id", i, j, id)
			}
			if ids[j] > id {
				t.Error("Older id is larger than new id", j, ids[j], i, id)
			}
		}
	}
}
