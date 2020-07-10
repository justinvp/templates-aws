import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const sizeConstraintSet = new aws.waf.SizeConstraintSet("size_constraint_set", {
    sizeConstraints: [{
        comparisonOperator: "EQ",
        fieldToMatch: {
            type: "BODY",
        },
        size: 4096,
        textTransformation: "NONE",
    }],
});

