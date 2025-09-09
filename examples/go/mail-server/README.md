# Mail Server Setup

After updating the credentials in `main.go:15`, follow the script:

```sh
# Build the app
GOOS=linux go build -o app .

# Modify the exec path of app as necessary in .service file
sudo cp mail-server.service /etc/systemd/system/

# Enable on login and start immediately
sudo systemctl daemon-reload
sudo systemctl enable --now go-mail-server

# Check logs to ensure it is running
sudo journalctl -u go-mail-server
```

To send an email:

```sh
curl -H "Authorization: some-random-string" -d "email body" localhost:8080
```
