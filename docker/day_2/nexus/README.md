# NEXUS AND NGINX CONFIGURATION
The code in this folder will allow us to build and configure docker-compose file that will run:
- NGINX As a reverse proxy
- Nexus Private repo 
---

#### First let's review the docker-compose file
##### (comments inline and details down below)

~~~
version: '3'

services:
  nginx: # OUR FIRST CONTAINER NAME
    image: nginx:1.15-alpine # NGINX IMAGE AND VERSION
    restart: unless-stopped # CONFIGURE THE CONTINER TO AUTOMATICLY RESTART UNLESS STOPPED
    volumes: # CONFIGURE LOCAL VOLUME THAT WILL STORE:
      - ./configuration/nginx/nexus.conf:/etc/nginx/conf.d/nexus.conf # NGINX CUSTOM CONFIGURATION FILE
      - ./configuration/nginx/certs:/etc/letsencrypt
      - ./configuration/certbot/www:/var/www/certbot
    ports: # PORTS TO BE EXPOSED FOR THE CONTAINER
      - "80:80"
      - "443:443"
      - "5000:5000"
    command: "/bin/sh -c 'while :; do sleep 6h & wait $${!}; nginx -s reload; done & nginx -g \"daemon off;\"'"
    networks: # CUSTOM NETWORK BETWEEN THE NGINX AND NEXUS CONTAINER
      - nexus
  
  nexus:
      image: sonatype/nexus3
      restart: unless-stopped
      volumes: 
        - ./volume/nexus/data:/nexus-data
      environment: # CUSTOM ENVIRONMENT VARIABLE 
        - NEXUS_CONTEXT=nexus
      networks:
       - nexus
   
  
networks:
  nexus:
~~~~

# Details 
Using Volumes we are mounting specific files into specific locations. this way we can change our application configuration without creating new images everytime we wish to change our configuration.

Do <b>note</b>! that in K8S we will use diffrent ways to accomplish this such as configmaps and KV mount
~~~
volumes:
      - ./configuration/nginx/nexus.conf:/etc/nginx/conf.d/nexus.conf
      - ./configuration/nginx/certs:/etc/certificates
~~~

- Custom NGINX Configuration ./configuration/nginx/nexus.conf:/etc/nginx/conf.d/nexus.conf
- Custom SSL certificates to be loaded by NGINX ./configuration/nginx/certs:/etc/certificates


We also use docker networks to allow DNS resolution using the name of the containers.

This will allow the NGINX container to "call" the NEXUS container by it's name


<b>Nexus</b> Continer spec uses the volume option to create a presistant storage for /nexus-data folder inside the continer and map it to your local folder .volume/nexus/data:
~~~
      volumes: 
        - ./volume/nexus/data:/nexus-data
~~~

This allow us to preserve all of Nexus data even if we  shutdown nexus container and start a new one later on.