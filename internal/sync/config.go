package sync

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/service/configservice"
	configTypes "github.com/aws/aws-sdk-go-v2/service/configservice/types"
	"log"
)

func (c *clients) getResourcesConfig(resourceIdentifiers []configTypes.AggregateResourceIdentifier) ([]configTypes.BaseConfigurationItem, error) {
	aggregatorName := "callsign-sub-accounts"
	var results []configTypes.BaseConfigurationItem

	batch := 100
	for i := 0; i < len(resourceIdentifiers); i += 100 {
		j := i + batch
		if j > len(resourceIdentifiers) {
			j = len(resourceIdentifiers)
		}

		input := &configservice.BatchGetAggregateResourceConfigInput{
			ConfigurationAggregatorName: &aggregatorName,
			ResourceIdentifiers:         resourceIdentifiers[i:j],
		}

		r, err := c.ConfigClient.BatchGetAggregateResourceConfig(context.TODO(), input)
		if err != nil {
			log.Fatalf("Error BatchGetAggregateResourceConfig: %v\n", err)
			return nil, err
		}
		results = append(results, r.BaseConfigurationItems...)
	}

	return results, nil
}


func (c *clients) getDiscoveredResources() ([]configTypes.AggregateResourceIdentifier, error) {
	aggregatorName := "callsign-sub-accounts"
	var results []configTypes.AggregateResourceIdentifier

	// List of resource types pulled from
	// github.com/sync/sync-sdk-go/models/apis/config/2014-11-12/api-2.json
	var resourceTypes = [...]configTypes.ResourceType{
		"AWS::EC2::Instance",
		"AWS::RDS::DBInstance",
		"AWS::EKS::Cluster",
		"AWS::ElastiCache::CacheCluster",
		"AWS::Elasticsearch::Domain",
		"AWS::MSK::Cluster",
	}

	for _, t := range &resourceTypes {
		input := &configservice.ListAggregateDiscoveredResourcesInput{
			ResourceType: t,
			ConfigurationAggregatorName: &aggregatorName,
		}

		result, err := c.ConfigClient.ListAggregateDiscoveredResources(context.TODO(), input)
		if err != nil {
			log.Fatalf("Error ListAggregateDiscoveredResources (ResourceType: %s): %v\n", t, err)
			return nil, err
		}
		results = append(results, result.ResourceIdentifiers...)
	}


	return results, nil
}