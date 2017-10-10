#coding=utf8
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
import requests
from Crypto.Cipher import AES
from binascii import a2b_hex
from Crypto.PublicKey import RSA
from Crypto.Cipher import PKCS1_v1_5
import random
import json
import binascii

site = 'http://127.0.0.1:9999'
aesKey = ''

def get_public_pem():
    global site, key
    url = site + '/public.pem'
    print url,
    r = requests.get(url)
    return r.status_code, r.text

def login(key):
    global site, aesKey
    url = site + '/auth/login'
    aesKey = str(int(1000000000000000 + random.random() * 8999999999999999))[:16]
    data = """{"a":"%s","p":"%s","k":"%s"}""" % ("leonardocaesarz@gmail.com", "123456", aesKey)

    cipher = PKCS1_v1_5.new(key)
    data = cipher.encrypt(data)

    print url,
    r = requests.post(url, data=data)
    return r.status_code, r.content

def checkErr(code):
    if code != 200:
        print '[FAIL]'
        exit()
    print ' [OK]'

def deAes(key, data):
    cryptor = AES.new(key, AES.MODE_CBC, key)
    plain_text = cryptor.decrypt(data)
    return plain_text.rstrip('\0')

code, text = get_public_pem()
checkErr(code)

code, text = login(RSA.importKey(text))
print code, len(text)

result = json.loads(deAes(aesKey, text))
print result
