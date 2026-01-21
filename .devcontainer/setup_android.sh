#!/bin/bash
set -e
if [ ! -d "/usr/lib/android-sdk" ]; then
    echo "📲 Installing Android SDK Command Line Tools..."
    sudo apt-get update && sudo apt-get install -y android-sdk openjdk-17-jdk
    echo 'export ANDROID_HOME=/usr/lib/android-sdk' >> ~/.zshrc
    echo 'export PATH=$PATH:$ANDROID_HOME/cmdline-tools/latest/bin:$ANDROID_HOME/platform-tools' >> ~/.zshrc
fi
echo "✅ Android Environment Ready."