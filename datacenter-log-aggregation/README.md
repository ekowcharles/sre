# Log Aggregation

## Design

## Before you start

1. Run the following in your favorite terminal to generate a key pair:
   ```sh
   ssh-keygen -m PEM -t rsa -b 4096
   ```
1. Place a copy of the **public key** in the path `datacenter-log-aggregation/provisioning/nginx/files/public_keys/www.pub`

   > Do not check in private key into version control!

## Usage

### Testing NGINX Log Collection

### Testing go Application Log Collection

### Testing Java Application Log Collection

## Stack

1. [vagrant](https://learn.hashicorp.com/collections/vagrant/getting-started)
1. [ansible](https://www.tutorialspoint.com/ansible/ansible_introduction.htm)
1. [systemd](https://www.linux.com/training-tutorials/understanding-and-using-systemd/)
1. [rsyslog](https://www.rsyslog.com/guides/)
1. [NGINX](https://www.netguru.com/codestories/nginx-tutorial-basics-concepts)
1. [go](https://golang.org/doc/tutorial/getting-started)
1. [java](https://www.dropwizard.io/en/latest/getting-started.html)

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
