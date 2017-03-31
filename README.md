# Based on ![Nginx](http://nginx.org/nginx.png)

# HOW TO INSTALL

## install dependencies

```Bash
yum groupinstall "Development Tools"
yum -y install pcre* zlib* openssl*
```
## install mongodb driver

```Bash
cd /AIOPath/mongo-c-driver
make & make install
```
## install aio

### entry sourcecode folder

```Bash
chmod 755 -R /AIOPath
cd /AIOPath/source
./esbcfg.sh
make & make install
vi /etc/ld.so.conf
```
### add /usr/local/lib and save

```Bash
/sbin/ldconfig
```
### start aio

```Bash
/usr/local/nginx/sbin/nginx
```
***

# AIO

![Architecture](https://github.com/a6205206z/AIO/blob/master/architecture.jpg)
## 2015/12/04

> add prometheus monitor data collecter
