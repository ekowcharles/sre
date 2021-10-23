import json
from flask import Response

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
