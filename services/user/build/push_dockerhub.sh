# docker buildしてdocker hubにpushするだけ
docker build -t takuyaatsugi9924/go-pf_user:v1 -f Dockerfile ../
docker push takuyaatsugi9924/go-pf_user:v1