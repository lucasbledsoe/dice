package dice

import (
	"fmt"
	"testing"
)

var groupProperties = []struct {
	name  string
	props GroupProperties
}{{"3d6",
	GroupProperties{
		Type:     TypePolyhedron,
		Size:     6,
		Count:    3,
		Unrolled: true,
	},
}, {"3dF",
	GroupProperties{
		Type:     TypeFate,
		Count:    3,
		Unrolled: true,
	},
},
}

func Benchmark(b *testing.B) {
	for _, bench := range groupProperties {
		b.Run(fmt.Sprintf("%s", bench.name), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				NewGroup(bench.props)
			}
		})
	}
}

func TestGroup_Total(t *testing.T) {
	tests := []struct {
		name string
		g    Group
		want float64
	}{
		{
			name: "basic",
			g: Group{
				&PolyhedralDie{Result: 2},
				&PolyhedralDie{Result: 3},
				&PolyhedralDie{Result: 4},
			},
			want: 9,
		},
		{
			name: "nested",
			g: Group{
				&PolyhedralDie{Result: 2},
				&Group{
					&PolyhedralDie{Result: 3},
				},
				&PolyhedralDie{Result: 4},
			},
			want: 9,
		},
		{
			name: "dropped",
			g: Group{
				&PolyhedralDie{Result: 2, Dropped: true},
				&PolyhedralDie{Result: 4},
			},
			want: 4,
		},
		{
			name: "mixed",
			g: Group{
				&PolyhedralDie{Result: 2, Dropped: true},
				&FateDie{Result: -1},
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Total(); got != tt.want {
				t.Errorf("Group.Total() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGroup_Expression(t *testing.T) {
	tests := []struct {
		name string
		g    Group
		want string
	}{
		{
			name: "basic",
			g: Group{
				&PolyhedralDie{Result: 2},
				&PolyhedralDie{Result: 3},
				&PolyhedralDie{Result: 4},
			},
			want: "2+3+4",
		},
		{
			name: "unrolled",
			g: Group{
				&PolyhedralDie{Type: "d3", Unrolled: true},
				&PolyhedralDie{Result: 3},
				&PolyhedralDie{Result: 4},
			},
			want: "d3+3+4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.g.Expression(); got != tt.want {
				t.Errorf("Group.Expression() = %v, want %v", got, tt.want)
			}
		})
	}
}
