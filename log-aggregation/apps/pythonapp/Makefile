SHELL = /bin/bash

default: run

install:
	python3 -m venv venv
	. venv/bin/activate
	python3 -m pip install -r ./requirements.pip

run:
	FLASK_APP=app \
	FLASK_ENV=development \
	APP_VERSION=0.1.0 \
		flask run --host=0.0.0.0 --port 8030

package:
	rm -f pythonapp.zip
	zip -vr pythonapp.zip ./ -x "*.DS_Store*" "*.zip*" "*__pycache__*" "*venv*" "*logs*"
