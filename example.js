import rsa from 'k6/x/rsa';

export default function () {

    // Example
    let privateKeyPEMFormat = `-----BEGIN PRIVATE KEY-----
        <your private key here>
        -----END PRIVATE KEY-----`

    // Careful of seperator differences on windows/linux. Remember double \\ for windows.
    let pathToPrivateKeyPEMFormat = "C:\\path\\to\\your\\rsakey.pem"

    // No real type assertion is done in the backend Go, so unix timestamps can behave oddly. Convert to seconds.
    let threeMinutes = 180
    let claims = {
        "sub": "1234567890",
        "name": "John Doe",
        "admin": true,
        "iat": Math.floor(Date.now() / 1000),
        "exp": Math.floor(Date.now() / 1000 + threeMinutes),
    }

    try {
        // Be wary, console.log can append newline characters based on terminal width.
        let signedJwtFromString = rsa.signFromString(claims, privateKeyPEMFormat)
        console.log(`signedJwtFromString: ${signedJwtFromString}`)

        let signedJwtFromPath = rsa.signFromFilePath(claims, pathToPrivateKeyPEMFormat)
        console.log(`signedJwtFromPath: ${signedJwtFromPath}`)
    } catch(error) {
        console.log(error.toString())
    }
}
