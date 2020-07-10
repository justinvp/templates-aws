import pulumi
import pulumi_aws as aws

foo = aws.kms.get_key(key_id="arn:aws:kms:us-east-1:111122223333:key/1234abcd-12ab-34cd-56ef-1234567890ab")

