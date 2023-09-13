#!/usr/bin/env bash

wifi=$(nmcli dev wifi | gum choose | sed -n 's/ \+/ /gp' | tee res.txt | cut -d " " -f 3)

password=$(gum input --password --placeholder="Password" --prompt="# ")

echo "Proof of concept with gum..."
echo "This command should be executed after choosing the ssid and entering the password"
echo "sudo nmcli dev wifi connect $wifi password $password"

