rm -rf admin-id_rsa admin-id_rsa-cert.pub
ssh-keygen -f admin-id_rsa
ssh-keygen -s users_ca -I admin -n admin admin-id_rsa.pub
rm admin-id_rsa.pub 
