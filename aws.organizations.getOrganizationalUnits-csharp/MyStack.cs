using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var org = Output.Create(Aws.Organizations.GetOrganization.InvokeAsync());
        var ou = org.Apply(org => Output.Create(Aws.Organizations.GetOrganizationalUnits.InvokeAsync(new Aws.Organizations.GetOrganizationalUnitsArgs
        {
            ParentId = org.Roots[0].Id,
        })));
    }

}

