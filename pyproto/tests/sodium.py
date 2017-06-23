#!/usr/bin/env python

'''
    LibSodium Tests
'''

from pysodium import *

pkf = open("../../RE/spublick", "rb")
pk = pkf.read()
pkf.close()

skf = open("../../RE/sprivk", "rb")
sk = skf.read()
skf.close()

f = open("../../RE/Packets/RawSample/Unencrypted.bin","rb")
data = f.read()
f.close()

f = open("../../RE/Packets/RawSample/Nounce.bin", "rb")
nounce = f.read()
f.close()

f = open("../../RE/Packets/RawSample/Encrypted.bin", "rb")
enctest = f.read()
f.close()

encrypted = crypto_box(data, nounce, pk, sk)
print "Encryption Test: %s" %(encrypted==enctest)

dectest = data

f = open("../../RE/Packets/RawSample/Encrypted.bin","rb")
data = f.read()
f.close()


decrypted = crypto_box_open(data, nounce, pk, sk)
print "Decryption Test: %s" %(dectest == decrypted)