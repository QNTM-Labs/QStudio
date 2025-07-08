#!/bin/bash

echo "Setting up QStudio dependencies..."

if command -v go &> /dev/null; then
    echo "Go is already installed"
else
    echo "Installing Go..."
    if [[ "$OSTYPE" == "linux-gnu"* ]]; then
        if command -v apt &> /dev/null; then
            sudo apt update
            sudo apt install -y golang-go
        elif command -v dnf &> /dev/null; then
            sudo dnf install -y golang
        elif command -v pacman &> /dev/null; then
            sudo pacman -S --noconfirm go
        elif command -v zypper &> /dev/null; then
            sudo zypper install -y go
        else
            echo "Please install Go manually from https://golang.org/doc/install"
            exit 1
        fi
    fi
fi

if command -v git &> /dev/null; then
    echo "Git is already installed"
else
    echo "Installing Git..."
    if command -v apt &> /dev/null; then
        sudo apt install -y git
    elif command -v dnf &> /dev/null; then
        sudo dnf install -y git
    elif command -v pacman &> /dev/null; then
        sudo pacman -S --noconfirm git
    elif command -v zypper &> /dev/null; then
        sudo zypper install -y git
    fi
fi

echo "Building QStudio..."
go mod init qstudio
go get github.com/BurntSushi/toml
go build -o qstudio main.go

echo "Installing QStudio..."
sudo mv qstudio /usr/local/bin/
sudo chmod +x /usr/local/bin/qstudio

echo ""
echo "QStudio installed successfully!"
echo "================================"
echo "Usage:"
echo "  qstudio install    - Install and setup Wine prefix with dependencies"
echo "  qstudio config     - Show current configuration"
echo "  qstudio launch     - Launch Roblox Studio"
echo "  qstudio            - Launch Roblox Studio (default)"
echo ""
echo "Configuration file: ~/.qstudio/config.toml"
echo "Overlays directory: ~/.qstudio/overlay/"
echo ""
echo "Run 'qstudio install' to set up Wine and dependencies automatically!"
