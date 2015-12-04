#HOW TO INSTALL
##install dependencies
 #yum groupinstall "Development Tools"
 #yum -y install pcre* zlib* openssl*
##install mongodb driver
 #cd /AIOPath/mongo-c-driver
 #make & make install
##install aio
###entry sourcecode folder
 #cd /AIOPath/source
 #./esbcfg.sh
 #make & make install
 #vi/etc/ld.so.conf
###add /usr/local/lib and save
 #/sbin/ldconfig
###start aio
 #/usr/local/nginx/sbin/nginx

# AIO
## 2015/12/04
add prometheus monitor data collecter
