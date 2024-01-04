#!/bin/bash

if [[ $EUID -ne 0 ]]; then
echo "This script must be run as root."
exit 1
fi
apt update -y
sudo apt-get install build-essential
apt install rustc -y
apt install cargo -y
apt-get install pkg-config libssl-dev -y
curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh

source "$HOME/.cargo/env"

