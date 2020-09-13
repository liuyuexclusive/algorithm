package heap

import "testing"

func TestMedianFinder(t *testing.T) {
	c := Constructor()
	c.AddNum(-1)
	c.AddNum(-2)
	c.AddNum(-3)
	c.AddNum(-4)
	c.AddNum(-5)
	// c.AddNum(3)

	var want float64 = -3
	if v := c.FindMedian(); v != want {
		t.Errorf("want: %v;got: %v", want, v)
	}
}
