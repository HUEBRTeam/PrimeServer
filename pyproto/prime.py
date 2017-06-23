#!/usr/bin/env python

'''
     ____       _                  _     _ _                          
    |  _ \ _ __(_)_ __ ___   ___  | |   (_) |__  _ __ __ _ _ __ _   _ 
    | |_) | '__| | '_ ` _ \ / _ \ | |   | | '_ \| '__/ _` | '__| | | |
    |  __/| |  | | | | | | |  __/ | |___| | |_) | | | (_| | |  | |_| |
    |_|   |_|  |_|_| |_| |_|\___| |_____|_|_.__/|_|  \__,_|_|   \__, |
                                                                |___/ 

    This is a Prime Library for Prime Server
'''

import struct
import binascii
import pysodium
import time
import socket

LAST_NOUNCE = "\x00" * 24
#OFICIAL_SERVER_IP = "115.68.108.183"
OFICIAL_SERVER_IP = "127.0.0.1"
def DownloadProfile(accesscode, pk, sk):
    accesscode = accesscode.lower().replace(" ","")

    login = LoginPacket()
    login.AccessCode = accesscode

    tcp = socket.socket(socket.AF_INET, socket.SOCK_STREAM)
    dest = (OFICIAL_SERVER_IP, 60000)
    tcp.connect(dest)
    data = EncryptPacket(login.ToBinary(), pk, sk)
    tcp.send (data)
    time.sleep(0.1)
    msg = ""
    profile = ""
    while True:
        msg += tcp.recv(4)
        if len(msg) > 0:
            size = struct.unpack("<I", msg[:4])[0]
            msg = msg[4:]
            msg += tcp.recv(size-4-len(msg))
            while len(msg) < size-4:
                msg += tcp.recv((size-4) - len(msg))

            data = DecryptPacket(msg, pk, sk)
            try:
                packtype = struct.unpack("<I",data[4:8])[0]
                if packtype == ProfilePacket.PacketType:
                    print "Got Profile Packet!"
                    profile = ProfilePacket()
                    profile.FromBinary(data)
                    print "NickName: %s" %profile.Nickname
                    break
                else:
                    print "PacketType: %s (%x)" %(packtype, packtype)
            except Exception,e:
                print "Error: %s" %e
    tcp.close()
    return profile

def SendSongPacket(accesscode, pk, sk):
    pass

def DecryptPacket(packet, pk, sk):
    nounce = packet[:24]
    data = packet[24:]
    LAST_NOUNCE = nounce
    return pysodium.crypto_box_open(data, nounce, pk, sk)

def EncryptPacket(packet, pk, sk):
    nounce = LAST_NOUNCE
    data = pysodium.crypto_box(packet, nounce, pk, sk)
    return struct.pack("<I",len(data)+24+4) + nounce + data

class RequestWorldBest:
    PacketHead = 2
    PacketType = 0x1000008

    def ToBinary(self):
        return struct.pack("<2I", self.PacketHead, self.PacketType)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType = struct.unpack("<2I", binary)

class RequestRankMode:
    PacketHead = 3
    PacketType = 0x100000A

    def ToBinary(self):
        return struct.pack("<2I", self.PacketHead, self.PacketType)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType = struct.unpack("<2I", binary)


class ACKPacket:
    PacketHead = 1
    PacketType = 0x1000002
    MachineID = 0xFFFFFFF

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.MachineID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.MachineID = struct.unpack("<3I", binary)

class KeepAlivePacket():
    PacketHead = 1
    PacketType = 0x3000000
    PacketTrail = 65535

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.PacketTrail)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.PacketTrail = struct.unpack("<3I", binary)

class ByePacket:
    PacketHead = 1
    PacketType = 0x1000010
    ProfileID = 3029

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.ProfileID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.ProfileID = struct.unpack("<3I", binary)


class ProfileBusyPacket():
    PacketHead = 1
    PacketType = 0x1000005
    unk0 = 0

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.unk0)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.unk0 = struct.unpack("<3I", binary)

