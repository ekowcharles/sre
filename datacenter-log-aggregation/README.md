# Log Aggregation

# Stack

1. vagrant
1. ansible
1. systemd
1. rsyslog
1. nginx
1. go
1. java


## Generate key-pair

```sh
ssh-keygen -m PEM -t rsa -b 4096
```

## References

1. [Multiple VMs, one Vagrantfile](https://www.thisprogrammingthing.com/2015/multiple-vagrant-vms-in-one-vagrantfile/)
1. [systemd.service Synopsis](https://www.freedesktop.org/software/systemd/man/systemd.service.html#)
1. [systemd.exec Synopsis](https://www.freedesktop.org/software/systemd/man/systemd.exec.html)
1. [Setting Up syslog](https://www.tecmint.com/install-rsyslog-centralized-logging-in-centos-ubuntu/)
1. [rsyslog Documentation](https://www.rsyslog.com/doc/v8-stable/)
1. [auditd Configuration](https://linux.die.net/man/5/auditd.conf)
1. [syslog Filters](https://kifarunix.com/a-basic-introduction-to-rsyslog-filters/2/)
1. [rsyslog Docs](https://rsyslog-mm.readthedocs.io/en/v7.4_stable/index.html)
1. [nginx syslog Configuration](http://nginx.org/en/docs/syslog.html)
1. [syslog Protocol RFC](https://tools.ietf.org/html/rfc3164#section-4.1.1)