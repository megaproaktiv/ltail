# Dual View Visual Guide

A visual tour of the `saw dual` command and its split-pane terminal interface with exact half-screen height panes.

## Interface Overview

```
┌────────────────────────────────────────────────────────────────────────┐
│                    Saw Dual View - Split Pane Interface                │
└────────────────────────────────────────────────────────────────────────┘

╔══════════════════════════════════════════════════════════════════════╗
║  production-api (145 messages)                                  ⬆    ║
╠──────────────────────────────────────────────────────────────────────╣
║ 14:23:01 [INFO] Request received: GET /api/users                     ║
║ 14:23:02 [INFO] Response sent: 200 OK                                ║
║ 14:23:03 [INFO] Request received: POST /api/orders                   ║
║ 14:23:04 [ERROR] Database connection timeout                         ║
║ 14:23:05 [INFO] Retrying connection...                               ║
║ 14:23:06 [INFO] Connection restored                                  ║
║ 14:23:07 [INFO] Request received: GET /api/products                  ║
║                                                                        ║
║                                                                        ║
╚══════════════════════════════════════════════════════════════════════╝
╔══════════════════════════════════════════════════════════════════════╗
║  production-worker (87) • q: quit • tab: switch • ↑↓/jk: scroll      ║
╠──────────────────────────────────────────────────────────────────────╣
║ 14:23:01 [INFO] Job started: process-order-123                       ║
║ 14:23:02 [INFO] Processing payment...                                ║
║ 14:23:03 [INFO] Payment confirmed                                    ║
║ 14:23:04 [INFO] Sending confirmation email                           ║
║ 14:23:05 [INFO] Job completed: process-order-123                     ║
║ 14:23:06 [INFO] Job started: process-order-124                       ║
║ 14:23:07 [ERROR] Failed to send email: SMTP timeout                  ║
║                                                                        ║
║                                                                        ║
╚══════════════════════════════════════════════════════════════════════╝

Note: Each pane takes exactly 50% of screen height. Help shown in bottom title.
```

## Pane States

### Active Pane (Bright Border)

```
╔══════════════════════════════════════╗  ← Highlighted border (Active)
║  production-api (145 messages)       ║
╠══════════════════════════════════════╣
║ 14:23:01 [INFO] Log message...       ║
║ 14:23:02 [ERROR] Error message...    ║
╚══════════════════════════════════════╝
      ↑
  Can scroll this pane
```

### Inactive Pane (Dim Border)

```
╔──────────────────────────────────────╗  ← Dim border (Inactive)
║  production-worker (87 messages)     ║
╠──────────────────────────────────────╣
║ 14:23:01 [INFO] Log message...       ║
║ 14:23:02 [INFO] Another message...   ║
╚──────────────────────────────────────╝
      ↑
  Cannot scroll until activated
```

## Keyboard Controls Diagram

```
┌─────────────────────────────────────────────────────────────┐
│                   KEYBOARD CONTROLS                          │
├─────────────────────────────────────────────────────────────┤
│                                                              │
│  QUIT APPLICATION                                            │
│  ┌───┐  ┌────────┐  ┌─────┐                                │
│  │ q │  │ Ctrl+C │  │ Esc │                                 │
│  └───┘  └────────┘  └─────┘                                 │
│                                                              │
│  SWITCH PANES                                                │
│  ┌─────┐                                                     │
│  │ Tab │  ←→  Toggle between top and bottom                 │
│  └─────┘                                                     │
│                                                              │
│  SCROLL (in active pane)                                     │
│  ┌───┐  ┌───┐                                               │
│  │ ↑ │  │ k │  ←  Scroll up                                │
│  └───┘  └───┘                                               │
│  ┌───┐  ┌───┐                                               │
│  │ ↓ │  │ j │  ←  Scroll down                              │
│  └───┘  └───┘                                               │
│                                                              │
│  JUMP                                                        │
│  ┌───┐                                                       │
│  │ g │  ←  Jump to top                                      │
│  └───┘                                                       │
│  ┌───┐                                                       │
│  │ G │  ←  Jump to bottom (shift + g)                       │
│  └───┘                                                       │
│                                                              │
└─────────────────────────────────────────────────────────────┘
```

## Color Highlighting

