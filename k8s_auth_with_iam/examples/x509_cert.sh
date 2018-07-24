openssl genrsa -out yago.pem 2048
openssl req -new -key yago.pem -out yago-csr.pem -subj "/CN=yago/O=app1/O=app2"
openssl x509 -req -days 360 -in yago-csr.pem \
	-CA /etc/kubernetes/pki/ca.crt \
	-CAkey /etc/kubernetes/pki/ca.key \
	-CAcreateserial -out yago.crt -sha256

