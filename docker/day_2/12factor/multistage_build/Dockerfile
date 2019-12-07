# Copies in our code and runs NPM Install
# Phase I
FROM node:latest as builder
WORKDIR /usr/src/app
COPY package* ./
COPY src/ src/
RUN [“npm”, “install”]
# Lints Code
# Phase II
FROM node:latest as linting
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/ .
RUN [“npm”, “lint”]
# Gets Sonarqube Scanner from Dockerhub and runs it
# Phase III
FROM newmitch/sonar-scanner:latest as sonarqube
COPY --from=builder /usr/src/app/src /root/src
# Runs Unit Tests
# Phase IV
FROM node:latest as unit-tests
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/ .
RUN [“npm”, “test”]
# Runs Accessibility Tests
# Phase V
FROM node:latest as access-tests
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/ .
RUN [“npm”, “access-tests”]
# Starts and Serves Web Page
# Phase VI
FROM node:latest as serve
WORKDIR /usr/src/app
COPY --from=builder /usr/src/app/dest ./
COPY --from=builder /usr/src/app/package* ./
RUN [“npm”, “start”]