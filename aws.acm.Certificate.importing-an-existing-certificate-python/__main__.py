import pulumi
import pulumi_aws as aws
import pulumi_tls as tls

example_private_key = tls.PrivateKey("examplePrivateKey", algorithm="RSA")
example_self_signed_cert = tls.SelfSignedCert("exampleSelfSignedCert",
    allowed_uses=[
        "key_encipherment",
        "digital_signature",
        "server_auth",
    ],
    key_algorithm="RSA",
    private_key_pem=example_private_key.private_key_pem,
    subjects=[{
        "commonName": "example.com",
        "organization": "ACME Examples, Inc",
    }],
    validity_period_hours=12)
cert = aws.acm.Certificate("cert",
    certificate_body=example_self_signed_cert.cert_pem,
    private_key=example_private_key.private_key_pem)

