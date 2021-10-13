install:
	python3 -m venv venv
	. venv/bin/activate
	python3 -m pip install -r ./requirements.pip

run:
	FLASK_APP=app \
	FLASK_ENV=development \
		flask run --host=0.0.0.0 --port 8030