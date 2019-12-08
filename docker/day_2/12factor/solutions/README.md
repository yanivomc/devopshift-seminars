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
curl -Ss http://localhost:9000/api/system/status | grep -e '"status":"UP"' 
OR
browse http://REMOTEIP:9000/
Login (on the top left) USER: admin , PASS: admin
~~~

### BUILD AND RUN CONTAINER PIPELINE
Copy the Dockerfile located in this folder to your project.
Make sure the Dockerfile sits ABOVE flask-api folder as show here:
~~~
.
├── Dockerfile
├── README.md
└── flask-api
    ├── Makefile
    ├── README.md
    ├── api
    │   ├── _01_manual_response_class.py
    │   ├── _02_make_response_helper.py
    │   ├── _03_post_method.py
    │   ├── _04_delete_method.py
    │   ├── _05_flask_restful_simple-fix
    │   ├── _05_flask_restful_simple.py
    │   ├── __init__.py
    │   └── utils.py
~~~
1. Build 
~~~
docker build --network sonarnet -t 127.0.0.1.nip.io/[imagename:tag]
~~~

<b>EXPECTED RESULTS</b> LINT FAILS
~~~
************* Module api._05_flask_restful_simple
api/_05_flask_restful_simple.py:37: [W0603(global-statement), BookListResource.post] Using the global statement

-----------------------------------
Your code has been rated at 9.96/10
~~~
<b> FIX LINT</b>
Follow Slides OR replace the failed file flask-api/api/_05_flask_restful_simple.py with the one in the solutions folder
~~~
cp ../solutions/_05_flask_restful_simple.py flask-api/api/
RERUN BUILD
docker build --network sonarnet -t 127.0.0.1.nip.io/[imagename:tag]
~~~
2. Run 
~~~
docker run -p 0.0.0.0:8080:8080 -ti 127.0.0.1.nip.io/[imagename:tag]
~~~

TEST
~~~
curl -f http://localhost:8080/book                                                                                                                                            
~~~
<b>EXPECTED RESULTS</b>
~~~
[{"id": 33, "title": "The Raven", "author_id": 1}]%                                                                                                                           ~~~   
3. Push to your repo
~~~
docker push 127.0.0.1.nip.io/[imagename:tag]
~~~