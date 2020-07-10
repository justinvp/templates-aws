import pulumi
import pulumi_aws as aws

example = aws.acm.get_certificate(domain="tf.example.com",
    key_types=["RSA_4096"])

