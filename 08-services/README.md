# 08-Services Demo

A demonstration of Tako's built-in Scheduler for managing background jobs safely without blocking the TUI thread.

## Features Demonstrated
- **Periodic Jobs (`Every`):** Runs an auto-incrementing background task every 1 second.
- **One-off Jobs (`Dispatch`):** Triggers a delayed, potentially failing task.
- **Error Handling (`OnError`):** Safely catches errors from background jobs and routes them back to the main UI.

## Use Case
Use the scheduler when you need to poll a REST API every 10 seconds, or perform a heavy filesystem scan (e.g., searching for files) that shouldn't freeze the user interface while it processes.

## Usage
1. Navigate to the directory: `cd demo/08-services`
2. Run the application: `go run main.go`
3. Watch the counter auto-increment.
4. Press `e` to dispatch a simulated failing background task.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
