import pulumi
import pulumi_aws as aws

main_vpc = aws.ec2.Vpc("mainVpc", cidr_block="10.0.0.0/16")
private_a = aws.ec2.Subnet("private-a",
    availability_zone="us-east-1a",
    cidr_block="10.0.0.0/24",
    vpc_id=main_vpc.id)
private_b = aws.ec2.Subnet("private-b",
    availability_zone="us-east-1b",
    cidr_block="10.0.1.0/24",
    vpc_id=main_vpc.id)
main_directory = aws.directoryservice.Directory("mainDirectory",
    password="#S1ncerely",
    size="Small",
    vpc_settings={
        "subnet_ids": [
            private_a.id,
            private_b.id,
        ],
        "vpc_id": main_vpc.id,
    })
main_workspaces_directory_directory = aws.workspaces.Directory("mainWorkspaces/directoryDirectory",
    directory_id=main_directory.id,
    self_service_permissions={
        "increaseVolumeSize": True,
        "rebuildWorkspace": True,
    })

