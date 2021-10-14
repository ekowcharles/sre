import json
import os
import random
from http import HTTPStatus

from flask import Flask, Response

HTTP_STATUS_CODES = list(map(lambda x: x.value, HTTPStatus))
HTTP_STATUS_CODES_COUNT = len(HTTP_STATUS_CODES)

APP_VERSION = os.environ.get('APP_VERSION')


app = Flask(__name__)


def create_response_payload(status):
    """ Create response payload in app response format
    """

    return json.dumps({
        'status_code': status.value,
        'description': status.phrase
    })

def create_response(status):
    """ Create response
    """

    return Response(
            create_response_payload(status),
            status=status.value,
            mimetype='application/json'
        )

@app.route("/")
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

@app.route("/random")
def randomize():
    """ Return with a random HTTP response code. Include HTTP status code
        description in response
    """

    ind = random.randint(0, HTTP_STATUS_CODES_COUNT)
    code = HTTP_STATUS_CODES[ind]
    status = HTTPStatus(code)

    return create_response(status)

@app.route("/http/<code>")
def http_code(code=HTTPStatus.OK.value):
    """ Return with provided HTTP response code. Return HTTP 400 - Bad request
        if code provided is invalid
    """
    status = HTTPStatus.BAD_REQUEST

    try:
        status = HTTPStatus(int(code))
    except ValueError as error:
        print(error)

    return create_response(status)
    

@app.route("/exception")
def exception():
    """ Raise exception
    """

    raise Exception('Doing this only for the logs')
