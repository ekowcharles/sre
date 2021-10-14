import json
import os
import random
from http import HTTPStatus

from flask import Flask, Response, escape
from werkzeug.routing import Rule

HTTP_STATUS_CODES = list(map(lambda x: x.value, HTTPStatus))
HTTP_STATUS_CODES_COUNT = len(HTTP_STATUS_CODES)

APP_VERSION = os.environ.get('APP_VERSION')


app = Flask(__name__)
app.url_map.add(Rule('/', endpoint='index'))
app.url_map.add(Rule('/http/<code>', endpoint='http_code'))
app.url_map.add(Rule('/random', endpoint='random'))
app.url_map.add(Rule('/exception', endpoint='exception'))



def build_response_payload(status):
    """ Build response payload in app response format
    """

    return json.dumps({
        'code': status.value,
        'description': status.phrase
    })

def build_response(status):
    """ Build response
    """

    return Response(
            build_response_payload(status),
            status=status.value,
            mimetype='application/json'
        )

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

    ind = random.randint(0, HTTP_STATUS_CODES_COUNT)
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
        status = HTTPStatus(int(escape(code)))
    except ValueError as error:
        print(error)

    return build_response(status)

@app.endpoint('exception')
def exception():
    """ Raise exception
    """

    raise Exception('Doing this only for the logs')
