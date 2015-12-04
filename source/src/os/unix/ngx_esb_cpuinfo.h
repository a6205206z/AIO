
/*
 * Copyright (C) Cean Cheng
 * Copyright (C) Vancl, Inc.
 */

#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>   

/*esb:ngx_esb_occupy*/
struct ngx_esb_occupy        
{
	char    	 name[20];      
	unsigned int user;  
	unsigned int nice; 
	unsigned int system;
	unsigned int idle;  
};


struct ngx_esb_memory
{
	char			user[20];
	unsigned int	pid;
    double cpu_used;
	double mem_used;
};

/*esb:ngx_esb_cal_occupy*/
double ngx_esb_cal_occupy(struct ngx_esb_occupy *o, struct ngx_esb_occupy *n);
/*esb:ngx_esb_get_occupy*/
void ngx_esb_get_occupy(struct ngx_esb_occupy *o);
/*esb:ngx_esb_get_memory*/
double ngx_esb_get_memory(struct ngx_esb_memory *o);
