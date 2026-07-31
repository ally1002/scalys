package ui

import (
	"strings"
	"testing"

	tea "github.com/charmbracelet/bubbletea"

	"scalys/internal/theory"
)

func press(m model, keys ...string) model {
	for _, k := range keys {
		var msg tea.KeyMsg
		switch k {
		case "enter":
			msg = tea.KeyMsg{Type: tea.KeyEnter}
		case "left":
			msg = tea.KeyMsg{Type: tea.KeyLeft}
		case "right":
			msg = tea.KeyMsg{Type: tea.KeyRight}
		case "backspace":
			msg = tea.KeyMsg{Type: tea.KeyBackspace}
		default:
			msg = tea.KeyMsg{Type: tea.KeyRunes, Runes: []rune(k)}
		}
		updated, _ := m.Update(msg)
		m = updated.(model)
	}
	return m
}

func TestSubmitCorrectAnswer(t *testing.T) {
	m := New()
	copy(m.slots[:], m.answer)
	m.submit()

	if m.state != stateSubmitted {
		t.Fatal("expected submitted state")
	}
	for i, r := range m.results {
		if !r {
			t.Errorf("slot %d marked wrong for a correct answer", i)
		}
	}
	if m.streak != 1 {
		t.Errorf("streak = %d, want 1", m.streak)
	}
}

func TestSubmitWrongAnswerResetsStreak(t *testing.T) {
	m := New()
	m.streak = 3
	copy(m.slots[:], m.answer)
	m.slots[3] = "nope"
	m.submit()

	if m.results[3] {
		t.Error("wrong slot marked correct")
	}
	if m.streak != 0 {
		t.Errorf("streak = %d, want 0", m.streak)
	}
}

func TestNewRoundPrefillsRoot(t *testing.T) {
	m := New()
	if m.slots[0] != theory.Notes[m.root] {
		t.Errorf("slot 0 = %q, want root %q", m.slots[0], theory.Notes[m.root])
	}
	for i := 1; i < 7; i++ {
		if m.slots[i] != "" {
			t.Errorf("slot %d = %q, want empty", i, m.slots[i])
		}
	}
	if m.cursor != 1 {
		t.Errorf("cursor = %d, want 1", m.cursor)
	}
}

func TestTypingDoesNotAdvance(t *testing.T) {
	m := press(New(), "g")
	if m.slots[1] != "G" {
		t.Fatalf("slot 1 = %q, want G", m.slots[1])
	}
	if m.cursor != 1 {
		t.Errorf("cursor = %d, want 1 (typing must not auto-advance)", m.cursor)
	}

	m = press(m, "a")
	if m.slots[1] != "A" {
		t.Errorf("slot 1 = %q, want A (retyping overwrites)", m.slots[1])
	}
}

func TestEnterAndArrowsAdvance(t *testing.T) {
	m := press(New(), "g", "enter")
	if m.cursor != 2 {
		t.Errorf("cursor = %d, want 2 after enter", m.cursor)
	}
	m = press(m, "a", "right")
	if m.cursor != 3 {
		t.Errorf("cursor = %d, want 3 after right arrow", m.cursor)
	}
	m = press(m, "left")
	if m.cursor != 2 {
		t.Errorf("cursor = %d, want 2 after left arrow", m.cursor)
	}
}

func TestSharpAdvancesAndToggles(t *testing.T) {
	m := press(New(), "c", "#")
	if m.slots[1] != "C#" {
		t.Fatalf("slot 1 = %q, want C#", m.slots[1])
	}
	if m.cursor != 2 {
		t.Errorf("cursor = %d, want 2 (sharp auto-advances)", m.cursor)
	}

	m = press(m, "left", "#")
	if m.slots[1] != "C" {
		t.Errorf("slot 1 = %q, want C after un-sharpen", m.slots[1])
	}
	if m.cursor != 1 {
		t.Errorf("cursor = %d, want 1 (un-sharpen must not advance)", m.cursor)
	}
}

