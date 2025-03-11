# deathbox - a dead man's switch

A monitoring service that listens to heartbeats from other devices and executes actions if no heartbeats are received after an inactive duration.

By default, sending a tweet is implemented as the action.

---

## Service Instructions

Intended for VPS hosting on Linux.

```shell
git clone https://github.com/echo4eva/deathbox.git
cd deathbox
cp .env.example .env
docker compose up --build
```

- Change shared secret in `.env`, this is what the client will use to authenticate with the service. Also change or remove Twitter keys/tokens.
- Refactor code as you see fit. `db.go` to change the duration, or `main.go` to remove tweeting action.

---

## Client Instructions

Intended for Linux devices. Shell script runs on startup to send heartbeat to service.

Important files in this repo's /scripts directory
- `deathbox.service` - systemd unit to run shell script on startup
- `deathbox.sh` - shell script to curl the service
- `deathbox.conf` - config file that holds environment variables for the shell script


1. Setting up `deathbox.service`
```shell
mkdir -p ~/.config/systemd/user
curl https://raw.githubusercontent.com/echo4eva/deathbox/refs/heads/main/scripts/deathbox.service -o ~/.config/systemd/user/deathbox.service
systemctl --user enable deathbox
```

2. Setting up `deathbox.sh`
```shell
curl https://raw.githubusercontent.com/echo4eva/deathbox/refs/heads/main/scripts/deathbox.sh -o ~/.local/bin/deathbox
chmod +x ~/.local/bin/deathbox
```

3. Setting up `deathbox.conf` AND EDIT
```shell
curl https://raw.githubusercontent.com/echo4eva/deathbox/refs/heads/main/scripts/.deathbox.conf -o ~/.deathbox.conf
```

Make sure to edit the configuration:
- `DEATHBOX_HOST` - address of your service to send heartbeat to
- `DEATHBOX_CLIENT_DEVICE` - name of the client device
- `DEATHBOX_SHARED_SECRET` - value for authenticating client to service, should be the same

---

## Acknowledgements

[5HT2B/heartbeat](https://github.com/5HT2B/heartbeat) - Used as reference for infrastructure and client setup! Shoutout to an oomf for telling me about this project.