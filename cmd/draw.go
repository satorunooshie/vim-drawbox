package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type where string

const (
	top    where = "top"
	middle where = "middle"
	bottom where = "bottom"
	center where = "center"
	left   where = "left"
	right  where = "right"
)

func main() {
	args := os.Args[1:]
	if len(args) < 5 {
		fmt.Fprintln(os.Stderr, "error:", "not enough argument")
		os.Exit(1)
	}
	cmd := args[0]
	y1, _ := strconv.Atoi(args[1])
	x1, _ := strconv.Atoi(args[2])
	y2, _ := strconv.Atoi(args[3])
	x2, _ := strconv.Atoi(args[4])

	var lines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
	for _, line := range exec(cmd, lines, y1, x1, y2, x2, args[5:]...) {
		fmt.Println(line)
	}
}

func exec(cmd string, lines []string, y1, x1, y2, x2 int, args ...string) []string {
	text := ""
	if len(args) != 0 {
		text = strings.Join(args, " ")
	}
	switch cmd {
	// Box drawing.
	case "+o": // rectangle.
		return drawBox(lines, y1, x1, y2, x2)
	case "+O", "+mcb": // labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, middle, center, text)

	case "+[O", "+mlb": // middle left labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, middle, left, text)
	case "+]O", "+mrb": // middle right labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, middle, right, text)

	case "+{O", "+tcb": // top center labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, top, middle, text)
	case "+}O", "+bcb": // bottom center labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, bottom, middle, text)

	case "+{[O", "+tlb": // top left labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, top, left, text)
	case "+{]O", "+trb": // top right labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, top, right, text)

	case "+}[O", "+blb": // bottom left labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, bottom, left, text)
	case "+}]O", "+brb": // bottom right labeled rectangle.
		return drawBoxWithLabel(lines, y1, x1, y2, x2, bottom, right, text)

		// Labeling.
	case "+c", "+mcl": // middle center labeling.
		return fillBox(lines, y1, x1, y2, x2, middle, center, text)
	case "+[c", "+mll": // middle left labeling.
		return fillBox(lines, y1, x1, y2, x2, middle, left, text)
	case "+]c", "+mrl": // middle right labeling.
		return fillBox(lines, y1, x1, y2, x2, middle, right, text)
	case "+{c", "+tcl": // top center labeling.
		return fillBox(lines, y1, x1, y2, x2, top, center, text)
	case "+}c", "+bcl": // bottom center labeling.
		return fillBox(lines, y1, x1, y2, x2, bottom, center, text)
	case "+{[c", "+tll": // top left labeling.
		return fillBox(lines, y1, x1, y2, x2, top, left, text)
	case "+{]c", "+trl": // top right labeling.
		return fillBox(lines, y1, x1, y2, x2, top, right, text)
	case "+}[c", "+bll": // bottom left labeling.
		return fillBox(lines, y1, x1, y2, x2, bottom, left, text)
	case "+}]c", "+brl": // bottom right labeling.
		return fillBox(lines, y1, x1, y2, x2, bottom, right, text)

		// Line drawing.
	case "+>", "+<": // horizontal line.
		return drawLineVH(lines, y1, x1, y2, x2, "-->")
	case "+V", "+v", "+^": // vertical line.
		return drawLineHV(lines, y1, x1, y2, x2, "-->")
	case "++>", "++<": // double horizontal line.
		return drawLineVH(lines, y1, x1, y2, x2, "<->")
	case "++V", "++v", "++^": // double vertical line.
		return drawLineHV(lines, y1, x1, y2, x2, "<->")
	case "+-", "+_": // horizontal line.
		return drawLineVH(lines, y1, x1, y2, x2, "---")
	case "+|": // vertical line.
		return drawLineHV(lines, y1, x1, y2, x2, "---")

		// Selection.
	case "ao", "ab": // outer box.
		return selectOuterBox(lines, y1, x1, y2, x2)
	case "io", "ib": // inner box.
		return selectInnerBox(lines, y1, x1, y2, x2)
	}
	return nil
}

