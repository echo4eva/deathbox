#!/bin/sh

if [ -f "$HOME/deathbox.conf" ]; then
    . "$HOME/deathbox.conf"
else
    echo "Exiting. Must have a deathbox.conf at $HOME/deathbox.conf"
    exit 1
fi

if [ -z "$DEATHBOX_HOST" ] || [ -z "$DEATHBOX_CLIENT_DEVICE" ]; then
    echo "Exiting. Must enter a host/client device in $HOME/deathbox.conf"
    exit 1
fi

echo "Sending heartbeat to Host"
curl -X POST -H "Device: $DEATHBOX_CLIENT_DEVICE" -H "Secret: $DEATHBOX_SHARED_SECRET" $DEATHBOX_HOST/heartbeat