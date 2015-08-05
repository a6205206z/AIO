
/*
 * Author: Cean Cheng(cean.ch@gmail.com)
 *
 */
 
#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>
#include <nginx.h>

#if (NGX_DEBUG)
#define VESB_DEBUG 1
#else
#define VESB_DEBUG 0
#endif

#define VESB_NORMAL_MODE 0
#define VESB_AIO_MODE 1
/*vesb:string ext*/
#define esb_str_set_strn(str, text, n)                                               \
    (str)->len = n; (str)->data = (u_char *) text
    


//static ngx_http_output_header_filter_pt  ngx_http_next_header_filter;
static ngx_int_t ngx_http_vesb_aio_process_request(ngx_http_request_t *r);
static ngx_int_t ngx_http_vesb_aio_init(ngx_conf_t *cf);
static void *ngx_http_vesb_aio_create_conf(ngx_conf_t *cf);
static char *ngx_http_vesb_aio_merge_conf(ngx_conf_t *cf, void *parent, void *child);

typedef struct {
    /*vesb:address list for services merge*/
	ngx_array_t				  *esb_merge_real_address;/*esb:*/
	ngx_int_t				  esb_mode;
    struct timeval            start_invoke_time,end_invoke_time;
}ngx_http_vesb_aio_main_conf_t;

static ngx_command_t  ngx_http_vesb_aio_commands[] = {
    /*esb:merge services*/
    { ngx_string("esb_merge_address"),
        NGX_HTTP_MAIN_CONF|NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_CONF_TAKE2,
        ngx_conf_set_keyval_slot,
        NGX_HTTP_MAIN_CONF_OFFSET,
        offsetof(ngx_http_vesb_aio_main_conf_t, esb_merge_real_address),
        NULL },
    ngx_null_command
};


static ngx_http_module_t  ngx_http_vesb_aio_module_ctx = {
    NULL,                                  /* preconfiguration */
    ngx_http_vesb_aio_init,                /* postconfiguration */
    ngx_http_vesb_aio_create_conf,         /* create main configuration */
    NULL,                                  /* init main configuration */

    NULL,                                  /* create server configuration */
    NULL,                                  /* merge server configuration */

    NULL,								   /* create location configuration */
    ngx_http_vesb_aio_merge_conf           /* merge location configuration */
};

ngx_module_t  ngx_http_vesb_aio_module = {
    NGX_MODULE_V1,
    &ngx_http_vesb_aio_module_ctx,         /* module context */
    ngx_http_vesb_aio_commands,            /* module directives */
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


static ngx_int_t 
ngx_http_vesb_aio_process_request(ngx_http_request_t *r)
{
	ngx_http_vesb_aio_main_conf_t *vacf;
	ngx_list_part_t 			  *part = &r->headers_in.headers.part;
	ngx_table_elt_t 			  *header = part->elts;
	ngx_keyval_t                  *merge_real_address_kv;/*vesb:merge*/
	ngx_uint_t				      i,j;
    
	vacf = ngx_http_get_module_main_conf(r, ngx_http_vesb_aio_module);
    
    gettimeofday(  &vacf->start_invoke_time, NULL );
    
	/*esb:start process*/
    gettimeofday( &r->esb_start_process_time, NULL );
    /*esb:esb mod url*/
	if(ngx_strstr(r->uri.data,"esbmerge") != NULL)//vacf->esb_mode == VESB_AIO_MODE)
	{
		for(i=0;/* void */;++i)
		{
			if(i >= part->nelts)
			{
				if(part->next == NULL)
				{
					break;
				}
				
				part = part->next;
				header = part->elts;
				i=0;
			}
			
			if(header[i].hash == 0)
			{
				continue;
			}
			
			if(ngx_strstr(header[i].key.data,"SOAPAction") != NULL
				||ngx_strstr(header[i].key.data,"Content-Type") != NULL)/*vesb:if soap action or content-type*/
			 {
				merge_real_address_kv = vacf->esb_merge_real_address->elts;
				for (j = 0; j < vacf->esb_merge_real_address->nelts; j++)
				{
					/*vesb:search the address table and find the real address*/
					if (ngx_strcasestrn(header[i].value.data
							 ,(char *)merge_real_address_kv[j].key.data
							 ,merge_real_address_kv[j].key.len - 1)) 
					{
                        esb_str_set_strn(&r->uri,
									merge_real_address_kv[j].value.data,
									merge_real_address_kv[j].value.len);
                                               
					 }
				}
			}
		}
	}
    /*esb:start process*/
    gettimeofday( &r->esb_end_process_time, NULL );
	return NGX_OK;
}

static void *
ngx_http_vesb_aio_create_conf(ngx_conf_t *cf)
{
    ngx_http_vesb_aio_main_conf_t  *vacf;

    vacf = ngx_pcalloc(cf->pool, sizeof(ngx_http_vesb_aio_main_conf_t));
    if (vacf == NULL) {
        return NGX_CONF_ERROR;
    }

    return vacf;
}

static char * 
ngx_http_vesb_aio_merge_conf(ngx_conf_t *cf, void *parent, void *child)
{
    //ngx_http_vesb_aio_loc_conf_t *prev = (ngx_http_vesb_aio_loc_conf_t *)parent;
    //ngx_http_vesb_aio_loc_conf_t *conf = (ngx_http_vesb_aio_loc_conf_t *)child;


    return NGX_CONF_OK;
}

static ngx_int_t
ngx_http_vesb_aio_init(ngx_conf_t *cf)
{
    ngx_http_handler_pt         *h;
    ngx_http_core_main_conf_t   *cmcf;
    
    cmcf = ngx_http_conf_get_module_main_conf(cf, ngx_http_core_module);
    
    h = ngx_array_push(&cmcf->phases[NGX_HTTP_POST_READ_PHASE].handlers);
    
    if(h == NULL)
    {
        return NGX_ERROR;
    }
    
    *h = ngx_http_vesb_aio_process_request;
    

    return NGX_OK;
}
