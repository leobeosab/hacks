## Resources
https://medium.com/@dwilkins/usb-gadget-mode-with-arch-linux-and-the-raspberry-pi-zero-e70a0f17730a#.lx2umjihv

https://archlinuxarm.org/platforms/armv6/raspberry-pi

## Getting netctl working without enable commands
systemd processes are controlled via symlinks

Example uses bridge conf
netctl copies `/etc/systemd/system/multi-user.target.wants/netctl@bridge.service -> /usr/lib/systemd/system/netctl@.service `
and generates `/etc/systemd/system/netctl@bridge.service.d/profile.conf`


## Enable dhcpd
/etc/systemd/system/multi-user.target.wants/dhcpcd.service -> /usr/lib/systemd/system/dhcpcd.service

sudo ln -s /usr/lib/systemd/system/dhcpcd.service root/etc/systemd/system/multi-user.target.wants/dhcpcd.service
