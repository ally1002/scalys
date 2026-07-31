package theory

import (
	"reflect"
	"testing"
)

func TestBuildScaleMajor(t *testing.T) {
	want := map[int][]string{
		0:  {"C", "D", "E", "F", "G", "A", "B"},
		1:  {"C#", "D#", "F", "F#", "G#", "A#", "C"},
		2:  {"D", "E", "F#", "G", "A", "B", "C#"},
		3:  {"D#", "F", "G", "G#", "A#", "C", "D"},
		4:  {"E", "F#", "G#", "A", "B", "C#", "D#"},
		5:  {"F", "G", "A", "A#", "C", "D", "E"},
		6:  {"F#", "G#", "A#", "B", "C#", "D#", "F"},
		7:  {"G", "A", "B", "C", "D", "E", "F#"},
		8:  {"G#", "A#", "C", "C#", "D#", "F", "G"},
		9:  {"A", "B", "C#", "D", "E", "F#", "G#"},
		10: {"A#", "C", "D", "D#", "F", "G", "A"},
		11: {"B", "C#", "D#", "E", "F#", "G#", "A#"},
	}
	for root, expected := range want {
		got := BuildScale(root, Major)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("%s major: got %v, want %v", Notes[root], got, expected)
		}
	}
}

func TestBuildScaleMinor(t *testing.T) {
	want := map[int][]string{
		0:  {"C", "D", "D#", "F", "G", "G#", "A#"},
		1:  {"C#", "D#", "E", "F#", "G#", "A", "B"},
		2:  {"D", "E", "F", "G", "A", "A#", "C"},
		3:  {"D#", "F", "F#", "G#", "A#", "B", "C#"},
		4:  {"E", "F#", "G", "A", "B", "C", "D"},
		5:  {"F", "G", "G#", "A#", "C", "C#", "D#"},
		6:  {"F#", "G#", "A", "B", "C#", "D", "E"},
		7:  {"G", "A", "A#", "C", "D", "D#", "F"},
		8:  {"G#", "A#", "B", "C#", "D#", "E", "F#"},
		9:  {"A", "B", "C", "D", "E", "F", "G"},
		10: {"A#", "C", "C#", "D#", "F", "F#", "G#"},
		11: {"B", "C#", "D", "E", "F#", "G", "A"},
	}
	for root, expected := range want {
		got := BuildScale(root, Minor)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("%s minor: got %v, want %v", Notes[root], got, expected)
		}
	}
}

func TestBuildTriad(t *testing.T) {
	want := map[int][]string{
		0: {"C", "E", "G"},
		9: {"A", "C#", "E"},
	}
	for root, expected := range want {
		got := Build(root, Major, TriadMode)
		if !reflect.DeepEqual(got, expected) {
			t.Errorf("%s major triad: got %v, want %v", Notes[root], got, expected)
		}
	}
	if got := Build(9, Minor, TriadMode); !reflect.DeepEqual(got, []string{"A", "C", "E"}) {
		t.Errorf("A minor triad: got %v", got)
	}
}

func TestBuildScaleModeMatchesBuildScale(t *testing.T) {
	if got := Build(3, Minor, ScaleMode); !reflect.DeepEqual(got, BuildScale(3, Minor)) {
		t.Errorf("ScaleMode Build diverged from BuildScale: got %v", got)
	}
}

func TestPatternSymbols(t *testing.T) {
	if got := Major.PatternSymbols(); !reflect.DeepEqual(got, []string{"W", "W", "H", "W", "W", "W", "H"}) {
		t.Errorf("major symbols: got %v", got)
	}
	if got := Minor.PatternSymbols(); !reflect.DeepEqual(got, []string{"W", "H", "W", "W", "H", "W", "W"}) {
		t.Errorf("minor symbols: got %v", got)
	}
}
