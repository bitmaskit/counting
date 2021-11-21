package counting

import (
	"reflect"
	"testing"
)

func Test_Sort(t *testing.T) {
	unsorted := []int{5,4,3,1,2}
	sorted := []int{1,2,3,4,5}

	Sort(unsorted)

	if !reflect.DeepEqual(sorted, unsorted) {
		t.Errorf("error sorting slice, got: %+v, expect: %+v", unsorted, sorted)
	}
}