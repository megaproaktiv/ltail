# Color Visualization Examples

This document shows examples of how log level colorization works in Saw.

## Overview

Saw automatically colorizes log levels in your CloudWatch Logs output to make them more readable and easier to scan visually.

## Colorization Rules

| Log Level | Background Color | Text Color | Example |
|-----------|-----------------|------------|---------|
| `INFO`    | Green           | Black      | <span style="background-color: #00ff00; color: black; padding: 2px 4px;">INFO</span> |
| `ERROR`   | Red             | White      | <span style="background-color: #ff0000; color: white; padding: 2px 4px;">ERROR</span> |

## How It Works

The colorization is applied using the [fatih/color](https://github.com/fatih/color) library:

- **INFO**: `color.New(color.FgBlack, color.BgGreen).SprintFunc()`
- **ERROR**: `color.New(color.FgWhite, color.BgRed).SprintFunc()`

The colorization is case-sensitive and only matches the exact string `INFO` or `ERROR` in all capital letters.

## Example Outputs

### Standard Log Format

```
2024-04-15 14:30:00 INFO Application started successfully
2024-04-15 14:30:01 INFO Database connection established
2024-04-15 14:30:05 ERROR Failed to connect to cache server
2024-04-15 14:30:06 INFO Retrying connection with backoff...
2024-04-15 14:30:07 ERROR Connection timeout after 5 seconds
```

In your terminal, `INFO` will appear with a green background and `ERROR` with a red background.

### CloudWatch Logs Format (with saw pretty mode)

```
[2024-04-15T14:30:15Z] (api-server-1) INFO Request received: GET /api/users
[2024-04-15T14:30:16Z] (api-server-1) INFO Response sent: 200 OK
[2024-04-15T14:30:20Z] (api-server-2) ERROR Database query failed: timeout
[2024-04-15T14:30:21Z] (api-server-2) INFO Falling back to cache
[2024-04-15T14:30:22Z] (worker-1) ERROR Worker crashed: out of memory
```

### JSON Log Format

```json
{"timestamp": "2024-04-15T14:30:20Z", "level": "INFO", "message": "API request processed", "duration_ms": 45}
{"timestamp": "2024-04-15T14:30:21Z", "level": "ERROR", "message": "Invalid request payload", "error": "missing field: user_id"}
{"timestamp": "2024-04-15T14:30:22Z", "level": "INFO", "message": "Cache hit", "key": "user:12345"}
{"timestamp": "2024-04-15T14:30:23Z", "level": "ERROR", "message": "Rate limit exceeded", "client_ip": "192.168.1.100"}
```

Even in JSON logs, the `INFO` and `ERROR` text will be colorized.

## Use Cases

### Quickly Spotting Errors

When watching logs in real-time, red `ERROR` backgrounds immediately catch your attention:

```bash
saw watch production --prefix api
```

Errors will stand out among the stream of INFO messages, making it easy to spot problems as they occur.

### Filtering with Color

Combine filtering with colorization for targeted monitoring:

```bash
# Watch only errors - they'll still be colored red
saw watch production --filter ERROR

# Get logs from the last hour - INFO and ERROR will be colored
saw get production --start -1h
```

## Testing the Colors

Run the included demo to see the colorization in action:

```bash
go run examples/color_demo.go
```

This will display various log formats with colorized INFO and ERROR levels.

## Disabling Colors

### Environment Variables

If you need plain text output (for redirecting to files or parsing), you can disable colors:

```bash
# Using NO_COLOR environment variable
NO_COLOR=1 saw get production > logs.txt

# Using TERM=dumb
TERM=dumb saw watch production | grep ERROR
```

### Color Support Detection

The `fatih/color` library automatically detects if your terminal supports colors:

- **TTY detected**: Colors are enabled
- **Pipe detected**: Colors are disabled automatically
- **Windows**: Uses appropriate Windows console APIs

## Advanced Examples

### Mixed Log Levels

```
2024-04-15 14:30:00 INFO Starting application
2024-04-15 14:30:01 DEBUG Loading configuration
2024-04-15 14:30:02 INFO Configuration loaded successfully
2024-04-15 14:30:03 WARN Deprecated API endpoint used
2024-04-15 14:30:04 ERROR Failed to initialize module: auth
2024-04-15 14:30:05 FATAL Application terminated
```

Only `INFO` and `ERROR` will be colorized. Other levels (DEBUG, WARN, FATAL) remain uncolored.

### Multiple Occurrences

```
ERROR: Connection failed - ERROR code: 500 - Retrying... INFO: Retry attempt 1
```

Both occurrences of `ERROR` will be colored red, and `INFO` will be colored green.

### Case Sensitivity

```
Info: This will NOT be colored (lowercase)
info: This will NOT be colored (lowercase)
INFO: This WILL be colored (all caps)
Error: This will NOT be colored (mixed case)
ERROR: This WILL be colored (all caps)
```

## Implementation Details

The colorization is implemented in `blade/blade.go` using a simple string replacement:

```go
func colorizeLogLevel(message string) string {
    infoColor := color.New(color.FgBlack, color.BgGreen).SprintFunc()
    errorColor := color.New(color.FgWhite, color.BgRed).SprintFunc()
    
    message = strings.ReplaceAll(message, "INFO", infoColor("INFO"))
    message = strings.ReplaceAll(message, "ERROR", errorColor("ERROR"))
    
    return message
}
```

This function is applied to all log output before printing, ensuring consistent colorization across all commands.
