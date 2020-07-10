import * as pulumi from "@pulumi/pulumi";
import * as aws from "@pulumi/aws";

const selectedRestApi = pulumi.output(aws.apigateway.getRestApi({
    name: var_api_gateway_name,
}, { async: true }));
const selectedUserPools = pulumi.output(aws.cognito.getUserPools({
    name: var_cognito_user_pool_name,
}, { async: true }));
const cognito = new aws.apigateway.Authorizer("cognito", {
    providerArns: selectedUserPools.arns,
    restApi: selectedRestApi.id,
    type: "COGNITO_USER_POOLS",
});

