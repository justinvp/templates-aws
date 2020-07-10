using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var rules = Output.Create(Aws.Inspector.GetRulesPackages.InvokeAsync());
        // e.g. Use in aws_inspector_assessment_template
        var @group = new Aws.Inspector.ResourceGroup("group", new Aws.Inspector.ResourceGroupArgs
        {
            Tags = 
            {
                { "test", "test" },
            },
        });
        var assessmentAssessmentTarget = new Aws.Inspector.AssessmentTarget("assessmentAssessmentTarget", new Aws.Inspector.AssessmentTargetArgs
        {
            ResourceGroupArn = @group.Arn,
        });
        var assessmentAssessmentTemplate = new Aws.Inspector.AssessmentTemplate("assessmentAssessmentTemplate", new Aws.Inspector.AssessmentTemplateArgs
        {
            Duration = 60,
            RulesPackageArns = rules.Apply(rules => rules.Arns),
            TargetArn = assessmentAssessmentTarget.Arn,
        });
    }

}

