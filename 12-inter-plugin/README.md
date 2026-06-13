# 12-Inter-Plugin Demo

A demonstration of multi-plugin architecture and how independent plugins communicate across the internal boundaries.

## Features Demonstrated
- **Multi-Provider Booting:** Registering `BackendProvider` and `FrontendProvider` independently.
- **Inter-Plugin Events:** `Frontend` publishes `frontend:ping`, `Backend` listens and publishes `backend:pong`.
- **Inter-Plugin RPC:** `Frontend` synchronously calls `weather:get` exposed exclusively by `Backend`.

## Use Case
Essential for building a marketplace or modular architecture where third-party developers can write a plugin that hooks into your core system's data streams safely without touching your core codebase.

## Usage
1. Navigate to the directory: `cd demo/12-inter-plugin`
2. Run the application: `go run main.go`
3. Press `p` to ping the backend asynchronously.
4. Press `w` to request weather data synchronously via RPC.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
