package ds

import (
	"fmt"
	"testing"
)

func TestScapegoatTree(t *testing.T) {
	st := NewScapegoatTree(0.75)

	values := []int{3, 1, 5, 0, 2, 4, 6}
	for _, v := range values {
		st.Insert(v)
	}

	st.InOrderTraversal(st.Root, func(value int) {
		fmt.Println(value)
	})
}