// wrap wraps the given text into lines of width w, without breaking words.
//
// Example:
// wrap("The quick brown fox jumps over the lazy dog.", 10) returns
// ["The quick", "brown fox", "jumps over", "the lazy", "dog."].
//
// wrap("not enough space", 1) returns
// ["not", "enough", "space"].
func wrap(text string, w int, breakingLongWords bool) []string {
	var line string
	var length int
	lines := make([]string, 0, len(text)/w+1)
	words := strings.Split(text, " ")
	for _, word := range words {
		if lw := len(word); lw+length <= w {
			// add word to current line.
			line += word + " "
			length += lw + 1
		} else if length == 0 && lw > w {
			// add long word to new line.
			lines = append(lines, string(word))
		} else if breakingLongWords {
			// finish current line and add word to next line.
			lines = append(lines, strings.TrimSpace(line))
			line = word + " "
			length = lw + 1
		} else {
			// add current line and start a new one.
			lines = append(lines, strings.TrimSpace(line))
			line = word + " "
			length = lw + 1
		}
	}
	if len(line) == 0 {
		return lines
	}
	if v := strings.TrimSpace(line); v != "" {
		return append(lines, v)
	}
	return lines
}

type BlockPosition struct {
	Y, X, Height, Width int
}

func blockPos(y1, x1, y2, x2 int) BlockPosition {
	return BlockPosition{
		Y:      min(y1, y2),
		X:      min(x1, x2),
		Height: max(y1, y2) - min(y1, y2) + 1,
		Width:  max(x1, x2) - min(x1, x2) + 1,
	}
}

// splitNewLine returns the contents of the line without the newline, and the newline if exists.
func splitNewLine(line string) (string, string) {
	if len(line) > 0 && line[len(line)-1] == '\n' {
		return line[:len(line)-1], line[len(line)-1:]
	}
	return line, ""
}

// expandLine returns the line padded with spaces to match the width, preserving any newline characters.
func expandLine(line string, width int) string {
	line, nl := splitNewLine(line)
	if len(line) < width {
		line += strings.Repeat(" ", width-len(line))
	}
	return line + nl
}

// replaceAt replaces part of the line starting at `pos`.
func replaceAt(line string, pos int, s string) string {
	line = expandLine(line, pos+len(s))
	return line[:pos] + s + line[pos+len(s):]
}

func overwriteAt(line string, pos int, s string) string {
	s = strings.TrimRight(s, " ")
	l := len(s)
	line = expandLine(line, pos+l)
	line, nl := splitNewLine(line)

	var sb strings.Builder
	sb.Grow(len(line))
	for i, r := range line {
		if pos > i || i >= pos+l {
			sb.WriteRune(r)
			continue
		}
		s := string(s[i-pos])
		if s == " " {
			sb.WriteRune(r)
		} else if rs := string(r); (strings.ContainsAny(s, "-+") && strings.ContainsAny(rs, "|+")) ||
			(strings.ContainsAny(s, "|+") && strings.ContainsAny(rs, "-+")) {
			sb.WriteString("+")
		} else {
			sb.WriteString(s)
		}
	}
	return sb.String() + nl
}

func mergeBlock(lines []string, y, x int, block []string, mergeFn func(string, int, string) string) []string {
	w := 0
	h := len(block)
	for _, line := range block {
		lw := len(line)
		if lw > w {
			w = lw
		}
	}

	mb := make([]string, len(lines))
	for l, line := range lines {
		line, nl := splitNewLine(line)
		if h > 0 && w > 0 && y <= l && l < y+h {
			mb[l] = mergeFn(line, x, block[l-y]) + nl
		} else {
			mb[l] = line + nl
		}
	}
	return mb
}

func replaceBlock(lines []string, y, x int, block []string) []string {
	return mergeBlock(lines, y, x, block, replaceAt)
}

// overwriteBlock overwrites a rectangular block into the given lines, except whitespaces.
func overwriteBlock(lines []string, y, x int, block []string) []string {
	return mergeBlock(lines, y, x, block, overwriteAt)
}

// amplifyPattern returns pattern amplified to fit width `w`.
func amplifyPattern(pattern string, w int) string {
	if len(pattern) == 0 || w == 0 {
		return ""
	}
	if len(pattern) == 1 {
		return string(pattern[0])
	}
	// concatenate the first and last characters and amplify the string to fit the width.
	amplified := string(pattern[0]) + strings.Repeat(string(pattern[1]), max(1, w-2)) + string(pattern[len(pattern)-1])
	// truncate to fit the width.
	if len(amplified) > w {
		return amplified[:w]
	}
	return amplified
}

// amplifyPatterns returns patterns amplified to fit width `w`.
func amplifyPatterns(patterns []string, w int) []string {
	amplified := make([]string, w)
	amplified[0] = amplifyPattern(patterns[0], w)
	amplified[w-1] = amplifyPattern(patterns[len(patterns)-1], w)
	for i := 1; i < w-1; i++ {
		amplified[i] = amplifyPattern(patterns[1], w)
	}
	return amplified
}

