import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const example = new aws.appsync.GraphQLApi("example", {
    authenticationType: "AMAZON_COGNITO_USER_POOLS",
    userPoolConfig: {
        awsRegion: aws_region_current.name,
        defaultAction: "DENY",
        userPoolId: aws_cognito_user_pool_example.id,
    },
});

