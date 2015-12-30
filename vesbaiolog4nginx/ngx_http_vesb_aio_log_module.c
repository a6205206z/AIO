
/*
 * Author: Cean Cheng(cean.ch@gmail.com)
 *
 */ 

#include <ngx_config.h>
#include <ngx_core.h>
#include <ngx_http.h>
#include <nginx.h>
#include "mongo.h"
#include <stdio.h>
#include <time.h>


#if (NGX_DEBUG)
#define VESB_DEBUG 1
#else
#define VESB_DEBUG 0
#endif

typedef struct {
	ngx_str_t             mongodb_conn_str;
	ngx_int_t             mongodb_conn_port;
	ngx_int_t             mongodb_enable;
}ngx_http_vesb_aio_log_main_conf_t;

typedef struct {
	/*esb:for customzie*/
	ngx_int_t                      invoke_max_time;
    ngx_str_t                      app_name;
    ngx_int_t                      enable_tracking;
}ngx_http_vesb_aio_log_loc_conf_t;

static ngx_int_t ngx_http_aio_log_header_filter(ngx_http_request_t *r);
static ngx_int_t ngx_http_aio_log_body_filter(ngx_http_request_t *r,
    ngx_chain_t *in);
static ngx_int_t ngx_http_vesb_aio_log_init(ngx_conf_t *cf);
static void *ngx_http_vesb_aio_log_main_create_conf(ngx_conf_t *cf);
static void *ngx_http_vesb_aio_log_loc_create_conf(ngx_conf_t *cf);
static char *ngx_http_vesb_aio_log_merge_conf(ngx_conf_t *cf, void *parent, void *child);
static char *ngx_http_mongodb_set_conn(ngx_conf_t *cf,ngx_command_t *cmd, void *conf);
static char *ngx_http_vesb_customize_info(ngx_conf_t *cf,ngx_command_t *cmd, void *conf);

static void ngx_wirte_log_into_mongodb(ngx_http_request_t *r,
	ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse);
/*write tracking log*/
static void ngx_wirte_tracking_into_mongodb(ngx_http_request_t *r,
                                            ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse);

/*esb:ngx_http_proxy_write_client_invoke_record*/
static void ngx_http_proxy_write_client_invoke_record(ngx_http_request_t *r,
	ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse);
static u_char *ngx_http_proxy_read_request_body(ngx_http_request_t *r);
static u_char *ngx_http_proxy_read_response_body(ngx_http_request_t *r);
static ngx_http_output_header_filter_pt  ngx_http_next_header_filter;
static ngx_http_output_body_filter_pt    ngx_http_next_body_filter;





static ngx_command_t  ngx_http_vesb_aio_log_commands[] = {
	  /*esb:mongodb connection string config*/
	  { ngx_string("mongodb_conn"),
      NGX_HTTP_MAIN_CONF|NGX_HTTP_SRV_CONF|NGX_HTTP_LOC_CONF|NGX_HTTP_LIF_CONF
                        |NGX_HTTP_LMT_CONF|NGX_CONF_TAKE123,
      ngx_http_mongodb_set_conn,
      NGX_HTTP_MAIN_CONF_OFFSET,
      0,
      NULL },
	  { ngx_string("proxy_pass_customize"),
      NGX_HTTP_LOC_CONF|NGX_HTTP_LIF_CONF|NGX_HTTP_LMT_CONF|NGX_CONF_TAKE123,
      ngx_http_vesb_customize_info,
      NGX_HTTP_LOC_CONF_OFFSET,
      0,
      NULL },
    ngx_null_command
};

static ngx_http_module_t  ngx_http_vesb_aio_log_module_ctx = {
    NULL,                                      /* preconfiguration */
    ngx_http_vesb_aio_log_init,                /* postconfiguration */
    ngx_http_vesb_aio_log_main_create_conf,	   /* create main configuration */
    NULL,                                      /* init main configuration */

    NULL,									   /* create server configuration */
    NULL,									   /* merge server configuration */

    ngx_http_vesb_aio_log_loc_create_conf,	   /* create location configuration */
    ngx_http_vesb_aio_log_merge_conf           /* merge location configuration */
};

