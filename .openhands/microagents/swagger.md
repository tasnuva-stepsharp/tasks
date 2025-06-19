---
trigger: ["swagger", "api docs", "openapi"]
---

# Swagger Guidelines

- Use `swaggo/swag` for Go.
- Annotate all handlers with Swagger comments.
- Update `docs/swagger.json` after every route change using:
  ```bash
  swag init
