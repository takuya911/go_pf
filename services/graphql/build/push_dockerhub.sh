# docker buildしてdocker hubにpushするだけ
docker login -u takuyaatsugi9924 -p takuya9924
docker build -t takuyaatsugi9924/go-pf_graphql:v1 -f Dockerfile ../
docker push takuyaatsugi9924/go-pf_graphql:v1