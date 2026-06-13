# 09-CLI Demo

A demonstration of Tako's Dual-Kernel architecture, proving that Tako applications do not need a graphical UI. 

## Features Demonstrated
- **CLI Commands:** Bypasses BubbleTea entirely.
- **Flag Parsing:** Automatic reading of CLI flags via `ctx.OptionString()`.
- **Badge Printing:** Using styled terminal stdout commands (`ctx.Info()`).

## Use Case
If your application provides a TUI, you might also want to provide pure CLI commands for CI/CD automation, database migrations, or cronjobs (e.g., `myapp db:migrate`). This allows both the TUI and CLI to share the exact same underlying service container.

## Usage
1. Navigate to the directory: `cd demo/09-cli`
2. Run the application with args: `go run main.go hello --name="Supian"`

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
