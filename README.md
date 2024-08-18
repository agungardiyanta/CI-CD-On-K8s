# CI-CD-On-K8s

* __CheckMySite__
This service checks if a domain is up or down by sending an HTTP request. Access it via:
```
http://url/checkmysite/{domain}
```
* __CheckMyIP__
This service detects the IP address of the client making the request. Access it via:
```
http://url/checkmyip/
```

### CI/CD Flow
This project uses Github Actions for the CI/CD platform, divided into several workflows/pipelines:
1. Unit Test\
The unit test workflow is triggered when a new Pull Request is made to the main branch and there are changes in specific paths for each service:
```
checkout > install dependencies > run unit test
```
2. Deployment\
The deployment workflow is triggered when there are changes or a push is merged into the main branch in specific paths for each service:
```
stage build
checkout > build docker image > push docker image
stage deploy
checkout > get GCP credentials > Set up GKE kubectl > deploy new image to GKE
```
The project uses Docker Hub's public repository to publish the built images with the Github SHA as the image tag.

### Access Service
The project uses the [ingress nginx controller](https://kubernetes.github.io/ingress-nginx/deploy/#quick-start) community to expose the services, which can be accessed through a public IP deployed using Helm.