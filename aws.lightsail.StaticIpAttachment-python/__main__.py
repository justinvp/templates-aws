import pulumi
import pulumi_aws as aws

test_static_ip = aws.lightsail.StaticIp("testStaticIp")
test_instance = aws.lightsail.Instance("testInstance",
    availability_zone="us-east-1b",
    blueprint_id="string",
    bundle_id="string",
    key_pair_name="some_key_name")
test_static_ip_attachment = aws.lightsail.StaticIpAttachment("testStaticIpAttachment",
    instance_name=test_instance.id,
    static_ip_name=test_static_ip.id)

