
### tag v1.0 - Simple Web App Template
  - simple server to listem on port 8080
  - creating a simple http handler to handle incoming request 
  - calling a remote http get from handler
  - creating packages and demonstrating function call

### tag v1.1 - Building a go Docker image

  - DOCKER_BUILDKIT=0 docker build --no-cache -t go-web-app  .
  - docker run --rm --name demo -p 8080:8080 demo