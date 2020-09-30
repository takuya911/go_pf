kubectl delete -f ./conf
kubectl delete -f ./mysql
kubectl delete -f ./phpmyadmin
kubectl delete -f ./user
kubectl delete -f ./graphql

kubectl apply -f ./conf
kubectl apply -f ./mysql
kubectl apply -f ./phpmyadmin
kubectl apply -f ./user
kubectl apply -f ./graphql