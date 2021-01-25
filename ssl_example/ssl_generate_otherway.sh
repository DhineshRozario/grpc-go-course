#!/bin/bash

################################################################################

# Output Files:
## ca.key - Certificate Authority private key file (This should not be shared in real-life)
## ca.crt - Certificate Authority trust certificate (this should be shared with users in real-life) 

## server.key - Server Private Key, password protected (this should not be shared)
## server.csr - Server Certificate Signing request (this should be shared with the CA owner)
## server.crt - Server Certificate Signed by the CA (this would be sent back by the CA Owner) - keep on server
## server.pem - Conversion of server.key into a format gRPC liles (this should not be shared)

#### SUMMARY ####
## Private Files : ca.key, server.key, server.pem, server.crt
## Public  Files : ca.crt (needed by the client), sever.csr (needed by the CA)
################################################################################


export SERVER_CN=localhost

# Step 1: Generate Certificate Authority + Truxst Certificate ca.crt
openssl genrsa -passout pass:1111 -des3 -out ca.key 4096

#Before running this command run
export MSYS_NO_PATHCONV=1 # from link : https://github.com/openssl/openssl/issues/8795
openssl req -passin pass:1111 -new -x509 --days 1825  -key ca.key -out ca.crt -subj "/CN=${SERVER_CN}"

# Step 2: Generate the Server Private Key (server.key)
openssl genrsa -passout pass:1111 -des3 -out server.key 4096

# Step 3: Get a certificate signing request from the CA (server.csr)
openssl req -passout pass:1111 -new -key server.key -out server.csr -subj "/CN=${SERVER_CN}"

#Step 4: Sign the certificat with the CA we are created (its called sekf signing) - server.crt
openssl x509 -req -passin pass:1111 -days 365 -in server.csr -CA ca.crt -CAkey ca.key -set_serial 01 -out server.crt

#Step 5: Convert the server certificate to .pem format (server.pem) - usable by gRPC
openssl pkcs8 -topk8 -nocrypt -passin pass:1111 -in server.key -out server.pem
