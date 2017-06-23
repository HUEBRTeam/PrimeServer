#!/usr/bin/env python
import os
import struct 
import sys

sys.path.insert(0, "../")

import prime
navfolder = "../../RE/Packets/BoughtSongs"
#navfolder = "../../DSUNA_Packets"
packettypes = [ prime.ACKPacket, prime.KeepAlivePacket, prime.ByePacket, prime.RequestWorldBest, prime.RequestRankMode, prime.RankModePacket, prime.EnterProfilePacket, prime.GameOverPacket, prime.LevelUpInfoPacket, prime.LoginPacket, prime.MachineInfoPacket, prime.MachineInfoPacket_v2, prime.ProfilePacket, prime.RequestLevelUpInfoPacket, prime.ScoreBoardPacket, prime.WorldBestPacket]

for dirname, dirnames, filenames in os.walk(navfolder):
    filenames.sort()
    for packet in filenames:
        f = open(os.path.join(dirname, packet), "rb")
        data = f.read(8)
        packettype = struct.unpack("<I",data[4:8])[0]
        #print "Type: %s" % hex(packettype)
        tt = "Unknown"
        for pt in packettypes:
            if pt.PacketType == packettype:
                tt = pt.__name__
                
                if pt == prime.ProfilePacket:
                    f.seek(0)
                    data = f.read()
                    p = pt()
                    p.FromBinary(data)
                    p.Print()
                
                break        
        f.close()
        print "%3s %30s (%s) - %s" %((">>>" if "_in_" in packet else "<<<"), tt, hex(packettype), packet)
