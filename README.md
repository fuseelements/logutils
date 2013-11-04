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

# Getting started

`logutils` works by replacing the output device used by the standard `log` package.

```go
package main

import(
  "log"
  "github.com/appio/logutils"
)

func main(){
  filter := &LevelFilter{
    Levels:   []LogLevel{"DEBUG", "INFO", "WARN", "ERROR", "CRIT"},
    MinLevel: "WARN",
    Writer:   os.Stdout,
    Color:    true,
  }

  logger := log.New(filter, "", 0)

  logger.Printf("[DEBUG] Started logging service")
}
```
