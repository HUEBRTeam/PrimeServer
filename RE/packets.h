typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x1000002 // 0x3000000
   uint32_t PacketTrail;        //   0x08 0xFFFFFFF // 0xFFFF
} ACKPacket;

typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x1000001
   uint32_t unk0;               //   0x08
   uint32_t DongleID;           //   0x0C
   uint32_t unk2;               //   0x10
   char MacAddress[20];         //   0x14
   char Version[12];            //   0x28
   char Processor[128];         //   0x34
   char MotherBoard[128];       //   0xb4
   char GraphicsCard[128];      //   0x134
   char HDDSerial[32];          //   0x1b4
   char USBMode[128];           //   0x1d4 //  Mode 1.0 / Mode 1.1 / Mode 2.0
   uint32_t Memory;             //   0x254
   uint32_t unk6;               //   0x258
   uint32_t unk7;               //   0x25c
   uint32_t unk8;               //   0x260
   uint32_t unk9;               //   0x264
   uint32_t unk10;              //   0x268
   uint32_t unk11;              //   0x26c
   uint32_t unk11;              //   0x270
   uint32_t unk12;              //   0x274
   uint32_t unk13;              //   0x278
   uint32_t unk14;              //   0x27c
   char unk15[104];             //   0x280
   
} MachineInfoPacket;
   
typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x100000E
   
   uint32_t SongID;             //   0x08
   uint16_t ChartLevel;         //   0x0C
   uint8_t Type;                //   0x0E
   uint8_t Flag;                //   0x0F
   uint32_t Score;              //   0x10
   uint32_t RealScore0;         //   0x14
   
   char     unk0[16];           //   0x18
   
   uint32_t RealScore1;         //   0x28   //   Same as SongScore0, dafuq?
   uint32_t Grade;              //   0x2C
   float    Kcal;               //   0x30
   
   uint32_t Perfect;            //   0x34
   uint32_t Great;              //   0x38
   uint32_t Good;               //   0x3c
   uint32_t Bad;                //   0x40
   uint32_t Miss;               //   0x44
   uint32_t MaxCombo;           //   0x48
   uint16_t EXP;                //   0x4c
   uint16_t PP;                 //   0x4e
   
   uint16_t unk1;               //   0x50
   uint16_t unk2;               //   0x52
   uint32_t unk3;               //   0x54
   uint32_t unk4;               //   0x58
   uint32_t unk5;               //   0x5c
   uint32_t unk6;               //   0x60
   
   char     GameVersion[12];    //   0x64
   
   uint32_t trailing0;          //   0x70   //   0xFFFFFF
   uint32_t trailing1;          //   0x74   //  0xB21
}   ScoreBoardPacket;


typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x1000003
   uint32_t unk0;               //   0x08
   uint32_t unk1;               //   0x0C
   char     AccessCode[32];     //   0x10  // Hex String
   uint32_t unk2;
} LoginPacket;

typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x1000004
   uint32_t Unk0;               //   0x08
   char     AccessCode[32];     //   0x0C
   uint32_t Unk1;               //   0x10
   char     Nickname[12];       //   0x30
   uint32_t Unk2;               //   0x3C
   uint16_t Unk3;               //   0x40
   uint16_t Level;              //   0x42
   uint32_t EXP;                //   0x44
   uint32_t Unk4;               //   0x48
   uint32_t PP;                 //   0x4C
   char     Unk5[20];           //   0x50
   uint32_t RunningStep;        //   0x64
   char     Unk6[32];           //   0x68
   uScore   Scores[4384];       //   0x88
} ProfilePacket;

typedef struct {
   uint32_t SongID;             //   0x00
   uint8_t  ChartLevel;         //   0x04
   uint8_t  Unk0;               //   0x05
   unit16_t Unk1;               //   0x06
   uint32_t Score;              //   0x08
   uint32_t RealScore;          //   0x0C   //  Maybe
   uint32_t Unk2;               //   0x10
} uScore;

typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x100000C
   uint32_t unk0;               //   0x08 0xBD5
} RequestLevelUpInfoPacket;

typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x100000D
   uint32_t unk0;               //   0x08 0xBD5
   uint32_t Level;              //   0x0C 
} LevelUpInfoPacket;

typedef struct {
   uint32_t PacketHead;         //   0x00 0x0000001
   uint32_t PacketType;         //   0x04 0x1000001
   uint32_t unk0;               //   0x08 0xBD5
} GameOverPacket;

typedef struct {
    uint32_t PacketHead;        //   0x00 0x00000002
    uint32_t PacketType;        //   0x04 0x10000009 
    uint32_t unk0;              //   0x08 5056
    uint32_t unk1;              //   0x0C 0x0000000F
    uint32_t unk2;              //   0x10 674200 
    uint32_t unk3;              //   0x14 0x00000000
    uint32_t unk4;              //   0x18 0x00000000
    WorldBestScore[4095];
    uint32_t unk5;              //   0x?? 0x00000000
    uint32_t unk6;              //   0x?? 0x00000000
    uint32_t PacketTrail;       //   0x?? 0x00000000
} WorldBest;

typedef struct {
    uint32_t SongID;
    uint16_t ChartLevel;
    uint16_t ChartMode;
    uint32_t Score;
    uint32_t unk0;
    unit32_t unk1;
    char Nickname[12];
} WorldBestScore;