package hello

import "testing"

func TestSayHello(t *testing.T) {
	subtests := []struct {
		items  []string
		result string
	}{
		{
			result: "Hello, world!",
		},
		{
			items:  []string{"Ashiq"},
			result: "Hello, Ashiq!",
		},
		{
			items:  []string{"Ashiq", "Yousuf"},
			result: "Hello, Ashiq, Yousuf!",
		},
	}

	for _, st := range subtests {
		if s := Say(st.items); s != st.result {
			t.Errorf("Wanted:%s (%v), got:%s", st.result, st.items, s)
		}
	}
}
