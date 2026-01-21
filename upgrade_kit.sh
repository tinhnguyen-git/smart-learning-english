#!/bin/bash

echo "🚀 Bắt đầu nâng cấp Antigravity Kit lên cấp độ 10/10..."

# ==========================================
# 1. BỘ NHỚ DÀI HẠN (PROJECT BRAIN)
# ==========================================
echo "🧠 Đang khởi tạo Project Brain..."

cat > project_brain.json <<EOF
{
  "meta_data": {
    "project_name": "English Learning SuperApp",
    "version": "1.0.0",
    "last_updated": "$(date -u +"%Y-%m-%dT%H:%M:%SZ")",
    "stack": {
      "mobile": "Flutter (Clean Arch + BLoC)",
      "web": "Angular + TailwindCSS",
      "backend": "Go (Microservices + DDD)",
      "database": "PostgreSQL"
    }
  },
  "context_state": {
    "current_phase": "Initialization",
    "active_sprint_goals":,
    "blocking_issues":
  },
  "architectural_decision_records":
}
EOF

# Tạo Rule bắt buộc Agent đọc não bộ
mkdir -p.agent/rules
cat >.agent/rules/99-memory-persistence.md <<EOF
# MEMORY PROTOCOL (FORCE READ)

## CRITICAL INSTRUCTION
Before writing any code or suggesting architecture, you MUST:
1. Read \`project_brain.json\` to understand the current phase and constraints.
2. Read \`docs/WORK_LOG.md\` to recall recent changes.

## SESSION END
When a task is done, you MUST using the \`memory_ops\` skill to update \`project_brain.json\` with new architectural decisions or phase changes.
EOF

# ==========================================
# 2. SANDBOX AN TOÀN (DEVCONTAINER)
# ==========================================
echo "🛡️ Đang thiết lập môi trường DevContainer an toàn..."
mkdir -p.devcontainer

# Tạo script cài đặt Android SDK headless (cho Mobile Agent)
cat >.devcontainer/setup_android.sh <<EOF
#!/bin/bash
set -e

# Chỉ chạy nếu chưa có Android SDK
if [! -d "/usr/lib/android-sdk" ]; then
    echo "📲 Installing Android SDK Command Line Tools..."
    sudo apt-get update && sudo apt-get install -y android-sdk openjdk-17-jdk
    
    # Thiết lập biến môi trường
    echo 'export ANDROID_HOME=/usr/lib/android-sdk' >> ~/.zshrc
    echo 'export PATH=\$PATH:\$ANDROID_HOME/cmdline-tools/latest/bin:\$ANDROID_HOME/platform-tools' >> ~/.zshrc
fi

echo "✅ Android Environment Ready."
EOF
chmod +x.devcontainer/setup_android.sh

# Tạo devcontainer.json đa năng (Monolith)
cat >.devcontainer/devcontainer.json <<EOF
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
      "extensions":
    }
  },
  "runArgs": [
    "--privileged", 
    "-v", "/dev/bus/usb:/dev/bus/usb" 
  ],
  "postCreateCommand": "bash.devcontainer/setup_android.sh",
  "remoteUser": "vscode"
}
EOF

# ==========================================
# 3. KỸ NĂNG THỰC THI (SKILLS)
# ==========================================
echo "🛠️ Đang cài đặt Skills nâng cao..."
mkdir -p.agent/skills/infra_ops
mkdir -p.agent/skills/memory_ops

# Skill: Memory Ops (Cập nhật não bộ)
cat >.agent/skills/memory_ops/SKILL.md <<EOF
# Memory Operations
Description: Updates the project_brain.json file safely.
Tools: python script
EOF

cat >.agent/skills/memory_ops/update_brain.py <<EOF
import json
import sys

def update_brain(key, value):
    with open('project_brain.json', 'r+') as f:
        data = json.load(f)
        # Logic cập nhật nested keys đơn giản
        keys = key.split('.')
        ref = data
        for k in keys[:-1]:
            ref = ref[k]
        ref[keys[-1]] = value
        
        f.seek(0)
        json.dump(data, f, indent=2)
        f.truncate()

if __name__ == "__main__":
    update_brain(sys.argv[1], sys.argv[2])
EOF

# Skill: Flutter Native Bridge (Tạo cầu nối Android/iOS)
mkdir -p.agent/skills/mobile_bridge
cat >.agent/skills/mobile_bridge/SKILL.md <<EOF
# Flutter Native Bridge
Description: Generates boilerplate for MethodChannel in Kotlin (Android) and Swift (iOS).
Usage: Use when user asks for "Native Camera" or "GPS Background".
EOF

# ==========================================
# 4. HẠ TẦNG (INFRASTRUCTURE)
# ==========================================
echo "🏗️ Đang dựng khung Infrastructure..."
mkdir -p infra/docker

# Docker Compose cho Local Dev
cat > infra/docker/docker-compose.yml <<EOF
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
EOF

echo "✅ ĐÃ NÂNG CẤP THÀNH CÔNG!"
echo "👉 Bước tiếp theo: Trong Antigravity, nhấn 'Reopen in Container' để kích hoạt môi trường an toàn."