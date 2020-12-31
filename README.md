# Salt256 - Securing binary block with a salted passphrase
The GoLANG function Salt256 operates on any arbitrary binary block and can be incorporated into your application so as to enable loading of salted SHA256 binary blobs.  The intended usage is for key storage, backups, and use in real time applications.  This function would layer inbetween the PEM decode and the DER/BER parse.

Both the GoLANG function and command line utility are provided for ease of use.

# Command line utility
Secure your PEM encoded files with a Salted Password SHA256 Hash, mainly for storage, transferring, or usage

# Basic usage
```
$ ./salt256pem orig.pem test > encoded.pem
$ head encoded.pem
-----BEGIN SALT256 RSA PRIVATE KEY-----
Salt: AzQTKUqEwhSPpAdS6AOI

5lPVcUEV7czROHRqOwWTZwM11tGrPpwsmCF+wtgsei7k75sVu4OMdJ/y5khkCp5T
X0WRirW1+3T3URkQHu/kjLUd8G0pZQfNLCUAy+/pCvMpAe3L703OrskHWOBELIor
2iEUb0hSumaX9mERqFawlnjCo5ELK6+nH8yiAbQHopzR5DjwhzDMLlSc3/4DWccL
AWSEd30W+7bAgC6ABhXvZDi/mbYYWI7xiK7xsmrZmjrFEBU5KAggnK7oiQNhE2UL
...
```

and to reverse the hash, apply use the same passphrase,
```
$ ./salt256pem encoded.pem test | head
-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQCuTH8zQlQjCn3LC8F11+1jL+T2WWaXypxn/YvJ00Fym4lH2rpB
fnUKSBU6LWvJFzExLCTEircBvZWQEbjm/VgfB+djo025+ljC2a/DMhXTE9AdGB+K
VqOKSvsSRSzM3l9sEM40oHetdnb2+SzlVbX5ZQ1vh1WpopzYvh5ow2ZByQIDAQAB
AoGAdHba+9ZTQALPoTAeue5pScCD/nQoUD7rGbVahaZqcfkp1uB5Wi/0ZsVZocEd
54VeslurycRimoAigtzEr1wTV7ZGfpmlKVpp6r5tYtFqOA2IjnYXsqspia54OBN1
...
```

and you'll get back your original key.  Any other passphrase will not result in a valid key.
