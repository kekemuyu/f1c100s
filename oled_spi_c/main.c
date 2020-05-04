#include"oled.h"
#include<stdio.h>
#include<stdlib.h>
#include<time.h>

void main(){
    OLED_Init();
    time_t *timep = malloc(sizeof(*timep));
    time(timep);
    char *s = ctime(timep);
    printf("%s",s);
	int flag = 0;
	for(int n=0;n<1000;n++) {
		if (flag==0) {
			flag = 1;
			OLED_Clear();
		} else {
			flag = 0;
			OLED_Unclear();
		}

	}
	time(timep);
    char *s2 = ctime(timep);
    printf("%s",s2);
    OLED_ShowString(0,3,"finish!");
}
