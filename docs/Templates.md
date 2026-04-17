
## Templates

PipeForge uses predefined CI/CD templates to generate pipelines.

---

### Template Structure

Each template follows this naming convention:

```plaintext
<provider>-<tier>.yml
```

### Examples:
- `github-basic.yml`
- `github-standard.yml`
- `github-advanced.yml`
- `gitlab-basic.yml`
- `jenkins-basic.groovy`

---

### Template Location

Templates are stored in:

```
internal/templates/
```

They are embedded into the binary using Go's embed feature.

---

### Template Selection Logic

Templates are selected based on user config:

```go
def fileName := fmt.Sprintf("%s-%s.yml", cfg.CI, cfg.Tier)
def Tiers:
def Basic: Build application, Minimal setup
def Standard: Build + test, Improved workflow structure
def Advanced: Build + test + Docker, Deployment steps (optional), Production-ready features```
```
**Template Responsibilities**
- Be production-ready  
- Follow best practices  
- Be minimal but extensible  
- Avoid hardcoded secrets  
- Use environment variables where needed  

**Customization**
Templates are designed to be:
- Editable after generation  
- Extendable by users  
- Compatible with common workflows  

**Adding New Templates**
to add a new template:
a. Create file: `internal/templates/<provider>-<tier>.yml`  
b. Follow naming convention  
c. Ensure compatibility with generator logic.

**Future Enhancements:**
d. Parameterized templates  
e. Conditional logic (Docker, language)  
f. Template validation  
g. User-defined templates