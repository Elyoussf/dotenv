# 🌍 universal-dotenv-reader (Go)

A blazing-fast `.env` parser written in **Go**, designed to be **language-agnostic**.

We're building **companion libraries for Python and Node.js** so developers can easily integrate `.env` file parsing into their favorite ecosystems.

---

## 🚀 Run It

```bash
go run . [state]
### 🔧 Parameters

* `state` → (optional) parsing mode:
  * `true` → **Multiline values supported**
  * `false` or omitted → **Multiline values NOT supported**


What to implement next ?

Reference Expansion

        Variable Expansion / Interpolation

        Example: URL=${HOST}:${PORT}

        Recursive Resolution Detection

        Prevents infinite loops if A=${B} and B=${A}.

Escape Characters Handling

    Handles escaped #, = inside quoted strings properly.