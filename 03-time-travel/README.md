# Tako Banking Simulation & Time Travel Example

A banking simulation application built with the Tako framework to demonstrate the powerful **Interactive Time-Travel Debugger** feature.

## Overview
This demo showcases:
- Using `ctx.State()` to store and broadcast reactive state variables.
- Using Tako's native Event Sourcing Tracer to record state changes.
- Using `tako replay` to replay a recorded session and inspect the state tree leading up to a system crash.

## Usage

### 1. Run and Record the Trace
To enable tracing, set the `TAKO_RECORD_TRACE` environment variable to `1`. This will automatically record all events and state changes to a `.tako/traces` directory.

```bash
TAKO_RECORD_TRACE=1 go run .
```

### 2. Simulate the Banking Crash
Once the application is running, interact with the system:
1. Press **`d`** multiple times to deposit funds.
2. Press **`w`** multiple times to withdraw funds.
3. Press **`c`** to trigger a **System Crash** (Go panic).

The application will abruptly exit, simulating a critical runtime failure where the application state might have been corrupted.

### 3. Replay with the Time-Travel Debugger
Now, use the built-in `replay` CLI command to analyze the recorded trace file step-by-step to understand the exact sequence of events that led to the crash.

```bash
go run . replay
```
*(If no file is provided, `replay` automatically selects the most recent trace file).*

Inside the Time-Travel Debugger:
- Press **Space** to play/pause the replay.
- Press **Right Arrow** / **Left Arrow** to step through each event one by one.
- Watch the **State Tree** on the right side highlight the changes (diffs) to the banking balance.

## Keybindings (App)
- **`d`**: Deposit $100.
- **`w`**: Withdraw $50.
- **`c`**: Simulate System Crash (Panic).
- **`Ctrl+C`**: Gracefully quit the application.

## Keybindings (Replay Debugger)
- **`Space`**: Play / Pause.
- **`Left / Right`**: Step backward / forward manually.
- **`f`**: Filter events.
- **`g`**: Jump to a specific event line.
- **`d`**: Toggle State Tree diffs.
- **`1, 2, 3, 4`**: Change playback speed.
- **`q` / `Ctrl+C`**: Exit debugger.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
