# Proposed pbeWithSHAAnd128BitAES-GCM method for securing certificates
The GoLANG function operates on any arbitrary binary block and can be incorporated into your application so as to enable 128BitAES-GCM encoding with a passphrase.  The intended usage is for key storage, backups, and use in real time applications.

This function would layer in between the PEM decode and the DER/BER parse.

Both the GoLANG function and command line utility are provided for ease of use.

# Command line utility
Secure your PEM encoded files with a pbeWithSHAAnd128BitAES-GCM cipher; for storage, transferring, or usage

# Basic usage
```
$ secure-pem orig.pem test > encoded.pem
$ head encoded.pem
-----BEGIN 128-AES-GCM ENCRYPTED RSA PRIVATE KEY-----
nonce: RKN6QyJZvjdnWF5m

BTG1uKDb9jJqFiVl1FI9/HQQMPclS9k/t4DP7794f5YTbOTugfn6ZQg5tqrCZXW1
NW1BFjq1kikZIQldF4xjvRZzZwEFnvk3fmxKXxaqaG4V/Dwb2UdD5pChzQwNlSxO
yk3epgCOb0tV7jgBcq9puKxK/tBn3e18H6cmHpvN0GJGhwSJtVoclJnD2jbWjcxI
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
