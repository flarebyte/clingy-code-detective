# Technical Design

> Guide for the implementation, including detailed design, priorities,
> coding conventions, and testing

Highlights:

## Code structure

-   **cmd**: Entry points to the application (main packages)
-   **internal**: Private application and library code
-   **pkg**: Public libraries intended for use by external applications
-   **api**: Protobuf definitions or OpenAPI specs if applicable
-   **config**: Configuration files and helpers
-   **test**: Go unit and integration tests (uses Go's built-in `testing`
    package)
-   **scripts**: Tooling scripts (bash, makefiles, etc.)
-   **build**: Packaging and distribution assets
-   **.github**: GitHub Actions workflows
-   **.vscode**: Visual Studio Code settings and tasks

## Useful links

-   [Effective Go](https://go.dev/doc/effective_go)
-   [Go Code Review Comments](https://go.dev/wiki/CodeReviewComments)
