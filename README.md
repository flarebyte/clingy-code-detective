# clingy-code-detective

![Go Report
Card](https://goreportcard.com/badge/github.com/flarebyte/clingy-code-detective)
![Build
status](https://github.com/flarebyte/clingy-code-detective/actions/workflows/go.yml/badge.svg)
![License](https://img.shields.io/github/license/flarebyte/clingy-code-detective)
![Experimental](https://img.shields.io/badge/status-experimental-blue)

> Exposing the code that's a little too attached

A command-line tool to scan project directories for dependencies across
multiple ecosystems, aggregating and reporting them.

![Hero image for clingy-code-detective](clingy-code-detective-hero-512.jpeg)

Highlights:

-   Scans multiple directories recursively for dependencies across various
    ecosystems.
-   Supports Node.js (package.json) and Dart (pubspec.yaml) projects.
-   Extracts dependency names, versions, and categories (dev, prod).
-   Provides a detailed list of dependencies along with their file paths.
-   Aggregates dependency data to count occurrences and identify version
    ranges.
-   Exports results to JSON or CSV for auditing and documentation.
-   Designed for monorepo and mixed-environment analysis.

A few examples of commands:

Show help:

```bash
clingy --help
```

Show version:

```bash
clingy --version
```

Scan a single directory with default output (JSON implied by context):

```bash
clingy ./my-project
```

Include specific ecosystems (e.g., Node.js and Dart):

```bash
clingy --include=node,dart ./my-project
```

Exclude specific path segments:

```bash
clingy --exclude=/node_modules/,/build/ ./my-project
```

Output results in CSV format:

```bash
clingy --csv ./my-project
```

Output results in Markdown format:

```bash
clingy --md ./my-project
```

Aggregate results across multiple paths:

```bash
clingy --aggregate ./proj-a ./proj-b
```

## Documentation and links

-   [Code Maintenance :wrench:](MAINTENANCE.md)
-   [Code Of Conduct](CODE_OF_CONDUCT.md)
-   [Api for clingy-code-detective](API.md)
-   [Contributing :busts\_in\_silhouette: :construction:](CONTRIBUTING.md)
-   [Diagram for the code base :triangular\_ruler:](INTERNAL.md)
-   [Vocabulary used in the code base :book:](CODE_VOCABULARY.md)
-   [Architectural Decision Records :memo:](DECISIONS.md)
-   [Contributors :busts\_in\_silhouette:
    ](https://github.com/flarebyte/clingy-code-detective/graphs/contributors)
-   [Dependencies](https://github.com/flarebyte/clingy-code-detective/network/dependencies)
-   [Glossary :book:
    ](https://github.com/flarebyte/overview/blob/main/GLOSSARY.md)
-   [Software engineering principles :gem:
    ](https://github.com/flarebyte/overview/blob/main/PRINCIPLES.md)
-   [Overview of Flarebyte.com ecosystem :factory:
    ](https://github.com/flarebyte/overview)
-   [Go dependencies](DEPENDENCIES.md)
-   [Usage](USAGE.md)

## Related

-   [npm ls](https://docs.npmjs.com/cli/v10/commands/npm-ls)
-   [pub deps](https://dart.dev/tools/pub/cmd/pub-deps)
-   [depcheck](https://github.com/depcheck/depcheck)
-   [snyk](https://snyk.io/)
-   [trivy](https://github.com/aquasecurity/trivy)

## Installation

To use as a module in your Go project:

```go
import "github.com/flarebyte/clingy-code-detective"
```

Install the module:

```bash
go get github.com/flarebyte/clingy-code-detective@latest
```
