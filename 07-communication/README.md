# 07-Communication Demo

A demonstration of Tako's single-plugin internal communication capabilities using EventBus, RPC, and Hooks.

## Features Demonstrated
- **EventBus (Pub/Sub):** Asynchronous messaging (Chat widget).
- **RPC (Sync):** Synchronous blocking request/response (Calculator widget).
- **Hooks:** Extensible UI injection without modifying component templates.

## Use Case
When you build a plugin that needs to separate its internal state processing from its UI, you can use the EventBus. When you need to provide a UI slot where other parts of your plugin can inject elements, you use Hooks.

## Usage
1. Navigate to the directory: `cd demo/07-communication`
2. Run the application: `go run main.go`
3. Press `enter` after typing to send a pub/sub chat message.
4. Press `c` to trigger a blocking RPC request.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
