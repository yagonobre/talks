go get -u -v github.com/kubernetes-sigs/aws-iam-authenticator/cmd/aws-iam-authenticator
mkdir /var/aws-iam-authenticator
mkdir /etc/kubernetes/aws-iam-authenticator
sudo /home/admin/go/bin/aws-iam-authenticator -i meetup
kubectl apply -f aws-iam-auth.yaml
