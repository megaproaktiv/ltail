# Dual Log View Feature

## Overview

The `dual` command provides a fullscreen split-pane terminal UI for watching two CloudWatch log groups simultaneously. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), it offers an interactive, real-time monitoring experience with two horizontal panes that each take exactly half the screen height.

## Features

- 🖥️ **Fullscreen TUI**: Immersive fullscreen terminal interface
- ↔️ **Exact Half-Screen Split**: Two panes stacked vertically, each taking exactly 50% of screen height
- 🔄 **Real-time Updates**: Auto-refreshes logs every second
- 🎯 **Independent Filtering**: Each pane can have its own filters and prefixes
- ⌨️ **Keyboard Navigation**: Vim-style navigation and controls
- 🎨 **Color Highlighting**: INFO and ERROR log levels are highlighted
- 📜 **Scrolling**: Independent scrolling for each pane
- 📊 **Message Counter**: Shows message count for each pane
- 🔚 **Tail Mode**: Automatically starts from recent logs (last 5 minutes) and auto-scrolls to bottom
- 📋 **Single-line Messages**: Newlines stripped from log messages for clean display

## Installation

The dual view feature is included in the main `saw` binary. Ensure you have the latest version:

```bash
# Build from source
task build

# Install to user bin
task install

# Verify dual command is available
saw dual --help
```

## Basic Usage

### Syntax

```bash
saw dual <log-group-1> <log-group-2> [flags]
```

### Simple Example

```bash
# Watch two production log groups
saw dual production-api production-worker
```

This opens a fullscreen interface with:
- **Top pane**: `production-api` logs (from last 5 minutes)
- **Bottom pane**: `production-worker` logs (from last 5 minutes)
- **Auto-scroll**: Automatically scrolls to show newest messages as they arrive

## Command Line Flags

### Per-Pane Flags

Configure each pane independently:

| Flag | Description |
|------|-------------|
| `--prefix1 <prefix>` | Stream prefix filter for first pane |
| `--prefix2 <prefix>` | Stream prefix filter for second pane |
| `--filter1 <pattern>` | CloudWatch filter pattern for first pane |
| `--filter2 <pattern>` | CloudWatch filter pattern for second pane |
| `--start1 <time>` | Start time for first pane |
| `--start2 <time>` | Start time for second pane |
| `--shorten1` | Shorten lines for first pane |
| `--shorten2` | Shorten lines for second pane |

### Global Flags

Apply to both panes:

| Flag | Short | Description |
|------|-------|-------------|
| `--shorten` | `-s` | Shorten lines for both panes |
| `--profile` | | AWS profile to use |
| `--region` | | AWS region override |
| `--endpoint-url` | | Custom endpoint (e.g., LocalStack) |

## Keyboard Controls

The dual view supports these keyboard shortcuts:

| Key | Action |
|-----|--------|
| `q` | Quit the application |
| `Ctrl+C` | Quit the application |
| `Esc` | Quit the application |
| `Tab` | Switch active pane |
| `↑` or `k` | Scroll up in active pane |
| `↓` or `j` | Scroll down in active pane |
| `g` | Jump to top of active pane |
| `G` | Jump to bottom of active pane |

**Note**: Keyboard shortcuts are displayed in the bottom pane's title bar.

### Active Pane Indicator

- **Active pane**: Highlighted border (bright blue)
- **Inactive pane**: Dimmed border (gray)

Only the active pane responds to scroll commands.

## Examples

### Example 1: Compare Two Environments

```bash
# Watch production vs staging
saw dual production-app staging-app
```

Compare behavior between environments in real-time.

### Example 2: Monitor Different Services

```bash
# API gateway vs Lambda functions
saw dual /aws/apigateway/api /aws/lambda/processor
```

Watch how API requests flow through to backend processing.

### Example 3: Filter Errors in Both Panes

```bash
# Show only errors from both services
saw dual production-api production-worker \
  --filter1 ERROR \
  --filter2 ERROR
```

Focus on errors across multiple services.

### Example 4: Different Filters Per Pane

```bash
# API errors vs Worker info
saw dual production-api production-worker \
  --filter1 ERROR \
  --filter2 INFO
```

Monitor errors in API while watching general worker activity.

### Example 5: Stream Prefix Filtering

```bash
# Specific Lambda functions
saw dual /aws/lambda/logs /aws/lambda/logs \
  --prefix1 order-processor \
  --prefix2 payment-handler
```

Watch two different Lambda functions from the same log group.

### Example 6: Shortened Lines

