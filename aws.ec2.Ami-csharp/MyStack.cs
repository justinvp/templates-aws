using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        // Create an AMI that will start a machine whose root device is backed by
        // an EBS volume populated from a snapshot. It is assumed that such a snapshot
        // already exists with the id "snap-xxxxxxxx".
        var example = new Aws.Ec2.Ami("example", new Aws.Ec2.AmiArgs
        {
            EbsBlockDevices = 
            {
                new Aws.Ec2.Inputs.AmiEbsBlockDeviceArgs
                {
                    DeviceName = "/dev/xvda",
                    SnapshotId = "snap-xxxxxxxx",
                    VolumeSize = 8,
                },
            },
            RootDeviceName = "/dev/xvda",
            VirtualizationType = "hvm",
        });
    }

}

