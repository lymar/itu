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

Only include laziness/eagerness notes when they add meaningful information for the API user. For trivial helpers where the behavior is obvious (e.g., an empty iterator), prefer a shorter comment without mentioning laziness.

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

# Git Commit Message Guidelines

This repository uses Conventional Commits for Git commit messages. All commit messages should be written in English.

## Format

Use this format:

`type(scope): summary`

- `scope` is optional, but recommended when it adds clarity.
- Use lowercase for `type` and `scope`.

## Types

Use one of:

- `feat`: add new functionality
- `fix`: fix a bug or incorrect behavior
- `docs`: documentation-only changes
- `refactor`: code change that neither fixes a bug nor adds a feature
- `test`: add or update tests/examples
- `perf`: performance improvements
- `ci`: CI/workflow changes
- `chore`: maintenance work (tooling, cleanup, etc.)

## Scopes

Use a scope that matches the area being changed. Common scopes in this repo:

- `range`, `zip`, `reduce`, `filter`, `map`, `count`, `chain`
- `docs`, `tests`, `dev`, `ci`, `agents`

## Summary line rules

- Write the summary in the imperative mood (e.g., "add", "fix", "refactor").
- Keep the summary concise (aim for ≤ 72 characters).
- Do not end the summary with a period.

## Body rules

Add a commit body when it helps explain the change.

- Focus on the "why" and any important semantic changes.
- For iterator utilities, call out behavior changes explicitly (laziness/eagerness, stopping conditions, empty inputs, overflow/panic behavior).

## Breaking changes

If the change is breaking, mark it explicitly:

- Use `!` after `type` or `scope`: `feat(range)!: change RangeBy step==0 behavior`
- If needed, also include a `BREAKING CHANGE:` paragraph in the body describing what changed and how to migrate.

## Examples

- `feat(zip): add Zip that stops when either input ends`
- `feat(reduce): add Reduce2 for Seq2 pairs`
- `fix(range): stop iteration on overflow instead of wrapping`
- `refactor(range): move Integer constraint to range.go`
- `docs: clarify laziness vs eagerness in package docs`
- `test(zip): add ExampleZip output`
- `ci: run tests on Go 1.23`
- `chore(dev): update doc.sh instructions`
