#!/bin/bash

if [[ $UID -ne 0 ]]; then
    sudo -p 'Restarting as root, password: ' bash $0 "$@"
    exit $?
fi

rm -rf /opt/runify
rm -rf /etc/xdg/autostart/runify-autostart.desktop
