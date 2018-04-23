#!/bin/bash
apt-get update && apt-get install -y x11vnc
ln -s vnc/startVnc.sh ~/startVnc.sh
