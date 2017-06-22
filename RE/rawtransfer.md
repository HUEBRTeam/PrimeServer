     ____      ___        __
    |  _ \    / \ \      / /
    | |_) |  / _ \ \ /\ / / 
    |  _ <  / ___ \ V  V /  
    |_| \_\/_/   \_\_/\_/   
                            

This document is here to explain how packets are made for Pump It Up: Prime. On this folder we have **sprivk** as `Private Key` and **spublick** as `Public Key` for traffic.

The Sodium Encrypter/Decrypter uses both Public Key and Private Key together with a Nounce. The keys has fixed length to 32 bytes, the nounce has 24 bytes and the final packet is also composed by a MAC of 16 bytes. In the start of the packet is the size of it. So for the packet size we will have `message length + 24 (nouce) + 16 (mac) + 4 (packet size) bytes`.

So here we have a sample packet:

    60 00 00 00 C6 A0 43 58 46 05 61 51 65 99 8E 55 72 A9 A4 EB 89 52 4D 30 04 0E D4 EE 63 A3 55 1C A7 BF 90 36 C5 D9 1C 89 C0 3F E2 68 B4 21 D3 BD 3D 50 C1 20 DA 68 DA 58 3E A6 7E 40 C2 94 89 B0 35 07 BC 25 5D AA 82 A2 1F 26 6B 62 B1 31 F9 86 A8 C7 80 87 F8 B1 CA DD A4 6C 36 89 48 4C 78 B4
    
It has 96 bytes. So the first 4 bytes (`uint32_t`) says 96 (0x60). Following we have the 24 byte nounce:

    C6 A0 43 58 46 05 61 51 65 99 8E 55 72 A9 A4 EB 89 52 4D 30 04 0E D4 EE
    
Then we have the encrypted message + MAC:

    63 A3 55 1C A7 BF 90 36 C5 D9 1C 89 C0 3F E2 68 B4 21 D3 BD 3D 50 C1 20 DA 68 DA 58 3E A6 7E 40 C2 94 89 B0 35 07 BC 25 5D AA 82 A2 1F 26 6B 62 B1 31 F9 86 A8 C7 80 87 F8 B1 CA DD A4 6C 36 89 48 4C 78 B4
    
The decrypted packet (as you can confirm with packet_fuck) is : 
    
    01 00 00 00 03 00 00 01 00 00 00 00 F7 00 00 00 39 66 32 38 35 37 65 38 33 66 32 66 65 63 37 64 30 62 62 34 66 38 38 66 32 32 66 31 34 37 35 37 00 00 00 00
    
You can see the dumps in folder `packets/rawsample`
    
