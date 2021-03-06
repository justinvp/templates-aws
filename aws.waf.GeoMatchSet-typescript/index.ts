import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const geoMatchSet = new aws.waf.GeoMatchSet("geo_match_set", {
    geoMatchConstraints: [
        {
            type: "Country",
            value: "US",
        },
        {
            type: "Country",
            value: "CA",
        },
    ],
});

