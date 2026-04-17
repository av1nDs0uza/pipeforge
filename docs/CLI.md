# 📄 `docs/CLI.md`

## CLI

PipeForge provides a simple command-line interface for generating CI/CD pipelines.

---

## Commands

### init

Initialize a new CI/CD setup.

```bash
pipeforge init
```

**Flow:**
- Prompts user for configuration
- Generates pipeline file in project directory

### add ci

Add CI/CD configuration to an existing project.

```bash
documentation: "pipeforge add ci"
```
**Flow:**
- Prompts user for configuration
- Injects pipeline into correct location

**Example output:**
- `.github/workflows/ci.yml`

## Command Behavior
| Command | Purpose |
|---------|---------|
| init | Create new pipeline |
| add ci | Inject into existing project |

## Interactive Mode
By default, commands run interactively:
- Select CI provider 
- Select tier 
- Optional features 
 
**Planned Flags (Future):**
pipeforge init --ci github --tier standard --docker  
purpose:
enable non-interactive usage, support automation.
Output Locations:
| CI Provider | Output Path |
|--------------|--------------|
github | .github/workflows/ci.yml |
gitlab | .gitlab-ci.yml |
jk | Jenkinsfile |
Error Handling:
Common errors:
template not found,
failure to write file,
invalid configuration.
Design Goals:
simple commands,
miminal learning curve,
fase execution,
predictable output.