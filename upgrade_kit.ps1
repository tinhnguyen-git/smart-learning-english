Write-Host "🚀 Bắt đầu nâng cấp Antigravity Kit lên cấp độ 10/10..." -ForegroundColor Cyan

# ==========================================
# 1. BỘ NHỚ DÀI HẠN (PROJECT BRAIN)
# ==========================================
Write-Host "🧠 Đang khởi tạo Project Brain..." -ForegroundColor Yellow

$projectBrain = @'
{
  "meta_data": {
    "project_name": "English Learning SuperApp",
    "version": "1.0.0",
    "last_updated": "2026-01-21T15:00:00Z",
    "stack": {
      "mobile": "Flutter (Clean Arch + BLoC)",
      "web": "Angular + TailwindCSS",
      "backend": "Go (Microservices + DDD)",
      "database": "PostgreSQL"
    }
  },
  "context_state": {
    "current_phase": "Initialization",
    "active_sprint_goals": [],
    "blocking_issues": []
  },
  "architectural_decision_records": []
}
'@
$projectBrain | Out-File -FilePath "project_brain.json" -Encoding utf8 -Force

# Tạo Rule bắt buộc Agent đọc não bộ
if (!(Test-Path ".agent/rules")) { New-Item -ItemType Directory -Path ".agent/rules" -Force }

$memoryProtocol = @'
# MEMORY PROTOCOL (FORCE READ)

## CRITICAL INSTRUCTION
Before writing any code or suggesting architecture, you MUST:
1. Read `project_brain.json` to understand the current phase and constraints.
2. Read `docs/WORK_LOG.md` to recall recent changes.

## SESSION END
When a task is done, you MUST using the `memory_ops` skill to update `project_brain.json` with new architectural decisions or phase changes.
'@
$memoryProtocol | Out-File -FilePath ".agent/rules/99-memory-persistence.md" -Encoding utf8 -Force

# ==========================================
# 2. SANDBOX AN TOÀN (DEVCONTAINER)
# ==========================================
Write-Host "🛡️ Đang thiết lập môi trường DevContainer an toàn..." -ForegroundColor Yellow
if (!(Test-Path ".devcontainer")) { New-Item -ItemType Directory -Path ".devcontainer" -Force }

# Sử dụng Here-String với nháy đơn để tránh lỗi Parser Variable
$setupAndroid = @'
#!/bin/bash
set -e
if [ ! -d "/usr/lib/android-sdk" ]; then
    echo "📲 Installing Android SDK Command Line Tools..."
    sudo apt-get update && sudo apt-get install -y android-sdk openjdk-17-jdk
    echo 'export ANDROID_HOME=/usr/lib/android-sdk' >> ~/.zshrc
    echo 'export PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools' >> ~/.zshrc
fi
echo "✅ Android Environment Ready."
'@
# Chuyển đổi sang LF để chạy trong Linux
[System.IO.File]::WriteAllText("$(Get-Location)/.devcontainer/setup_android.sh", $setupAndroid.Replace("`r`n", "`n"), [System.Text.Encoding]::UTF8)

# Tạo devcontainer.json
$devContainerJson = @'
{
  "name": "Antigravity Super-Environment (Go/Flutter/Angular)",
  "image": "mcr.microsoft.com/devcontainers/base:ubuntu-22.04",
  "features": {
    "ghcr.io/devcontainers/features/docker-outside-of-docker:1": {},
    "ghcr.io/devcontainers/features/go:1": { "version": "1.21" },
    "ghcr.io/devcontainers/features/node:1": { "version": "18" },
    "ghcr.io/devcontainers/features/python:1": { "version": "3.11" },
    "ghcr.io/devcontainers/features/java:1": { "version": "17" }
  },
  "customizations": {
    "vscode": {
      "extensions": [
        "golang.go",
        "Dart-Code.flutter",
        "angular.ng-template",
        "ms-azuretools.vscode-docker"
      ]
    }
  },
  "runArgs": [
    "--privileged",
    "-v", "/dev/bus/usb:/dev/bus/usb"
  ],
  "postCreateCommand": "bash .devcontainer/setup_android.sh",
  "remoteUser": "vscode"
}
'@
$devContainerJson | Out-File -FilePath ".devcontainer/devcontainer.json" -Encoding utf8 -Force

# ==========================================
# 3. KỸ NĂNG THỰC THI (SKILLS)
# ==========================================
Write-Host "🛠️ Đang cài đặt Skills nâng cao..." -ForegroundColor Yellow
$skillsDirs = @(".agent/skills/infra_ops", ".agent/skills/memory_ops", ".agent/skills/mobile_bridge")
foreach ($dir in $skillsDirs) {
    if (!(Test-Path $dir)) { New-Item -ItemType Directory -Path $dir -Force }
}

$updateBrainPy = @'
import json
import sys

def update_brain(key, value):
    try:
        with open('project_brain.json', 'r+') as f:
            data = json.load(f)
            keys = key.split('.')
            ref = data
            for k in keys[:-1]:
                ref = ref[k]
            ref[keys[-1]] = value
            f.seek(0)
            json.dump(data, f, indent=2)
            f.truncate()
    except Exception as e:
        print(f"Error updating brain: {e}")

if __name__ == "__main__":
    if len(sys.argv) > 2:
        update_brain(sys.argv[1], sys.argv[2])
'@
$updateBrainPy | Out-File -FilePath ".agent/skills/memory_ops/update_brain.py" -Encoding utf8 -Force

# ==========================================
# 4. HẠ TẦNG (INFRASTRUCTURE)
# ==========================================
Write-Host "🏗️ Đang dựng khung Infrastructure..." -ForegroundColor Yellow
if (!(Test-Path "infra/docker")) { New-Item -ItemType Directory -Path "infra/docker" -Force }

$dockerCompose = @'
version: '3.8'
services:
  postgres:
    image: postgres:15-alpine
    environment:
      POSTGRES_USER: user
      POSTGRES_PASSWORD: password
      POSTGRES_DB: english_app
    ports:
      - "5432:5432"
    volumes:
      - pgdata:/var/lib/postgresql/data
  redis:
    image: redis:alpine
    ports:
      - "6379:6379"
volumes:
  pgdata:
'@
$dockerCompose | Out-File -FilePath "infra/docker/docker-compose.yml" -Encoding utf8 -Force

Write-Host "✅ ĐÃ NÂNG CẤP THÀNH CÔNG!" -ForegroundColor Green
Write-Host "👉 Bước tiếp theo: Trong VS Code, nhấn 'Reopen in Container'." -ForegroundColor White