```
┌──────────────────────────────────────────────────────────┐
│                   LOG LEVEL COLORS                        │
├──────────────────────────────────────────────────────────┤
│                                                           │
│  14:23:01 ┌──────┐ Application started                   │
│           │ INFO │  ← Green background, black text       │
│           └──────┘                                        │
│                                                           │
│  14:23:02 ┌───────┐ Connection failed                    │
│           │ ERROR │  ← Red background, white text        │
│           └───────┘                                       │
│                                                           │
│  14:23:03  DEBUG  Log message  ← No color (not INFO/ERROR)│
│                                                           │
└──────────────────────────────────────────────────────────┘
```

## Split Layout Visualization

### Horizontal Split (Exact 50/50)

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃    TOP PANE (Exactly 50% height)     ┃
┃                                      ┃
┃         Log Group 1                  ┃
┃         Messages here                ┃
┃                                      ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫  ← No gap
┃  BOTTOM PANE (Exactly 50% height)    ┃
┃                                      ┃
┃         Log Group 2                  ┃
┃    (Help text in title bar)          ┃
┃                                      ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛

Each pane = terminal_height / 2
```

## Workflow Examples

### Example 1: Deployment Monitoring

```
BEFORE:                          AFTER DUAL VIEW:
                                 
Two separate terminals      →    One terminal, split view
                                 
Terminal 1:                      ╔═══════════════════════╗
$ saw watch old-version          ║ old-version (v1.2.3)  ║
[INFO] Processing...             ║ [INFO] Processing...  ║
[INFO] Request...                ║ [INFO] Request...     ║
                                 ╠═══════════════════════╣
Terminal 2:                      ║ new-version (v1.2.4)  ║
$ saw watch new-version          ║ [INFO] Starting...    ║
[INFO] Starting...               ║ [INFO] Ready!         ║
[INFO] Ready!                    ╚═══════════════════════╝
                                 
   Hard to compare          →    Easy side-by-side view!
```

### Example 2: Error Correlation

```
SCENARIO: API errors causing worker failures

┌────────────────────────────────────────────────────────────┐
│ Time    │  API Pane              │  Worker Pane            │
├────────────────────────────────────────────────────────────┤
│ 14:23:01│ [INFO] Request received │ [INFO] Job started     │
│ 14:23:02│ [ERROR] DB timeout      │ [INFO] Processing...   │
│ 14:23:03│ [ERROR] Request failed  │ [ERROR] Job failed!    │
│         │           ↓             │           ↑            │
│         └─────────  Correlation  ─────────────┘            │
└────────────────────────────────────────────────────────────┘

See the relationship instantly with dual view!
```

### Example 3: Tab Switching Flow

```
Step 1: Start with top pane active
╔══════════════════════╗  ← Active (bright border)
║  production-api      ║
╠══════════════════════╣
║ Can scroll here      ║
╚══════════════════════╝

╔──────────────────────╗  ← Inactive (dim border)
║  production-worker   ║
╠──────────────────────╣
║ Cannot scroll        ║
╚──────────────────────╝

        Press [Tab]
           ↓

Step 2: Bottom pane now active
╔──────────────────────╗  ← Inactive (dim border)
║  production-api      ║
╠──────────────────────╣
║ Cannot scroll        ║
╚──────────────────────╝

╔══════════════════════╗  ← Active (bright border)
║  production-worker   ║
╠══════════════════════╣
║ Can scroll here      ║
╚══════════════════════╝
```

## Scrolling Visualization

```
┌──────────────────────────────────────────────────────────┐
│  BEFORE SCROLLING               AFTER SCROLLING UP       │
├──────────────────────────────────────────────────────────┤
│                                                           │
│  ╔════════════════╗            ╔════════════════╗        │
│  ║ Messages (100) ║            ║ Messages (100) ║        │
│  ╠════════════════╣            ╠════════════════╣        │
│  ║ Msg 91         ║            ║ Msg 81         ║  ⬆     │
│  ║ Msg 92         ║            ║ Msg 82         ║  ⬆     │
│  ║ Msg 93         ║    [k]     ║ Msg 83         ║  ⬆     │
│  ║ Msg 94         ║     or     ║ Msg 84         ║        │
│  ║ Msg 95         ║    [↑]     ║ Msg 85         ║        │
│  ║ Msg 96  ←────  ║   ────→    ║ Msg 86  ←────  ║        │
│  ║ Msg 97  View   ║  Scroll    ║ Msg 87  View   ║        │
│  ║ Msg 98         ║    Up      ║ Msg 88         ║        │
│  ║ Msg 99         ║            ║ Msg 89         ║        │
│  ║ Msg 100        ║            ║ Msg 90         ║        │
│  ╚════════════════╝            ╚════════════════╝        │
│                                                           │
└──────────────────────────────────────────────────────────┘
```

## Message Counter and Help Display

```
╔══════════════════════════════════════════════╗
║  production-api (145 messages)          ⬆   ║  ← Top pane title
║                     ^^^                      ║
║              Message counter                 ║
╚══════════════════════════════════════════════╝
╔══════════════════════════════════════════════╗
║  production-worker (87) • q: quit • tab...   ║  ← Bottom pane title
║                            ^^^^^^^^^^^^^^^^  ║
║                            Help text here    ║
╚══════════════════════════════════════════════╝

