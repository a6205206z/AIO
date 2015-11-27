
/*
 * Copyright (C) Cean Cheng
 * Copyright (C) Vancl, Inc.
 */

#include <ngx_config.h>
#include <ngx_core.h>

double
ngx_esb_cal_occupy(struct ngx_esb_occupy *o, struct ngx_esb_occupy *n)
{
	double							g_cpu_used;
	double							od, nd;   
	double							id, sd;  
	//double							scale;   
	od = (double) (o->user + o->nice + o->system +o->idle);
	nd = (double) (n->user + n->nice + n->system +n->idle);
	//scale = 100.0 / (float)(nd-od);      
	id = (double) (n->user - o->user);  
	sd = (double) (n->system - o->system);
	g_cpu_used = ((sd+id)*100.0)/(nd-od);
	return g_cpu_used;
}
void
ngx_esb_get_occupy(struct ngx_esb_occupy *o)
{
	FILE *fd;               
	char buff[1024];  

	fd = fopen ("/proc/stat", "r");
	fgets (buff, sizeof(buff), fd); 
	sscanf (buff, "%s %u %u %u %u", o->name, &o->user, &o->nice,&o->system, &o->idle);
	
	fclose(fd);   
}
double 
ngx_esb_get_memory(struct ngx_esb_memory *m)
{
	FILE			*fd;               
	char			buff[1024];  
	double			mem;

	mem = 0;
	fd = popen ("ps aux|grep nginx", "r");

	for( ; ; )
	{
		if( fgets (buff, sizeof(buff), fd) == NULL)
		{
			break;
		}
		sscanf (buff, "%s %u %lf %lf", m->user, &m->pid, &m->cpu_used, &m->mem_used);
		mem += m->mem_used;
	}

	fclose(fd);   
	return mem;
}
