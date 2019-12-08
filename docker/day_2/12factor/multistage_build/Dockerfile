# Copies in our code and runs NPM Install
# Phase I
FROM python:latest as builder
WORKDIR /usr/src/app
COPY flask-api/ ./
RUN pip install -r requirements.txt

# Lints Code
# Phase II
FROM eeacms/pylint:latest as linting
WORKDIR /code
COPY --from=builder /usr/src/app/pylint.cfg /etc/pylint.cfg
COPY --from=builder /usr/src/app/*.py ./
COPY --from=builder /usr/src/app/api ./api
RUN ["/docker-entrypoint.sh", "pylint"]

# Phase III Running Sonarqube scanner (Sonarqube server also required)
FROM newtmitch/sonar-scanner as sonarqube
WORKDIR /usr/src
COPY --from=builder /usr/src/app/*.py ./
COPY --from=builder /usr/src/app/api ./api
RUN sonar-scanner -Dsonar.projectBaseDir=/usr/src
# Runs Unit Tests
# Phase IV
FROM python:latest as unit-tests
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/*.py ./
COPY --from=builder /usr/src/app/api ./api
RUN [“make”, “test”]

# Starts and Serves Web Page
# Phase VI
FROM python:latest as serve
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/*.py ./
COPY --from=builder /usr/src/app/api ./api
RUN [“python”, “run_app.py”]