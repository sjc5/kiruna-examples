# Kiruna Examples

Examples repository for the [Kiruna](https://github.com/sjc5/kiruna) project.

## Fullstack Minimal Example

To run the fullstack minimal example, run the following commands:

```sh
cd fullstack-minimal
go mod tidy
make dev
```

Then open your browser to `http://localhost:8080`. You should see a styled html entry page.

- To test the instant CSS hot reloading, try editing any stylesheet in the `styles` folder.
- To test the browser refresh when you edit a Go template, try editing `static/private/templates/index.go.html`.
- To test the browser refresh when you rebuild your Go binary, edit any `.go` file in the `fullstack-minimal` directory (e.g., `kiruna.config.go` or `cmd/app/main.go`).

## Backend Minimal Example

To run the backend minimal example, run the following commands:

```sh
cd backend-minimal
go mod tidy
make dev
```

Then open your browser to `http://localhost:8080`. You'll see the "Hello, World!" response.
