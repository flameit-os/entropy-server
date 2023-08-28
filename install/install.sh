#!/bin/bash
# Author: Paweł 'felixd' Wojciechowski (c) FlameIT - Immersion Cooling - 2023

# Designed for RaspberryPi

BINARY="flameit-entropy-server"

# Building Tank Monitoring app from source
echo ""
echo "####### FlameIT - Immersion Cooling #######"
echo "Building Tank Monitoring app from source"
echo "####### FlameIT - Immersion Cooling #######"

cd ..
git pull
go build
cd install

if [ -f "../$BINARY" ]; then
    # Install Tank-Monitoring service
    echo ""
    echo "####### FlameIT - Immersion Cooling #######"
    echo "Installing Entropy Server for RaspberryPi"
    echo "####### FlameIT - Immersion Cooling #######"
    sudo cp flameit-entropy-server.service /lib/systemd/system/flameit-entropy-server.service
    sudo systemctl enable flameit-entropy-server.service
    sudo service flameit-entropy-server restart
    journalctl -u flameit-entropy-server.service --no-pager
else
    echo ""
    echo "####### FlameIT - Immersion Cooling #######"
    echo "I was not able to find application binary file."
    echo "####### FlameIT - Immersion Cooling #######"
    ls -al ..
fi
