#include"devmem.h"


int tempfd;
void *map_base,*virt_addr;

void Openfile(long  target){
    int fd;
    
    if((tempfd = open("/dev/mem", O_RDWR | O_SYNC)) == -1) FATAL; 
    fd = fcntl(tempfd, F_DUPFD, 0);
    if(fd<0){
	    FATAL;
    } 
    
    /* Map one page */
    map_base = mmap(0, MAP_SIZE, PROT_READ | PROT_WRITE, MAP_SHARED, fd, target & ~MAP_MASK);
    if(map_base == (void *) -1) FATAL;
    
    virt_addr = map_base + (target & MAP_MASK);
}

void Closefile(){
     if(munmap(map_base, MAP_SIZE) == -1) FATAL;
     close(tempfd);
}

void Writebit(int offset,int bitsize ,char value){
    unsigned long read_result, writeval;
       
//    if((fd = open("/dev/mem", O_RDWR | O_SYNC)) == -1) FATAL;
//    printf("/dev/mem opened.\n");
//    fflush(stdout);

    
//    printf("Memory mapped at address %p.\n", map_base);
//    fflush(stdout);
    
    
    
    read_result = *((unsigned long *) (virt_addr+offset));
    
//    printf("Value at address 0x%X (%p): 0x%X\n", target, virt_addr, read_result);
//    fflush(stdout);
    

//    printf("bitsize:%d;value:%d\n",bitsize,value);
    if(value==0){
       read_result&=~(1<<bitsize); 
    }else{
        read_result|=1<<bitsize;
    }

    writeval=read_result;
    *((unsigned long *) (virt_addr+offset))=writeval;
//   read_result=*((unsigned long *) virt_addr);
    
//    printf("Written 0x%X; readback 0x%X\n", writeval, read_result);
//    fflush(stdout);
    
   
}
