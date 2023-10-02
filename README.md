# Skeleton for a GO CRUD app with MySQL

This is a project that I use as example when I have to make a CRUD API for a Code Challenge, a PoC, scafolding, etc.

## API usage

```bash
# Get dev
curl http://localhost:8000/get\?email\=dev@test.com
# Create dev
curl -d '{"email":"dev@test.com", "expertise":2}' -H "Content-Type: application/json" -X POST http://localhost:8000/create
# Update dev
curl -d '{"email":"dev@test.com", "expertise":2, "languages": ["go","ruby"]}' -H "Content-Type: application/json" -X PATCH http://localhost:8000/update
# Delete dev
curl -X DELETE http://localhost:8000/delete\?email\=dev@test.com
```
