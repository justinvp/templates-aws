using Pulumi;
using Aws = Pulumi.Aws;

class MyStack : Stack
{
    public MyStack()
    {
        var geoMatchSet = new Aws.Waf.GeoMatchSet("geoMatchSet", new Aws.Waf.GeoMatchSetArgs
        {
            GeoMatchConstraints = 
            {
                new Aws.Waf.Inputs.GeoMatchSetGeoMatchConstraintArgs
                {
                    Type = "Country",
                    Value = "US",
                },
                new Aws.Waf.Inputs.GeoMatchSetGeoMatchConstraintArgs
                {
                    Type = "Country",
                    Value = "CA",
                },
            },
        });
    }

}

