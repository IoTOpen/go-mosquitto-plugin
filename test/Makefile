all: server-certs client-certs

server-certs:
	CAROOT=`pwd`/ca mkcert -cert-file server/server.crt -key-file server/server.key localhost
client-certs:
	CAROOT=`pwd`/ca mkcert -cert-file client/client.crt -key-file client/client.key --client localhost 
