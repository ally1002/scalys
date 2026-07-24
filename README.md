# Scalys

A terminal game for practicing **major and minor scales**, built with Go and
[Bubble Tea](https://github.com/charmbracelet/bubbletea).

Each round, Scalys picks a random root note and a scale type (major or natural
minor). Your job is to fill in the remaining six degrees of the scale from
memory, following the correct tone/semitone pattern, and submit your answer.

```
  SCALYS — Build the scale

  Challenge:  A# MINOR          Streak: 3

  ╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮
  │  A#  ││  C   ││  C#  ││  __  ││  __  ││  __  ││  __  │
  ╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯
     1       2       3       4       5       6       7
  Pattern: W  H  W  W  H  W

  type a-g write note · # sharpen it · enter/→ next · ← back · backspace clear · s submit · n new · q quit
```

After submitting, correct slots turn green and wrong ones turn red showing
`yours→expected`. A fully correct round increases your streak; a miss resets it.

## Install & Run

Requires Go 1.22+.

```sh
git clone <repo-url>
cd scalys
go run .
```

Or build a binary:

```sh
go build -o scalys .
./scalys
```

## Controls

| Key | Action |
|---|---|
| `a`–`g` | write a natural note into the focused slot |
| `#` | sharpen the note (and advance) — press again to remove |
| `enter` / `→` | move to the next slot |
| `←` | move to the previous slot |
| `backspace` | clear the slot (or step back and clear) |
| `s` | submit (when all slots are filled) |
| `n` | new round |
| `q` / `ctrl+c` | quit |

Notes are entered by typing on purpose — there is no note picker to cycle
through, so you can't count your way to the answer. You have to know the scale.

## Notes on the music theory

- Scales are built on the chromatic circle with **sharps-only spelling**
  (`C C# D D# E F F# G G# A A# B`), so F major is spelled with `A#` rather
  than `Bb`.
- Minor means **natural minor** (`W H W W H W W`); major is `W W H W W W H`.
- Slot 1 is pre-filled with the root; you fill degrees 2–7.

## Project layout

```
main.go               entry point
internal/theory/      note set, interval patterns, BuildScale (+ tests)
internal/ui/          Bubble Tea model, update, view, styles (+ tests)
```

Run the tests with:

```sh
go test ./...
```

## Roadmap ideas

- Harmonic and melodic minor modes
- Proper flat spellings (F major → Bb)
- Practice a specific key via a CLI flag
- Hide the W/H pattern hint for a harder mode
