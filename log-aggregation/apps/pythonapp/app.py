import json
import logging
import logging.config
import os
import random
import traceback
from http import HTTPStatus
from logging import config

import yaml
from flask import Response, escape, request

from . import create_app
from .util import build_response

HTTP_STATUS_CODES = list(map(lambda x: x.value, HTTPStatus))
HTTP_STATUS_CODES_COUNT = len(HTTP_STATUS_CODES)

APP_VERSION = os.environ.get('APP_VERSION')

with open('./logging.yaml', 'rt') as file:
    config = yaml.safe_load(file.read())
    logging.config.dictConfig(config)

app = create_app()


@app.endpoint('index')
def index():
    """ Return app version
    """

    data = json.dumps({
        'version': APP_VERSION
    })

    return Response(
        data,
        status=HTTPStatus.OK.value,
        mimetype='application/json'
    )


@app.endpoint('random')
def randomize():
    """ Return with a random HTTP response code. Include HTTP status code
        description in response
    """

    app.logger.debug('Retrieving random http response ...')
    ind = random.randint(0, HTTP_STATUS_CODES_COUNT)

    # trying to access the last position throws IndexError error which is OK
    code = HTTP_STATUS_CODES[ind]
    status = HTTPStatus(code)

    return build_response(status)


@app.endpoint('http_code')
def http_code(code=HTTPStatus.OK.value):
    """ Return with provided HTTP response code. Return HTTP 400 - Bad request
        if code provided is invalid
    """
    status = HTTPStatus.BAD_REQUEST

    try:
        app.logger.info('Fetching status for %s ...', code)
        status = HTTPStatus(int(escape(code)))
    except ValueError as error:
        app.logger.error(error)

    return build_response(status)


@app.endpoint('exception')
def exception():
    """ Raise exception
    """

    raise Exception('Doing this only for the logs')


@app.before_first_request
def before_first_request_func():
    app.logger.info('pythonapp.first')


@app.before_request
def before_request():
    app.logger.info('%s %s%s %s %s %s', request.method,
                    request.full_path,
                    request.query_string.decode("utf-8"),
                    request.scheme,
                    request.host,
                    request.remote_addr)


@app.errorhandler(Exception)
def exceptions(error):
    app.logger.error('%s \n%s', error, traceback.format_exc())
    return build_response(HTTPStatus.INTERNAL_SERVER_ERROR)
