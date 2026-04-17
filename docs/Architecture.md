# Architecture

PipeForge is a modular CLI tool designed to generate CI/CD pipelines using predefined templates.

---

## Overview

The system follows a simple flow:

`CLI → Config → Generator → Templates → Output`

---

## Core Components

### 1. CLI (`cmd/`)
Handles user interaction and command execution using Cobra.

**Responsibilities:**
- Parse commands (`init`, `add ci`)
- Collect user input (via prompts)
- Pass configuration to generator

---

### 2. Config (`internal/config/`)
Defines the structure of user choices.

**Example:**
```go
type Config struct {
    CI       string // github, gitlab, jenkins
    Tier     string // basic, standard, advanced
    Docker   bool
    Language string
}
```
**Responsibilities:**
- Store user-selected options
- Act as a contract between CLI and generator

---

### 3. Prompts (`internal/prompts/`)
Handles interactive input from the user.

**Responsibilities:**
- Ask for CI provider
- Ask for tier selection
- Ask for optional features (Docker, etc.)

---

### 4. Generator (`internal/generator/`)
Core engine responsible for file generation.

**Responsibilities:**
- Select correct template<br>- Read template from embedded filesystem<br>- Write output files to target directory</br>
---
# 5. Templates (`internal/templates/`)

Contains embedded CI/CD templates.
Uses Go embed:

```go
//go:embed *.yml 
avar Files embed.FS
```

## Responsibilities
- Display reusable pipeline definitions
- Provide template data to generator
- Detect project type (Node, Python, etc.)
- Assist in auto-config selection

## Data flow
1. User runs command: `pipeforge init`
2. CLI collects input → creates Config
3. Generator: selects template (`<ci>-<tier>.yml`), reads from embedded templates, writes to project directory.
4. Output is created at `.github/workflows/ci.yml`

## Design Principles
- **Separation of concerns:** Each module has a single responsibility.
- **No global state:** All inputs passed explicitly via config.
- **Deterministic output:** Same input always produces same result.
- **Testability:** File output is controlled via injected directory.
- **Portability:** Templates embedded → no runtime dependencies.

## Future Extensions
- Plugin system for custom templates.
- Multi-language support.
- Advanced pipeline configurations.
- CLI flags for non-interactive mode.