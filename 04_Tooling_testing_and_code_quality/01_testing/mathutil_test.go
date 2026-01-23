package mathutil

import "testing"

func TestAdd(t *testing.T) {
	tests := []struct {
		name string
		a, b int
		want int
	}{
		{"both positive", 2, 3, 5},
		{"zero", 0, 5, 5},
		{"negatives", -1, -2, -3},
	}

	for _, tc := range tests {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			if got := Add(tc.a, tc.b); got != tc.want {
				t.Fatalf("Add(%d,%d) = %d; want %d", tc.a, tc.b, got, tc.want)
			}
		})
	}
}

func TestMul(t *testing.T) {
	cases := []struct{ a, b, w int }{{2, 3, 6}, {0, 5, 0}, {-2, 3, -6}}
	for i, c := range cases {
		if got := Mul(c.a, c.b); got != c.w {
			t.Fatalf("case %d: Mul(%d,%d)=%d want=%d", i, c.a, c.b, got, c.w)
		}
	}
}