class RankModePacket:
    PacketHead = 3
    PacketType = 0x100000B
    rankdata = {}

    _TOTAL_SIZE = 16008

    def ToBinary(self):
        packet = struct.pack("<2I", self.PacketHeader, self.PacketType)
        for key in self.rankdata.keys():
            packet += struct.pack("<I12s12s12s",key,self.rankdata[key][0],self.rankdata[key][1],self.rankdata[key][2])    
        packet += "\x00" * (self._TOTAL_SIZE - len(packet))
        return packet

    def FromBinary(self, binary):
        self.PacketHead, self.PacketType = struct.unpack("<2I", binary[:8])
        data = binary[8:]
        i = 0
        while i < len(data):
            songid, first, second, third = struct.unpack("<I12s12s12s",data[i:i+40])
            if songid == 0x00:
                break
            i+= 40
            first = first.split("\x00")[0]
            second = second.split("\x00")[0]
            third = third.split("\x00")[0]
            self.rankdata[songid] = [first,second,third]

    def Print(self):
        print "RankModePacket"
        for key in self.rankdata.keys():
            print "0x%04x: %s -> %s -> %s" % (key,self.rankdata[key][0],self.rankdata[key][1],self.rankdata[key][2])


class MachineInfoPacket:
    '''
        Send from client when game starts
    '''
    PacketHead = 1
    PacketType = 0x1000001
    unk0 = 0
    unk1 = 0
    unk2 = 0
    MacAddress = ""         #   20 bytes
    Version = ""            #   12 bytes
    Processor = ""          #   128 bytes
    MotherBoard = ""        #   128 bytes
    GraphicsCard = ""       #   128 bytes
    HDDSerial = ""          #   32 bytes
    USBMode = ""            #   128 bytes
    Memory = 0
    unk6 = 0
    unk7 = 0
    unk8 = 0
    unk9 = 0
    unk10 = 0 
    unk11 = 0
    unk12 = 0
    unk13 = 0
    unk14 = 0
    unk15 = ""              #   104 bytes
    PacketTrail = 0xFFFFFFFF

    def ToBinary(self):
        return struct.pack("<5I20s12s128s128s128s32s128s10I104sI", self.PacketHead, self.PacketType, self.unk0, self.DongleID, self.unk2, self.MacAddress, self.Version, self.Processor, self.MotherBoard, self.GraphicsCard,  self.HDDSerial, self.USBMode, self.Memory, self.unk6, self.unk7, self.unk8, self.unk9, self.unk10, self.unk11, self.unk12, self.unk13, self.unk14, self.unk15, self.PacketTrail)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.unk0, self.DongleID, self.unk2, self.MacAddress, self.Version, self.Processor, self.MotherBoard, self.GraphicsCard,  self.HDDSerial, self.USBMode, self.Memory, self.unk6, self.unk7, self.unk8, self.unk9, self.unk10, self.unk11, self.unk12, self.unk13, self.unk14, self.unk15, self.PacketTrail = struct.unpack("<5I20s12s128s128s128s32s128s10I104sI", binary)
        self.MacAddress = self.MacAddress.split("\x00")[0] 
        self.Version = self.Version.split("\x00")[0]
        self.Processor = self.Processor.split("\x00")[0] 
        self.MotherBoard = self.MotherBoard.split("\x00")[0]
        self.GraphicsCard = self.GraphicsCard.split("\x00")[0]
        self.HDDSerial = self.HDDSerial.split("\x00")[0]
        self.USBMode = self.USBMode.split("\x00")[0]

    def Print(self):
        print "Machine Packet"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tUnknown uint32_t 0: %s (0x%x)" %(self.unk0,self.unk0)
        print "\tDongle ID: %s (0x%x)" %(self.DongleID,self.DongleID)
        print "\tUnknown uint32_t 2: %s (0x%x)" %(self.unk2,self.unk2)
        print "\tMac Address: %s" %self.MacAddress
        print "\tVersion: %s" %self.Version
        print "\tProcessor: %s" %self.Processor
        print "\tMother Board: %s" %self.MotherBoard
        print "\tGraphics Card: %s" %self.GraphicsCard
        print "\tHDD Serial: %s" %self.HDDSerial
        print "\tUSB Mode: %s" %self.USBMode
        print "\tMemory: %s" %self.Memory
        print "\tUnknown uint32_t 6: %s (0x%x)" %(self.unk6,self.unk6)
        print "\tUnknown uint32_t 7: %s (0x%x)" %(self.unk7,self.unk7)
        print "\tUnknown uint32_t 8: %s (0x%x)" %(self.unk8,self.unk8)
        print "\tUnknown uint32_t 9: %s (0x%x)" %(self.unk9,self.unk9)
        print "\tUnknown uint32_t 10: %s (0x%x)" %(self.unk10,self.unk10)
        print "\tUnknown uint32_t 11: %s (0x%x)" %(self.unk11,self.unk11)
        print "\tUnknown uint32_t 12: %s (0x%x)" %(self.unk12,self.unk12)
        print "\tUnknown uint32_t 13: %s (0x%x)" %(self.unk13,self.unk13)
        print "\tUnknown uint32_t 14: %s (0x%x)" %(self.unk14,self.unk14)
        print "\tUnknown String: %s" %(binascii.hexlify(self.unk15))
        print "\tPacket Trail: %s (0x%x)" %(self.PacketTrail, self.PacketTrail);

