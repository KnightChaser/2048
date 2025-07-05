package engine

import (
	"reflect"
	"testing"
)

func TestSlideLine(t *testing.T) {
	cases := []struct {
		name  string
		input []int
		want  []int
		moved bool
	}{
		{"no change", []int{2, 0, 2, 0}, []int{2, 2, 0, 0}, true},
		{"already slid", []int{4, 2, 0, 0}, []int{4, 2, 0, 0}, false},
		{"all zeros", []int{0, 0, 0, 0}, []int{0, 0, 0, 0}, false},
		{"mixed", []int{0, 4, 0, 2}, []int{4, 2, 0, 0}, true},
	}

	for _, c := range cases {
		got, moved := slideLine(c.input)
		if !reflect.DeepEqual(got, c.want) || moved != c.moved {
			t.Errorf("%s: slideLine(%v) = %v, %v; want %v, %v",
				c.name, c.input, got, moved, c.want, c.moved)
		}
	}
}

func TestMergeLine(t *testing.T) {
	cases := []struct {
		name       string
		input      []int // A pre-slid line
		want       []int // The line after merging, but before the final slide
		scoreGain  int
		wantMerged bool
	}{
		{
			"single merge start",
			[]int{2, 2, 4, 8},
			[]int{4, 0, 4, 8},
			4,
			true,
		},
		{
			"single merge end",
			[]int{4, 8, 2, 2},
			[]int{4, 8, 4, 0},
			4,
			true,
		},
		{
			"double merge",
			[]int{2, 2, 2, 2},
			[]int{4, 0, 4, 0}, // NOTE: becomes [4, 4, 0, 0] only AFTER a final slide
			8,
			true,
		},
		{
			"no merge",
			[]int{2, 4, 8, 16},
			[]int{2, 4, 8, 16},
			0,
			false,
		},
		{
			"no merge with zeros",
			[]int{4, 2, 0, 0},
			[]int{4, 2, 0, 0},
			0,
			false,
		},
		{
			"no chain reaction merge (important!)",
			[]int{4, 4, 8, 0},
			[]int{8, 0, 8, 0}, // NOTE: Should not become [16, 0, 0, 0]
			8,
			true,
		},
	}

	for _, c := range cases {
		// We no longer call slideLine here. We call mergeLine directly.
		got, gain, merged := mergeLine(c.input)
		if !reflect.DeepEqual(got, c.want) || gain != c.scoreGain || merged != c.wantMerged {
			t.Errorf("%s: mergeLine(%v) = %v, %d, %v; want %v, %d, %v",
				c.name, c.input, got, gain, merged, c.want, c.scoreGain, c.wantMerged)
		}
	}
}

func TestSlideMergeLine(t *testing.T) {
	cases := []struct {
		name      string
		input     []int
		want      []int
		moved     bool
		scoreGain int
	}{
		{"slide only", []int{0, 2, 0, 0}, []int{2, 0, 0, 0}, true, 0},
		{"adjacent merge", []int{2, 2, 0, 0}, []int{4, 0, 0, 0}, true, 4},
		{"slide & merge", []int{0, 2, 2, 0}, []int{4, 0, 0, 0}, true, 4},
		{"gap merge", []int{2, 0, 2, 0}, []int{4, 0, 0, 0}, true, 4},
		{"no move", []int{2, 4, 8, 16}, []int{2, 4, 8, 16}, false, 0},
	}

	for _, c := range cases {
		got, moved, gain := slideMergeLine(c.input)
		if !reflect.DeepEqual(got, c.want) || moved != c.moved || gain != c.scoreGain {
			t.Errorf("%s: slideMergeLine(%v) = %v, %v, %d; want %v, %v, %d",
				c.name, c.input, got, moved, gain, c.want, c.moved, c.scoreGain)
		}
	}
}
