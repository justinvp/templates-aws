using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var geoMatchSet = new Aws.WafRegional.GeoMatchSet("geoMatchSet", new Aws.WafRegional.GeoMatchSetArgs
        {
            GeoMatchConstraints = 
            {
                new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
                {
                    Type = "Country",
                    Value = "US",
                },
                new Aws.WafRegional.Inputs.GeoMatchSetGeoMatchConstraintArgs
                {
                    Type = "Country",
                    Value = "CA",
                },
            },
        });
    }

}

