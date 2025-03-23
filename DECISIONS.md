# Architecture decision records

An [architecture
decision](https://cloud.google.com/architecture/architecture-decision-records)
is a software design choice that evaluates:

- a functional requirement (features).
- a non-functional requirement (technologies, methodologies, libraries).

The purpose is to understand the reasons behind the current architecture, so
they can be carried-on or re-visited in the future.

**Problem Specification: CLI Tool to Scan Multiple Projects for Dependencies**

**Objective**  
Create a command-line tool in Go that scans multiple project directories (recursively) to identify supported dependency files (e.g., `package.json` for Node.js, `pubspec.yaml` for Dart), extracts all declared dependencies with their versions and categories (e.g., dev, prod), and returns a detailed list along with the file path where each dependency was found. The tool must support mixed environments and ecosystems in a single scan.

**Use Cases**

- Scanning multiple directories (e.g., `./apps`, `./libs`) that include both Dart and Node.js projects to produce a unified list of dependencies.
- Running the tool across a monorepo that contains heterogeneous project types and aggregating all dependencies into a single report.
- Identifying dependency reuse across projects regardless of version differences.
- Exporting all dependency data to a JSON or CSV file for auditing or documentation.
- Using the aggregation option to count how often a dependency appears and determine the range of versions used.

**Edge Cases**

- Projects that contain both `package.json` and `pubspec.yaml` in the same directory.
- Nested project folders with their own dependency files.
- Mixed ecosystem files that declare the same dependency name with different versioning schemes.
- Invalid or empty dependency files.
- Missing or malformed version constraints.
- Dependency categories that differ slightly by ecosystem (e.g., `devDependencies` in Node.js vs `dev_dependencies` in Dart).

**Limits and Exclusions**

- The CLI should not perform installation, validation, or resolution of dependencies.
- It should not attempt to interpret or align versioning schemes across ecosystems.
- It should not support additional ecosystems (e.g., Go) unless explicitly added later.
- It should not require any configuration file; functionality is entirely controlled via CLI flags and arguments.

**Output Options**

- Display results in the console in a readable tabular format.
- Export results in JSON or optionally CSV, with structured entries: dependency name, version, type (dev, prod, etc.), file path.
- Support an `--aggregate` flag to summarize dependency occurrences across projects and show min/max version where applicable.

**CLI Syntax Requirements**

- Follows conventions similar to tools like `scc`, `fd`, or `ripgrep`.
- Accepts one or more directory paths to scan.
- Supports `--json`, `--csv`, and `--aggregate` flags.
- Supports filtering by ecosystem with `--include=node,dart`.
