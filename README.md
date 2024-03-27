# Display Test

**Purpose**: just test the i2c display ssd1306 128x32 with orange pi zero 3

building:
```bash
make build
```

## systemd

1. `make copy_remote`
2. loggin into the remote host: `sudo cp /tmp/display-test /usr/local/bin`
3. create the service:
```bash
echo '[Unit]
Description=Display Test
After=network.target

[Service]
ExecStart=/usr/local/bin/display-test
Restart=always

[Install]
WantedBy=multi-user.target' | sudo tee /etc/systemd/system/display_test.service >/dev/null
```
4. `sudo systemctl daemon-reload`
5. `sudo systemctl enable display_test.service`
6. `sudo systemctl start display_test.service`
7. `sudo systemctl status display_test.service`
