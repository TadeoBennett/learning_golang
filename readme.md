# Finding the go file that generates certifications 

## IN WINDOWS
go run A:\Go\src\crypto\tls\generate_cert.go --rsa-bits=2048 --host=localhost

--note that this is relative to my go installation directory


## IN UBUNTU
go run /usr/local/go/src/crypto//tls/generate_cert.go --rsa-bits=2048 --host=localhost


### Certifcate Formats (PEM)
PEM - Privacy Enhanced Mail


### Note
After accessing the page via https, you should see a warning. Click on advanced options and find the page link to proceed. 
Then, see the Page Information using Ctrl + I and going in the security tab. It tells you about the certificate