import pycurl
import time
import StringIO
import json
import random
from random import randint
import requests



response = StringIO.StringIO()


all = '''
while True:
    c = pycurl.Curl()
    c.setopt(c.URL, 'http://localhost:1949/all')
    c.setopt(pycurl.HTTPHEADER, ['X-Postmark-Server-Token: API_TOKEN_HERE', 'Accept: application/json'])
    c.setopt(pycurl.POST, 1)
    data = json.dumps([{"Key": "real", "Val": "24", "File": "ini.ini", "File_type": "ini"}])
    c.setopt(pycurl.POSTFIELDS, data)
    print (data)
    try:
        c.perform()
    except:
        print "the server is down ... "
    c.close()
    try:
        print response.getvalue()
    except:
        print "no server ..."
    response.close()
    time.sleep(5)
    print(randint(0, 50))


'''


while True:
    url = 'http://192.168.65.138:1949/change_this'
    randy_newman = randint(0, 50)  # i find this man random
    file_type = 'JANK'
    if randy_newman > 25:
        file_type = 'ini'
    else:
        file_type = 'XML'
    data = json.dumps([{"Key": "real", "Val": str(randy_newman), "File": file_type + "." + file_type, "File_type": file_type}])
#    data = '{"query":{"bool":{"must":[{"text":{"record.document":"SOME_JOURNAL"}},{"text":{"record.articleTitle":"farmers"}}],"must_not":[],"should":[]}},"from":0,"size":50,"sort":[],"facets":{}}'
    try:
        response = requests.get(url, data=data)
        time.sleep(1)
        print (response.json())
    except:
        print("server is downe ... \nsleeping \n")
        time.sleep(5)


joe = '''


while True:
    c = pycurl.Curl()
    c.setopt(c.URL, 'http://localhost:8081/change_this')
    c.setopt(pycurl.HTTPHEADER, ['X-Postmark-Server-Token: API_TOKEN_HERE', 'Accept: application/json'])
    c.setopt(pycurl.POST, 1)
    randy_newman = randint(0, 50)     # i find this man random
    data = json.dumps([{"Key": "real", "Val": str(randy_newman), "File": "ini.ini", "File_type": "ini"}])
    c.setopt(pycurl.POSTFIELDS, data)
    print (data)
    try:
        c.perform()
    except:
        print "the server is down ... "
    c.close()
    try:
        print response.getvalue()
    except:
        print "no server ..."
    response.close()
    time.sleep(2)
    print(randint(0, 50))

'''
