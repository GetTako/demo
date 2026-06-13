# 10-Theming Demo

A demonstration of dynamic run-time switching of Language (i18n) and Color Themes.

## Features Demonstrated
- **Language Manager:** `Lang().Register()` and `Lang().SetLocale()`.
- **Theme Manager:** `Theme().Register()` and `Theme().Use()`.
- **Dynamic Re-render:** The BubbleTea interface instantly reacts when the underlying context state changes.

## Use Case
Useful for applications that require multiple language support or user-selectable dark/light/custom accessibility themes without requiring a full application restart.

## Usage
1. Navigate to the directory: `cd demo/10-theming`
2. Run the application: `go run main.go`
3. Press `t` to toggle between dark/light themes.
4. Press `l` to toggle between EN/ID languages.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
