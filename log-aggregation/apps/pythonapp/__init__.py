from flask import Flask
from werkzeug.routing import Rule


def create_app():
    app = Flask(__name__)

    app.debug = False
    app.use_reloader = False

    app.url_map.add(Rule('/', endpoint='index'))
    app.url_map.add(Rule('/http/<code>', endpoint='http_code'))
    app.url_map.add(Rule('/random', endpoint='random'))
    app.url_map.add(Rule('/exception', endpoint='exception'))

    app.logger.info('pythonapp.started')

    return app