// horizontalAlign returns the line aligned horizontally within the width,
// according to the alignment direction.
// If the line contains a newline character, it is split into two lines and aligned
// separately. If the line is longer than the width, it is truncated.
func horizontalAlign(line string, width int, direction where) string {
	line, nl := splitNewLine(line)
	l := len(line)
	if l > width { // truncate.
		line = line[:width]
	}
	switch direction {
	case left:
		return fmt.Sprintf("%-*s%s", width, line, nl)
	case right:
		return fmt.Sprintf("%*s%s", width, line, nl)
	default:
		leftWidth := (width + l) / 2
		rightWidth := width - leftWidth
		return fmt.Sprintf("%*s%-*s", leftWidth, line, rightWidth, nl)
	}
}

// verticalAlign returns the lines aligned vertically within the height,
// according to the alignment direction.
// If the lines are longer than the height, it is truncated.
// The empty string provided will be used to fill missing lines.
func verticalAlign(lines []string, height int, direction where, empty string) []string {
	l := len(lines)
	if l > height {
		lines = lines[:height]
	}
	switch direction {
	case top:
		// add empty lines at the bottom.
		for i := l; i < height; i++ {
			lines = append(lines, empty)
		}
	case bottom:
		// add empty lines at the top.
		for i := l; i < height; i++ {
			lines = append([]string{empty}, lines...)
		}
	default:
		// add empty lines both at the top and the bottom.
		top := (height - l) / 2
		bottom := height - l - top
		for i := 0; i < top; i++ {
			lines = append([]string{empty}, lines...)
		}
		for i := 0; i < bottom; i++ {
			lines = append(lines, empty)
		}
	}
	return lines
}

// charAt returns the character at the given position, or default if it is out of bounds.
func charAt(lines []string, y, x int, defaultValue string) string {
	if !(0 <= y && y < len(lines)) {
		return defaultValue
	}
	line := lines[y]
	if !(0 <= x && x < len(line)) {
		return defaultValue
	}
	return string(line[x])
}

// drawBox returns a box and clears its contents with spaces.
func drawBox(l []string, y1, x1, y2, x2 int) []string {
	bp := blockPos(y1, x1, y2, x2)
	box := []string{amplifyPattern("+-+", bp.Width)}
	l = replaceBlock(l, bp.Y, bp.X, box)

	for i := bp.Y + 1; i < bp.Y+bp.Height-1; i++ {
		box = []string{amplifyPattern("| |", bp.Width)}
		l = replaceBlock(l, i, bp.X, box)
	}
	box = []string{amplifyPattern("+-+", bp.Width)}
	return replaceBlock(l, bp.Y+bp.Height-1, bp.X, box)
}

// drawBoxWithLable returns a box and fills it with text.
func drawBoxWithLabel(lines []string, y1, x1, y2, x2 int, yalign, xalign where, text string) []string {
	if y1 > y2 {
		y1, y2 = y2, y1
	}
	if x1 > x2 {
		x1, x2 = x2, x1
	}
	lines = drawBox(lines, y1, x1, y2, x2)
	if x2-x1 > 2 {
		lines = fillBox(lines, y1+1, x1+2, y2-1, x2-2, yalign, xalign, text)
	}
	return lines
}

var replacer = strings.NewReplacer("\n", "", "\r", "", "\t", "    ")

// fillBox returns a rectangular area with text.
func fillBox(lines []string, y1, x1, y2, x2 int, yalign, xalign where, text string) []string {
	bp := blockPos(y1, x1, y2, x2)
	if bp.Height > 0 && bp.Width > 0 {
		text = strings.TrimSpace(replacer.Replace(text))
		linesToFill := wrap(text, bp.Width, false)
		linesToFill = verticalAlign(linesToFill, bp.Height, yalign, amplifyPattern(strings.Repeat(" ", bp.Width), bp.Width))
		for i := range linesToFill {
			linesToFill[i] = horizontalAlign(linesToFill[i], bp.Width, xalign)
		}
		return replaceBlock(lines, bp.Y, bp.X, linesToFill)
	}
	return lines
}

// -------- Line drawing --------

// reverseArrow returns the reverse of an arrow: +-> becomes <-+.
func reverseArrow(arrow string) string {
	var builder strings.Builder
	l := len(arrow)
	builder.Grow(l)
	for i := l - 1; i >= 0; i-- {
		switch arrow[i] {
		case '>':
			builder.WriteByte('<')
		case '<':
			builder.WriteByte('>')
		case '^':
			builder.WriteByte('v')
		case 'v':
			builder.WriteByte('^')
		default:
			builder.WriteByte(arrow[i])
		}
	}
	return builder.String()
}

