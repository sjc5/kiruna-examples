# Kiruna Examples

Examples repository for the [Kiruna](https://github.com/sjc5/kiruna) project.

## Fullstack Minimal

To run the backend minimal example, run the following commands:

```sh
cd fullstack-minimal
go mod tidy
make dev
```

Then open your browser to localhost:8080. You'll see the web page.

### CSS Hot Reload

To try the instant CSS reload, edit any stylesheet in the `styles` folder.

### Template Browser Refresh

To try the Go template reloading, edit `static/private/templates/index.go.html`.

### Go Server Browser Refresh

To try the full Go server reloading, edit any `.go` file (e.g., `kiruna.config.go` or `cmd/app/main.go`).

## Backend Minimal

To run the backend minimal example, run the following commands:

```sh
cd backend-minimal
go mod tidy
make dev
```

Then open your browser to localhost:8080. You'll see the "Hello world" response.
