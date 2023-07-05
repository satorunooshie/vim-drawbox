package main

import (
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_expandLine(t *testing.T) {
	type args struct {
		line  string
		width int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{line: "", width: 0},
			want: "",
		},
		{
			args: args{line: "\n", width: 0},
			want: "\n",
		},
		{
			args: args{line: "xx", width: 1},
			want: "xx",
		},
		{
			args: args{line: "xx\n", width: 1},
			want: "xx\n",
		},
		{
			args: args{line: "xxx\n", width: 6},
			want: "xxx   \n",
		},
		{
			args: args{line: "xxx", width: 6},
			want: "xxx   ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := expandLine(tt.args.line, tt.args.width); got != tt.want {
				t.Errorf("expandLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceAt(t *testing.T) {
	type args struct {
		line string
		pos  int
		s    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{line: "----", pos: 0, s: "xx"},
			want: "xx--",
		},
		{
			args: args{line: "----", pos: 1, s: "xx"},
			want: "-xx-",
		},
		{
			args: args{line: "----", pos: 3, s: "xx"},
			want: "---xx",
		},
		{
			args: args{line: "----", pos: 3, s: "xxx"},
			want: "---xxx",
		},
		{
			args: args{line: "----", pos: 3, s: "xx\n"},
			want: "---xx\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := replaceAt(tt.args.line, tt.args.pos, tt.args.s); got != tt.want {
				t.Errorf("replaceAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_overwriteAt(t *testing.T) {
	type args struct {
		line string
		pos  int
		s    string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{line: "----", pos: 0, s: "x x "},
			want: "x-x-",
		},
		{
			args: args{line: "----", pos: 1, s: "x x "},
			want: "-x-x",
		},
		{
			args: args{line: "----", pos: 2, s: "x x "},
			want: "--x-x",
		},
		{
			args: args{line: "----", pos: 3, s: "x x "},
			want: "---x x",
		},
		{
			args: args{line: "----\n", pos: 3, s: "x x "},
			want: "---x x\n",
		},
		{
			args: args{line: "---", pos: 0, s: " | "},
			want: "-+-",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := overwriteAt(tt.args.line, tt.args.pos, tt.args.s); got != tt.want {
				t.Errorf("overwriteAt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_replaceBlock(t *testing.T) {
	type args struct {
		lines []string
		y     int
		x     int
		block []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				lines: []string{"foo", "bar", "b"},
				y:     1,
				x:     2,
				block: []string{"1234", "5678"},
			},
			want: []string{"foo", "ba1234", "b 5678"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := replaceBlock(tt.args.lines, tt.args.y, tt.args.x, tt.args.block)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("replaceBlock() is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_amplifyPattern(t *testing.T) {
	type args struct {
		pattern string
		w       int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{pattern: "<->", w: 0},
			want: "",
		},
		{
			args: args{pattern: "<->", w: 1},
			want: "<",
		},
		{
			args: args{pattern: "<->", w: 2},
			want: "<-",
		},
		{
			args: args{pattern: "<->", w: 3},
			want: "<->",
		},
		{
			args: args{pattern: "<->", w: 4},
			want: "<-->",
		},
		{
			args: args{pattern: "<->", w: 5},
			want: "<--->",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := amplifyPattern(tt.args.pattern, tt.args.w); got != tt.want {
				t.Errorf("amplifyPattern() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_amplifyPatterns(t *testing.T) {
	type args struct {
		patterns []string
		w        int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				patterns: []string{
					"+---+",
					"|   |",
					"+---+",
				},
				w: 5,
			},
			want: []string{
				"+---+",
				"|   |",
				"|   |",
				"|   |",
				"+---+",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := amplifyPatterns(tt.args.patterns, tt.args.w)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("amplifyPatterns() is mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_reverseArrow(t *testing.T) {
	type args struct {
		arrow string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			args: args{arrow: "---"},
			want: "---",
		},
		{
			args: args{arrow: "<->"},
			want: "<->",
		},
		{
			args: args{arrow: "-->"},
			want: "<--",
		},
		{
			args: args{arrow: "<--"},
			want: "-->",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseArrow(tt.args.arrow); got != tt.want {
				t.Errorf("reverseArrow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_drawLineHV(t *testing.T) {
	type args struct {
		lines []string
		arrow string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				lines: []string{" 1    2 "},
				arrow: "o->",
			},
			want: []string{" o----> "},
		},
		{
			args: args{
				lines: []string{" 2    1 "},
				arrow: "o->",
			},
			want: []string{" <----o "},
		},
		{
			args: args{
				lines: []string{
					" ",
					"1",
					" ",
					" ",
					"2",
				},
				arrow: "o->",
			},
			want: []string{
				" ",
				"o",
				"|",
				"|",
				"v",
			},
		},
		{
			args: args{
				lines: []string{
					" ",
					"2",
					" ",
					" ",
					"1",
				},
				arrow: "o->",
			},
			want: []string{
				" ",
				"^",
				"|",
				"|",
				"o",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					" 1      ",
					"        ",
					"      2 ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" o----+ ",
				"      | ",
				"      v ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					" 2      ",
					"        ",
					"      1 ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" ^      ",
				" |      ",
				" +----o ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					"      1 ",
					"        ",
					" 2      ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" +----o ",
				" |      ",
				" v      ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					"      2 ",
					"        ",
					" 1      ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				"      ^ ",
				"      | ",
				" o----+ ",
				"        ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y1, x1, y2, x2 := pos(t, tt.args.lines)
			t.Log(y1, x1, y2, x2)
			got := drawLineHV(tt.args.lines, y1, x1, y2, x2, tt.args.arrow)
			t.Log(t.Name(), "got:")
			for _, v := range got {
				t.Log(v)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("drawLineHV() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_drawLineVH(t *testing.T) {
	type args struct {
		lines []string
		arrow string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				lines: []string{" 1    2 "},
				arrow: "o->",
			},
			want: []string{" o----> "},
		},
		{
			args: args{
				lines: []string{" 2    1 "},
				arrow: "o->",
			},
			want: []string{" <----o "},
		},
		{
			args: args{
				lines: []string{
					" ",
					"1",
					" ",
					" ",
					"2",
				},
				arrow: "o->",
			},
			want: []string{
				" ",
				"o",
				"|",
				"|",
				"v",
			},
		},
		{
			args: args{
				lines: []string{
					" ",
					"2",
					" ",
					" ",
					"1",
				},
				arrow: "o->",
			},
			want: []string{
				" ",
				"^",
				"|",
				"|",
				"o",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					" 1      ",
					"        ",
					"      2 ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" o      ",
				" |      ",
				" +----> ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					" 2      ",
					"        ",
					"      1 ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" <----+ ",
				"      | ",
				"      o ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					"      1 ",
					"        ",
					" 2      ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				"      o ",
				"      | ",
				" <----+ ",
				"        ",
			},
		},
		{
			args: args{
				lines: []string{
					"        ",
					"      2 ",
					"        ",
					" 1      ",
					"        ",
				},
				arrow: "o->",
			},
			want: []string{
				"        ",
				" +----> ",
				" |      ",
				" o      ",
				"        ",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			y1, x1, y2, x2 := pos(t, tt.args.lines)
			t.Log(y1, x1, y2, x2)
			got := drawLineVH(tt.args.lines, y1, x1, y2, x2, tt.args.arrow)
			t.Log(t.Name(), "got:")
			for _, v := range got {
				t.Log(v)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("drawLineVH() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_selectOuterBox(t *testing.T) {
	type args struct {
		lines []string
		y1    int
		x1    int
		y2    int
		x2    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				lines: []string{
					"              ",
					"   +-------+  ",
					"   | #     |  ",
					"   +-------+  ",
					"              ",
					"              ",
				},
				y1: 2,
				x1: 5,
				y2: 2,
				x2: 5,
			},
			want: []string{"1,3,3,11"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectOuterBox(tt.args.lines, tt.args.y1, tt.args.x1, tt.args.y2, tt.args.x2)
			t.Log(t.Name(), "got:")
			for _, v := range got {
				t.Log(v)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("selectOuterBox() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_selectInnerBox(t *testing.T) {
	type args struct {
		lines []string
		y1    int
		x1    int
		y2    int
		x2    int
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			args: args{
				lines: []string{
					"              ",
					"   +-------+  ",
					"   | #     |  ",
					"   +-------+  ",
					"              ",
					"              ",
				},
				y1: 2,
				x1: 5,
				y2: 2,
				x2: 5,
			},
			want: []string{"2,4,2,10"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := selectInnerBox(tt.args.lines, tt.args.y1, tt.args.x1, tt.args.y2, tt.args.x2)
			t.Log(t.Name(), "got:")
			for _, v := range got {
				t.Log(v)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("selectInnerBox() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

func Test_exec(t *testing.T) {
	type args struct {
		cmd   string
		lines []string
		args  []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "draw box nothing to do",
			args: args{
				cmd: "+o",
				lines: []string{
					"+---+",
					"|   |",
					"|   |",
					"|   |",
					"+---+",
				},
				args: []string{},
			},
			want: []string{
				"+---+",
				"|   |",
				"|   |",
				"|   |",
				"+---+",
			},
		},
		{
			name: "draw box basic pattern",
			args: args{
				cmd: "+o",
				lines: []string{
					"........",
					"..1.....",
					"........",
					"......2.",
					"........",
				},
				args: []string{},
			},
			want: []string{
				"........",
				"..+---+.",
				"..|   |.",
				"..+---+.",
				"........",
			},
		},
		{
			name: "draw box after line end",
			args: args{
				cmd: "+o",
				lines: []string{
					".......",
					"..1.",
					"",
					"......2",
				},
				args: []string{},
			},
			want: []string{
				".......",
				"..+---+",
				"  |   |",
				"..+---+",
			},
		},
		{
			name: "fill box alignments",
			args: args{
				cmd: "+{[c",
				lines: []string{
					"+------------+",
					"|1...........|",
					"|....FOO.....|",
					"|............|",
					"|...........2|",
					"+------------+",
				},
				args: []string{"This is a test."},
			},
			want: []string{
				"+------------+",
				"|This is a   |",
				"|test.       |",
				"|            |",
				"|            |",
				"+------------+",
			},
		},
		{
			name: "fill box alignments",
			args: args{
				cmd: "+{c",
				lines: []string{
					"+------------+",
					"|1...........|",
					"|....FOO.....|",
					"|............|",
					"|...........2|",
					"+------------+",
				},
				args: []string{"This is a test."},
			},
			want: []string{
				"+------------+",
				"| This is a  |",
				"|   test.    |",
				"|            |",
				"|            |",
				"+------------+",
			},
		},
		{
			name: "fill box alignments",
			args: args{
				cmd: "+{]c",
				lines: []string{
					"+------------+",
					"|1...........|",
					"|....FOO.....|",
					"|............|",
					"|...........2|",
					"+------------+",
				},
				args: []string{"This is a test."},
			},
			want: []string{
				"+------------+",
				"|   This is a|",
				"|       test.|",
				"|            |",
				"|            |",
				"+------------+",
			},
		},
		{
			name: "fill box alignments",
			args: args{
				cmd: "+c",
				lines: []string{
					"+------------+",
					"|1...........|",
					"|....FOO.....|",
					"|............|",
					"|...........2|",
					"+------------+",
				},
				args: []string{"This is a test."},
			},
			want: []string{
				"+------------+",
				"|            |",
				"| This is a  |",
				"|   test.    |",
				"|            |",
				"+------------+",
			},
		},
		{
			name: "fill box alignments",
			args: args{
				cmd: "+}]c",
				lines: []string{
					"+------------+",
					"|1...........|",
					"|....FOO.....|",
					"|............|",
					"|...........2|",
					"+------------+",
				},
				args: []string{"This is a test."},
			},
			want: []string{
				"+------------+",
				"|            |",
				"|            |",
				"|   This is a|",
				"|       test.|",
				"+------------+",
			},
		},
		{
			name: "fill box short of space",
			args: args{
				cmd: "+{[c",
				lines: []string{
					"+-----+",
					"|1    |",
					"|    2|",
					"+-----+",
				},
				args: []string{"not enough space"},
			},
			want: []string{
				"+-----+",
				"|not  |",
				"|enoug|",
				"+-----+",
			},
		},
		{
			name: "fill box short of space",
			args: args{
				cmd: "+{[c",
				lines: []string{
					"+-+",
					"|1|",
					"|.|",
					"|2|",
					"+-+",
				},
				args: []string{"not enough space"},
			},
			want: []string{
				"+-+",
				"|n|",
				"|e|",
				"|s|",
				"+-+",
			},
		},
		{
			name: "draw box with label",
			args: args{
				cmd: "+O",
				lines: []string{
					".........",
					".1.......",
					".........",
					".........",
					".......2.",
					".........",
				},
				args: []string{"foo bar"},
			},
			want: []string{
				".........",
				".+-----+.",
				".| foo |.",
				".| bar |.",
				".+-----+.",
				".........",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.args.cmd+"_"+tt.name, func(t *testing.T) {
			y1, x1, y2, x2 := pos(t, tt.args.lines)
			t.Log(pos(t, tt.args.lines))
			got := exec(tt.args.cmd, tt.args.lines, y1, x1, y2, x2, tt.args.args...)
			t.Log(t.Name(), "got:")
			for _, v := range got {
				t.Log(v)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("exec() mismatch (-got +want):\n%s", diff)
			}
		})
	}
}

// "1", "2" の場所から y, x を返す.
func pos(t *testing.T, lines []string) (y1, x1, y2, x2 int) {
	t.Helper()
	for i, l := range lines {
		if strings.Contains(l, "1") {
			y1 = i
			x1 = strings.Index(l, "1")
		}
		if strings.Contains(l, "2") {
			y2 = i
			x2 = strings.Index(l, "2")
		}
	}
	return y1, x1, y2, x2
}
