# Log Aggregation

## Design

### Log file locations

The location of the logs that are being aggregated are as follows:

|Application|Log Files|
| :----: | :---- |
|nginx|/var/log/nginx/error.log, /var/log/nginx/access.log|
|javaapp|/var/log/javaapp/error.log|
|golangapp|/var/log/golangapp/error.log|
|pythonapp|/var/log/pythonapp/error.log|

## Before you start

## Before you start

1. Install [golang](https://golang.org/doc/install)
1. Install [mvn](https://maven.apache.org/install.html)
1. Run the following in your favorite terminal to generate a key pair:
   ```sh
   ssh-keygen -m PEM -t rsa -b 4096
   ```
1. Place a copy of the **public key** in the path `log-aggregation/provisioning/nginx/files/public_keys/www.pub`

   > Do not check in private key into version control!

## Usage

1. Run `make all`
1. Run `curl http://localhost:8080` to use the nginx load balancer to route to any of the application endpoints
1. Run `curl http://localhost:8080/python/` to use pythonapp endpoints
1. Run `curl http://localhost:8080/java/` to use javaapp endpoints
1. Run `curl http://localhost:8080/golang/` to use golangapp endpoints

See [here](../apps/README.md) for application endpoints and descriptions.

## Stack

1. [virtualbox](https://www.virtualbox.org/)
1. [vagrant](https://learn.hashicorp.com/collections/vagrant/getting-started)
1. [ansible](https://www.tutorialspoint.com/ansible/ansible_introduction.htm)
1. [systemd](https://www.linux.com/training-tutorials/understanding-and-using-systemd/)
1. [NGINX](https://www.netguru.com/codestories/nginx-tutorial-basics-concepts)

## References

1. [Multiple VMs, one Vagrantfile](https://www.thisprogrammingthing.com/2015/multiple-vagrant-vms-in-one-vagrantfile/)
1. [systemd.service Synopsis](https://www.freedesktop.org/software/systemd/man/systemd.service.html#)
1. [systemd.exec Synopsis](https://www.freedesktop.org/software/systemd/man/systemd.exec.html)
1. [auditd Configuration](https://linux.die.net/man/5/auditd.conf)
1. [Ansible Loops](https://docs.ansible.com/ansible/2.5/user_guide/playbooks_loops.html)
1. [Ansbile Playbook Loops](https://docs.ansible.com/ansible/latest/user_guide/playbooks_loops.html)
