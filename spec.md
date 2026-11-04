# layzy — Execution/View Refactor Spec

Split the single `package main` into domain packages so that space concerns and
task concerns live in their own folders.

Module: `layzy` (from `go.mod`) → all import paths are `layzy/src/<folder>`.

---

## 1. The hard constraint

In Go **a folder is a package**. `src/space/` cannot hold methods on the current
`model` struct, because `model` lives in `package main` and its fields are
unexported.

So the split only works if **each folder owns its own state**. The folders are not
"files moved into subdirectories" — they are self-contained Bubble Tea child
models, each with its own cursor, its own data, its own `Update` and its own `View`.

This is the standard Elm/Bubble Tea nested-component pattern.

---

## 2. Target structure

```
src/
  main.go                    package main
  app/
    model.go                 package app
    update.go                package app
    view.go                  package app
  space/
    space.go                 package space
    update.go                package space
    view.go                  package space
  task/
    task.go                  package task
    update.go                package task
    view.go                  package task
  ui/
    styles.go                package ui
    overlay.go               package ui
  components/
    inputs.go                package components   (unchanged)
```

Deleted at the end: `execution-engine.go`, `view-engine.go`, `model.go`,
`styles.go`, `helpers.go` (all at `src/` root).

---

## 3. Dependency direction

```
main ──> app ──> space ──> task
             │      │        │
             └──────┴────────┴──> ui
             │
             └──> components
```

Rules that must hold:

- `task` never imports `space`. One direction only.
- `ui` imports nothing from the project. Leaf package.
- `space` and `task` never import `app`. Children never know their parent.
- `app` is the only package that knows about focus, window size and the modal.

Any import that violates these is a design error, not a convenience.

---

## 4. Package detail

### `package main` — `src/main.go`

Import path: `layzy` (root command).

**Responsibility:** process entry point only.

- Builds the root model via `app.New()` and hands it to `tea.NewProgram`.
- Keeps the existing error handling and `os.Exit(1)`.
- Contains nothing else. `Init` moves to `app`.

---

### `package app` — `src/app/`

Import path: `layzy/src/app`.

**Responsibility:** the root Bubble Tea model. Owns everything that is *not*
a space or a task: window size, which pane has focus, and the modal.
Routes messages to the focused child. Composes the two panes into one screen.

**Files**

| File | Holds |
|---|---|
| `model.go` | `Focus` type + `FocusSpaces` / `FocusTasks` constants, the root `Model` struct, `New`, `Init` |
| `update.go` | `Update` (satisfies `tea.Model`), plus unexported helpers for modal keys and app-level keys |
| `view.go` | `View` (satisfies `tea.Model`), plus an unexported modal compositor |

**State it owns**

- `Spaces` — a `space.Model`
- `Focus` — which pane receives navigation keys (replaces `isTaskPaneFocused`)
- `Inputs` — the `components.Inputs` used by the modal
- `IsAdding`, `IsDeleting` — modal flags
- `Width`, `Height` — window size

**State it must NOT own**

- No `taskCursor`. The task cursor belongs to `task.Model`.
- No `spaces.cursor` logic. That belongs to `space.Model`.
- No `selectedSpace` / `selectedTask` maps — see §6.

**Key routing (the whole point of this package)**

1. `tea.WindowSizeMsg` → store size, stop.
2. If `IsAdding`, the modal swallows every key (`enter` saves, `esc` cancels,
   everything else goes to the text input). Nothing reaches the panes.
3. App-level keys: `ctrl+c`/`q` quit, `1`/`left` focus spaces, `2`/`right`
   focus tasks, `N` opens the modal.
4. Everything else is forwarded to the focused child and the result reassigned.

---

### `package space` — `src/space/`

Import path: `layzy/src/space`.

**Responsibility:** the space list and the spaces pane. A space owns its tasks,
so this package is also the only place tasks are reached.

**Files**

