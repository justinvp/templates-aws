using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var exampleDetector = new Aws.GuardDuty.Detector("exampleDetector", new Aws.GuardDuty.DetectorArgs
        {
            Enable = true,
        });
        var exampleOrganizationConfiguration = new Aws.GuardDuty.OrganizationConfiguration("exampleOrganizationConfiguration", new Aws.GuardDuty.OrganizationConfigurationArgs
        {
            AutoEnable = true,
            DetectorId = exampleDetector.Id,
        });
    }

}

