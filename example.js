import rsa from 'k6/x/rsa';

export default function () {

    // Example
    let privateKeyPEMFormat = `-----BEGIN PRIVATE KEY-----
        <your private key here>
        -----END PRIVATE KEY-----`

    let claims = {
        "sub": "1234567890",
        "name": "John Doe",
        "admin": true,
        "iat": Date.now(),
        "exp": Date.now() + (180 * 1000),
    }

    // k6 will also throw an error if you don't use a try/catch
    try {
        let signedJwt = rsa.sign(claims, privateKeyPEMFormat);
        console.log(signedJwt)
    } catch(error) {
        console.log(error.toString())
    }
}
