# Gokite

```bash
❯ kitex -gen-path kitgen idl/echo.proto

❯ kitex -gen-path kitgen -module gokite -service echo idl/echo.proto

# Generate With Module Embedded
❯ kitex -I proto -gen-path kitgen -module github.com/kelein/cookbook -service echo proto/echo.proto
```
