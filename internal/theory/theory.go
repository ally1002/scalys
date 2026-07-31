package theory

// Notes is the chromatic circle using sharps-only spelling.
var Notes = [12]string{"C", "C#", "D", "D#", "E", "F", "F#", "G", "G#", "A", "A#", "B"}

type ScaleType int

const (
	Major ScaleType = iota
	Minor
)

func (s ScaleType) String() string {
	if s == Major {
		return "MAJOR"
	}
	return "MINOR"
}

// Pattern returns the interval pattern in semitones for the scale type.
func (s ScaleType) Pattern() []int {
	if s == Major {
		return []int{2, 2, 1, 2, 2, 2, 1} // W W H W W W H
	}
	return []int{2, 1, 2, 2, 1, 2, 2} // W H W W H W W (natural minor)
}

// PatternSymbols returns the pattern as W/H symbols for display.
func (s ScaleType) PatternSymbols() []string {
	syms := make([]string, 0, 7)
	for _, step := range s.Pattern() {
		if step == 2 {
			syms = append(syms, "W")
		} else {
			syms = append(syms, "H")
		}
	}
	return syms
}

// BuildScale returns the 7 note names of the scale starting at root
// (an index into Notes).
func BuildScale(root int, scaleType ScaleType) []string {
	scale := make([]string, 0, 7)
	idx := root % 12
	scale = append(scale, Notes[idx])
	for _, step := range scaleType.Pattern()[:6] {
		idx = (idx + step) % 12
		scale = append(scale, Notes[idx])
	}
	return scale
}

// PracticeMode selects what a round asks the player to build.
type PracticeMode int

const (
	ScaleMode PracticeMode = iota
	TriadMode
)

func (p PracticeMode) String() string {
	if p == TriadMode {
		return "TRIAD"
	}
	return "SCALE"
}

// Build returns the notes for a round in the given mode: the full 7-note
// scale, or for TriadMode just the root, third, and fifth (scale degrees
// 1, 3, 5).
func Build(root int, scaleType ScaleType, mode PracticeMode) []string {
	scale := BuildScale(root, scaleType)
	if mode == TriadMode {
		return []string{scale[0], scale[2], scale[4]}
	}
	return scale
}
