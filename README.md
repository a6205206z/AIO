#Based on ![Nginx](http://nginx.org/nginx.png)
#HOW TO INSTALL
##install dependencies
```Bash
yum groupinstall "Development Tools"
yum -y install pcre* zlib* openssl*
```
##install mongodb driver
```Bash
cd /AIOPath/mongo-c-driver
make & make install
```
##install aio
###entry sourcecode folder
```Bash
cd /AIOPath/source<br>
./esbcfg.sh<br>
make & make install<br>
vi/etc/ld.so.conf<br>
```
###add /usr/local/lib and save
```Bash
/sbin/ldconfig
```
###start aio
```Bash
/usr/local/nginx/sbin/nginx
```
***
# AIO
## 2015/12/04
add prometheus monitor data collecter
