import pulumi
import pulumi_aws as aws

# Map public IP on launch must be enabled for public (Internet accessible) subnets
example_subnet = aws.ec2.Subnet("exampleSubnet", map_public_ip_on_launch=True)
example_cluster = aws.emr.Cluster("exampleCluster",
    core_instance_group={},
    ec2_attributes={
        "subnet_id": example_subnet.id,
    },
    master_instance_group={
        "instance_count": 3,
    },
    release_label="emr-5.24.1",
    termination_protection=True)

