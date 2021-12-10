package azure

import (
	"github.com/infracost/infracost/internal/config"
	"github.com/infracost/infracost/internal/schema"
	"github.com/tidwall/gjson"
)

func GetAzureRMHDInsightSparkClusterRegistryItem() *schema.RegistryItem {
	return &schema.RegistryItem{
		Name:  "azurerm_hdinsight_spark_cluster", //nolint:misspell
		RFunc: NewAzureRMHDInsightSparkCluster,
	}
}

func NewAzureRMHDInsightSparkCluster(ctx *config.ProjectContext, d *schema.ResourceData, u *schema.UsageData) *schema.Resource {
	region := lookupRegion(d, []string{})

	costComponents := []*schema.CostComponent{}

	headNodeVM := d.Get("roles.0.head_node.0.vm_size").String()
	workerNodeVM := d.Get("roles.0.worker_node.0.vm_size").String()
	var workerInstances int64
	if d.Get("roles.0.worker_node.0.target_instance_count").Type != gjson.Null {
		workerInstances = d.Get("roles.0.worker_node.0.target_instance_count").Int()
	}
	zookeeperNodeVM := d.Get("roles.0.zookeeper_node.0.vm_size").String()

	costComponents = append(costComponents, hdInsightVMCostComponent(region, "Head", headNodeVM, 2))
	costComponents = append(costComponents, hdInsightVMCostComponent(region, "Worker", workerNodeVM, workerInstances))
	costComponents = append(costComponents, hdInsightVMCostComponent(region, "Zookeeper", zookeeperNodeVM, 3))

	return &schema.Resource{
		Name:           d.Address,
		CostComponents: costComponents,
	}
}
