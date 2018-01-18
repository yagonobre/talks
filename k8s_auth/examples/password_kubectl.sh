kubectl config set-credentials yago-password \
	--username admin --password admin
kubectl config set-context yago-password \
	--cluster kubernetes \
	--user yago-password

