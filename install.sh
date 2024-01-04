#!/bin/bash
if [ "$EUID" -ne 0 ]; then
    echo "Run script as root, sudo -i."
    exit 1
fi
sudo apt update -y

architecture=$(uname -m)
if [ "$architecture" = "x86_64" ]; then
    if [ ! -f "go1.21.5.linux-amd64.tar.gz" ]; then
        wget https://github.com/Azumi67/UDP2RAW_FEC/releases/download/go/go1.21.5.linux-amd64.tar.gz
        sudo tar -C /usr/local -xzf go1.21.5.linux-amd64.tar.gz
    fi
elif [ "$architecture" = "aarch64" ]; then
    if [ ! -f "go1.21.5.linux-arm64.tar.gz" ]; then
        wget https://github.com/Azumi67/UDP2RAW_FEC/releases/download/go/go1.21.5.linux-arm64.tar.gz
        sudo tar -C /usr/local -xzf go1.21.5.linux-arm64.tar.gz
    fi
else
    echo "Unsupported arch: $architecture"
    exit 1
fi

echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bash_profile

source ~/.bash_profile
go mod init mymodule
go mod tidy

if [ -f "install.go" ]; then
    rm install.go
    echo "deleted previous version!"
fi

wget https://raw.githubusercontent.com/Azumi67/Rathole_reverseTunnel/main/install.go

go run install.go
