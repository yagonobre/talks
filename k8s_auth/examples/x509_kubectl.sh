kubectl config set-credentials yago-x509 \
	--client-certificate yago.crt \
	--client-key yago.pem
kubectl config set-context yago-x509 \
	--cluster kubernetes \
	--user yago-x509
