### Example: Microservices Deploy With Skaffold

In this example:

* Deploy multiple applications with skaffold
* In development, only rebuild and redeploy the artifacts that have changed
* Deploy multiple applications outside the working directory

In the real world, Kubernetes deployments will consist of multiple applications that work together.
In this example, we'll walk through using skaffold to develop and deploy two applications, an exposed "web" frontend which calls an unexposed "app" backend.


From this directory, run

```bash
skaffold dev
```