# 11-State Demo

A demonstration of reactive state bindings between decoupled components using the `StateManager`.

## Features Demonstrated
- **State Setting:** Modifying global application state.
- **State Observing:** `Observe().OnUpdate().Subscribe()` to automatically react to changes.
- **Decoupled Components:** The `Dashboard` handles input, while `DisplayComponent` simply listens to the global state.

## Use Case
Whenever you have deep UI trees and you need a component at the bottom of the tree to react to a change made by a component at the top of the tree, without passing variables manually down the hierarchy.

## Usage
1. Navigate to the directory: `cd demo/11-state`
2. Run the application: `go run main.go`
3. Press `up` arrow to increase the state score.
4. Press `down` arrow to decrease the state score.

## 🛠 Pro-Tip: Built-in Developer Tools

Tako comes with powerful tools out-of-the-box. You don't need to install anything extra to use them!

- **Hot Reloading:** Instead of `go run main.go`, run `go run main.go dev`. Whenever you save a `.go` file, the app will instantly rebuild and restart.
- **Live DevTools:** While the app is running in one terminal, open a second terminal and run `go run main.go inspect`. You'll get a real-time dashboard showing the focus stack, active keybindings, event streams, and performance metrics (like FPS and memory usage).
