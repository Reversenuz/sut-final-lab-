package entity

import (
	"testing"
)

func TestPositive(t *testing.T) {
	g := NewWithT(t)

	c := title.NewWithT([]string{"abcd"})

	g.Expect(c.title()).To(BeTrue(), "Title must be between 3 and 100")
}
