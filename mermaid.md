---
modified: Sat Jul 4 14:18:11 2026 -0400
---

# Mermaid.js Diagrams

- [Mermaid Intro](https://mermaid.ai/open-source/intro/)

## State Chart

```mermaid
---
title: Simple sample
---
stateDiagram-v2
    [*] --> Still
    Still --> [*]

    Still --> Moving
    Moving --> Still
    Moving --> Crash
    Crash --> [*]
```