```bash
# Large payloads - shorten both panes
saw dual api-gateway data-pipeline -s
```

Keep output clean when logs contain large JSON payloads.

### Example 7: Time Range

```bash
# Start from 1 hour ago on both
saw dual production-api production-db \
  --start1 -1h \
  --start2 -1h
```

Review historical logs in both panes. Without `--start` flags, defaults to last 5 minutes (tail mode).

### Example 8: Different Time Ranges

```bash
# Different start times per pane
saw dual app-logs app-logs \
  --prefix1 server-1 \
  --start1 -30m \
  --prefix2 server-2 \
  --start2 -15m
```

Compare logs from different time periods.

## Use Cases

### 1. Deployment Monitoring

Monitor old and new versions during deployment:

```bash
saw dual production-app production-app \
  --prefix1 v1.2.3 \
  --prefix2 v1.2.4
```

### 2. Load Balancer Comparison

Compare traffic across different servers:

```bash
saw dual nginx-logs nginx-logs \
  --prefix1 server-east \
  --prefix2 server-west
```

### 3. Debugging Microservices

Watch request flow through multiple services:

```bash
saw dual api-service order-service \
  --filter1 "request_id=abc-123" \
  --filter2 "request_id=abc-123"
```

### 4. Error Correlation

Compare errors between related services:

```bash
saw dual frontend-logs backend-logs \
  --filter1 ERROR \
  --filter2 ERROR
```

### 5. Performance Monitoring

Watch performance metrics from different components:

```bash
saw dual app-performance db-performance \
  --filter1 "duration >" \
  --filter2 "query_time >"
```

### 6. A/B Testing

Compare behavior of different variants:

```bash
saw dual experiment-logs experiment-logs \
  --prefix1 variant-a \
  --prefix2 variant-b
```

### 7. Multi-Region Monitoring

Watch same service across regions:

```bash
# Terminal 1 (us-east-1)
saw dual app-logs app-logs --region us-east-1 \
  --prefix1 api \
  --prefix2 worker

# Terminal 2 (eu-west-1)
saw dual app-logs app-logs --region eu-west-1 \
  --prefix1 api \
  --prefix2 worker
```

### 8. Development vs Production

Compare same endpoints:

```bash
saw dual dev-api prod-api \
  --filter1 "/users" \
  --filter2 "/users"
```

## Technical Details

### Architecture

- **Bubble Tea TUI**: Modern terminal user interface framework
- **Exact Half-Screen Split**: Two panes stacked vertically, each exactly 50% of terminal height
- **Independent Scrolling**: Each pane maintains its own scroll position
- **Auto-refresh**: Fetches new logs every second
- **Message Limit**: Keeps last 1000 messages per pane to prevent memory issues
- **Tail Mode**: Defaults to last 5 minutes of logs, auto-scrolls to bottom on new messages
- **Single-line Display**: Newlines in log messages are replaced with spaces for clean rendering

### Display Format

Each pane shows:
- **Title bar**: Log group name and message count (bottom pane also shows keyboard shortcuts)
- **Messages**: Timestamped log entries
- **Scrollbar indicator**: Current scroll position
- **Border**: Active/inactive state

The layout uses exactly half the terminal height for each pane with no gaps.

### Message Format

```
HH:MM:SS [INFO] Log message content
HH:MM:SS [ERROR] Error message content
```

**Note**: Multi-line log messages are automatically converted to single lines (newlines replaced with spaces) for consistent display.

### Performance

- **Polling interval**: 1 second
- **Message buffer**: 1000 messages per pane
- **Memory efficient**: Auto-trims old messages
- **Network efficient**: Incremental fetches using last seen timestamp
- **Tail mode**: Starts from last 5 minutes by default (can override with `--start1/2`)
- **Auto-scroll**: Automatically scrolls to bottom when new messages arrive

## Comparison: `dual` vs `watch`

| Feature | `dual` | `watch` |
|---------|--------|---------|
| Log Groups | 2 | 1 |
| Interface | Fullscreen TUI | Streaming text |
| Scrolling | Interactive | Terminal buffer |
| Switching | Tab key | N/A |
| Layout | Horizontal split | Full width |
| Navigation | Vim-style keys | Terminal scroll |
| Best For | Comparison | Piping/scripting |

### When to Use `dual`

✅ Comparing two services  
✅ Monitoring during deployments  
✅ Debugging distributed systems  
✅ Correlation analysis  
✅ Interactive exploration  

### When to Use `watch`

