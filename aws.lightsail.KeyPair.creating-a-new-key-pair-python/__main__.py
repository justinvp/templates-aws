import pulumi
import pulumi_aws as aws

# Create a new Lightsail Key Pair
lg_key_pair = aws.lightsail.KeyPair("lgKeyPair")

