# Tako Overlays & Dialogs Example

A demonstration of Tako's Z-Index stacking and overlay management capabilities.

## Overview
This demo showcases:
- How to access the framework's `OverlayManager`.
- Rendering built-in primitive modals like `Confirm()` dialogs.
- Pushing and closing custom component overlays (`ShowComponent`).
- The automatic routing of keyboard focus to the topmost layer in the stack.

## Usage
Run the demo using Go:
```bash
go run .
```

## Keybindings
- **`d`**: Opens the native framework `Confirm` dialog popup. Inside the dialog, use `y/n` or `enter` to choose an option.
- **`p`**: Opens a custom Modal component over the base layout.
- **`esc`**: Closes the topmost active modal/overlay and returns focus to the underlying layer.
- **Ctrl+C**: Exits the application.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
