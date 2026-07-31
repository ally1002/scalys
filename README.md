# Scalys

A terminal game for practicing **scales, triads, and pentatonics**, built with
Go and [Bubble Tea](https://github.com/charmbracelet/bubbletea).

Each round, Scalys picks a random root note and quality (major or natural
minor). Your job is to fill in the remaining degrees from memory — the full
scale, the triad, or the pentatonic, depending on the practice mode — and
submit your answer.

```
  SCALYS — Build the scale

  Challenge:  A# MINOR          Streak: 3

  ╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮╭──────╮
  │  A#  ││  C   ││  C#  ││  __  ││  __  ││  __  ││  __  │
  ╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯╰──────╯
     1       2       3       4       5       6       7
  Pattern: W  H  W  W  H  W

  type a-g write note · # sharpen it · enter/→ next · ← back · backspace clear · s submit · m mode settings · n new · q quit
```

After submitting, correct slots turn green and wrong ones turn red showing
`yours→expected`. A fully correct round increases your streak; a miss resets
it. `enter` (or `n`) deals the next round.

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

## Practice modes

Press `m` to open the mode-settings modal:

- **Scale** — all 7 degrees. Slot 1 is pre-filled with the root; the W/H
  pattern hint is shown.
- **Triad** — degrees 1, 3, 5. No pattern hint.
- **Pentatonic** — major omits degrees 4 and 7 (1 2 3 5 6); minor omits 2
  and 6 (1 3 4 5 7). No pattern hint.

The modal also has a **quality** setting — `Random` (default), `Major`, or
`Minor` — controlling which scale type future rounds deal.

Inside the modal: `s`/`t`/`p` pick a mode and start a round immediately,
`↑`/`↓` + `enter` browse and confirm, `←`/`→` cycle the quality, `esc`
closes.

## Controls

| Key | Action |
|---|---|
| `a`–`g` | write a natural note into the focused slot |
| `#` | sharpen the note (and advance) — press again to remove |
| `enter` / `→` | move to the next slot |
| `←` | move to the previous slot |
| `backspace` | clear the slot (or step back and clear) |
| `s` | submit (when all slots are filled) |
| `m` | open the mode-settings modal |
| `n` | new round (`enter` also works after submitting) |
| `q` / `ctrl+c` | quit |

Notes are entered by typing on purpose — there is no note picker to cycle
through, so you can't count your way to the answer. You have to know the scale.

## Notes on the music theory

- Scales are built on the chromatic circle with **sharps-only spelling**
  (`C C# D D# E F F# G G# A A# B`), so F major is spelled with `A#` rather
  than `Bb`.
- Minor means **natural minor** (`W H W W H W W`); major is `W W H W W W H`.
- Triads are the 1st, 3rd, and 5th degrees of the scale; pentatonics follow
  the standard major/minor pentatonic degree selections.
- Slot 1 is pre-filled with the root; you fill the rest.

## Web version

A standalone browser version lives in [`docs/index.html`](docs/index.html) —
one self-contained file (no build step, no dependencies) with the same rules,
practice modes, and keyboard-only input. On top of the TUI it adds:

- a **best-streak** record (persisted in `localStorage`)
- `?` — help overlay listing all commands
- `h` — music theory overlay
- `m` — mode-settings modal (same as the TUI)
- `z` — zen mode (disables the CRT flicker/blink effects)
- `t` — cycle theme: system → light → dark (persisted)

To share it: send the file directly, open it locally, or enable **GitHub
Pages** on this repo (Settings → Pages → deploy from branch `master`, folder
`/docs`) and share the URL.

## Project layout

```
main.go               entry point
internal/theory/      note set, interval patterns, scale/triad/pentatonic builders (+ tests)
internal/ui/          Bubble Tea model, update, view, styles (+ tests)
docs/index.html       standalone web version (GitHub Pages ready)
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
