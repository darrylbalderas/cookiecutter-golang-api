# cookiecutter-golang-api


# Setup

1. Install [pipx](https://github.com/pypa/pipx)  e.g. `brew imstall pipx`


1. Install cookiecutter with pipx e.g. `pipx install cookiecutter`

1. Verify you have cookiecutter installed `pipx run cookiecutter --version`

```
$: pipx run cookiecutter --version
Cookiecutter 2.5.0
```

1. Clone golang api cookie cutter template `git clone https://github.com/darrylbalderas/cookiecutter-golang-api`

1. Run cookiecutter to generate api folder `pipx run cookiecutter cookiecutter-golang-api/`

```
$: pipx run cookiecutter cookiecutter-golang-api/
  [1/4] project_name (): api
  [2/4] project_slug (api):
  [3/4] version (0.0.1):
  [4/4] golang_version (1.21):
```

1. Change directory generated cookiecutter template `cd api`

1. Execute makefile step Build docker image `make build`

```
$: make build
[+] Building 0.9s (15/15) FINISHED                                                                 docker:desktop-linux
 => [internal] load build definition from Dockerfile                                                               0.0s
 => => transferring dockerfile: 574B                                                                               0.0s
 => [internal] load .dockerignore                                                                                  0.0s
 => => transferring context: 2B                                                                                    0.0s
 => [internal] load metadata for gcr.io/distroless/static-debian11:latest                                          0.5s
 => [internal] load metadata for docker.io/library/golang:1.21                                                     0.9s
 => [builder 1/6] FROM docker.io/library/golang:1.21@sha256:549dd88a1a53715f177b41ab5fee25f7a376a6bb5322ac7abe263  0.0s
 => [internal] load build context                                                                                  0.0s
 => => transferring context: 3.77kB                                                                                0.0s
 => [stage-1 1/3] FROM gcr.io/distroless/static-debian11@sha256:a43abc840a7168c833a8b3e4eae0f715f7532111c9227ba17  0.0s
 => CACHED [stage-1 2/3] WORKDIR /app                                                                              0.0s
 => CACHED [builder 2/6] WORKDIR /app                                                                              0.0s
 => CACHED [builder 3/6] COPY go.mod ./                                                                            0.0s
 => CACHED [builder 4/6] RUN go mod download                                                                       0.0s
 => CACHED [builder 5/6] COPY . .                                                                                  0.0s
 => CACHED [builder 6/6] RUN CGO_ENABLED=0 GOOS=linux go build -o main .                                           0.0s
 => CACHED [stage-1 3/3] COPY --from=builder /app/main .                                                           0.0s
 => exporting to image                                                                                             0.0s
 => => exporting layers                                                                                            0.0s
 => => writing image sha256:81c67e436c527e71a23183b479ae77b788fb65ddbe44b8b9a2b289236d044ca1                       0.0s
 => => naming to docker.io/library/api:0.0.1                                                                       0.0s

What's Next?
  View summary of image vulnerabilities and recommendations → docker scout quickview
```

1. Execute makefile step to start server `make run`

```
$: make run
2024/02/19 00:22:52 Server is running on http://0.0.0.0:8000
```

1. In a different terminal window you can execute curl, to verify api server is running and responding
```
$: curl http://0.0.0.0:8000
{"message":"Hello world"}%
```