# Tako Todo App Example

A basic Todo list application built with the Tako framework.

## Overview
This demo showcases:
- Component implementation (`contracts.Component`) using raw `lipgloss/v2` for rendering.
- Text input handling via the `KeyManager`.
- Rendering dynamic lists.
- Managing simple component state.

## Usage
Run the demo using Go:
```bash
go run .
```

## Keybindings
- **Any Character**: Types into the input box.
- **Space**: Adds a space to the input box.
- **Backspace**: Removes the last character from the input box.
- **Enter**: Adds the typed item to the todo list and clears the input.
- **Ctrl+C**: Exits the application.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
