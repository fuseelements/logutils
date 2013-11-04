# logutils

logutils is a Go package that augments the standard library "log" package
to make logging a bit more modern, without fragmenting the Go ecosystem
with new logging packages.

## Automatic Colorisation

If you use standard filter arguments in your log messages `logutils` will automatically colorise the output.

### Example:

```go
log.Print("[WARN] This is going to be red")
```

Result:

```bash
\x1b[33m[WARN] This is going to be red\n\x1b[0m
```
