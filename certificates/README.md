# TLS Client Authentication

Using TLS certificates to authenticate server and client

[This guide](https://www.makethenmakeinstall.com/2014/05/ssl-client-authentication-step-by-step/) was used.

For some more information on cert generation, look [here](https://gist.github.com/jafossum/638587074ad0e187f147882fa88e23e2), or Google it :) 

## Generate a certificate authority (CA) cert

If you do not have a CA cert, you can generate one using OpenSSL.
Generate your CA certificate using this command:

    $ openssl req -newkey rsa:4096 -keyform PEM -keyout ca.key -x509 -days 3650 -outform PEM -out ca.cer

NB: It is important the the CommonName (CN) parameter for the CA matches your domain. For this test `localhost` will do.

>When connecting to an HTTPS server, browsers will check the CN value and it should be conforming to the domain. Wildcard certificates usually start with a `*` in CN to allow any subdomain. e.g. `CN=*.example.com`

> Note that browsers will reject the wilcard for the naked domain, i.e. `example.com` is not conforming to `*.example.com`

Then keep them secret – keep them safe. If someone were to get a hold of these files they would be able to generate server and client certs that would be trusted by our web server as it will be configured below.

## Generate your Server SSL key and certificate

Now that we have our CA cert, we can generate the SSL certificate that will be used by the Server

1. Generate a server private key:

    > $ openssl genrsa -out server.key 4096

2. Use the server private key to generate a certificate generation request:

    > $ openssl req -new -key server.key -out server.req -sha256

3. Use the certificate generation request and the CA cert to generate the server cert:

    > $ openssl x509 -req -in server.req -CA ca.cer -CAkey ca.key -set_serial 100 -extensions server -days 1460 -outform PEM -out server.cer -sha256

4. Clean up – now that the cert has been created, we no longer need the request:

    > $ rm server.req

## Generate a client SSL certificate

1. Generate a private key for the SSL client

    > $ openssl genrsa -out client.key 4096

2. Use the client’s private key to generate a cert request

    > $ openssl req -new -key client.key -out client.req

3. Issue the client certificate using the cert request and the CA cert/key.

    > & openssl x509 -req -in client.req -CA ca.cer -CAkey ca.key -set_serial 101 -extensions client -days 365 -outform PEM -out client.cer

4. (Optional) Convert the client certificate and private key to pkcs#12 format for use by browsers.

    > & openssl pkcs12 -export -inkey client.key -in client.cer -out client.p12

5. Clean up – now that the cert has been created, we no longer need the request.

    > $ rm client.req
