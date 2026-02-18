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

## Notes

- Stack order shown is `bottom -> top`.
- The top element is the last item in the slice.
