kubectl create rolebinding bob-admin-binding \
	--clusterrole=admin --user=bob \
	--namespace=acme

kubectl create clusterrolebinding root-cluster-admin-binding \
	--clusterrole=cluster-admin --user=root
