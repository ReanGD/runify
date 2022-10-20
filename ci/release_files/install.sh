#!/bin/bash

if [[ $UID -ne 0 ]]; then
    sudo -p 'Restarting as root, password: ' bash $0 "$@"
    exit $?
fi

SCRIPT_DIR=`cd -- "$( dirname -- "${BASH_SOURCE[0]}" )" &> /dev/null && pwd`

rm -rf /opt/runify
mkdir -p /opt/runify
cp -r ${SCRIPT_DIR}/opt/runify/* /opt/runify/
cp ${SCRIPT_DIR}/runify-autostart.desktop /etc/xdg/autostart/
chmod 0755 /opt/runify/runify-server /opt/runify/runify-ui
find /opt/runify/ -type d -exec chmod 755 {} \;
