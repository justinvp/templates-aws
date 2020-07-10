using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Map public IP on launch must be enabled for public (Internet accessible) subnets
        var exampleSubnet = new Aws.Ec2.Subnet("exampleSubnet", new Aws.Ec2.SubnetArgs
        {
            MapPublicIpOnLaunch = true,
        });
        var exampleCluster = new Aws.Emr.Cluster("exampleCluster", new Aws.Emr.ClusterArgs
        {
            CoreInstanceGroup = ,
            Ec2Attributes = new Aws.Emr.Inputs.ClusterEc2AttributesArgs
            {
                SubnetId = exampleSubnet.Id,
            },
            MasterInstanceGroup = new Aws.Emr.Inputs.ClusterMasterInstanceGroupArgs
            {
                InstanceCount = 3,
            },
            ReleaseLabel = "emr-5.24.1",
            TerminationProtection = true,
        });
    }

}

