/*
 * Copyright (C) Cean Cheng
 * Copyright (C) Vancl, Inc.
 */

#include <ngx_config.h>
#include <ngx_core.h>




/*esb:search node by key*/
ngx_esb_kv_node_t *ngx_find_ngx_esb_kv_node(ngx_esb_kv_node_t *head,ngx_str_t key)
{
	ngx_esb_kv_node_t				*p;
	p = head;

	if( key.data != NULL )
	{
		while(p != NULL)
		{
		
			if( p->key.data!=NULL
				&& ngx_strcmp( p->key.data, key.data ) == 0 )
			{
				return p;
			}
			p = p->next;
		}
	}

	return NULL;
}

/*esb:append at end*/
ngx_esb_kv_node_t *ngx_append_ngx_esb_kv_node(ngx_esb_kv_node_t *head, ngx_esb_kv_node_t *n)
{
	ngx_esb_kv_node_t				*p;
	p = head;

	if(n == NULL)
		return head;

	if(n->next||head == NULL)
		return NULL;

	while(p)
	{
		if(p->next)
		{
			p = p->next;
		}
		else
		{
			p->next = n;
			break;
		}
	}
   
	return head;
}

/*esb:create*/
ngx_esb_kv_node_t *ngx_create_ngx_esb_kv_node(ngx_log_t *log,ngx_int_t data,ngx_str_t key)
{
	 ngx_esb_kv_node_t        *node;
	 ngx_shm_t                shm;

	 if(key.data != NULL)
	 {
		shm.size = sizeof(ngx_esb_kv_node_t);
		shm.name.len = key.len;
		shm.name.data = (u_char *) key.data;
		shm.log = log;

		if (ngx_shm_alloc(&shm) != NGX_OK) {
		ngx_log_error(NGX_LOG_EMERG, log, 0,
			"create ngx kv %s node failed",(u_char *)key.data);
		return NULL;
		}

		node = ( ngx_esb_kv_node_t * )shm.addr;
		node->data = data;
		node->key = key;
		node->invoke_max_time = 0;
		node->invoke_avg_time = 0;
		node->esb_proc_max_time = 0;
		node->next = NULL;
		node->shm = shm;
		return node;
	 }
	 else
	 {
		return NULL;
	 }
	
}

/*esb:get count*/
ngx_int_t ngx_get_ngx_kv_count(ngx_esb_kv_node_t *head)
{
	ngx_esb_kv_node_t        *p;
	ngx_int_t                count;

	count = 0;
	p = head;
	while(p)
	{
		count++;
		p = p->next;
	}

	return count;
}


/*esb:destroy*/
ngx_int_t ngx_destroy_ngx_kv(ngx_esb_kv_node_t *head)
{
	ngx_esb_kv_node_t				*p,*pi;
	p = head;


	while(p != NULL)
	{
		pi = p->client_invoke;
		while(pi != NULL)
		{
			ngx_shm_free(&pi->shm);
			pi = pi->next;
		}
		ngx_shm_free(&p->shm);
		p = p->next;
	}

	return NGX_OK;
}

ngx_esb_kv_node_t  *ngx_server_invoke_info;

