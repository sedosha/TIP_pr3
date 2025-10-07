param(
    [string]$Command = "run"
)

switch ($Command) {
    "run" {
        Write-Host "Запуск сервера..." -ForegroundColor Green
        go run ./cmd/server
    }
    "build" {
        Write-Host "Сборка проекта..." -ForegroundColor Yellow
        go build -o server.exe ./cmd/server
        Write-Host "Собран файл: server.exe" -ForegroundColor Green
    }
    "test" {
        Write-Host "Запуск тестов..." -ForegroundColor Cyan
        go test ./internal/api/...
    }
    default {
        Write-Host "Доступные команды:" -ForegroundColor White
        Write-Host "  .\run.ps1 run    - запуск сервера"
        Write-Host "  .\run.ps1 build  - сборка проекта" 
        Write-Host "  .\run.ps1 test   - запуск тестов"
    }
}