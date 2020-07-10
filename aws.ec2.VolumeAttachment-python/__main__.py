import pulumi
import pulumi_aws as aws

web = aws.ec2.Instance("web",
    ami="ami-21f78e11",
    availability_zone="us-west-2a",
    instance_type="t1.micro",
    tags={
        "Name": "HelloWorld",
    })
example = aws.ebs.Volume("example",
    availability_zone="us-west-2a",
    size=1)
ebs_att = aws.ec2.VolumeAttachment("ebsAtt",
    device_name="/dev/sdh",
    instance_id=web.id,
    volume_id=example.id)

