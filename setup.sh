#rm ~/.docker/config.json
docker login https://localhost:5001 -u HerMBujLXjt7PhUcENjmyGT8PJmdCYcY -p K6WSUNxFELBM2WZRF86Lx72g5FREUrym
docker pull nginx:1.19
docker pull nginx:1.18
docker pull busybox:1.32
docker pull busybox:1.33
docker pull memcached:1.6
docker pull alpine:3.13
docker pull alpine:3.12
docker tag nginx:1.19 localhost:5001/nginx:1.19
docker tag nginx:1.18 localhost:5001/nginx:1.18
docker tag busybox:1.32 localhost:5001/busybox:1.32
docker tag busybox:1.33 localhost:5001/busybox:1.33
docker tag memcached:1.6 localhost:5001/cached/memcached:1.6
docker tag alpine:3.13 localhost:5001/cached/test/alpine:3.13
docker tag alpine:3.12 localhost:5001/cached/test/alpine:3.12
docker push localhost:5001/nginx:1.19
docker push localhost:5001/nginx:1.18
docker push localhost:5001/busybox:1.32
docker push localhost:5001/busybox:1.33
docker push localhost:5001/cached/memcached:1.6
docker push localhost:5001/cached/test/alpine:3.13
docker push localhost:5001/cached/test/alpine:3.12

