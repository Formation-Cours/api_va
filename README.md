```bash
$ docker buildx build --platform linux/arm/v7,linux/arm64,linux/amd64,darwin/amd64,darwin/arm64 -t olprog/api_va:latest --push .
```