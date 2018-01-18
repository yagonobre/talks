kubectl config set-credentials yago-token \
	--token MEUTOKEN
kubectl config set-context yago-password \
	--cluster kubernetes \
	--user yago-token