✅ Single service monitoring  
✅ Piping to other tools  
✅ Background monitoring  
✅ Simple log streaming  
✅ Scripting/automation  

## Troubleshooting

### Issue: Pane Not Updating

**Symptom**: One pane shows no new messages

**Solutions**:
1. Check log group name is correct
2. Verify AWS credentials and permissions
3. Check if log group has recent logs
4. Verify filter pattern isn't too restrictive

### Issue: Terminal Size Too Small

**Symptom**: Display is garbled or overlapping

**Solution**: Resize terminal to at least 80x24 characters. The panes will automatically adjust to take exactly half the screen height each.

```bash
# Check terminal size
tput cols; tput lines

# Recommended: 120x40 or larger for comfortable viewing
```

### Issue: Logs Not Scrolling

**Symptom**: Cannot scroll up/down

**Solutions**:
1. Ensure pane is active (switch with `Tab`)
2. Check if there are enough messages to scroll
3. Try `g` to jump to top, `G` to jump to bottom
4. Note: New messages auto-scroll to bottom - manually scroll up if needed

### Issue: Colors Not Showing

**Symptom**: No color highlighting for INFO/ERROR

**Solutions**:
1. Ensure terminal supports colors
2. Check `TERM` environment variable
3. Try different terminal emulator

### Issue: Performance Slow

**Symptom**: Sluggish updates or high CPU

**Solutions**:
1. Use filters to reduce log volume
2. Use stream prefixes to narrow scope
3. Reduce terminal size if extremely large
4. Check network latency to AWS

### Issue: Cannot Quit

**Symptom**: App doesn't respond to `q`

**Solutions**:
1. Try `Ctrl+C`
2. Try `Esc`
3. Force quit with terminal kill command

## Advanced Usage

### Using with LocalStack

```bash
# Test with local CloudWatch Logs
saw dual test-group-1 test-group-2 \
  --endpoint-url http://localhost:4566 \
  --region us-east-1
```

### Combining with Shell Scripts

```bash
#!/bin/bash
# Monitor deployment
LOG_GROUP="production-app"
OLD_VERSION="v1.0.0"
NEW_VERSION="v1.1.0"

saw dual "$LOG_GROUP" "$LOG_GROUP" \
  --prefix1 "$OLD_VERSION" \
  --prefix2 "$NEW_VERSION" \
  --filter1 ERROR \
  --filter2 ERROR
```

### Multiple Profiles

```bash
# Compare across AWS accounts
# Terminal 1
saw dual app-logs app-logs --profile account-a

# Terminal 2  
saw dual app-logs app-logs --profile account-b
```

## Limitations

1. **Two panes only**: Cannot display more than 2 log groups simultaneously
2. **No search**: No built-in search functionality (use `watch` + `grep` for that)
3. **No export**: Cannot save logs from TUI (use `get` command instead)
4. **Terminal dependent**: Requires ANSI-compatible terminal
5. **Memory limit**: 1000 messages per pane (prevents unbounded growth)

## Tips & Tricks

### Tip 1: Quick Comparison Template

Create an alias for common comparisons:

```bash
alias sawcmp='saw dual'
sawcmp prod-api stage-api --filter1 ERROR --filter2 ERROR
```

### Tip 2: Focus on Recent

Start from recent time to avoid loading old logs:

```bash
saw dual app1 app2 --start1 -5m --start2 -5m
```

### Tip 3: Use Tab Completion

Enable bash/zsh completion for faster commands.

### Tip 4: Terminal Multiplexer

Use with tmux/screen for persistent sessions:

```bash
tmux new-session 'saw dual api worker'
```

### Tip 5: Save Configuration

Create shell functions for frequent monitoring:

```bash
function monitor-prod() {
  saw dual production-api production-worker \
    --filter1 ERROR \
    --filter2 ERROR \
    -s
}
```

## Future Enhancements

Potential future features:
- [ ] Vertical split option
- [ ] More than 2 panes
- [ ] Built-in search/filtering
- [ ] Export visible logs
- [ ] Custom color schemes
- [ ] Adjustable pane sizes (currently fixed at 50/50)
- [ ] Save/load sessions
- [ ] Mouse support
- [ ] Configurable help text position

## See Also

- [Main README](../README.md) - Complete saw documentation
- [Quick Reference](QUICK_REFERENCE.md) - All features at a glance
- [Shorten Feature](SHORTEN.md) - Line shortening details
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lipgloss](https://github.com/charmbracelet/lipgloss) - Styling library

---

**Version**: Saw with Dual View TUI  
**Last Updated**: 2024-04-15