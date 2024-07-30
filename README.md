# xk6-rsa

An incredibly simple k6 extension for signing a collection of claims with a provided private key (PKCS8 format) using RS256.

### Example

Look within example.js for an example of how to use the extension. You can pass a private key from string, or from a file path.

### How To Add The Extension To Your k6 Binary.

Run the standard xk6 build command using the with flag.

```
xk6 build --with github.com/jmkraut/xk6-rsa
```

### How To Compile The Custom k6 Binary Locally

Ensure you are within the xk6-rsa folder and then run the following:

```
xk6 build --with xk6-rsa=.
```
This will produce a k6 binary within the current folder. Ensure you use this binary and not an existing one on your path.

