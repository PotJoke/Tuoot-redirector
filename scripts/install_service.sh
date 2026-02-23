sudo bash << 'SCRIPT'
cat > /etc/systemd/system/redirection.service << 'EOF'
[Unit]
Description=http trafic redirection service
After=network.target

[Service]
Type=simple
User=www-data
WorkingDirectory=/opt/Tuoot-redirector
ExecStart=go run /opt/Tuoot-redirector/.
Restart=on-failure
RestartSec=5s

[Install]
WantedBy=multi-user.target
EOF

SCRIPT

apt install nginx -y
if [ -f "nginx.config" ]; then
    echo "nginx.config found, copying to /etc/nginx/sites-available/default"
else
    echo "nginx.config not found, please make sure it is in the current directory"
    exit 1
fi

sudo systemctl daemon-reload
sudo systemctl enable redirection.service
sudo systemctl start redirection.service
sudo systemctl restart nginx