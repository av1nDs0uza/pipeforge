# PipeForge ⚡

A Vite-like CLI for DevOps — instantly generate production-ready CI/CD pipelines.

---

## 🚀 What is PipeForge?

PipeForge helps you:

- Generate CI/CD pipelines in seconds
- Support GitHub Actions, GitLab CI, Jenkins
- Use predefined, production-ready templates
- Plug-and-play into new or existing projects

---

## 🧱 Features

- 🔹 CI/CD template generation
- 🔹 Multi-provider support (GitHub, GitLab, Jenkins)
- 🔹 Tier-based setups (basic, standard, advanced)
- 🔹 Works for new + existing projects
- 🔹 No external dependencies (single binary)

---

## ⚙️ Installation

```bash
git clone https://github.com/yourusername/pipeforge
cd pipeforge
go build -o pipeforge

# 🚀 Usage

## Initialize CI/CD
```bash
./pipeforge init
```

## Add CI/CD to existing project
```bash
./pipeforge add ci
```

# 📦 Example Output
- `.github/workflows/ci.yml`

# 🧪 Running Tests
```bash
go test ./...
```

# 🧠 How it works
- Templates are embedded into the binary.
- CLI generates pipelines based on user input.
- Files are injected into your project.

# 📌 Roadmap
- GitHub Actions templates.
- GitLab CI templates.
- Jenkins pipelines.
- Docker + Kubernetes support.
- Interactive CLI (Vite-style).