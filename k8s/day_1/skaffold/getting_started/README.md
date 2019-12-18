# SKAFFOLD SINGLE POD TEST
Deploy a single pod using skaffold.

Please update the k8s-pod.yaml and skaffold.yaml to match your registry and image name.
ex: k8s-pod.yaml
~~~~
 ...
 containers:
  - name: getting-started
    image: registry.ynot.work/skaffold/skaffold-example
~~~~


ex: skaffold.yaml
~~~~
 ...
build:
  artifacts:
  - image: registry.ynot.work/skaffold/skaffold-example
~~~~

Note the "registry.ynot.work"

Before running skaffold dev
run:
~~~
docker login registry.ynot.work -u admin -p admin
