---
modified: "Sat Apr 25 11:25:14 EDT 2026"
---

# ollama

## Run a model with custom context

```bash
# Run ollama normally
ollama serve

# Check model context
ollama show gemma4:latest
# >> context length 131072

# Restart ollama with proper context
OLLAMA_CONTEXT_LENGTH=131072 ollama serve
```

To check if your model is running with corrext context:

```bash
ollama ps
# >> CONTEXT 131072
```
