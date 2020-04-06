#include"devmem.h"


int tempfd;

void Openfile(){
    if((tempfd = open("/dev/mem", O_RDWR | O_SYNC)) == -1) FATAL; 
    
}

void Closefile(){
     close(tempfd);
}

void Writebit(long  target,int bitsize ,char value){
 
    void *map_base, *virt_addr;
    unsigned long read_result, writeval;
    int fd;

    fd = fcntl(tempfd, F_DUPFD, 0);
    if(fd<0){
	FATAL;
    }    
//    if((fd = open("/dev/mem", O_RDWR | O_SYNC)) == -1) FATAL;
//    printf("/dev/mem opened.\n");
//    fflush(stdout);

    /* Map one page */
    map_base = mmap(0, MAP_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, target & ~MAP_MASK);
    if(map_base == (void *) -1) FATAL;
//    printf("Memory mapped at address %p.\n", map_base);
//    fflush(stdout);
    
    virt_addr = map_base + (target & MAP_MASK);
    
    read_result = *((unsigned long *) virt_addr);
    
//    printf("Value at address 0x%X (%p): 0x%X\n", target, virt_addr, read_result);
//    fflush(stdout);
    

//    printf("bitsize:%d;value:%d\n",bitsize,value);
    if(value==0){
       read_result&=~(1<<bitsize); 
    }else{
        read_result|=1<<bitsize;
    }

    writeval=read_result;
    *((unsigned long *) virt_addr)=writeval;
    read_result=*((unsigned long *) virt_addr);
    
//    printf("Written 0x%X; readback 0x%X\n", writeval, read_result);
//    fflush(stdout);
    
    if(munmap(map_base, MAP_SIZE) == -1) FATAL;
    close(fd);
}
