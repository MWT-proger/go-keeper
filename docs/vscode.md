# Настройки для проекта в VS Code

## launch.json
```
{
    "version": "0.2.0",
    "configurations": [
      {
        "name": "Launch Server Go",
        "type": "go",
        "request": "launch",
        "mode": "debug",
        "program": "${workspaceFolder}/cmd/server",
        "env": {
          "RUN_ADDRESS": ":8000",
          "DATABASE_URI": "user=postgres password=postgres host=localhost port=5432 dbname=testDB sslmode=disable"
        },
        "args": ["-a=${env.RUN_ADDRESS}", "-d=${env.DATABASE_URI}", "-l=debug"]
      }
    ]
  }
```
