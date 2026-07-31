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
	PentatonicMode
)

func (p PracticeMode) String() string {
	switch p {
	case TriadMode:
		return "TRIAD"
	case PentatonicMode:
		return "PENTATONIC"
	default:
		return "SCALE"
	}
}

// DegreeIndices returns the 0-based scale-degree indices a practice mode
// asks for, or nil for ScaleMode (the full 7-note scale). Pentatonic
// degrees depend on scaleType: major pentatonic omits 4 and 7 (degrees
// 1,2,3,5,6), minor pentatonic omits 2 and 6 (degrees 1,3,4,5,7).
func DegreeIndices(mode PracticeMode, scaleType ScaleType) []int {
	switch mode {
	case TriadMode:
		return []int{0, 2, 4}
	case PentatonicMode:
		if scaleType == Major {
			return []int{0, 1, 2, 4, 5}
		}
		return []int{0, 2, 3, 4, 6}
	default:
		return nil
	}
}

// DegreeNumbers returns the 1-based scale degree numbers for a practice
// mode, for labeling slots (e.g. 1,3,5 for a triad).
func DegreeNumbers(mode PracticeMode, scaleType ScaleType) []int {
	idx := DegreeIndices(mode, scaleType)
	if idx == nil {
		return []int{1, 2, 3, 4, 5, 6, 7}
	}
	nums := make([]int, len(idx))
	for i, d := range idx {
		nums[i] = d + 1
	}
	return nums
}

// Build returns the notes for a round in the given mode: the full 7-note
// scale, or the notes at the mode's degree indices (triad, pentatonic).
func Build(root int, scaleType ScaleType, mode PracticeMode) []string {
	scale := BuildScale(root, scaleType)
	idx := DegreeIndices(mode, scaleType)
	if idx == nil {
		return scale
	}
	notes := make([]string, len(idx))
	for i, d := range idx {
		notes[i] = scale[d]
	}
	return notes
}
