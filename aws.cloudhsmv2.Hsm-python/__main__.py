import pulumi
import pulumi_aws as aws

cluster = aws.cloudhsmv2.get_cluster(cluster_id=var["cloudhsm_cluster_id"])
cloudhsm_v2_hsm = aws.cloudhsmv2.Hsm("cloudhsmV2Hsm",
    cluster_id=cluster.cluster_id,
    subnet_id=cluster.subnet_ids[0])

