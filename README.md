# xk6-rsa

An incredibly simple k6 extension for signing a collection of claims with a provided private key (PKCS8 format) using RS256.

### Example

Look within example.js for an example of how to use the extension.

### How To Compile The Custom k6 Binary

Ensure you are within the xk6-rsa folder and then run the following:

```
xk6 build --with xk6-rsa=.
```
This will produce a k6 binary within the current folder. Ensure you use this binary and not an existing one on your path.

