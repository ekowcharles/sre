base:
	vagrant up base
	vagrant package --base base --output base.box

prepare:
	cd ../apps/golangapp && $(MAKE) package && \
							cp golangapp ../../virtualization/provisioning/apps/files
	cd ../apps/pythonapp && $(MAKE) package && \
							cp pythonapp.zip ../../virtualization/provisioning/apps/files
	cd ../apps/javaapp && $(MAKE) package && \
						  cp target/javapp-1.0-SNAPSHOT.jar ../../virtualization/provisioning/apps/files && \
						  cp config.yaml ../../virtualization/provisioning/apps/files

all:
	$(MAKE) prepare
	vagrant up app

destroy:
	vagrant destroy --force