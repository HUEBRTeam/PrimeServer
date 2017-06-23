#!/usr/bin/env python

'''
    Client for getting profiles from original servers
'''

import socket
import struct
import time
import sys
import binascii
from prime import *

if len(sys.argv) == 2:
    
    if  ".bin" in sys.argv[1]:
        print "File Access Key Mode"
        f = open(sys.argv[1], "rb")
        data = f.read()
        f.close()
        ACCESS_KEY = binascii.hexlify(data)
        if len(ACCESS_KEY) != 32:
            print "Invalid Access Key!"
            exit(1)
        print "Read access key from %s is %s" %(sys.argv[1],ACCESS_KEY)
    else:
        ACCESS_KEY = sys.argv[1].lower().replace(" ","")
        if len(ACCESS_KEY) != 32:
            print "Invalid Access Key: %s (parsed to: %s)" %(sys.argv[1], ACCESS_KEY) 
            exit(1)

    print "Loading Public Key"
    pkf = open("spublick", "rb")
    pk = pkf.read()
    pkf.close()

    print "Loading Private Key"
    skf = open("sprivk", "rb")
    sk = skf.read()
    skf.close()


    #ACCESS_KEY = "56729a6f738a628b647dabaa3a56f1dd"
    #ACCESS_KEY = "e8da79a5b000e910030b3bca3ff97ab1"

    ACCESS_KEY = ACCESS_KEY.lower().replace(" ","")
    profile = DownloadProfile(ACCESS_KEY, pk, sk)
    print profile.Nickname
    f = open("profiles/"+ACCESS_KEY+".bin","wb")
    f.write(profile.ToBinary())
    f.close()
    print "Profile saved to %s" %("profiles/"+ACCESS_KEY+".bin")
else:
    print "Use: python client.py ACCESS_KEY"
    print "The ACCESS_KEY can be a prime.bin file or an key like: "
    print "\te8da79a5b000e910030b3bca3ff97ab1"
    print "\t\"e8 da 79 a5 b0 00 e9 10 03 0b 3b ca 3f f9 7a b1\""
