# Tako Plugins Example

A demonstration of the strict Plugin architectural pattern in the Tako framework.

## Overview
This demo showcases:
- Defining a full `plugin.Manifest` struct with explicit metadata (`ID`, `Name`, `Version`).
- Registering a Plugin factory using `plugin.Register()`.
- Embedding `plugin.NoopLifecycle` to easily satisfy the `Lifecycle` interface.
- Using the explicit `OnInit(ctx)` hook to wire dependencies and mount components at boot time.

## Usage
Run the demo using Go:
```bash
go run .
```

## Keybindings
- **Ctrl+C**: Exits the application.
*(The main purpose of this demo is to inspect the `app/plugin.go` file architecture rather than complex UI interactions).*

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
