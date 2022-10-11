package linux

/*
lscpu or less /proc/cpuinfo

See free memory
cat /proc/meminfo
cat /proc/cpuinfo

// find kernel version
cat /proc/version

Finding releale datasets
lsb_release -a

//find hardware nfo
dmidecode

// list running services
service --status-all
pstree
systemctl

To show all installed unit files use
  systemctl list-unit-files
//check for services enabled on boot
chkconfig --list


//List system services enabled on boot up
systemctl list-unit-files --state=enabled

// list system running services
systemctl list-units --type=service --state=running

systemd-cgtop shows top control groups by their resource usage such as tasks, CPU, Memory, Input, and Output:
systemd-cgtop

//check service status
service name status

*/
