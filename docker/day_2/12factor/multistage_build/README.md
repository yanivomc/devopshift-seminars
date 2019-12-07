# MULTISTAGE BUILD 
#### Docker container pipelines 

The application we are using as a POC can be downloaded from:  https://github.com/yohangz/ornamentum

##### Running tests and validation in the applicaiton

~~~
- run TS lint
npm run lint:demo:ts

- run SCSS lint
npm run lint:demo:scss

- run test stute on watch mode
npm run test:demo

- run test suite and generate coverage report
npm run test:demo:coverage

- serve demo app on watch mode
npm run start

- serve demo app with server side rendering build
npm run start:ssr
~~~

We are going to build this application inside a docker container and run all of the above pipelines in it.

Dockerfile exmpale:
~~~
FROM node:10.15.0
WORKDIR /app
COPY dist/ornamentum-demo /app
EXPOSE 8080
CMD ["node", "server.js"]
~~~
