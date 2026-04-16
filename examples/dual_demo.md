# Dual View Demo

This document provides hands-on examples for using the `saw dual` command to monitor two CloudWatch log groups simultaneously in a split-pane terminal UI.

## Quick Start

The simplest way to start:

```bash
saw dual log-group-1 log-group-2
```

This opens a fullscreen interface with two horizontal panes showing real-time logs from both groups.

## Basic Examples

### Example 1: Monitor Two Services

Watch API and Worker services side-by-side:

```bash
saw dual production-api production-worker
```

**What you'll see:**
- Top pane: Live logs from `production-api`
- Bottom pane: Live logs from `production-worker`
- Real-time updates every second
- Message count in each title bar

### Example 2: Compare Errors

Focus on errors from both services:

```bash
saw dual production-api production-worker \
  --filter1 ERROR \
  --filter2 ERROR
```

**Use case:** Quickly spot if errors are correlated between services.

### Example 3: Different Log Groups, Different Filters

```bash
saw dual frontend-logs backend-logs \
  --filter1 "status=404" \
  --filter2 "database error"
```

**Use case:** Monitor specific issues in each service.

## Deployment Monitoring

### Example 4: Blue-Green Deployment

Monitor old (blue) and new (green) versions:

```bash
saw dual production production \
  --prefix1 blue-v1.2.3 \
  --prefix2 green-v1.2.4
```

**Keyboard navigation:**
1. Press `Tab` to switch between panes
2. Use `↑`/`↓` or `k`/`j` to scroll through messages
3. Press `g` to jump to top, `G` to bottom
4. Press `q` to quit

### Example 5: Canary Deployment

Compare canary vs stable:

```bash
saw dual production-app production-app \
  --prefix1 stable \
  --prefix2 canary \
  --filter1 ERROR \
  --filter2 ERROR \
  -s
```

The `-s` flag shortens long lines in both panes for cleaner output.

## AWS Lambda Monitoring

### Example 6: Multiple Lambda Functions

```bash
saw dual /aws/lambda/my-functions /aws/lambda/my-functions \
  --prefix1 order-processor \
  --prefix2 payment-handler
```

**Scenario:** Track order flow from processing to payment.

### Example 7: Lambda with Different Time Ranges

```bash
saw dual /aws/lambda/processor /aws/lambda/processor \
  --prefix1 recent-run \
  --start1 -10m \
  --prefix2 previous-run \
  --start2 -30m
```

**Use case:** Compare current behavior with past behavior.

## Microservices Debugging

### Example 8: Request Tracing

Follow a request through multiple services:

```bash
saw dual api-gateway order-service \
  --filter1 "request_id=abc-xyz-123" \
  --filter2 "request_id=abc-xyz-123"
```

**Tip:** Switch between panes with `Tab` to follow the request flow.

### Example 9: API + Database

```bash
saw dual api-server database-logs \
  --filter1 "POST /orders" \
  --filter2 "INSERT INTO orders"
```

**Pattern:** See API calls and corresponding database operations.

## Performance Monitoring

### Example 10: Slow Queries

```bash
saw dual app-logs db-logs \
  --filter1 "duration > 1000" \
  --filter2 "slow query"
```

**Goal:** Correlate slow application requests with database performance.

### Example 11: Error Rate Comparison

```bash
saw dual us-east-1-api us-west-2-api \
  --filter1 ERROR \
  --filter2 ERROR \
  --region us-east-1
```

**Multi-region:** Compare error rates across regions.

## Development Workflows

### Example 12: Dev vs Prod

```bash
saw dual development-api production-api \
  --filter1 "/users/123" \
  --filter2 "/users/123"
```

**Verify:** Ensure dev environment matches production behavior.

### Example 13: Testing Feature Flags

```bash
saw dual app-logs app-logs \
  --prefix1 feature-enabled \
  --prefix2 feature-disabled \
  --filter1 "checkout process" \
  --filter2 "checkout process"
```

**A/B testing:** Compare feature flag variants.

## Advanced Usage

### Example 14: Shortened Lines for Large Payloads

```bash
saw dual api-gateway data-pipeline \
  --shorten1 \
  --shorten2
```

Or simply:

```bash
saw dual api-gateway data-pipeline -s
```

**When:** Logs contain large JSON payloads or base64 data.

### Example 15: Individual Pane Configuration

```bash
saw dual service-a service-b \
  --prefix1 production \
  --prefix2 staging \
  --filter1 ERROR \
  --filter2 "INFO|WARN|ERROR" \
  --start1 -1h \
  --start2 -30m \
  --shorten1
```

**Full control:** Each pane with completely different settings.

## Real-World Scenarios

### Scenario 1: Incident Response

**Problem:** Production API showing errors

```bash
saw dual production-api production-worker \
  --filter1 ERROR \
  --filter2 ERROR
```

**Actions:**
1. Open dual view with error filters
2. Switch between panes with `Tab` to see error patterns
3. Scroll through history to find when errors started
4. Identify if worker errors correlate with API errors

### Scenario 2: Post-Deployment Verification

**Problem:** Just deployed new version, need to verify

