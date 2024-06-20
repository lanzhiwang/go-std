```bash
$ docker pull docker-mirrors.alauda.cn/library/golang:1.22.4-bullseye

$ docker run -ti --rm \
-v ~/work/code/go_code/go-std:/go-std \
-w /go-std \
docker-mirrors.alauda.cn/library/golang:1.22.4-bullseye \
bash

GO111MODULE='on' GOPROXY='https://goproxy.cn,direct' go build -o main -v -work -x -a []

```