ngx_module_t  ngx_http_vesb_aio_log_module = {
    NGX_MODULE_V1,
    &ngx_http_vesb_aio_log_module_ctx,     /* module context */
    ngx_http_vesb_aio_log_commands,        /* module directives */
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



static 
ngx_int_t ngx_http_aio_log_header_filter(ngx_http_request_t *r)
{
	ngx_http_vesb_aio_log_main_conf_t				*valcf;
	ngx_http_vesb_aio_log_loc_conf_t				*vallcf;
	int												timeuse;
	struct timeval									endinvoketime;

	valcf = ngx_http_get_module_main_conf(r, ngx_http_vesb_aio_log_module);
	vallcf = ngx_http_get_module_loc_conf(r, ngx_http_vesb_aio_log_module);

	gettimeofday( &endinvoketime, NULL );
	timeuse = 1000000 * ( endinvoketime.tv_sec - r->esb_start_process_time.tv_sec ) 
		+ endinvoketime.tv_usec - r->esb_start_process_time.tv_usec; 

	/*esb: ngx_stat_proxy_request*/
	(void) ngx_atomic_fetch_add(ngx_stat_proxy_request, 1);

	/*esb:write invoke record*/
    ngx_http_proxy_write_client_invoke_record(r, valcf,vallcf,timeuse);

    
    if(valcf->mongodb_enable >0 && vallcf->enable_tracking >0)
    {
        ngx_wirte_tracking_into_mongodb(r,valcf,vallcf,timeuse);
    }

    
	return ngx_http_next_header_filter(r);
}
static 
ngx_int_t ngx_http_aio_log_body_filter(ngx_http_request_t *r,ngx_chain_t *in)
{
  ngx_http_vesb_aio_log_main_conf_t				*valcf;
	ngx_http_vesb_aio_log_loc_conf_t				*vallcf;
	int												timeuse;
	struct timeval									endinvoketime;
    
	valcf = ngx_http_get_module_main_conf(r, ngx_http_vesb_aio_log_module);
	vallcf = ngx_http_get_module_loc_conf(r, ngx_http_vesb_aio_log_module);
    
	gettimeofday( &endinvoketime, NULL );
	timeuse = 1000000 * ( endinvoketime.tv_sec - r->esb_start_process_time.tv_sec )
    + endinvoketime.tv_usec - r->esb_start_process_time.tv_usec;
    
    
  if(valcf->mongodb_enable > 0)
	{
		ngx_wirte_log_into_mongodb(r,valcf,vallcf,timeuse);
	}
	else
	{
		ngx_log_error(NGX_LOG_NOTICE, r->connection->log, 0, "mongodb is unenable!!! request log will be shut down!");
	}
    //empty method
    return ngx_http_next_body_filter(r, in);
}

static void *
ngx_http_vesb_aio_log_main_create_conf(ngx_conf_t *cf)
{
    ngx_http_vesb_aio_log_main_conf_t  *valcf;

    valcf = ngx_pcalloc(cf->pool, sizeof(ngx_http_vesb_aio_log_main_conf_t));
    if (valcf == NULL) {
        return NGX_CONF_ERROR;
    }

    return valcf;
}

static void *
ngx_http_vesb_aio_log_loc_create_conf(ngx_conf_t *cf)
{
    ngx_http_vesb_aio_log_loc_conf_t  *valcf;

    valcf = ngx_pcalloc(cf->pool, sizeof(ngx_http_vesb_aio_log_loc_conf_t));
    if (valcf == NULL) {
        return NGX_CONF_ERROR;
    }

    return valcf;
}

static char * 
ngx_http_vesb_aio_log_merge_conf(ngx_conf_t *cf, void *parent, void *child)
{
    //ngx_http_vesb_aio_loc_conf_t *prev = (ngx_http_vesb_aio_loc_conf_t *)parent;
    //ngx_http_vesb_aio_loc_conf_t *conf = (ngx_http_vesb_aio_loc_conf_t *)child;


    return NGX_CONF_OK;
}

static ngx_int_t
ngx_http_vesb_aio_log_init(ngx_conf_t *cf)
{
    ngx_http_next_header_filter = ngx_http_top_header_filter;
    ngx_http_top_header_filter = ngx_http_aio_log_header_filter;

    ngx_http_next_body_filter = ngx_http_top_body_filter;
    ngx_http_top_body_filter = ngx_http_aio_log_body_filter;

    return NGX_OK;
}

/*esb:mongodb connection string config*/
static char * 
ngx_http_mongodb_set_conn(ngx_conf_t *cf,ngx_command_t *cmd, void *conf)
{
	ngx_http_vesb_aio_log_main_conf_t *valcf = conf;
	mongo							  conn[1];
	ngx_str_t                         *value;
	char                              *conn_str;
  int                               port,status;

	
	value = cf->args->elts;

	valcf->mongodb_conn_str = value[1];

	valcf->mongodb_conn_port = ngx_atoi(value[2].data, value[2].len);

	if (valcf->mongodb_conn_port == (ngx_int_t) NGX_ERROR) {
        ngx_conf_log_error(NGX_LOG_EMERG, cf, 0,
                           "invalid number \"%V\"", &value[2]);

        return NGX_CONF_ERROR;
    }

    conn_str = (char *)valcf->mongodb_conn_str.data;
    port = (int)valcf->mongodb_conn_port;

    status = mongo_client( conn, conn_str, port );

    if( status != MONGO_OK ) {
	    ngx_conf_log_error(NGX_LOG_EMERG, cf, 0,
                           "can't connect server %V %V", &value[1],&value[2]);

        return NGX_CONF_ERROR;
    }



	mongo_destroy( conn );

	valcf->mongodb_enable = 1;

	return NGX_CONF_OK;
}

static char * 
ngx_http_vesb_customize_info(ngx_conf_t *cf,ngx_command_t *cmd, void *conf)
{
	ngx_http_vesb_aio_log_loc_conf_t  *vallcf = conf;
	ngx_str_t						  *value;

	value = cf->args->elts;

	vallcf->app_name = value[1];

	vallcf->invoke_max_time = ngx_atoi(value[2].data, value[2].len);

    if (vallcf->invoke_max_time == -1) {
        vallcf->enable_tracking = 1;
    }


	return NGX_CONF_OK;
}

/*
*esb:write log into mongodb 
*/
static void
ngx_wirte_log_into_mongodb(ngx_http_request_t *r,
	ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse)
{
  mongo                     conn[1];
  u_char                    *request_body,*response_body;
  char                      *conn_str;
  int                       port,status,insertstatus;
  time_t                    timenow;
  ngx_list_part_t 		      *part = &r->headers_in.headers.part;
  ngx_table_elt_t 		      *header = part->elts;
  ngx_uint_t				        i;

  conn_str = (char *)valcf->mongodb_conn_str.data;
  port = (int)valcf->mongodb_conn_port;

  
  if((timeuse >= vallcf->invoke_max_time
	  &&	0 < vallcf->invoke_max_time)
	  || (r->headers_out.status >= NGX_HTTP_INTERNAL_SERVER_ERROR
	  && r->headers_out.status <= NGX_HTTP_INSUFFICIENT_STORAGE)) //timeusr over invoke_max_time or request error
  {
	  status = mongo_client( conn, conn_str, port );
	  if( status == MONGO_OK ) {
		 mongo_set_op_timeout( conn, 10000 );//time oust 10000 ms
		 time ( &timenow );
		 bson b[1];
		 bson_init( b );
		 bson_append_new_oid( b, "_id" );
		 bson_append_string_n( b, "location", (char *)r->request_line.data,r->request_line.len);
		 bson_append_string_n( b, "url",  (char *)r->uri.data,r->uri.len);
		 bson_append_string_n( b, "clientip",  (char *)r->connection->addr_text.data, r->connection->addr_text.len);

		 if( r->headers_in.x_forwarded_for == NULL )
		 {
			bson_append_string_n( b, "realclientip" , (char *)r->connection->addr_text.data, r->connection->addr_text.len);
		 }
		 else
		 {
			bson_append_string_n( b, "realclientip" , (char *)r->headers_in.x_forwarded_for->value.data,r->headers_in.x_forwarded_for->value.len);
		 }

		 bson_append_int( b,"statuscode",r->headers_out.status);
		 bson_append_int( b,"usetime",timeuse);
     bson_append_int( b,"requestsize",r.request_length);
     bson_append_int( b,"responsesize",r->headers_out.content_length_n);
		 bson_append_time_t( b,"invoketime",timenow );

		 
		 if(vallcf->app_name.data != NULL)
		 {
			bson_append_string_n( b,"appname",(char *)vallcf->app_name.data,vallcf->app_name.len);
		 }
		 else
		 {
			 bson_append_string( b,"appname","undefine");
		 }
		 
		 request_body = ngx_http_proxy_read_request_body(r);
		 

	   if(request_body != NULL)
		 {
			bson_append_string( b, "requestbody", (char *)request_body);
		 }

		 if((r->headers_out.status >= NGX_HTTP_INTERNAL_SERVER_ERROR
			&& r->headers_out.status <= NGX_HTTP_INSUFFICIENT_STORAGE))
		 {
			 response_body = ngx_http_proxy_read_response_body(r);
	 		 if(response_body != NULL)
			 {
				bson_append_string( b, "responsebody", (char *)response_body);
			 }
		 }
		 
      /*get method name*/
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
          
          if(ngx_strstr(header[i].key.data,"SOAPAction") != NULL)
          {
              bson_append_string_n( b,"SOAPAction",(char *)header[i].value.data,header[i].value.len);
          }
          else if(ngx_strstr(header[i].key.data,"Content-Type") != NULL)
          {
              bson_append_string_n( b,"Content-Type",(char *)header[i].value.data,header[i].value.len);
          }
      }
		

		 bson_finish( b );
		 insertstatus = mongo_insert( conn,"vesb.tracking", b , NULL );
		 
		 if( insertstatus != MONGO_OK ) {
  			ngx_log_error(NGX_LOG_NOTICE, r->connection->log, 0, "insert request log in mongodb is failed!(error:%d)",conn[0].err);
  		 }
  		 bson_destroy( b );
  	  }
  	  else
  	  {
  		  ngx_log_error(NGX_LOG_NOTICE, r->connection->log, 0, "mongodb is unconnection!(error:%d)",status);
  	  }

	   mongo_destroy( conn );
  }
}

