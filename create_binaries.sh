#!/bin/bash -e
mkdir -p ../binaries
GOOS=linux GOARCH=amd64 go build -o ../binaries/check-shib3idp-login-linuxamd64  
GOOS=linux GOARCH=386 go build -o ../binaries/check-shib3idp-login-linux386  
GOOS=windows GOARCH=386 go build -o ../binaries/check-shib3idp-login-windows386  
GOOS=windows GOARCH=amd64 go build -o ../binaries/check-shib3idp-login-windowsamd64  
GOOS=darwin GOARCH=amd64 go build -o ../binaries/check-shib3idp-login-darwinamd64  
GOOS=darwin GOARCH=386 go build -o ../binaries/check-shib3idp-login-darwin386  
cd ../binaries
for i in check-shib3idp-login-*; do sha512sum $i > $i.sha512; done
