# NGINX REVERSE PROXY CONFIGURATION
As this is not an NGINX session, we wont dive too deep the rabit hole BUT at least let's understand the nginx configuration file:


The main idea behind the nginx configuration is that we have a configuration where every request that opens to our NGINX container under the URI:
~~~
 /nexus
~~~

will be proxied to our Nexus container using a proxy_pass feature.

~~~
location /nexus {
  ....
  proxy_pass http://nexus/nexus;
~~~

Nginx knows the IP of our Nexus Container due to the CUSTOM network we created between them in our Docker-compose file.

** Using Custom network cross containers allow DNS resolving inside each continer.


