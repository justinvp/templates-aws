import pulumi
import pulumi_aws as aws

service_image = aws.ecr.get_image(image_tag="latest",
    repository_name="my/service")

