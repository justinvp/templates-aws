using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var bootstrapScript = Output.Create(Aws.S3.GetBucketObject.InvokeAsync(new Aws.S3.GetBucketObjectArgs
        {
            Bucket = "ourcorp-deploy-config",
            Key = "ec2-bootstrap-script.sh",
        }));
        var example = new Aws.Ec2.Instance("example", new Aws.Ec2.InstanceArgs
        {
            Ami = "ami-2757f631",
            InstanceType = "t2.micro",
            UserData = bootstrapScript.Apply(bootstrapScript => bootstrapScript.Body),
        });
    }

}

