package util

import "testing"

func TestArray(t *testing.T) {
	// find
	arr := []int{10, 30, 50, 20, 60, 70, 80, 40}

	finds := []int{10, 30, 50, -50, -20, -30}

	for _, v := range finds {
		expection := -1
		for i, a := range arr {
			if v == a {
				expection = i
				break
			}
		}

		idx := Find(arr, v)

		if idx != expection {
			t.Errorf("Find failed: %v shoud be at %v, %v.", v, expection, arr)
		}
	}

	// insert
	inserts := []int{25, 65, 35}

	for i, v := range inserts {
		Insert(&arr, i, v)

		idx := Find(arr, v)

		if idx != i {
			t.Errorf("Insert failed: %v shoud be at %v, %v", v, i, arr)
		}
	}

	// delete
	for i := len(inserts) - 1; i >= 0; i-- {
		v := inserts[i]

		n := Delete(&arr, v)
		if n != i {
			t.Errorf("Delete failed: %v shoud be deleted %v, %v.", v, n, arr)
		}

		idx := Find(arr, v)
		if idx != -1 {
			t.Errorf("Delete failed: %v shoud not be here %v, %v.", v, idx, arr)
		}
	}

	// reverse
	b := make([]int, len(arr))
	copy(b, arr)

	Reverse(b)
	for i, v := range arr {
		j := len(arr) - 1 - i

		if b[j] != v {
			t.Errorf("Reverse failed: %v ,%v not match. %v, %v.", i, j, arr, b)
		}
	}

	// remove
	for i := len(arr) - 1; i >= 0; i-- {
		v := arr[i]

		Remove(&arr, i)

		idx := Find(arr, v)
		if idx != -1 {
			t.Errorf("Remove failed: %v shoud not be here %v, %v.", v, idx, arr)
		}
	}

}
