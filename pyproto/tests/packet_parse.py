#!/usr/bin/env python

import sys
sys.path.insert(0, "../")

import prime

MACHINE_PACKET = "../../RE/Packets/Start/sodium_53_out_pkt.bin"
SCORE_PACKET = "../../RE/Packets/MusicScore/sodium_16_out_pkt.bin"
LOGIN_PACKET = "../../RE/Packets/Login/sodium_56_out_pkt.bin"
PROFILE_PACKET = "../../RE/Packets/Login/sodium_12_in_pkt.bin"
WORLD_BEST_PACKET = "../../RE/Packets/World Best/sodium_12_in_pkt.bin"

print "Parsing MachineInfoPacket %s" %MACHINE_PACKET

f = open(MACHINE_PACKET,"rb")
data = f.read()
f.close()

mp = prime.MachineInfoPacket()
mp.FromBinary(data)
mp.Print()

print ""
print "Parsing ScoreBoardPacket %s" %SCORE_PACKET

f = open(SCORE_PACKET, "rb")
data = f.read()
f.close()

sp = prime.ScoreBoardPacket()
sp.FromBinary(data)
sp.Print()

print ""
print "Parsing LoginPacket %s" %LOGIN_PACKET

f = open(LOGIN_PACKET, "rb")
data = f.read()
f.close()

lp = prime.LoginPacket()
lp.FromBinary(data)
lp.Print()

print ""
print "Parsing Profile Packet %s" %PROFILE_PACKET

f = open(PROFILE_PACKET, "rb")
data = f.read()
f.close()

pp = prime.ProfilePacket()
pp.FromBinary(data)
pp.Print()

print ""
print "Parsing World Best Packet %s" %WORLD_BEST_PACKET

f = open(WORLD_BEST_PACKET, "rb")
data = f.read()
f.close()

pp = prime.WorldBestPacket()
pp.FromBinary(data)
pp.Print()