class MachineInfoPacket_v2:
    '''
        Send from client when game starts
    '''
    PacketHead = 1
    PacketType = 0x1000011
    MachineID = 1130
    DongleID = 0xC0FEBABE
    CountryID = 24
    MacAddress = ""         #   20 bytes
    Version = ""            #   12 bytes
    Processor = ""          #   128 bytes
    MotherBoard = ""        #   128 bytes
    GraphicsCard = ""       #   128 bytes
    HDDSerial = ""          #   32 bytes
    USBMode = ""            #   128 bytes
    Memory = 0
    ConfigMagic = 345396
    unk7 = 0xFFFFFFFF
    unk8 = 0xFFFFFF
    unk9 = 516
    unk10 = 0 
    unk11 = 262144
    unk12 = 16777216
    unk13 = 256
    unk14 = 16777217
    unk15 = 1052672
    unk16 = 0
    unk17 = 0
    unk18 = 0
    unk19 = 0
    unk20 = 0
    unk21 = 0
    unk22 = 145840
    unk23 = ""              #   76 bytes
    netaddr = ""            #   16 bytes
    def ToBinary(self):
        return struct.pack("<5I20s12s128s128s128s32s128s18I76s16s", self.PacketHead, self.PacketType, self.MachineID, self.DongleID, self.CountryID, self.MacAddress, self.Version, self.Processor, self.MotherBoard, self.GraphicsCard,  self.HDDSerial, self.USBMode, self.Memory, self.ConfigMagic, self.unk7, self.unk8, self.unk9, self.unk10, self.unk11, self.unk12, self.unk13, self.unk14, self.unk15, self.unk16, self.unk17, self.unk18, self.unk19, self.unk20, self.unk21, self.unk22, self.unk23, self.netaddr)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.MachineID, self.DongleID, self.CountryID, self.MacAddress, self.Version, self.Processor, self.MotherBoard, self.GraphicsCard,  self.HDDSerial, self.USBMode, self.Memory, self.ConfigMagic, self.unk7, self.unk8, self.unk9, self.unk10, self.unk11, self.unk12, self.unk13, self.unk14, self.unk15, self.unk16, self.unk17, self.unk18, self.unk19, self.unk20, self.unk21, self.unk22, self.unk23, self.netaddr = struct.unpack("<5I20s12s128s128s128s32s128s18I76s16s", binary)
        self.MacAddress = self.MacAddress.split("\x00")[0] 
        self.Version = self.Version.split("\x00")[0]
        self.Processor = self.Processor.split("\x00")[0] 
        self.MotherBoard = self.MotherBoard.split("\x00")[0]
        self.GraphicsCard = self.GraphicsCard.split("\x00")[0]
        self.HDDSerial = self.HDDSerial.split("\x00")[0]
        self.USBMode = self.USBMode.split("\x00")[0]

    def Print(self):
        print "Machine Packet"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tMachineID: %s (0x%x)" %(self.MachineID,self.MachineID)
        print "\tDongle ID: %s (0x%x)" %(self.DongleID,self.DongleID)
        print "\tCountry ID: %s (0x%x)" %(self.CountryID,self.CountryID)
        print "\tMac Address: %s" %self.MacAddress
        print "\tVersion: %s" %self.Version
        print "\tProcessor: %s" %self.Processor
        print "\tMother Board: %s" %self.MotherBoard
        print "\tGraphics Card: %s" %self.GraphicsCard
        print "\tHDD Serial: %s" %self.HDDSerial
        print "\tUSB Mode: %s" %self.USBMode
        print "\tMemory: %s" %self.Memory
        print "\tConfig Magic: %s (0x%x)" %(self.ConfigMagic,self.ConfigMagic)
        print "\tUnknown uint32_t 7: %s (0x%x)" %(self.unk7,self.unk7)
        print "\tUnknown uint32_t 8: %s (0x%x)" %(self.unk8,self.unk8)
        print "\tUnknown uint32_t 9: %s (0x%x)" %(self.unk9,self.unk9)
        print "\tUnknown uint32_t 10: %s (0x%x)" %(self.unk10,self.unk10)
        print "\tUnknown uint32_t 11: %s (0x%x)" %(self.unk11,self.unk11)
        print "\tUnknown uint32_t 12: %s (0x%x)" %(self.unk12,self.unk12)
        print "\tUnknown uint32_t 13: %s (0x%x)" %(self.unk13,self.unk13)
        print "\tUnknown uint32_t 14: %s (0x%x)" %(self.unk14,self.unk14)
        print "\tUnknown uint32_t 15: %s (0x%x)" %(self.unk14,self.unk15)
        print "\tUnknown uint32_t 16: %s (0x%x)" %(self.unk14,self.unk16)
        print "\tUnknown uint32_t 17: %s (0x%x)" %(self.unk14,self.unk17)
        print "\tUnknown uint32_t 18: %s (0x%x)" %(self.unk14,self.unk18)
        print "\tUnknown uint32_t 19: %s (0x%x)" %(self.unk14,self.unk19)
        print "\tUnknown uint32_t 20: %s (0x%x)" %(self.unk14,self.unk20)
        print "\tUnknown uint32_t 21: %s (0x%x)" %(self.unk14,self.unk21)
        print "\tUnknown uint32_t 22: %s (0x%x)" %(self.unk14,self.unk22)
        print "\tUnknown String: %s" %(binascii.hexlify(self.unk23))
        print "\tNet Address: %s" %(self.netaddr);

