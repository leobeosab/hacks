# Netctl
Copy these files and folders to the required locations or generate them yourselves on another computer (easier ways sure but this is the hacks repo)

## output of enable command for reference
▶ sudo netctl enable usbpizero              
'/etc/systemd/system/multi-user.target.wants/netctl@usbpizero.service' -> '/usr/lib/systemd/system/netctl@.service'
generated '/etc/systemd/system/netctl@usbpizero.service.d/profile.conf'

## Copy commands for the lazy assuming your mounted root folder is in cwd
```
sudo cp /etc/systemd/system/multi-user.target.wants/netctl@usbpizero.service root/etc/systemd/system/multi-user.target.wants/netctl@usbpizero.service
sudo cp /etc/systemd/system/multi-user.target.wants/netctl@usbpizero.service root/usr/lib/systemd/system/netctl@.service
sudo cp /etc/systemd/system/netctl@usbpizero.service.d root/etc/systemd/system/netctl@usbpizero.service.d -r
```
