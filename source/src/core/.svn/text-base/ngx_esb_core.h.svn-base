/*
 * Copyright (C) Cean Cheng
 * Copyright (C) Vancl, Inc.
 */

#include <ngx_config.h>
#include <ngx_core.h>

#ifndef _NGX_ESB_CORE_H_INCLUDED_
#define _NGX_ESB_CORE_H_INCLUDED_

typedef struct ngx_esb_kv_node_s       ngx_esb_kv_node_t;
typedef volatile long             ngx_int_v_t;
/*
esb:key value list
*/
struct ngx_esb_kv_node_s 
{
	ngx_str_t            key;
    ngx_int_v_t          data;
	ngx_int_v_t          invoke_max_time;
	ngx_int_v_t          invoke_avg_time;
	ngx_int_v_t          esb_proc_max_time;
	ngx_shm_t            shm;
	ngx_esb_kv_node_t    *next;
	ngx_esb_kv_node_t    *client_invoke;
};

/*esb:search node by key*/
ngx_esb_kv_node_t *ngx_find_ngx_esb_kv_node(ngx_esb_kv_node_t *head, ngx_str_t key);


/*esb:append at end*/
ngx_esb_kv_node_t *ngx_append_ngx_esb_kv_node(ngx_esb_kv_node_t *head, ngx_esb_kv_node_t *n);


/*esb:create*/
ngx_esb_kv_node_t *ngx_create_ngx_esb_kv_node(ngx_log_t *log,ngx_int_t data,ngx_str_t key);

/*esb:get count*/
ngx_int_t ngx_get_ngx_kv_count(ngx_esb_kv_node_t *head);

/*esb:destroy*/
ngx_int_t ngx_destroy_ngx_kv(ngx_esb_kv_node_t *head);

extern ngx_esb_kv_node_t  *ngx_server_invoke_info;

#endif /* _NGX_ESB_CORE_H_INCLUDED_ */
