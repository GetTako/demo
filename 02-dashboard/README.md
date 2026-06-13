# Tako Dashboard Example

A system dashboard demo built with the Tako framework.

## Overview
This demo showcases:
- Emitting events from background Goroutines using `time.Ticker`.
- Subscribing to events via the internal `EventBus` (`ctx.Events()`).
- Decoupling background data processing from the UI component layer.
- Reactive UI updates when system data changes.

## Usage
Run the demo using Go:
```bash
go run .
```

## Keybindings
- **Ctrl+C**: Exits the application.
*(The dashboard updates automatically every second without user input).*

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
