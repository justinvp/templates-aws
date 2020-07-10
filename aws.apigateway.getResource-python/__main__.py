import pulumi
import pulumi_aws as aws

my_rest_api = aws.apigateway.get_rest_api(name="my-rest-api")
my_resource = aws.apigateway.get_resource(path="/endpoint/path",
    rest_api_id=my_rest_api.id)

