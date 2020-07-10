import pulumi
import pulumi_aws as aws

geo_match_set = aws.wafregional.GeoMatchSet("geoMatchSet", geo_match_constraints=[
    {
        "type": "Country",
        "value": "US",
    },
    {
        "type": "Country",
        "value": "CA",
    },
])

