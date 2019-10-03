username=$1
principals=$2
rm -rf $username-id_rsa $username-id_rsa-cert.pub
ssh-keygen -f $username-id_rsa
ssh-keygen -s users_ca -I $username -n $principals $username-id_rsa.pub
rm $username-id_rsa.pub 
