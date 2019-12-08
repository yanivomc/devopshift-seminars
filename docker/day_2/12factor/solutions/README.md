# MULTISTAGE BUILD LAB SOLUTION
### Start SonarQube server
1. Create custom network
~~~
docker network create sonarnet
~~~

2. Run sonarqube server
~~~
docker run -d --rm --name sonarqube --network sonarnet -p 0.0.0.0:9000:9000 sonarqube
~~~
3. Check that Sonarqube is running and ready to get connections 
~~~
curl -Ss  http://localhost:9000/issues | grep  UP
OR
browse http://REMOTEIP:9000/
USER: admin , PASS: admin
~~~

### BUILD AND RUN CONTAINER PIPELINE
1. Build 
~~~
docker build -t 127.0.0.1.nip.io/[imagename:tag]
~~~
2. Run 
~~~
docker run -p 8080:8080 -ti 127.0.0.1.nip.io/[imagename:tag]
TEST
curl -f http://localhost:8080/book                                                                                                                                            
~~~
<b>EXPECTED RESULTS</b>
~~~
[{"id": 33, "title": "The Raven", "author_id": 1}]%                                                                                                                           ~~~   
3. Push to your repo
~~~
docker push 127.0.0.1.nip.io/[imagename:tag]
~~~