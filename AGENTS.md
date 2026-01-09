# Documentation Comment Guidelines (GoDoc)

This repository uses standard GoDoc conventions. When adding new exported APIs or package-level docs, follow these rules so that `pkg.go.dev` and `pkgsite` render consistent, helpful documentation.

## Scope

- **Exported identifiers** (packages, types, constants, variables, functions, methods) should have a GoDoc comment.
- **Unexported helpers** generally do not need doc comments unless the code is non-obvious.

## Format

- Use line comments (`//`), not block comments.
- The first line **must start with the identifier name** (or `Package <name>` for package docs).
  - Good: `// RangeBy returns a lazy iterator ...`
  - Good: `// Package itu provides ...`
- Wrap the comment as a short paragraph (multiple `//` lines). Use a blank `//` line to separate paragraphs.
- Prefer clear, direct sentences. Avoid marketing language.

## Content rules for this repo

Most functions here operate on `iter.Seq` / `iter.Seq2` and have important semantic differences. Document them explicitly:

### Laziness vs eagerness

- If the function returns an iterator, say that it is **lazy** and that values are produced only as the returned iterator is consumed.
- If the function consumes the input iterator to produce a result, say it **consumes eagerly**.

### Iteration behavior

- Describe **what each yielded element/pair represents**.
- For `Seq2`-based utilities, use consistent wording such as “pairs (k, v)” or “pairs (a, b)”.
- If an operation truncates or stops early, state the stopping condition precisely (e.g., “stops as soon as either sequence runs out of values”).

### Edge cases and special values

- Document notable boundary behaviors:
  - Empty inputs (what is returned/yielded)
  - Inclusive vs half-open ranges (`[start, end)` vs `[start, end]`)
  - Directional behavior when steps can be negative
  - Special-case handling (e.g., `step == 0` producing no values vs an infinite sequence until the consumer stops)

### Panics, overflow, and “infinite” sequences

- If a function can panic, document **when** and **why** (e.g., overflow). Keep the panic condition in the comment even if the panic message is self-explanatory.
- If overflow stops an iteration (rather than wrapping), state that explicitly.
- When an iterator may be effectively infinite, prefer describing behavior in terms of “until the consumer stops” rather than claiming it is truly infinite.

### Reference-type mutation note

- If an API takes/returns accumulators or values that may be mutated, add a short “Note:” paragraph when it is relevant (e.g., reference types like map/slice/pointer).

## Examples

- Add Go examples in `*_example_test.go` files using `ExampleXxx` functions.
- Examples should be small and show typical usage.
- Include `// Output:` blocks for deterministic output.
- If an example uses a standard library helper whose behavior isn’t obvious (e.g., `slices.All`, `maps.All`), add a short explanatory comment.

## Style checklist (quick)

Before submitting, verify:

- The comment starts with the exported identifier name.
- The first sentence reads well on `pkg.go.dev` (it should stand alone).
- Laziness/eagerness is explicitly stated where applicable.
- Edge cases and stopping conditions are covered when non-trivial.
- Panics/overflow behavior is documented if applicable.
