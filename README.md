# Tako Framework Demos

A comprehensive collection of demo applications showcasing the robust features, architecture, and capabilities of the **Tako** terminal UI framework. 

Each demo is a standalone Go module designed to highlight a specific architectural concept or feature. Feel free to explore the source code, run them locally, and use them as reference implementations for your own Tako applications.

---

## 🚀 Getting Started

To run any of the demos, simply navigate into its directory and use `go run main.go`:

```bash
cd 01-todo
go run main.go
```

> **Pro-Tip:** Append `dev` or `inspect` to your run command to utilize Tako's built-in developer tools!
> - **Hot Reloading:** `go run main.go dev` (Auto-restarts when you save `.go` files)
> - **Live DevTools:** `go run main.go inspect` (Opens a real-time diagnostics dashboard in a second terminal window)

---

## 📂 Demos Catalog

### 🏗 Foundational Demos
- **[`01-todo/`](./01-todo/)**: A classic Todo list application. Demonstrates layout structure, text inputs, list rendering, keyboard event binding, and basic state management.
- **[`02-dashboard/`](./02-dashboard/)**: A monitoring dashboard. Showcases complex Grid/Flex layouts, asynchronous data updates, and utilizing the EventBus for decoupled communication.

### 🎯 Feature-Specific Demos
- **[`03-time-travel/`](./03-time-travel/)**: Demonstrates advanced *State Management* with history recording, enabling visual undo/redo (time travel debugging).
- **[`04-mouse/`](./04-mouse/)**: Highlights precise *Mouse Routing*, click area detection, hover states, and drag-and-drop mechanics inside the terminal.
- **[`05-overlays/`](./05-overlays/)**: Covers *Z-Index* manipulation for rendering pop-ups, modal dialogs, context menus, and tooltips seamlessly over base components.
- **[`06-plugins/`](./06-plugins/)**: Showcases the *Plugin Lifecycle Ecosystem*, demonstrating how external plugins can modularly inject functionality into a host application.

### 🧠 Advanced Architecture Demos
- **[`07-communication/`](./07-communication/)**: A deep dive into internal plugin communication using Asynchronous *EventBus* (Pub/Sub), Synchronous *RPC* (Request/Response), and *Hooks* (UI Injection).
- **[`08-services/`](./08-services/)**: Demonstrates the built-in *Scheduler* for managing background jobs safely (`Every`, `Dispatch`) without blocking or freezing the TUI thread, complete with error handling.
- **[`09-cli/`](./09-cli/)**: Showcases Tako's *Dual-Kernel Architecture*. Proves that Tako applications can bypass the graphical TUI entirely to act as standard CLI tools with flag parsing.
- **[`10-theming/`](./10-theming/)**: Demonstrates dynamic run-time switching. Change *Color Themes* (e.g. Dark/Light) and *Languages* (i18n) on the fly without restarting the app.
- **[`11-state/`](./11-state/)**: Showcases extreme *Reactivity*. Completely decoupled components bind to global state via `Observe().OnUpdate()`, instantly reacting to changes made elsewhere in the app.
- **[`12-inter-plugin/`](./12-inter-plugin/)**: Demonstrates *Multi-Plugin Architecture*. Two completely independent plugins (Frontend UI and Backend Data) boot together and communicate securely across internal boundaries.
- **[`13-native-bubbles/`](./13-native-bubbles/)**: Demonstrates *Bubble Tea Interoperability*. Shows how to use the `NativeComponent` interface to receive raw `tea.Msg` events, enabling seamless integration with the existing Charm ecosystem.

---

## 🤝 Contributing

If you've built something cool with Tako and want to add it as a demo, feel free to submit a Pull Request! Please ensure your demo follows the established `app/provider.go` structure and includes a detailed `README.md`.