Note: Help text displayed in bottom pane title to maintain exact 50/50 split
```

## Filter Visualization

### Without Filters

```
╔════════════════════════════════╗
║  All Messages                  ║
╠════════════════════════════════╣
║ [INFO] Message 1               ║
║ [DEBUG] Message 2              ║
║ [INFO] Message 3               ║
║ [ERROR] Message 4              ║
║ [WARN] Message 5               ║
║ [INFO] Message 6               ║
╚════════════════════════════════╝
```

### With --filter1 ERROR

```
╔════════════════════════════════╗
║  Filtered: ERROR only          ║
╠════════════════════════════════╣
║ [ERROR] Message 4              ║
║                                ║
║     (Only errors shown)        ║
║                                ║
║                                ║
║                                ║
╚════════════════════════════════╝
```

## Real-World Layout Examples

### Example: Microservices Debugging

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ api-gateway (request_id=abc-123)                    ⬆   ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 14:23:00 [INFO] Received POST /orders request_id=abc-123┃
┃ 14:23:01 [INFO] Routing to order-service              ┃
┃ 14:23:05 [INFO] Response received: 200 OK             ┃
┃ 14:23:06 [INFO] Returning to client                   ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ order-service (request_id=abc-123)                      ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ 14:23:01 [INFO] Processing order request_id=abc-123   ┃
┃ 14:23:02 [INFO] Validating order data                 ┃
┃ 14:23:03 [INFO] Saving to database                    ┃
┃ 14:23:04 [INFO] Order created: order-456              ┃
┃ 14:23:05 [INFO] Sending response                      ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
        ↑ Follow the request through the system ↑
```

### Example: A/B Testing

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ experiment-variant-a (conversion tracking)               ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ [INFO] User viewed product: conversion=false            ┃
┃ [INFO] User added to cart: conversion=false             ┃
┃ [INFO] User completed purchase: conversion=true   ✓     ┃
┃ [INFO] User viewed product: conversion=false            ┃
┃ Conversion Rate: 25%                                     ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ experiment-variant-b (conversion tracking)               ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ [INFO] User viewed product: conversion=false            ┃
┃ [INFO] User completed purchase: conversion=true   ✓     ┃
┃ [INFO] User completed purchase: conversion=true   ✓     ┃
┃ [INFO] User viewed product: conversion=false            ┃
┃ Conversion Rate: 50%                                     ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
        ↑ Compare variants side-by-side ↑
```

## Command Examples with Visual Output

### Basic Command

```bash
saw dual production-api production-worker
```

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ production-api (234 messages)     ┃  ← Pane 1
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ Real-time logs streaming...       ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ production-worker (156 messages)  ┃  ← Pane 2
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ Real-time logs streaming...       ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
```

### With Filters

```bash
saw dual api worker --filter1 ERROR --filter2 ERROR
```

```
┏━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┓
┃ api (12 messages) [ERROR only]    ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ [ERROR] Database timeout          ┃
┃ [ERROR] Request failed            ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ worker (8 messages) [ERROR only]  ┃
┣━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┫
┃ [ERROR] Job failed                ┃
┃ [ERROR] SMTP connection lost      ┃
┗━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┛
      Clean, focused error view
```

## Navigation Flow Diagram

```
                    Start Application
                           │
                           ▼
                  ┌─────────────────┐
                  │ Load Log Groups │
                  └─────────────────┘
                           │
                           ▼
              ┌────────────────────────┐
              │  Display Split View    │
              │  (Top pane active)     │
              └────────────────────────┘
                           │
          ┌────────────────┼────────────────┐
          │                │                │
          ▼                ▼                ▼
     ┌────────┐      ┌─────────┐      ┌────────┐
     │ Scroll │      │   Tab   │      │  Quit  │
     │  ↑/↓   │      │  Switch │      │   q    │
     └────────┘      └─────────┘      └────────┘
          │                │                │
          │                ▼                ▼
          │       ┌──────────────┐      Exit App
          │       │ Toggle Pane  │
          │       │  Active      │
          │       └──────────────┘
          │                │
          └────────────────┘
                  │
                  ▼
         Continue Monitoring
```

