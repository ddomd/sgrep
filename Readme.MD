# ⚡️ SGREP
<br>
A simple multithreaded grep implementation using Go.

### Usage:

```bash
sgrep [-vrwb] <query> [path]
```

### Options:
- '-v' shows line numbers and file path
- '-r' use query as a regex
- '-w' sets the amount of search workers(default: 10)
- '-b' sets the buffer of the work channel(default: 100)