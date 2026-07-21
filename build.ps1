# Build script for Golang + Svelte PWA Monolith
Write-Host "Building Svelte Frontend with PWA support..." -ForegroundColor Cyan
Set-Location -Path "$PSScriptRoot\frontend"
npm run build

if ($LASTEXITCODE -ne 0) {
    Write-Host "Frontend build failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}

Set-Location -Path "$PSScriptRoot"
Write-Host "Building Golang Monolith Backend..." -ForegroundColor Cyan
go build -o server.exe .

if ($LASTEXITCODE -ne 0) {
    Write-Host "Go build failed!" -ForegroundColor Red
    exit $LASTEXITCODE
}

Write-Host "Build complete! Running server.exe..." -ForegroundColor Green
.\server.exe
