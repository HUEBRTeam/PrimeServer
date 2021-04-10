#!/usr/bin/env python

'''
    TODO
    On Piu Exec the IP is at 0x16633e
    115.68.108.183_60000

    You can also make a host entry on /etc/hosts:
    127.0.0.1       localhost prime_daemon.piugame.com
'''

import socket
import struct
from prime import *

print "Loading Public Key"
pkf = open("../RE/spublick", "rb")
pk = pkf.read()
pkf.close()

print "Loading Private Key"
skf = open("../RE/sprivk", "rb")
sk = skf.read()
skf.close()

print "Loading Test Profile Packet"
f = open("../RE/Packets/Login/sodium_12_in_pkt.bin", "rb")
data = f.read()
f.close()

profpack = ProfilePacket()
profpack.FromBinary(data)

def ProcessPacket(packet, socket):
    packtype = struct.unpack("<I",packet[4:8])[0]
    print "Received Packet %x"%packtype
    if packtype == MachineInfoPacket.PacketType:
        data = MachineInfoPacket()
        data.FromBinary(packet)
        data.Print()
        ack = ACKPacket()
        ba = EncryptPacket(ack.ToBinary(), pk, sk)
        print "Sending ACK"
        socket.send(ba)
    elif packtype == MachineInfoPacket_v2.PacketType:
        data = MachineInfoPacket_v2()
        data.FromBinary(packet)
        data.Print()
        ack = ACKPacket()
        ba = EncryptPacket(ack.ToBinary(), pk, sk)
        print "Sending ACK"
        socket.send(ba)
    elif packtype == LoginPacket.PacketType:
        data = LoginPacket()
        data.FromBinary(packet)
        data.Print()
        print "Login Request from %s" %data.AccessCode
        profpack.AccessCode = data.AccessCode
        ba = EncryptPacket(profpack.ToBinary(), pk, sk)
        print "Sending %s Profile" %profpack.Nickname
        socket.send(ba)
    elif packtype == EnterProfilePacket.PacketType:
        print "EnterProfile Received"
        data = EnterProfilePacket()
        data.FromBinary(packet)
        data.Print()
    elif packtype == ScoreBoardPacket.PacketType:
        print "Score Received"
        data = ScoreBoardPacket()
        data.FromBinary(packet)
        data.Print()
    elif packtype == RequestLevelUpInfoPacket.PacketType:
        print "Received Request for Level"
        data = RequestLevelUpInfoPacket()
        data.FromBinary(packet)
        data.Print()
        out = LevelUpInfoPacket()
        data.Level = profpack.Level
        print "Sending same level"
        ba = EncryptPacket(out.ToBinary(), pk, sk)
        socket.send(ba)
    elif packtype == GameOverPacket.PacketType:
        print "Received a GameOver Packet"
        data = GameOverPacket()
        data.FromBinary(packet)
        data.Print()


HOST = ''               # Endereco IP do Servidor
PORT = 60010            # Porta que o Servidor esta
tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
tcp.setsockopt(socket.SOL_SOCKET, socket.SO_REUSEADDR, 1)
orig = (HOST, PORT)
tcp.bind(orig)
tcp.listen(1)
msg = ''

print 'Starting server'
while True:
    con, cliente = tcp.accept()
    print 'Connected by', cliente
    while True:
        msg += con.recv(4)
        if len(msg) > 0:
            size = struct.unpack("<I", msg[:4])[0]
            msg = msg[4:]
            if (size-4-len(msg) <= 0):
              break
            msg += con.recv(size-4-len(msg))
            while len(msg) < size-4:
                msg += con.recv((size-4) - len(msg))

            data = DecryptPacket(msg, pk, sk)
            ProcessPacket(data, con)
        if not msg: break
    print 'Closing client connection', cliente
    con.close()