func arrowH2V(arrow string) []string {
	// Converts an arrow to vertical. Returns a block that can be merged into lines.
	arrow = strings.ReplaceAll(strings.ReplaceAll(strings.ReplaceAll(arrow, "-", "|"), "<", "^"), ">", "v")
	block := make([]string, 0, len(arrow))
	for _, c := range arrow {
		block = append(block, string(c))
	}
	return block
}

// arrowStart replaces the beginning of an arrow with '+' if necessary.
func arrowStart(lines []string, y1, x1 int, arrow string) string {
	c := charAt(lines, y1, x1, "")
	if c == "+" || c == "|" || c == "-" {
		return "+" + arrow[1:]
	}
	return arrow
}

// drawLineHV draws an arrow between two points, always starting with the horizontal line.
func drawLineHV(lines []string, y1, x1, y2, x2 int, arrow string) []string {
	bp := blockPos(y1, x1, y2, x2)
	arrow = arrowStart(lines, y1, x1, arrow)
	if bp.Width > 1 {
		a := arrow
		if x2 < x1 {
			a = reverseArrow(arrow)
		}
		lines = overwriteBlock(lines, y1, bp.X, []string{amplifyPattern(a, bp.Width)})
	}
	if bp.Height > 1 {
		if bp.Width > 1 {
			arrow = "+" + arrow[1:]
		}
		a := arrowH2V(arrow)
		if y2 < y1 {
			a = arrowH2V(reverseArrow(arrow))
		}
		lines = overwriteBlock(lines, bp.Y, x2, amplifyPatterns(a, bp.Height))
	}
	return lines
}

// drawLineVH draws an arrow between two points, always starting with the vertical line.
func drawLineVH(lines []string, y1, x1, y2, x2 int, arrow string) []string {
	bp := blockPos(y1, x1, y2, x2)
	arrow = arrowStart(lines, y1, x1, arrow)
	if bp.Height > 1 {
		a := arrowH2V(arrow)
		if y2 < y1 {
			a = arrowH2V(reverseArrow(arrow))
		}
		lines = overwriteBlock(lines, bp.Y, x1, amplifyPatterns(a, bp.Height))
	}
	if bp.Width > 1 {
		if bp.Height > 1 {
			arrow = "+" + arrow[1:]
		}
		a := arrow
		if x2 < x1 {
			a = reverseArrow(arrow)
		}
		lines = overwriteBlock(lines, y2, bp.X, []string{amplifyPattern(a, bp.Width)})
	}
	return lines
}

// -------- Selection --------

type Box struct {
	Left, Right, Top, Bottom int
}

func (b Box) String() string {
	return fmt.Sprintf("%d,%d,%d,%d", b.Left, b.Right, b.Top, b.Bottom)
}

func findBox(lines []string, y1, x1, y2, x2 int) Box {
	// select left |.
	sx := min(x1, x2)
	for charAt(lines, y1, sx, "\n") != "|" && charAt(lines, y1, sx, "\n") != "+" && charAt(lines, y1, sx, "\n") != "\n" {
		sx--
	}
	// select right |.
	ex := max(x1, x2)
	for charAt(lines, y2, ex, "\n") != "|" && charAt(lines, y2, ex, "\n") != "+" && charAt(lines, y2, ex, "\n") != "\n" {
		ex++
	}
	// select top -.
	sy := min(y1, y2)
	for charAt(lines, sy, sx, "\n") != "-" && charAt(lines, sy, sx, "\n") != "+" && charAt(lines, sy, sx, "\n") != "\n" {
		sy--
	}
	// select bottom -.
	ey := max(y1, y2)
	for charAt(lines, ey, ex, "\n") != "-" && charAt(lines, ey, ex, "\n") != "+" && charAt(lines, ey, ex, "\n") != "\n" {
		ey++
	}
	return Box{
		Left:   sy,
		Right:  sx,
		Top:    ey,
		Bottom: ex,
	}
}

func selectOuterBox(lines []string, y1, x1, y2, x2 int) []string {
	return []string{findBox(lines, y1, x1, y2, x2).String()}
}

func selectInnerBox(lines []string, y1, x1, y2, x2 int) []string {
	box := findBox(lines, y1, x1, y2, x2)
	return []string{
		Box{
			Left:   min(box.Left+1, box.Top),
			Right:  min(box.Right+1, box.Bottom),
			Top:    max(box.Top-1, box.Left),
			Bottom: max(box.Bottom-1, box.Right),
		}.String(),
	}
}

func min[T ~int](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func max[T ~int](a, b T) T {
	if a > b {
		return a
	}
	return b
}
