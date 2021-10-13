base:
	vagrant up base
	vagrant package --base base --output base.box

all:
	vagrant up log
	vagrant up ping
	vagrant up nginx

provision:
	vagrant provision log
	vagrant provision ping
	vagrant provision nginx

destroy:
	vagrant destroy --force