## Size Requirements Diagram

```
┌─────────────────────────────────────────────────────┐
│            TERMINAL SIZE GUIDE                      │
├─────────────────────────────────────────────────────┤
│                                                     │
│  MINIMUM (80x24)                                    │
│  ┌────────────────────┐                            │
│  │ Each pane: 12 lines│                            │
│  │ Cramped but works  │                            │
│  └────────────────────┘                            │
│                                                     │
│  RECOMMENDED (120x40)                              │
│  ┌──────────────────────────────────┐              │
│  │ Each pane: 20 lines              │              │
│  │ Comfortable viewing              │              │
│  │ Good message visibility          │              │
│  └──────────────────────────────────┘              │
│                                                     │
│  IDEAL (140x50+)                                   │
│  ┌─────────────────────────────────────────┐       │
│  │ Each pane: 25+ lines                    │       │
│  │ Excellent visibility                    │       │
│  │ No truncation for most messages         │       │
│  └─────────────────────────────────────────┘       │
│                                                     │
└─────────────────────────────────────────────────────┘
```

## Quick Reference Card

```
╔═══════════════════════════════════════════════════════════╗
║                 SAW DUAL VIEW QUICK REF                   ║
╠═══════════════════════════════════════════════════════════╣
║                                                           ║
║  START:  saw dual <group1> <group2>                       ║
║                                                           ║
║  LAYOUT: Exact 50/50 split (each pane = half screen)     ║
║  QUIT:   q / Ctrl+C / Esc                                 ║
║  SWITCH: Tab                                              ║
║  SCROLL: ↑↓ or jk                                         ║
║  JUMP:   g (top) / G (bottom)                             ║
║                                                           ║
║  FILTERS:                                                 ║
║    --filter1 ERROR   (top pane errors only)               ║
║    --filter2 INFO    (bottom pane info only)              ║
║                                                           ║
║  PREFIX:                                                  ║
║    --prefix1 api     (top pane streams matching "api")    ║
║    --prefix2 worker  (bottom pane streams with "worker")  ║
║                                                           ║
║  SHORTEN:                                                 ║
║    -s                (shorten both panes)                 ║
║    --shorten1        (shorten top only)                   ║
║                                                           ║
║  ACTIVE PANE:  Bright border                              ║
║  INACTIVE:     Dim border                                 ║
║                                                           ║
╚═══════════════════════════════════════════════════════════╝
```

## Tips & Tricks Visual

```
┌──────────────────────────────────────────────────────────┐
│                    PRO TIPS                               │
├──────────────────────────────────────────────────────────┤
│                                                           │
│  💡 TIP 1: Start Broad                                   │
│     saw dual api worker                                   │
│           ↓                                               │
│     (See everything first)                                │
│           ↓                                               │
│     saw dual api worker --filter1 ERROR --filter2 ERROR   │
│     (Then narrow down)                                    │
│                                                           │
│  💡 TIP 2: Use Aliases                                   │
│     alias sawprod='saw dual prod-api prod-worker'         │
│     sawprod   ← Quick access!                            │
│                                                           │
│  💡 TIP 3: Jump to Latest                                │
│     Press [G] to jump to bottom (newest logs)             │
│     Press [g] to jump to top (oldest logs)                │
│                                                           │
│  💡 TIP 4: Tab Like a Pro                                │
│     Tab → Scroll top → Tab → Scroll bottom → Repeat      │
│                                                           │
└──────────────────────────────────────────────────────────┘
```

## Summary Diagram

```
              ┌─────────────────────────────┐
              │    SAW DUAL VIEW            │
              │    Two Logs, One Screen     │
              └─────────────────────────────┘
                         │
        ┌────────────────┼────────────────┐
        │                │                │
        ▼                ▼                ▼
   ┌─────────┐     ┌──────────┐    ┌──────────┐
   │Real-time│     │Interactive│    │ Compare  │
   │Streaming│     │ Keyboard  │    │Side-by   │
   │         │     │ Controls  │    │  -Side   │
   └─────────┘     └──────────┘    └──────────┘
        │                │                │
        └────────────────┼────────────────┘
                         ▼
              ┌─────────────────────┐
              │  Better Monitoring  │
              │   Better Debugging  │
              │  Better Insights!   │
              │  Exact 50/50 Split! │
              └─────────────────────┘
```

---

**Happy log monitoring with dual view!** 🪵✨

Each pane takes exactly half your screen height - no gaps, perfect split!
Press `q` to quit, `Tab` to switch, and arrow keys to scroll!