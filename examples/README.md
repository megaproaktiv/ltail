# Saw Examples

This directory contains example programs and demonstrations for the Saw CloudWatch Logs tool.

## Color Demo

The `color_demo.go` program demonstrates the log level colorization feature in Saw.

### Running the Demo

```bash
go run examples/color_demo.go
```

### Features Demonstrated

- **INFO colorization**: All text matching "INFO" (all caps) is displayed with a green background and black text
- **ERROR colorization**: All text matching "ERROR" (all caps) is displayed with a red background and white text

### Color Examples

The demo shows colorization across different log formats:

1. **Standard timestamps**: `2024-04-15 14:30:00 INFO Application started`
2. **RFC3339 timestamps**: `[2024-04-15T14:30:15Z] (stream-1) INFO Starting worker`
3. **JSON logs**: `{"level": "INFO", "message": "API request processed"}`

### Implementation

The colorization is implemented using the [fatih/color](https://github.com/fatih/color) package:

- `color.New(color.FgBlack, color.BgGreen)` - Green background, black foreground for INFO
- `color.New(color.FgWhite, color.BgRed)` - Red background, white foreground for ERROR

### Integration with Saw

This colorization is automatically applied to all log output in the Saw tool when using:

- `saw get <log-group>` - Get log events
- `saw watch <log-group>` - Stream log events in real-time

### Disabling Colors

If you need to disable colorization (e.g., for piping to files), you can set the `NO_COLOR` environment variable:

```bash
NO_COLOR=1 saw get my-log-group
```

Or use the fatih/color's built-in support by setting:

```bash
TERM=dumb saw get my-log-group
```