static void
ngx_wirte_tracking_into_mongodb(ngx_http_request_t *r,
        ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse)
{
    mongo                     conn[1];
    char                      *conn_str;
    int                       port,status,insertstatus;
    time_t                    timenow;
    ngx_list_part_t 		      *part = &r->headers_in.headers.part;
	  ngx_table_elt_t 		      *header = part->elts;
    ngx_uint_t				        i;
    
    if(r->headers_out.status == NGX_HTTP_OK)
    {
        conn_str = (char *)valcf->mongodb_conn_str.data;
        port = (int)valcf->mongodb_conn_port;
        status = mongo_client( conn, conn_str, port );
        if( status == MONGO_OK ) {
            mongo_set_op_timeout( conn, 10000 );//time oust 10000 ms
            time ( &timenow );
            bson b[1];
            bson_init( b );
            bson_append_new_oid( b, "_id" );
        
            if( r->headers_in.x_forwarded_for == NULL )
            {
                bson_append_string_n( b, "realclientip" , (char *)r->connection->addr_text.data, r->connection->addr_text.len);
            }
            else
            {
                bson_append_string_n( b, "realclientip" , (char *)r->headers_in.x_forwarded_for->value.data,r->headers_in.x_forwarded_for->value.len);
            }
            
            bson_append_int( b,"statuscode",r->headers_out.status);
            bson_append_int( b,"usetime",timeuse);
            bson_append_int( b,"requestsize",r.request_length);
            bson_append_int( b,"responsesize",r->headers_out.content_length_n);
            bson_append_time_t( b,"invoketime",timenow );
            
            
            if(vallcf->app_name.data != NULL)
            {
                bson_append_string_n( b,"appname",(char *)vallcf->app_name.data,vallcf->app_name.len);
            }
            else
            {
                bson_append_string( b,"appname","undefine");
            }
       
            /*get method name*/
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
			
                if(ngx_strstr(header[i].key.data,"SOAPAction") != NULL)
                {
                    bson_append_string_n( b,"SOAPAction",(char *)header[i].value.data,header[i].value.len);
                }
                else if(ngx_strstr(header[i].key.data,"Content-Type") != NULL)
                {
                    bson_append_string_n( b,"Content-Type",(char *)header[i].value.data,header[i].value.len);
                }
            }
            
            
            bson_finish( b );
            insertstatus = mongo_insert( conn,"vesb.tracking", b , NULL );
            
            if( insertstatus != MONGO_OK ) {
                ngx_log_error(NGX_LOG_NOTICE, r->connection->log, 0, "insert tracking log in mongodb is failed!(error:%d)",conn[0].err);
            }
            bson_destroy( b );
        }
        else
        {
                ngx_log_error(NGX_LOG_NOTICE, r->connection->log, 0, "mongodb is unconnection!(error:%d)",status);
        }
    
        mongo_destroy( conn );

    }
}

