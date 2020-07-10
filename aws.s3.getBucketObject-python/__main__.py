import pulumi
import pulumi_aws as aws

bootstrap_script = aws.s3.get_bucket_object(bucket="ourcorp-deploy-config",
    key="ec2-bootstrap-script.sh")
example = aws.ec2.Instance("example",
    ami="ami-2757f631",
    instance_type="t2.micro",
    user_data=bootstrap_script.body)

