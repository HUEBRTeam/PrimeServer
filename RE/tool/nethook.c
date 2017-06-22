#define __USE_GNU
#include <stdio.h>
#include <dlfcn.h>
#include <string.h>
#include <sodium/crypto_box.h>
#include <sys/socket.h>
#include <netinet/in.h>
#include <arpa/inet.h>
#include <curl/curl.h>

static int (*real_crypto_box_easy)(unsigned char *c, const unsigned char *m,unsigned long long mlen, const unsigned char *n,const unsigned char *pk, const unsigned char *sk)=NULL;
static ssize_t (*real_send)(int socket, const void *buffer, size_t length, int flags) = NULL;
static int (*real__vsprintf_chk)(char * s, int flag, size_t slen, const char * format, va_list args) = NULL;
static int (*real_crypto_box_open_easy)(unsigned char *m, const unsigned char *c,unsigned long long clen, const unsigned char *n,const unsigned char *pk, const unsigned char *sk) = NULL;
 
static ssize_t (*real_recv)(int sockfd, void *buf, size_t len, int flags) = NULL;

static void __mtrace_init(void)
{
    void *handle = dlopen("libsodium.so.13",RTLD_LAZY);
    real_crypto_box_easy = dlsym(handle, "crypto_box_easy");
    real_crypto_box_open_easy = dlsym(handle, "crypto_box_open_easy");
    real_send = dlsym(RTLD_NEXT, "send");
    real_recv = dlsym(RTLD_NEXT, "recv");
    real__vsprintf_chk = dlsym(RTLD_NEXT, "__vsprintf_chk");
}

static int pkt_count = 0;

void WritePacket(const unsigned char *pkt, unsigned long long len, int dir)    {
    char filename[60];
    pkt_count++;
    if(!dir)
        sprintf (filename, "packets/sodium_%d_out_pkt.bin", pkt_count);
    else
        sprintf (filename, "packets/sodium_%d_in_pkt.bin", pkt_count);
    printf("Writting packet #%d in %s\n",pkt_count,filename);
    FILE *f = fopen(filename,"wb");
    fwrite ( pkt, len, 1, f );
    fclose(f);
}

void WritePacketRaw(const unsigned char *pkt, unsigned long long len, int dir, const char *ip, int port)    {
    char filename[60];
    pkt_count++;
    if(!dir)
        sprintf (filename, "packets/raw_%d_out_%s_%d.bin", pkt_count,ip,port);
    else
        sprintf (filename, "packets/raw_%d_in_%s_%d.bin", pkt_count,ip,port);
    printf("Writting raw packet #%d in %s\n",pkt_count,filename);
    FILE *f = fopen(filename,"wb");
    fwrite ( pkt, len, 1, f );
    fclose(f);
}

int crypto_box_easy(unsigned char *c, const unsigned char *m,
                    unsigned long long mlen, const unsigned char *n,
                    const unsigned char *pk, const unsigned char *sk)  {
    WritePacket(m, mlen, 0);
    
    if(real_crypto_box_easy==NULL)
        __mtrace_init();
	return real_crypto_box_easy(c,m,mlen,n,pk,sk);
}

int crypto_box_open_easy(unsigned char *m, const unsigned char *c,
                         unsigned long long clen, const unsigned char *n,
                         const unsigned char *pk, const unsigned char *sk)  {
    int h = real_crypto_box_open_easy(m,c,clen,n,pk,sk);
    WritePacket(m, clen-crypto_box_MACBYTES, 1);
    
    if(real_crypto_box_open_easy==NULL)
        __mtrace_init();
        
	return h;                         
}


ssize_t send(int socket, const void *buffer, size_t length, int flags)  {
    struct sockaddr_storage addr;
    char ipstr[20];
    int port;
    socklen_t len;
    
    if(real_send==NULL)
        __mtrace_init();
       
    getpeername(socket, (struct sockaddr*)&addr, &len);    
    struct sockaddr_in *s = (struct sockaddr_in *)&addr;
    if(s != NULL && s->sin_family == AF_INET)    {
        port = ntohs(s->sin_port);
        inet_ntop(AF_INET, &s->sin_addr, ipstr, sizeof ipstr);

        if(port != 0 && port != 53)   {
            WritePacketRaw(buffer, length, 0, ipstr, port);
        }
    }
    return real_send(socket,buffer,length,flags);
}

ssize_t recv(int socketfd, void *buffer, size_t length, int flags)  {
    struct sockaddr_storage addr;
    struct sockaddr_in pa;
    char ipstr[20];
    int port;
    socklen_t len;
    
    if(real_recv==NULL)
        __mtrace_init();
        
    getpeername(socketfd, (struct sockaddr*)&addr, &len);    
    struct sockaddr_in *s = (struct sockaddr_in *)&addr;
    
    ssize_t n = real_recv(socketfd, buffer, length, flags);
    
    if(s != NULL && s->sin_family == AF_INET)    {
        port = ntohs(s->sin_port);
        inet_ntop(AF_INET, &s->sin_addr, ipstr, sizeof ipstr);

        if(port != 0 && port != 53)   {
            WritePacketRaw(buffer, length, 1, ipstr, port);
        }  
    }
    return n;
}

//Debug Print
int __vsprintf_chk(char * s, int flag, size_t slen, const char * format, va_list args)  {
    if(real__vsprintf_chk==NULL)
        __mtrace_init(); 
    int ret = real__vsprintf_chk(s,flag,slen,format,args);
    for(int i=0;i<slen;i++) {
        if(s[i] == '\n')    {
            printf("%s\n",s);
            break;
        }
    }
    return ret;
       
}
