
/*
 * Copyright (C) Igor Sysoev
 * Copyright (C) Nginx, Inc.
 */


#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>

/*esb: status*/
ngx_str_t                      ngx_status_callback_name;

static char *ngx_http_set_status(ngx_conf_t *cf, ngx_command_t *cmd,
                                 void *conf);

/*esb:ngx_http_set_status_callback_name*/
static char *ngx_http_set_status_callback_name(ngx_conf_t *cf, ngx_command_t *cmd,
                                 void *conf);

static ngx_command_t  ngx_http_status_commands[] = {

    { ngx_string("stub_status"),
      NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_CONF_FLAG,
      ngx_http_set_status,
      0,
      0,
      NULL },
	{ ngx_string("status_callback_name"),
      NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_CONF_FLAG,
      ngx_http_set_status_callback_name,
      0,
      0,
      NULL },

      ngx_null_command
};



static ngx_http_module_t  ngx_http_stub_status_module_ctx = {
    NULL,                                  /* preconfiguration */
    NULL,                                  /* postconfiguration */

    NULL,                                  /* create main configuration */
    NULL,                                  /* init main configuration */

    NULL,                                  /* create server configuration */
    NULL,                                  /* merge server configuration */

    NULL,                                  /* create location configuration */
    NULL                                   /* merge location configuration */
};


ngx_module_t  ngx_http_stub_status_module = {
    NGX_MODULE_V1,
    &ngx_http_stub_status_module_ctx,      /* module context */
    ngx_http_status_commands,              /* module directives */
    NGX_HTTP_MODULE,                       /* module type */
    NULL,                                  /* init master */
    NULL,                                  /* init module */
    NULL,                                  /* init process */
    NULL,                                  /* init thread */
    NULL,                                  /* exit thread */
    NULL,                                  /* exit process */
    NULL,                                  /* exit master */
    NGX_MODULE_V1_PADDING
};


static ngx_int_t ngx_http_status_handler(ngx_http_request_t *r)
{
    size_t             size;
    ngx_int_t          rc;
    ngx_buf_t         *b;
    ngx_chain_t        out;
    ngx_atomic_int_t   ap, hn, ac, rq, rd, wr,prq;

    if (r->method != NGX_HTTP_GET && r->method != NGX_HTTP_HEAD) {
        return NGX_HTTP_NOT_ALLOWED;
    }

    rc = ngx_http_discard_request_body(r);

    if (rc != NGX_OK) {
        return rc;
    }

    ngx_str_set(&r->headers_out.content_type, "text/plain");

    if (r->method == NGX_HTTP_HEAD) {
        r->headers_out.status = NGX_HTTP_OK;

        rc = ngx_http_send_header(r);

        if (rc == NGX_ERROR || rc > NGX_OK || r->header_only) {
            return rc;
        }
    }

    size = ngx_status_callback_name.len + sizeof("([{\"activeconnections\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"accepts\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"handled\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"requests\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"proxyrequest\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"reading\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"writing\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\",")
		   + sizeof("\"waiting\":") + sizeof("\"") + NGX_ATOMIC_T_LEN + sizeof("\"}])");

    b = ngx_create_temp_buf(r->pool, size);
    if (b == NULL) {
        return NGX_HTTP_INTERNAL_SERVER_ERROR;
    }

    out.buf = b;
    out.next = NULL;

    ap = *ngx_stat_accepted;
    hn = *ngx_stat_handled;
    ac = *ngx_stat_active;
    rq = *ngx_stat_requests;
    rd = *ngx_stat_reading;
    wr = *ngx_stat_writing;
	/*esb ngx_stat_proxy_request*/
	prq = *ngx_stat_proxy_request;
	b->last = ngx_sprintf(b->last, "%s", ngx_status_callback_name.data);
    b->last = ngx_sprintf(b->last, "([{\"activeconnections\":\"%uA\",", ac);
	b->last = ngx_sprintf(b->last, "\"accepts\":\"%uA\",", ap);
	b->last = ngx_sprintf(b->last, "\"handled\":\"%uA\",", hn);
	b->last = ngx_sprintf(b->last, "\"requests\":\"%uA\",", rq);
	b->last = ngx_sprintf(b->last, "\"proxyrequest\":\"%uA\",", prq);
	b->last = ngx_sprintf(b->last, "\"Reading\":\"%uA\",", rd);
	b->last = ngx_sprintf(b->last, "\"Writing\":\"%uA\",", wr);
	b->last = ngx_sprintf(b->last, "\"Waiting\":\"%uA\"}])", ac - (rd + wr));


    r->headers_out.status = NGX_HTTP_OK;
    r->headers_out.content_length_n = b->last - b->pos;

    b->last_buf = (r == r->main) ? 1 : 0;

    rc = ngx_http_send_header(r);

    if (rc == NGX_ERROR || rc > NGX_OK || r->header_only) {
        return rc;
    }

    return ngx_http_output_filter(r, &out);
}


static char *ngx_http_set_status(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
{
    ngx_http_core_loc_conf_t  *clcf;
	


    clcf = ngx_http_conf_get_module_loc_conf(cf, ngx_http_core_module);
    clcf->handler = ngx_http_status_handler;

	

    return NGX_CONF_OK;
}

/*esb:ngx_http_set_status_callback_name*/
static char *ngx_http_set_status_callback_name(ngx_conf_t *cf, ngx_command_t *cmd, void *conf)
{
	ngx_str_t                 *value;

	/*esb: status*/
	value = cf->args->elts;
	ngx_status_callback_name = value[1];


	return NGX_CONF_OK;
}
