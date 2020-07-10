import pulumi
import pulumi_aws as aws

# Create an AMI that will start a machine whose root device is backed by
# an EBS volume populated from a snapshot. It is assumed that such a snapshot
# already exists with the id "snap-xxxxxxxx".
example = aws.ec2.Ami("example",
    ebs_block_devices=[{
        "device_name": "/dev/xvda",
        "snapshot_id": "snap-xxxxxxxx",
        "volume_size": 8,
    }],
    root_device_name="/dev/xvda",
    virtualization_type="hvm")

