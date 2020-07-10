import pulumi
import pulumi_aws as aws

# Request a spot instance at $0.03
cheap_worker = aws.ec2.SpotInstanceRequest("cheapWorker",
    ami="ami-1234",
    instance_type="c4.xlarge",
    spot_price="0.03",
    tags={
        "Name": "CheapWorker",
    })

