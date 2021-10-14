import json
import os
from http import HTTPStatus

from flask import Flask
from flask import Response
import random

HTTP_STATUS_CODES = list(map(lambda x: x.value, HTTPStatus))
HTTP_STATUS_CODES_COUNT = len(HTTP_STATUS_CODES)

APP_VERSION = os.environ.get('APP_VERSION')


app = Flask(__name__)

def create_bad_request_response():
    status = HTTPStatus.BAD_REQUEST

    data = json.dumps({
                'status_code': status.value,
                'description': status.phrase
            })

    return Response(
            data, 
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
    index = random.randint(0, HTTP_STATUS_CODES_COUNT)
    code = HTTP_STATUS_CODES[index]
    status = list(filter(lambda x: x.value == code, HTTPStatus))[0]
    data = json.dumps({
            'status_code': code,
            'description': status.phrase
        })

    return Response(
        data, 
        status=status.value, 
        mimetype='application/json'
    )

@app.route("/http/<code>")
def http_code(code=HTTPStatus.OK.value):
    """ Return with provided HTTP response code. Return HTTP 400 - Bad request
        if code provided is invalid
    """
    if not code.isnumeric():
        return create_bad_request_response()
    
    try:
        status = HTTPStatus(int(code))
        data = json.dumps({
                'status_code': status.value,
                'description': status.phrase
            })
    except ValueError as e:
        print(e)

        return create_bad_request_response()
    
    return Response(
            data, 
            status=status.value, 
            mimetype='application/json'
        )
    

@app.route("/exception")
def exception():
    """ Raise exception
    """

    raise Exception('Doing this only for the logs')