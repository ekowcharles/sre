prepare:
	mkdir -p ./provisioning/apps/files/golangapp ./provisioning/apps/files/javaapp ./provisioning/apps/files/pythonapp
	cd ../apps/golangapp && $(MAKE) package && \
							cp golangapp ../../virtualization/provisioning/apps/files/golangapp
	cd ../apps/pythonapp && $(MAKE) package && \
							unzip -o pythonapp.zip -d ../../virtualization/provisioning/apps/files/pythonapp
	cd ../apps/javaapp && $(MAKE) package && \
						  cp target/javapp-1.0-SNAPSHOT.jar ../../virtualization/provisioning/apps/files/javaapp && \
						  cp config.yml ../../virtualization/provisioning/apps/files/javaapp

all:
	$(MAKE) prepare
	vagrant provision app
	vagrant up app

destroy:
	vagrant destroy --force
	rm -f *.box
	rm -rf .vagrant

clean:
	$(MAKE) destroy

	rm -rf virtualization/provisioning/apps/files/pythonapp
	rm -rf virtualization/provisioning/apps/files/golangapp
	rm -rf virtualization/provisioning/apps/files/javaapp

	rm -rf "/Users/ekowcharles/VirtualBox VMs/app"