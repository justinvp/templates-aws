import pulumi
import pulumi_aws as aws

config = pulumi.Config()
vpc = config.require_object("vpc")
domain = config.get("domain")
if domain is None:
    domain = "tf-test"
selected_vpc = aws.ec2.get_vpc(tags={
    "Name": vpc,
})
selected_subnet_ids = aws.ec2.get_subnet_ids(tags={
        "Tier": "private",
    },
    vpc_id=selected_vpc.id)
current_region = aws.get_region()
current_caller_identity = aws.get_caller_identity()
es_security_group = aws.ec2.SecurityGroup("esSecurityGroup",
    description="Managed by Pulumi",
    ingress=[{
        "cidr_blocks": [selected_vpc.cidr_block],
        "from_port": 443,
        "protocol": "tcp",
        "to_port": 443,
    }],
    vpc_id=selected_vpc.id)
es_service_linked_role = aws.iam.ServiceLinkedRole("esServiceLinkedRole", aws_service_name="es.amazonaws.com")
es_domain = aws.elasticsearch.Domain("esDomain",
    access_policies=f"""{{
	"Version": "2012-10-17",
	"Statement": [
		{{
			"Action": "es:*",
			"Principal": "*",
			"Effect": "Allow",
			"Resource": "arn:aws:es:{current_region.name}:{current_caller_identity.account_id}:domain/{domain}/*"
		}}
	]
}}

""",
    advanced_options={
        "rest.action.multi.allow_explicit_index": "true",
    },
    cluster_config={
        "cluster_config": "m4.large.elasticsearch",
    },
    elasticsearch_version="6.3",
    snapshot_options={
        "snapshot_options": 23,
    },
    tags={
        "Domain": "TestDomain",
    },
    vpc_options={
        "security_group_ids": [es_security_group.id],
        "subnet_ids": [
            selected_subnet_ids.ids[0],
            selected_subnet_ids.ids[1],
        ],
    },
    opts=ResourceOptions(depends_on=["aws_iam_service_linked_role.es"]))

