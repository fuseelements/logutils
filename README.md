# logutils

logutils is a Go package that augments the standard library "log" package
to make logging a bit more modern, without fragmenting the Go ecosystem
with new logging packages.

## Colorization

If you use standard filter arguments in your log messages `logutils` will automatically colorise the output.

### Example:

```go
log.Print("[WARN] This is going to be red")
```

Result:

```bash
\x1b[33m[WARN] This is going to be red\n\x1b[0m
```

## Filtering

Log messages are automatically filtered based on the log message filter tags you use. Example:

```go
filter := &LevelFilter{
  Levels:   []LogLevel{"DEBUG", "INFO", "WARN", "ERROR", "CRIT"},
  MinLevel: "WARN",
  Writer:   os.Stdout,
  Color:    false,
}

logger := log.New(filter, "", 0)

logger.Print("[DEBUG] This will be filtered out")
logger.Print("[ERROR] This is an error")
```

Result:

```
[ERROR] This is an error
```

## Alignment

Messages will automatically be aligned to match the longest level code.

```
logger.Print("[WARN] foo")
logger.Print("[DEBUG] bar")
logger.Print("[ERROR] baz")
logger.Print("[INFO] baz")
```

Result:

```
[WARN]  foo
[DEBUG] bar
[ERROR] baz
[INFO]  baz
```

# Getting started

`logutils` works by replacing the output device used by the standard `log` package.

```go
package main

import (
  "gopkg.in/fuseelements/logutils.v1"
  "log"
)

func main() {
  filter := logutils.NewFilter(nil, true)

  logger := log.New(filter, "", 0)

  logger.Printf("[DEBUG] This is a debug message")
  logger.Printf("[INFO] This is an info message")
  logger.Printf("[WARN] This is a warning")
  logger.Printf("[ERROR] This is an error message")
  logger.Printf("[CRIT] Sh!t just got real")
}
```

Also check out the [examples directory](/examples).