| File | Holds |
|---|---|
| `space.go` | `Space` struct, `Model` struct, `New`, `Current`, `CurrentTasks`, `Add`, `Remove`, `AddTask` |
| `update.go` | `Model.Update` (space cursor keys) and `Model.UpdateTasks` (routes a message into the active space's task list) |
| `view.go` | `Model.View` — renders the spaces pane |

**Types**

- `Space` — `Name` plus a `task.Model`. Each space carries its own task list.
- `Model` — the slice of spaces plus the cursor that picks the active one.

**Why `UpdateTasks` lives here:** it is the single choke point where
`List[Cursor].Tasks` is written. Keeping it in one method is what makes
"a space owns its tasks" an enforceable invariant rather than a comment.

**`CurrentTasks` returns a usable zero value** when there are no spaces, so
`app/view.go` can render the task pane unconditionally with no nil check.

**Does not know:** focus, window size, the modal, or that an `app` package exists.
`View` receives width, height and a `focused bool` as arguments.

---

### `package task` — `src/task/`

Import path: `layzy/src/task`.

**Responsibility:** one space's task list and the tasks pane. The most isolated
package in the project — it has no idea spaces exist.

**Files**

| File | Holds |
|---|---|
| `task.go` | `Task` struct, `Model` struct, `New`, `Current`, `Add`, `Remove` |
| `update.go` | `Model.Update` — task cursor keys and done-toggle |
| `view.go` | `Model.View` — renders the tasks pane, including the empty state |

**Types**

- `Task` — replaces the current bare `string`. Carries `Title` and `Done`.
- `Model` — the slice of tasks plus **this list's own cursor**.

**The important consequence:** every space has its own `task.Model`, therefore
its own cursor. Moving between spaces no longer resets the task cursor to 0 —
each space remembers where you were. The `m.taskCursor = 0` lines in the current
`execution-engine.go` are deleted, not ported.

**Must never import `space`.** If it needs something from a space, the dependency
is pointing the wrong way.

---

### `package ui` — `src/ui/`

Import path: `layzy/src/ui`.

**Responsibility:** presentation primitives shared by `app`, `space` and `task`.
A leaf package — imports only `lipgloss` and `ansi`.

**Files**

| File | Holds |
|---|---|
| `styles.go` | the current `styles.go` with names exported: `TitleStyle`, `BasePaneStyle`, `FocusedPaneStyle`, `ModalStyle`, `ModalHeaderStyle`, `ItemStyle`, `SelectedItemStyle`, `HelpStyle`, plus a `PaneStyle(focused bool)` helper |
| `overlay.go` | the current `helpers.go` with `placeOverlay` → `PlaceOverlay` and `dimBackground` → `DimBackground` |

`PaneStyle(focused)` exists so `space` and `task` stop duplicating the
base-vs-focused branch.

Two things to fix while moving `styles.go`:
- `modalStyle` sets `BorderStyle(lipgloss.RoundedBorder())` twice — drop one.
- Its border colour is `"FF0055"`, missing the leading `#`.

---

### `package components` — `src/components/`

Import path: `layzy/src/components`. **Unchanged.**

Holds `Inputs`. Only `app` imports it, since only `app` owns the modal.

---

## 5. Naming conventions

- **No `-engine` suffix inside the packages.** Once the package scopes the name,
  `space.Update` reads better than `space.ExecutionEngine`. Use `update.go` and
  `view.go`.
- **No hyphens in Go filenames.** Also avoid underscores — `_test.go`,
  `_darwin.go` and `_linux.go` are build-constraint suffixes, so
  `space_update.go` is a name to stay away from.
- **Do not stutter.** Inside `package space` the state type is `Model`, not
  `SpaceModel` — the call site reads `space.Model`. Same for `task.Model`.
  Today's `SpaceList` becomes `space.Model`; `Space` stays `space.Space`.
- Child `Update` returns its **own** concrete type, not `tea.Model`. Only
  `app.Model.Update` returns `tea.Model`, because only it satisfies the
  Bubble Tea interface.
- `View` on a child takes `(width, height int, focused bool)` and returns a
  `string`. Only `app.Model.View` returns a `tea.View`.

---

## 6. Deliberate behaviour changes

Flagging these so they are choices, not accidents. Reject any you disagree with
before starting.

1. **`selectedSpace` and `selectedTask` are dropped.**
   - `selectedTask` is replaced by `Done` on the `Task` struct. Storing selection
     as `map[int]struct{}` means every delete has to re-index the map; a field on
     the item has no such problem.
   - `selectedSpace` is removed outright — it is never read by the current view,
     and "the selected space" is already expressed by the space cursor.
   - If multi-select on spaces is actually wanted later, add a `Selected bool`
     to `Space`, not a map on `Model`.

2. **`enter` / `space` now acts on the focused pane.** Today it toggles
   `selectedSpace` regardless of which pane has focus. After routing, it toggles
   a task's `Done` when the task pane is focused.

3. **The task cursor is per-space** and no longer reset on space change (§4,
   `package task`).

4. **`isTaskPaneFocused bool` becomes a `Focus` enum** in `app`. A bool is fine
   for two panes but has no room for a third.

5. **Tasks become `Task` structs instead of `[]string`**, which is what the
   commented-out code in `execution-engine.go` was already reaching for.

Unchanged on purpose: `Width`/`Height` stay `uint` with casts at the lipgloss
call sites, exactly as today.

---

## 7. Migration steps

Go will not compile the intermediate states, so do the whole move before building.

**Step 1 — `ui` (no dependencies, do it first)**
1. `src/ui/styles.go`: move the contents of `src/styles.go`, export every name,
   add `PaneStyle`.
2. `src/ui/overlay.go`: move the contents of `src/helpers.go`, export both
   functions.

**Step 2 — `task` (depends only on `ui`)**
3. `src/task/task.go`: define `Task` and `Model`, plus `New`, `Current`, `Add`,
   `Remove`.
4. `src/task/update.go`: port the `up` / `down` branches that sit under
   `if m.isTaskPaneFocused` in `execution-engine.go`, plus the done-toggle.
   Drop the focus check — the parent does that now.
5. `src/task/view.go`: port the tasks-pane block from `view-engine.go`, including
   the `No tasks available.` empty state.

**Step 3 — `space` (depends on `task` and `ui`)**
6. `src/space/space.go`: define `Space` and `Model`, plus `New`, `Current`,
   `CurrentTasks`, `Add`, `Remove`, `AddTask`.
7. `src/space/update.go`: port the `else` branches of `up` / `down`, minus the
   `m.taskCursor = 0` resets. Add `UpdateTasks`.
8. `src/space/view.go`: port the spaces-pane block from `view-engine.go`.

**Step 4 — `app` (depends on everything)**
9. `src/app/model.go`: define `Focus`, the root `Model`, `New` (the seed data
   from today's `initialModel`), and `Init`.
10. `src/app/update.go`: `Update` handling `WindowSizeMsg`, then the modal, then
    app-level keys, then delegation. Un-comment and restore the modal save/cancel
    logic that is currently commented out in `execution-engine.go`.
11. `src/app/view.go`: the size guard, the pane-width arithmetic, the
    `JoinHorizontal`, the help line, and the modal overlay branch.

**Step 5 — entry point**
12. `src/main.go`: reduce to `main` only; it calls `app.New()`.

**Step 6 — remove the old files**
13. Delete `src/execution-engine.go`, `src/view-engine.go`, `src/model.go`,
    `src/styles.go`, `src/helpers.go`.

**Step 7 — verify**
14. `go build ./...`
15. `go vet ./...`
16. Run it and check: pane switching with `1`/`2`/`←`/`→`, `↑`/`↓` in each pane,
    `N` opens the modal and `enter` adds to the *current* space, `esc` cancels,
    task cursor is remembered per space, `q` quits.

---

## 8. Done when

- [ ] `go build ./...` and `go vet ./...` are clean.
- [ ] No file at `src/` root except `main.go`.
- [ ] `grep -r "layzy/src/app" src/space src/task src/ui` returns nothing.
- [ ] `grep -r "layzy/src/space" src/task` returns nothing.
- [ ] `src/ui` imports no `layzy/...` path.
- [ ] Every behaviour in §7 step 16 works as it does today, except the
      deliberate changes in §6.
