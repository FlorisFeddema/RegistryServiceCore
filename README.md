# CoreService

## Setup development environment



``` shell 
docker-compose up

docker-compose down -v

sh setup.sh
 ```

## API Setup

/api/ -> Generic API

/v2/ -> Registry

/v2/_catalog -> gets catalog

/v2/cache/<>/tags/list -> docker images without registry

/v2/cache/<>/<>/tags/list -> docker images with registry

/v2/<>/<>/tags/list -> images with registry

registry name cannot be:
- blobs
- manifests
- tags
- _catalog
- cache


## Resources
- https://florisfeddema.atlassian.net/secure/RapidBoard.jspa?rapidView=11&projectKey=RC&selectedIssue=RC-27
- https://github.com/RegistryProxy/CoreService
- https://docs.docker.com/registry/spec/api/#detail
- https://drive.google.com/drive/u/0/folders/1QMxo0gMCh8LOyluLfp1IUthm6lVuHI05
- https://app.diagrams.net/#G1qABI3FnopF_MHzjZEH0F9-L3KxMdoHIT
- https://github.com/emirpasic/gods
- https://gorm.io/docs/index.html
- https://sonar.cloud.feddema.dev/dashboard?id=CoreService