class ScoreBoardPacket():
    '''
        Send from client when finishes a Song 
    '''
    PacketHead = 0x0000001
    PacketType = 0x100000E

    SongID = 0
    ChartLevel = 0
    Type = 0
    Flag = 0 
    Score = 0
    RealScore0 = 0

    unk0 = ""                   #   16 bytes

    RealScore1 = 0
    Grade = 0 
    Kcal = 0

    Perfect = 0
    Great = 0
    Good = 0
    Bad = 0 
    Miss = 0
    MaxCombo = 0
    EXP = 0
    PP = 0

    RunningStep = 0
    unk2 = 0
    unk3 = 0
    unk4 = 0
    unk5 = 0
    RushSpeed = 0

    GameVersion = ""            #   12 Bytes

    MachineID   = 0xFFFFFFFF
    ProfileID   = 0x00

    def ToBinary(self):
        return struct.pack("<3IH2B2I16s2If6I4H3If12s2I", self.PacketHead, self.PacketType, self.SongID, self.ChartLevel, self.Type, self.Flag, self.Score, self.RealScore0, self.unk0, self.RealScore1, self.Grade, self.Kcal, self.Perfect, self.Great, self.Good, self.Bad, self.Miss, self.MaxCombo, self.EXP, self.PP, self.RunningStep, self.unk2, self.unk3, self.unk4, self.unk5, self.RushSpeed, self.GameVersion, self.MachineID, self.ProfileID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.SongID, self.ChartLevel, self.Type, self.Flag, self.Score, self.RealScore0, self.unk0, self.RealScore1, self.Grade, self.Kcal, self.Perfect, self.Great, self.Good, self.Bad, self.Miss, self.MaxCombo, self.EXP, self.PP, self.RunningStep, self.unk2, self.unk3, self.unk4, self.unk5, self.RushSpeed, self.GameVersion, self.MachineID, self.ProfileID = struct.unpack("<3IH2B2I16s2If6I4h3If12s2I", binary)

    def Print(self):
        print "Score Board Packet"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tSongID: %s (0x%x)" %(self.SongID,self.SongID)
        print "\tChart Level: %s" %self.ChartLevel
        print "\tType: %s (0x%x)" %(self.Type,self.Type)
        print "\tFlag: %s (0x%x)" %(self.Flag, self.Flag)
        print "\tScore: %s" %self.Score
        print "\tRealScore0: %s" %self.RealScore0
        print "\tUnknown String 0: %s" %binascii.hexlify(self.unk0)
        print "\tRealScore1: %s" %self.RealScore1
        print "\t\tGrade: %s (0x%x)" %(self.Grade, self.Grade)
        print "\t\tKcal: %s" %self.Kcal
        print "\t\tPerfect: %s" %self.Perfect
        print "\t\tGreat: %s" %self.Great
        print "\t\tGood: %s" %self.Good
        print "\t\tBad: %s" %self.Bad
        print "\t\tMiss: %s" %self.Miss
        print "\t\tMaxCombo: %s" %self.MaxCombo
        print "\t\tEXP: %s" %self.EXP
        print "\t\tPP: %s" %self.PP
        print ""
        print "\tRunning Step: %s" %(self.RunningStep)
        print "\tUnknown uint16_t 2: %s (0x%x)" %(self.unk2,self.unk2)
        print "\tUnknown uint32_t 3: %s (0x%x)" %(self.unk3,self.unk3)
        print "\tUnknown uint32_t 4: %s (0x%x)" %(self.unk4,self.unk4)
        print "\tUnknown uint32_t 5: %s (0x%x)" %(self.unk5,self.unk5)
        print "\tRushSpeed: %s" %(self.RushSpeed)
        print "\tMachineID: %s (0x%x)" %(self.MachineID, self.MachineID);
        print "\tProfileID: %s (0x%x)" %(self.ProfileID, self.ProfileID);


class LoginPacket():
    '''
        Packet that Client sends to Server to Receive Login Stuff
    '''
    PacketHead = 1
    PacketType = 0x1000003
    PlayerID = 0
    MachineID = 1130
    AccessCode = ""             #   32 bytes
    PacketTrail = 0x0

    def ToBinary(self):
        return struct.pack("<4I32sI", self.PacketHead, self.PacketType, self.PlayerID, self.MachineID, self.AccessCode, self.PacketTrail)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.PlayerID, self.MachineID, self.AccessCode, self.PacketTrail = struct.unpack("<4I32sI", binary)

    def Print(self):
        print "LoginPacket"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tPlayerID: %s (0x%x)" %(self.PlayerID, self.PlayerID)
        print "\tMachineID: %s (0x%x)" %(self.MachineID, self.MachineID)
        print "\tAccessCode: %s" %self.AccessCode
        print "\tPacketTrail: %s (0x%x)" %(self.PacketTrail,self.PacketTrail)

class uScore():
    '''
        Song Score Entry on Profile Data
    '''
    flags = {
        "UNLOCK" : 0x4000,
        "NORMAL" : 0x8000,
        "MISSION" : 0x100,
        "DOUBLE" :  0x80
    }

    SongID = 0
    ChartLevel = 0
    unk0 = 0
    GameDataFlag = 0
    Score = 0
    RealScore = 0
    unk2 = 0

    def ToBinary(self):
        return struct.pack("<I2BH3I", self.SongID, self.ChartLevel, self.unk0, self.GameDataFlag, self.Score, self.RealScore, self.unk2)

    def FromBinary(self,binary):
        self.SongID, self.ChartLevel, self.unk0, self.GameDataFlag, self.Score, self.RealScore, self.unk2 = struct.unpack("<I2BH3I", binary)

    def Print(self):
        FLAGS = "|"
        for flag in self.flags:
            if self.GameDataFlag & self.flags[flag] > 0:
                FLAGS += flag + "|"
        FLAGS = FLAGS if len(FLAGS) != 1 else ""
        print "\t\tSongID: %X \tat level %s%s\t Score: %s - RealScore: %s - [Unk0: %s(0x%x), GameDataFlag: %s(0x%x) %s, Unk2: %s(0x%x)]" %(self.SongID, "D" if (self.GameDataFlag & self.flags["DOUBLE"] > 0) else "S", self.ChartLevel, self.Score, self.RealScore, self.unk0, self.unk0, self.GameDataFlag, self.GameDataFlag, FLAGS, self.unk2, self.unk2)

class EnterProfilePacket():
    '''
        Packet Received when Someone enters ingame
    '''
    PacketHead = 2
    PacketType = 0x100000F
    PlayerID = 0
    MachineID = 1130
    ProfileID = 0
    def ToBinary(self):
        return struct.pack("<5I", self.PacketHead, self.PacketType, self.PlayerID, self.MachineID, self.ProfileID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.PlayerID, self.MachineID, self.ProfileID = struct.unpack("<5I", binary)

    def Print(self):
        print "EnterProfilePacket: "
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tPlayerID: %s (0x%x)" %(self.PlayerID, self.PlayerID)
        print "\tMachineID: %s (0x%x)" %(self.MachineID, self.MachineID)
        print "\tProfileID: %s (0x%x)" %(self.ProfileID, self.ProfileID)

class ProfilePacket():
    '''
        Packet send by server to client after LoginPacket
    '''
    __USCORE_PACK_LENGTH = 20  #  Used for packing/unpacking uScore packets 

    PacketHead = 1
    PacketType = 0x1000004
    PlayerID  = 0            #   32 bit
    AccessCode  = ""         #   36 bytes
    Nickname = ""            #   12 Bytes
    ProfileID = 0            #   32 bit
    CountryID = 0            #   8 bit
    Avatar = 0               #   8 bit
    Level = 0                #   8 bit
    unk2 = 0                 #   8 bit
    EXP = 0                  #   64 bit
    PP = 0                   #   64 bit
    RankSingle = 0           #   64 bit
    RankDouble = 0           #   64 bit
    RunningStep = 0          #   64 bit
    PlayCount = 0            #   32 bit
    Kcal = 0                 #   Float
    Modifiers = 0            #   64 bit
    unk3 = 0                 #   32 bit
    RushSpeed = 0            #   Float
    unk4 = 0                 #   32 bit
    Scores  = []             #   4384 uScore Items 

    def ToBinary(self):
        data0 = struct.pack("<3I36s12sI4B5QIfQIfI", self.PacketHead, self.PacketType, self.PlayerID, self.AccessCode, self.Nickname, self.ProfileID, self.CountryID, self.Avatar, self.Level, self.unk2, self.EXP, self.PP, self.RankSingle, self.RankDouble, self.RunningStep, self.PlayCount, self.Kcal, self.Modifiers, self.unk3, self.RushSpeed, self.unk4)
        data1 = ""
        for i in self.Scores:
            data1 += i.ToBinary()       

        data1 += "\x00" * (4384 - len(self.Scores)) * self.__USCORE_PACK_LENGTH
        return data0 + data1

    def FromBinary(self,binary):
        data0 = binary[:0x88]
        data1 = binary[0x88:]
        self.PacketHead, self.PacketType, self.PlayerID, self.AccessCode, self.Nickname, self.ProfileID, self.CountryID, self.Avatar, self.Level, self.unk2, self.EXP, self.PP, self.RankSingle, self.RankDouble, self.RunningStep, self.PlayCount, self.Kcal, self.Modifiers, self.unk3, self.RushSpeed, self.unk4 = struct.unpack("<3I36s12sI4B5QIfQIfI", data0)
        i = 0
        c = 0
        self.Nickname = self.Nickname.split("\x00")[0]
        while i < len(data1):
            scoretmp = uScore()
            scoretmp.FromBinary(data1[i:i+self.__USCORE_PACK_LENGTH])
            if scoretmp.SongID != 0:
                self.Scores.append(scoretmp)
                c += 1
            i += self.__USCORE_PACK_LENGTH
        #print "Loaded %s uScores" %c

    def Print(self):
        print "Profile Packet"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tPlayerID: %s (0x%x)" %(self.PlayerID, self.PlayerID)
        print "\tAcessCode: %s" %self.AccessCode    
        print "\tNickName: %s" %self.Nickname   
        print "\tProfileID: %s (0x%x)" %(self.ProfileID, self.ProfileID)
        print "\tCountryID: %s (0x%x)" %(self.CountryID, self.CountryID)
        print "\tAvatar: %s (0x%x)" %(self.Avatar, self.Avatar)
        print "\tLevel: %s " %(self.Level) 
        print "\tUnknown uint32_t 2: %s (0x%x)" %(self.unk2, self.unk2)
        print "\tExperience: %s" %self.EXP
        print "\tPP: %s" %self.PP
        print "\tRank Single: %s (0x%x)" %(self.RankSingle, self.RankSingle)
        print "\tRank Double: %s (0x%x)" %(self.RankDouble, self.RankDouble)
        print "\tRunning Steps: %s" % self.RunningStep
        print "\tPlay Count: %s" % self.PlayCount
        print "\tKcal: %s" % self.Kcal
        print "\tModifiers: %s (0x%x)" % (self.Modifiers, self.Modifiers)
        print "\tUnknown uint32_t 3: %s (0x%s)" % (self.unk3, self.unk3)
        print "\tRush Speed: %s" % self.RushSpeed
        print "\tUnknown uint32_t 3: %s (0x%s)" % (self.unk3, self.unk4)
        print "\tScore Board: "
        for i in self.Scores:
            i.Print()

class RequestLevelUpInfoPacket():
    '''
        Packet send by client after all musics has been played.
    '''

    PacketHead = 1
    PacketType = 0x100000C
    ProfileID = 3029

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.ProfileID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.ProfileID = struct.unpack("<3I", binary)

    def Print(self):
        print "RequestLevelUpInfoPacket: "
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tProfileID: %s (0x%x)" %(self.ProfileID, self.ProfileID)

class LevelUpInfoPacket():
    '''
        Packet send by server after receiving a RequestLevelUpInfoPacket
    '''
    PacketHead = 1
    PacketType = 0x100000D
    ProfileID = 3029   
    Level = 0

    def ToBinary(self):
        return struct.pack("<4I", self.PacketHead, self.PacketType, self.ProfileID, self.Level)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.ProfileID, self.Level = struct.unpack("<4I", binary)

    def Print(self):
        print "EnterProfilePacket: "
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tProfileID: %s (0x%x)" %(self.ProfileID, self.ProfileID)
        print "\tLevel: %s" %(self.Level)


class GameOverPacket():
    '''
        Packet send by client on GameOver
    '''

    PacketHead = 1
    PacketType = 0x1000001
    ProfileID = 3029

    def ToBinary(self):
        return struct.pack("<3I", self.PacketHead, self.PacketType, self.ProfileID)

    def FromBinary(self,binary):
        self.PacketHead, self.PacketType, self.ProfileID = struct.unpack("<3I", binary)

    def Print(self):
        print "EnterProfilePacket: "
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tProfile ID: %s (0x%x)" %(self.ProfileID, self.ProfileID)

class WorldBestScore():
    '''
        WorldBest Entry
    '''
    SongID = 0
    ChartLevel = 0
    ChartMode = 0
    Score = 0
    unk0 = 0
    unk1 = 0
    Nickname = ""

    def ToBinary(self):
        return struct.pack("<I2H3I12s", self.SongID, self.ChartLevel, self.ChartMode, self.Score, self.unk0, self.unk1, self.Nickname)

    def FromBinary(self,binary):
        self.SongID, self.ChartLevel, self.ChartMode, self.Score, self.unk0, self.unk1, self.Nickname = struct.unpack("<I2H3I12s", binary)
        self.Nickname = self.Nickname.split("\x00")[0]

    def Print(self):
        print "\tNickname: %s - SongID: %x (lvl %s mode 0x%x) - Score: %s - Unk(%s[0x%x],%s[0x%x])" %(self.Nickname, self.SongID, self.ChartLevel, self.ChartMode, self.Score, self.unk0, self.unk0, self.unk1, self.unk1)

class WorldBestPacket():
    '''
        World Best Packet
    '''
    __WORLDBESTSCORE_PACK_LENGTH    =   32

    PacketHead = 2
    PacketType = 0x1000009
    Scores = []         #   4096 

    def ToBinary(self):
        data0 = struct.pack("<2I", self.PacketHead, self.PacketType)
        data1 = ""
        for i in self.Scores:
            data1 += i.ToBinary()       

        data1 += "\x00" * (4096 - len(self.Scores)) * self.__USCORE_PACK_LENGTH
        return data0 + data1 

    def FromBinary(self,binary):
        data0 = binary[:0x8]
        data1 = binary[0x8:]
        self.PacketHead, self.PacketType = struct.unpack("<2I", data0)
        i = 0
        c = 0
        while i < len(data1):
            scoretmp = WorldBestScore()
            scoretmp.FromBinary(data1[i:i+self.__WORLDBESTSCORE_PACK_LENGTH ])
            if scoretmp.SongID != 0:
                self.Scores.append(scoretmp)
                c += 1
            i += self.__WORLDBESTSCORE_PACK_LENGTH 
        print "Loaded %s World Best Scores" %c

    def Print(self):
        print "World Best Score Packet"
        print "\tPacketHead: %s (0x%x)" %(self.PacketHead,self.PacketHead)
        print "\tPacketType: %s (0x%x)" %(self.PacketType,self.PacketType)
        print "\tScores: "
        for i in self.Scores:
            i.Print()
