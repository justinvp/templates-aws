import pulumi
import pulumi_aws as aws

# Create a new certificate
test = aws.dms.Certificate("test",
    certificate_id="test-dms-certificate-tf",
    certificate_pem="...")

