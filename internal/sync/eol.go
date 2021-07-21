package sync

import (
	"github.com/crashdump/birdview/internal/model"
	"time"
)

// TODO: Is there an API I can consume to find this?

func (c *clients) GetServicesEOL() model.EOLs {
	return model.EOLs{
		/*
		RDS
		 */
		model.EOL{
			Type:    "AWS::RDS::DBInstance",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::RDS::DBInstance",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::RDS::DBInstance",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::RDS::DBInstance",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},

		/*
		Elasticache
		 */
		model.EOL{
			Type:    "AWS::ElastiCache::CacheCluster",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::ElastiCache::CacheCluster",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::ElastiCache::CacheCluster",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::ElastiCache::CacheCluster",
			Version: "",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},

		/*
		Elasticsearch
		 */
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "6.8",
			EOL:     time.Date(2020, time.October,20,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.1",
			EOL:     time.Date(2020, time.November,20,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.4",
			EOL:     time.Date(2021, time.April,20,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.7",
			EOL:     time.Date(2021, time.November,13,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.8",
			EOL:     time.Date(2021, time.December,18,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.9",
			EOL:     time.Date(2022, time.February,18,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::Elasticsearch::Domain",
			Version: "7.10",
			EOL:     time.Date(2022, time.March,18,0,0,0,0,time.UTC),
		},

		/*
		EKS
		 */
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.14",
			EOL:     time.Date(2020, time.July,3,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.15",
			EOL:     time.Date(2020, time.December,18,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.16",
			EOL:     time.Date(2021, time.July,25,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.17",
			EOL:     time.Date(2021, time.October,4,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.18",
			EOL:     time.Date(2021, time.November,1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.19",
			EOL:     time.Date(2022, time.April,1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.20",
			EOL:     time.Date(2022, time.June,1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::EKS::Cluster",
			Version: "1.21",
			EOL:     time.Date(2022, time.September,1,0,0,0,0,time.UTC),
		},


		// MSK
		model.EOL{
			Type:    "AWS::MSK::Cluster",
			Version: "1",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},
		model.EOL{
			Type:    "AWS::MSK::Cluster",
			Version: "2",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},		model.EOL{
			Type:    "AWS::MSK::Cluster",
			Version: "3",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},		model.EOL{
			Type:    "AWS::MSK::Cluster",
			Version: "4",
			EOL:     time.Date(0, time.January, 1,0,0,0,0,time.UTC),
		},

	}
}
