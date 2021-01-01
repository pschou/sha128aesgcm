# Proposed AES-128-GCM method for securing certificates
The GoLANG function operates on any arbitrary binary block and can be incorporated into your application so as to enable 128BitAES-GCM encoding with a passphrase.  The intended usage is for key storage, backups, and use in real time applications.

This function would layer in between the PEM decode and the DER/BER parse.

Both the GoLANG function and command line utility are provided for ease of use.

This encoding scheme should become a good option for encoding of p12 certificates, ie: pbeWithSHAAnd128BitAES-GCM.

# Your own app
The included example.go shows how to include this into your own library.

# Command line utility
Secure your PEM encoded files for storage, transferring, or usage.

```
$ secure-pem orig.pem test > encoded.pem
$ head encoded.pem
-----BEGIN RSA PRIVATE KEY-----
Proc-Type: 4,ENCRYPTED
DEK-Info: AES-128-GCM,4B013446431C9A10DD2DDF06

7unNI7DHqFAGkOo5jA3ygMSZoM0wPNjUeu6eFxH1nKpnFR3Se5acaax2/U2kaxMY
PR4vFuvo593qTBM9emBQ4sN1waDgODleqLdIqpnaun72RnIRiaVzikJOO152/DJE
...
```

and to reverse the hash, apply use the same passphrase,
```
$ secure-pem encoded.pem test
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCuTH8zQlQjCn3LC8F11+1jL+T2WWaXypxn/YvJ00Fym4lH2rpB
fnUKSBU6LWvJFzExLCTEircBvZWQEbjm/VgfB+djo025+ljC2a/DMhXTE9AdGB+K
VqOKSvsSRSzM3l9sEM40oHetdnb2+SzlVbX5ZQ1vh1WpopzYvh5ow2ZByQIDAQAB
...
```

and you'll get back your original key.

Password notice:  As with any encryption method, the longer the passphrase the better; use what you can remember but don't make it obviously simple like repeated numbers or commonly used phrases, but in the end length matters so make it long.