func TestSharpenSkipsEAndB(t *testing.T) {
	for _, note := range []string{"e", "b"} {
		m := press(New(), note, "#")
		want := strings.ToUpper(note)
		if m.slots[1] != want {
			t.Errorf("slot 1 = %q, want %q (no %s# in sharps-only set)", m.slots[1], want, want)
		}
		if m.cursor != 1 {
			t.Errorf("cursor = %d, want 1 (%s can't be sharpened)", m.cursor, want)
		}
	}
}

func TestSharpOnEmptySlotIsNoop(t *testing.T) {
	m := press(New(), "#")
	if m.slots[1] != "" {
		t.Errorf("slot 1 = %q, want empty", m.slots[1])
	}
	if m.cursor != 1 {
		t.Errorf("cursor = %d, want 1", m.cursor)
	}
}

func TestBackspaceClearsAndMovesBack(t *testing.T) {
	m := press(New(), "c", "enter", "d", "backspace")
	if m.slots[2] != "" {
		t.Errorf("slot 2 = %q, want empty", m.slots[2])
	}
	m = press(m, "backspace")
	if m.cursor != 1 || m.slots[1] != "" {
		t.Errorf("cursor = %d, slot 1 = %q; want cursor 1 and empty slot", m.cursor, m.slots[1])
	}
}

func TestFullRoundViaKeys(t *testing.T) {
	m := New()
	for _, note := range m.answer[1:] {
		keys := []string{strings.ToLower(note[:1])}
		if strings.HasSuffix(note, "#") {
			keys = append(keys, "#")
		} else {
			keys = append(keys, "enter")
		}
		m = press(m, keys...)
	}
	m = press(m, "s")
	if m.state != stateSubmitted {
		t.Fatal("expected submitted state after s")
	}
	if m.streak != 1 {
		t.Errorf("streak = %d, want 1 (typed the correct scale)", m.streak)
	}
}

func TestToggleModeSwitchesToTriad(t *testing.T) {
	m := press(New(), "m")
	if m.mode != theory.TriadMode {
		t.Fatal("expected triad mode after pressing m")
	}
	if len(m.answer) != 3 || len(m.slots) != 3 || len(m.results) != 3 {
		t.Errorf("triad round should have 3 notes, got answer=%d slots=%d results=%d", len(m.answer), len(m.slots), len(m.results))
	}
	if m.slots[0] != m.answer[0] {
		t.Errorf("slot 0 = %q, want root %q", m.slots[0], m.answer[0])
	}

	m = press(m, "m")
	if m.mode != theory.ScaleMode {
		t.Fatal("expected scale mode after toggling back")
	}
	if len(m.answer) != 7 {
		t.Errorf("scale round should have 7 notes, got %d", len(m.answer))
	}
}

func TestFullTriadRoundViaKeys(t *testing.T) {
	m := press(New(), "m")
	for _, note := range m.answer[1:] {
		keys := []string{strings.ToLower(note[:1])}
		if strings.HasSuffix(note, "#") {
			keys = append(keys, "#")
		} else {
			keys = append(keys, "enter")
		}
		m = press(m, keys...)
	}
	m = press(m, "s")
	if m.state != stateSubmitted {
		t.Fatal("expected submitted state after s")
	}
	if m.streak != 1 {
		t.Errorf("streak = %d, want 1 (typed the correct triad)", m.streak)
	}
}

func TestViewRenders(t *testing.T) {
	m := New()
	out := m.View()
	for _, want := range []string{"SCALYS", "Challenge:", "Pattern:"} {
		if !strings.Contains(out, want) {
			t.Errorf("view missing %q", want)
		}
	}

	copy(m.slots[:], m.answer)
	m.submit()
	if !strings.Contains(m.View(), "Correct!") {
		t.Error("submitted view missing success message")
	}
}