/*esb:ngx_http_proxy_write_client_invoke_record*/
static void
ngx_http_proxy_write_client_invoke_record(ngx_http_request_t *r,ngx_http_vesb_aio_log_main_conf_t *valcf,ngx_http_vesb_aio_log_loc_conf_t *vallcf,int timeuse)
{
    ngx_esb_kv_node_t                      *server_invoke_info,*client_invoke_info;
    ngx_str_t                              client_ip_address,unregisterclient=ngx_string("unregisterclient");
    int									   esb_prcess_timeuse;

    /*esb:write server log*/
    server_invoke_info = ngx_find_ngx_esb_kv_node(ngx_server_invoke_info,vallcf->app_name);
    
    if(server_invoke_info != NULL)
    {
		server_invoke_info->data++;
		if(server_invoke_info->invoke_max_time < timeuse)
		{
			server_invoke_info->invoke_max_time = timeuse;
		}

		esb_prcess_timeuse = 1000000 * ( r->esb_end_process_time.tv_sec - r->esb_start_process_time.tv_sec ) + r->esb_end_process_time.tv_usec - r->esb_start_process_time.tv_usec; 

		
		if(server_invoke_info->esb_proc_max_time < esb_prcess_timeuse)
		{
			server_invoke_info->esb_proc_max_time = esb_prcess_timeuse;
		}

		if(server_invoke_info->invoke_avg_time == 0)
		{
			server_invoke_info->invoke_avg_time = timeuse;
		}
		else
		{
			server_invoke_info->invoke_avg_time = (server_invoke_info->invoke_avg_time + timeuse) / 2;
		}
        
        
        if( r->headers_in.x_forwarded_for == NULL )
        {
            client_ip_address = r->connection->addr_text;
        }
        else
        {
            client_ip_address = r->headers_in.x_forwarded_for->value;
        }
        
        /*esb:write client invoke record*/
        /*esb: about client invoke*/
        if(client_ip_address.data != NULL)
        {
            client_invoke_info = ngx_find_ngx_esb_kv_node(server_invoke_info->client_invoke,client_ip_address);
        
            if (client_invoke_info == NULL) {

                client_invoke_info = ngx_find_ngx_esb_kv_node(server_invoke_info->client_invoke,unregisterclient);
            }
            
            if(client_invoke_info != NULL)
            {
                client_invoke_info->data++;
            }
        }
    }
}


/*
*esb:read request body
*/
static u_char *
ngx_http_proxy_read_request_body(ngx_http_request_t *r)
{
	ngx_http_request_body_t            *rb;
  u_char                             *post_content = NULL;
    
	if (r->request_body != NULL
		&& r->request_body->buf != NULL) { 
		rb = r->request_body; 


	  post_content = ngx_palloc(r->pool, rb->buf->last-rb->buf->pos + 2);
		ngx_cpystrn(post_content, rb->buf->pos, rb->buf->last-rb->buf->pos + 1);
	}
		
	return post_content;
}
/*
*esb:read response body
*/
static u_char *
ngx_http_proxy_read_response_body(ngx_http_request_t *r)
{
	u_char                             *response_content = NULL;
	ngx_chain_t                        *out;
	out = r->out;

	
	//ngx_http_output_filter(r, out);
	while(out != NULL)
	{
		if (out->buf != NULL) { 
			response_content = ngx_palloc(r->pool, out->buf->last-out->buf->pos + 2);
			ngx_cpystrn(response_content, out->buf->pos, out->buf->last-out->buf->pos + 1);
		}		
	    out = out->next;
	}

    return response_content;
}
