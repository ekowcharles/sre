# Log Aggregation

This SRE module shows how metrics and logs are aggregated from multiple services into Graphite and Elasticsearch respectively.

## Design

![Flow of logs](./docs/images/flow.png)

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
1. Run `curl http://localhost:8080/python/` to use [pythonapp](./apps/pythonapp) endpoints
1. Run `curl http://localhost:8080/java/` to use [javaapp](./apps/javaapp) endpoints
1. Run `curl http://localhost:8080/golang/` to use [golangapp](./apps/golangapp) endpoints

See [here](./apps/README.md) for application endpoints and descriptions.

## Stack

1. [Flask](https://flask.palletsprojects.com/en/2.0.x/quickstart/#a-minimal-application)
1. [Dropwizard](https://www.dropwizard.io/en/latest/getting-started.html)
1. [net/http](https://pkg.go.dev/net/http)
1. [Vagrant](https://www.vagrantup.com/docs)
1. [Ansible](https://docs.ansible.com/ansible-core/devel/index.html)
1. [nginx](https://docs.nginx.com/nginx/admin-guide/monitoring/logging)