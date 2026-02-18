# Go Stack Visualization

Simple Go example to demonstrate stack operations using a slice.

## What it does

- Creates an input slice with 10 numbers (`1..10`).
- Pushes each number to the stack.
- After each `push`:
  - Prints the current stack.
  - Waits for 1 second.
- Pops items one by one until the stack is empty.
- After each `pop`:
  - Prints the popped value.
  - Prints the current stack.
  - Waits for 1 second.

## Run

```bash
go run main.go
```

## Run with Docker

Build the image:

```bash
docker build -t go-stack .
```

Run the container:

```bash
docker run --rm go-stack
```

## Practice releases

Use this quick loop to practice versioning and release notes:

1. Use Conventional Commits in PR titles/commits (examples: `feat: add stack peek`, `fix: handle empty pop`).
2. Merge to `main`.
3. Wait for the `release-please` workflow to open or update a Release PR.
4. Review the generated version bump and `CHANGELOG.md` changes in that PR.
5. Merge the Release PR.

After merge, GitHub will automatically create the tag and GitHub Release.

Versioning details: see `VERSIONING.md`.

## Notes

- Stack order shown is `bottom -> top`.
- The top element is the last item in the slice.
