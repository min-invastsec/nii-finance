# -*- coding: utf-8 -*-
"""
Created on Thu Jun 02 22:36:21 2016

@author: Jacob
"""

import test_pb2
import requests

payload = {"jsonrpc": "2.0",
           "method": "postMessage",
           "params": [{"name": "Jacob"}],
           "id": 99}

person = test_pb2.Person()
person.id = 1234
person.name = "John Doe"
person.email = "jdoe@example.com"

url = "http://127.0.0.1:4000"

# POST with protobuf 
r = requests.post(url, data=person.SerializeToString())

print r.text
print r.headers['content-type']
print r.status_code

