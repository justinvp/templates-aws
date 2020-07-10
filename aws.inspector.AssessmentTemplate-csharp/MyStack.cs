using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var example = new Aws.Inspector.AssessmentTemplate("example", new Aws.Inspector.AssessmentTemplateArgs
        {
            Duration = 3600,
            RulesPackageArns = 
            {
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-9hgA516p",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-H5hpSawc",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-JJOtZiqQ",
                "arn:aws:inspector:us-west-2:758058086616:rulespackage/0-vg5GGHSD",
            },
            TargetArn = aws_inspector_assessment_target.Example.Arn,
        });
    }

}

