using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var mainVpc = new Aws.Ec2.Vpc("mainVpc", new Aws.Ec2.VpcArgs
        {
            CidrBlock = "10.0.0.0/16",
        });
        var private_a = new Aws.Ec2.Subnet("private-a", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-east-1a",
            CidrBlock = "10.0.0.0/24",
            VpcId = mainVpc.Id,
        });
        var private_b = new Aws.Ec2.Subnet("private-b", new Aws.Ec2.SubnetArgs
        {
            AvailabilityZone = "us-east-1b",
            CidrBlock = "10.0.1.0/24",
            VpcId = mainVpc.Id,
        });
        var mainDirectory = new Aws.DirectoryService.Directory("mainDirectory", new Aws.DirectoryService.DirectoryArgs
        {
            Password = "#S1ncerely",
            Size = "Small",
            VpcSettings = new Aws.DirectoryService.Inputs.DirectoryVpcSettingsArgs
            {
                SubnetIds = 
                {
                    private_a.Id,
                    private_b.Id,
                },
                VpcId = mainVpc.Id,
            },
        });
        var mainWorkspaces_directoryDirectory = new Aws.Workspaces.Directory("mainWorkspaces/directoryDirectory", new Aws.Workspaces.DirectoryArgs
        {
            DirectoryId = mainDirectory.Id,
            SelfServicePermissions = new Aws.Workspaces.Inputs.DirectorySelfServicePermissionsArgs
            {
                IncreaseVolumeSize = true,
                RebuildWorkspace = true,
            },
        });
    }

}

