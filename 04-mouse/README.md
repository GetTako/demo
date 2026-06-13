# Tako Mouse Example

A demonstration of mouse interactivity built with the Tako framework.

## Overview
This demo showcases:
- Implementing `contracts.MouseComponent` to handle mouse inputs.
- Registering a coordinate hitbox using `mouseMgr.UpdateHitbox(...)` dynamically on every render cycle.
- Attaching `OnClick` event listeners to specific UI zones.

## Usage
Run the demo using Go:
```bash
go run .
```

## Interactions
- **Mouse Click**: Click anywhere inside the rendered boundary box to trigger the `OnClick` listener. The coordinates of your click will be displayed.
- **Ctrl+C**: Exits the application.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
