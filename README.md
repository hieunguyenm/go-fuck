# go-fuck

Brainfuck interpreter in Go.

## Usage

```bash
go run gofuck.go [args...]
```

| Argument  |                   Description                    | Default |
| :-------: | :----------------------------------------------: | :-----: |
|  `-file`  |             Path to Brainfuck source             | `stdin` |
| `-length` | Length of tape, expanded automatically if needed |  30000  |
| `-pause`  |            Delay after console output            |  200ms  |
