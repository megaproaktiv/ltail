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

## Shorten Demo

The `shorten_demo.go` program demonstrates the line shortening feature in Saw.

### Running the Demo

```bash
go run examples/shorten_demo.go
```

### Features Demonstrated

- **Line truncation**: Lines exceeding 512 characters are automatically truncated
- **Visual marker**: Truncated lines are marked with "..." at the end
- **Multiple formats**: Shows shortening across different log formats (standard logs, JSON, etc.)

### Use Cases

The `--shorten` (or `-s`) flag is particularly useful when:

1. **Dealing with large payloads**: Logs containing large JSON objects or base64-encoded data
2. **Terminal readability**: Preventing line wrapping that clutters the terminal
3. **Quick scanning**: Making it easier to scan through logs without excessive scrolling
4. **Stack traces**: Truncating very long stack traces while keeping the beginning visible

### Integration with Saw

The shortening feature works with both commands:

```bash
# Get logs with line shortening
saw get production --shorten
saw get production -s

# Watch logs with line shortening
saw watch production --shorten
saw watch production -s

# Combine with other flags
saw get production -s --pretty --filter ERROR
saw watch production -s --prefix api
```

### Implementation

Lines are checked after colorization but before output. If a line exceeds 512 characters, it's truncated:

```go
func shortenLine(line string) string {
    const maxLength = 512
    if len(line) > maxLength {
        return line[:maxLength] + "..."
    }
    return line
}
```

This ensures you can still see the beginning of each log message while keeping the output clean and manageable.

## Dual View Demo

The `dual_demo.md` document provides comprehensive examples for using the dual split-pane view feature.

### Viewing the Demo

```bash
# Read the documentation
cat examples/dual_demo.md

# Or open in your browser/editor
open examples/dual_demo.md
```

### Features Demonstrated

- **Split-pane TUI**: Fullscreen terminal interface with two horizontal panes
- **Real-time monitoring**: Watch two log groups simultaneously
- **Interactive navigation**: Vim-style keyboard controls
- **Independent configuration**: Each pane can have different filters and settings

### Use Cases

The dual view is particularly useful when:

1. **Comparing deployments**: Monitor old vs new version side-by-side
2. **Debugging microservices**: Watch request flow through multiple services
3. **Error correlation**: Compare errors between related services
4. **Performance monitoring**: Track metrics across different components
5. **A/B testing**: Compare behavior of different variants

### Integration with Saw

The dual view feature is accessed through the `dual` command:

```bash
# Basic usage
saw dual log-group-1 log-group-2

# With filters per pane
saw dual production-api production-worker \
  --filter1 ERROR \
  --filter2 ERROR

# With different prefixes
saw dual /aws/lambda/logs /aws/lambda/logs \
  --prefix1 function-a \
  --prefix2 function-b

# With line shortening
saw dual api-gateway data-pipeline -s
```

### Keyboard Controls

- `q`, `Ctrl+C`, or `Esc` - Quit the application
- `Tab` - Switch between panes
- `↑`/`↓` or `k`/`j` - Scroll up/down in active pane
- `g` - Jump to top of active pane
- `G` - Jump to bottom of active pane

### Implementation

The dual view is built with:
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Terminal styling

Each pane independently:
- Fetches logs from CloudWatch
- Maintains scroll position
- Applies filters and formatting
- Updates every second

### Tips

1. **Start with no filters** to get an overview, then add filters as needed
2. **Use time ranges** to skip old logs: `--start1 -15m --start2 -15m`
3. **Create aliases** for common monitoring tasks
4. **Minimum terminal size**: 80x24 (recommended: 120x40 or larger)
5. **Switch panes with Tab** before trying to scroll

See [dual_demo.md](dual_demo.md) for detailed examples and scenarios.