```bash
saw dual production production \
  --prefix1 old-version-1.2.3 \
  --prefix2 new-version-1.2.4 \
  --filter1 ERROR \
  --filter2 ERROR
```

**Checklist:**
- [ ] Watch for 5-10 minutes
- [ ] Verify new version has no new errors
- [ ] Compare error rates between versions
- [ ] Check for warning patterns

### Scenario 3: Performance Investigation

**Problem:** Users reporting slow responses

```bash
saw dual nginx-access app-logs \
  --filter1 "response_time > 2000" \
  --filter2 "duration"
```

**Investigation:**
1. Top pane shows slow nginx responses
2. Bottom pane shows corresponding app duration logs
3. Use timestamps to correlate slow requests
4. Scroll to find patterns

### Scenario 4: Database Connection Issues

**Problem:** Intermittent database errors

```bash
saw dual app-server db-proxy \
  --filter1 "connection" \
  --filter2 "connection" \
  -s
```

**Debug flow:**
1. Watch both application and database proxy
2. Look for connection errors
3. Check timing - do app errors precede DB errors?
4. Identify connection pool exhaustion

## Tips for Effective Use

### Tip 1: Start Broad, Then Filter

```bash
# Start with no filters to get overview
saw dual service-a service-b

# Then add filters if too noisy
saw dual service-a service-b --filter1 ERROR --filter2 ERROR
```

### Tip 2: Use Time Ranges to Skip Old Logs

```bash
# Start from 15 minutes ago
saw dual api worker --start1 -15m --start2 -15m
```

Avoids loading thousands of old messages.

### Tip 3: Keyboard Shortcuts Save Time

- `Tab` - Quick pane switching
- `G` - Jump to latest logs
- `g` - Jump to earliest logs
- `j`/`k` - Smooth scrolling

### Tip 4: Combine with Shell Aliases

```bash
# Add to ~/.bashrc or ~/.zshrc
alias sawprod='saw dual production-api production-worker'
alias sawerr='saw dual production-api production-worker --filter1 ERROR --filter2 ERROR'

# Then use:
sawprod
sawerr
```

### Tip 5: Use with Different AWS Profiles

```bash
saw dual account-a-logs account-b-logs --profile default
```

## Troubleshooting

### Logs Not Appearing?

**Check:**
1. Log group names are correct
2. AWS credentials are valid
3. Logs exist in the time range
4. Filters aren't too restrictive

**Test:**
```bash
# Verify log groups exist
saw groups | grep production-api

# Test without filters first
saw dual production-api production-worker
```

### Terminal Too Small?

**Minimum size:** 80x24
**Recommended:** 120x40 or larger

**Check your size:**
```bash
echo "Cols: $(tput cols) Lines: $(tput lines)"
```

### Pane Not Scrolling?

**Remember:**
- Only active pane scrolls
- Switch panes with `Tab`
- Need enough messages to scroll

## Exit and Navigation

### How to Exit

Three ways to quit:
1. Press `q`
2. Press `Ctrl+C`
3. Press `Esc`

### Navigation Cheat Sheet

```
┌─────────────────────────────────────┐
│  q, Ctrl+C, Esc  →  Quit            │
│  Tab             →  Switch pane     │
│  ↑, k            →  Scroll up       │
│  ↓, j            →  Scroll down     │
│  g               →  Jump to top     │
│  G               →  Jump to bottom  │
└─────────────────────────────────────┘
```

## Comparison: When to Use What

### Use `dual` when:
✅ Comparing two services
✅ Monitoring deployments
✅ Debugging distributed systems
✅ Correlating events
✅ Interactive exploration

### Use `watch` when:
✅ Single service monitoring
✅ Piping to other tools (grep, awk, etc.)
✅ Background monitoring
✅ Scripting/automation

### Use `get` when:
✅ Historical log analysis
✅ Exporting logs
✅ Time-range queries
✅ Batch processing

## Practice Exercise

Try this progression to learn dual view:

1. **Basic viewing:**
   ```bash
   saw dual production-api production-worker
   ```

2. **Add scrolling:**
   - Press `↓` several times
   - Press `G` to jump to bottom
   - Press `g` to jump to top

3. **Switch panes:**
   - Press `Tab` to switch
   - Scroll in second pane
   - Press `Tab` to return

4. **Add filtering:**
   ```bash
   saw dual production-api production-worker --filter1 ERROR --filter2 ERROR
   ```

5. **Practice quitting:**
   - Press `q` to quit
   - Restart and try `Ctrl+C`
   - Restart and try `Esc`

## Summary

The `saw dual` command is perfect for:
- 👀 **Visual comparison** of two log streams
- 🔄 **Real-time monitoring** during deployments
- 🐛 **Debugging** distributed systems
- 📊 **Correlation analysis** between services
- 🎯 **Interactive exploration** of logs

**Remember:** Press `q` to quit, `Tab` to switch panes, and arrow keys to scroll!

## See Also

- [Dual View Documentation](../docs/DUAL_VIEW.md) - Complete feature guide
- [Quick Reference](../docs/QUICK_REFERENCE.md) - All commands
- [Main README](../README.md) - Saw overview

---

Happy log monitoring! 🪵✨