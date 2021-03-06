package v1alpha1

import (
	"testing"

	jsonTypes "github.com/appscode/go/encoding/json/types"
	"github.com/appscode/go/types"
	core "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestMongoDB_HostAddress(t *testing.T) {
	mongodb := &MongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "demo-name",
			Namespace: "demo",
			Labels: map[string]string{
				"app": "kubedb",
			},
		},
		Spec: MongoDBSpec{
			Version: jsonTypes.StrYo("3.6-v2"),
			ShardTopology: &MongoDBShardingTopology{
				Shard: MongoDBShardNode{
					Shards: 3,
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				ConfigServer: MongoDBConfigNode{
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				Mongos: MongoDBMongosNode{
					MongoDBNode: MongoDBNode{
						Replicas: 2,
					},
				},
			},
		},
	}

	shardDSN := mongodb.HostAddress()
	t.Log(shardDSN)

	mongodb.Spec.ShardTopology = nil
	mongodb.Spec.Replicas = types.Int32P(3)
	mongodb.Spec.ReplicaSet = &MongoDBReplicaSet{
		Name: "mgo-rs",
	}

	repsetDSN := mongodb.HostAddress()
	t.Log(repsetDSN)

}

func TestMongoDB_ShardDSN(t *testing.T) {
	mongodb := &MongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "demo-name",
			Namespace: "demo",
			Labels: map[string]string{
				"app": "kubedb",
			},
		},
		Spec: MongoDBSpec{
			Version: jsonTypes.StrYo("3.6-v2"),
			ShardTopology: &MongoDBShardingTopology{
				Shard: MongoDBShardNode{
					Shards: 3,
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				ConfigServer: MongoDBConfigNode{
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				Mongos: MongoDBMongosNode{
					MongoDBNode: MongoDBNode{
						Replicas: 2,
					},
				},
			},
		},
	}

	shardDSN := mongodb.ShardDSN(0)
	t.Log(shardDSN)

	mongodb.Spec.SetDefaults()
}

func TestMongoDB_ConfigSvrDSN(t *testing.T) {
	mongodb := &MongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "demo-name",
			Namespace: "demo",
			Labels: map[string]string{
				"app": "kubedb",
			},
		},
		Spec: MongoDBSpec{
			Version: jsonTypes.StrYo("3.6-v2"),
			ShardTopology: &MongoDBShardingTopology{
				Shard: MongoDBShardNode{
					Shards: 3,
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				ConfigServer: MongoDBConfigNode{
					MongoDBNode: MongoDBNode{
						Replicas: 3,
					},
					Storage: &core.PersistentVolumeClaimSpec{
						Resources: core.ResourceRequirements{
							Requests: core.ResourceList{
								core.ResourceStorage: resource.MustParse("1Gi"),
							},
						},
						StorageClassName: types.StringP("standard"),
					},
				},
				Mongos: MongoDBMongosNode{
					MongoDBNode: MongoDBNode{
						Replicas: 2,
					},
				},
			},
		},
	}

	configDSN := mongodb.ConfigSvrDSN()
	t.Log(configDSN)
}

func TestMongoDB_SetDefaults(t *testing.T) {
	mongodb := &MongoDB{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "demo-sample",
			Namespace: "demo",
		},
		Spec: MongoDBSpec{
			Version: "3.6-v2",
			Storage: &core.PersistentVolumeClaimSpec{
				Resources: core.ResourceRequirements{
					Requests: core.ResourceList{
						core.ResourceStorage: resource.MustParse("1Gi"),
					},
				},
				StorageClassName: types.StringP("standard"),
			},
		},
	}

	mongodb.Spec.SetDefaults()
}
