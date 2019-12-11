/*
* Tencent is pleased to support the open source community by making TKEStack
* available.
*
* Copyright (C) 2012-2019 Tencent. All Rights Reserved.
*
* Licensed under the Apache License, Version 2.0 (the “License”); you may not use
* this file except in compliance with the License. You may obtain a copy of the
* License at
*
* https://opensource.org/licenses/Apache-2.0
*
* Unless required by applicable law or agreed to in writing, software
* distributed under the License is distributed on an “AS IS” BASIS, WITHOUT
* WARRANTIES OF ANY KIND, either express or implied.  See the License for the
* specific language governing permissions and limitations under the License.
 */
package fake

import "tkestack.io/kube-jarvis/pkg/plugins/cluster"

type Cluster struct {
	Res *cluster.Resources
}

func NewCluster() *Cluster {
	return &Cluster{Res: cluster.NewResources()}
}

// Init Instantiation for cluster, it will fetch Resources
func (c *Cluster) Init() error {
	return nil
}

// SyncResources fetch all resource from cluster
func (c *Cluster) SyncResources() error {
	c.Res = cluster.NewResources()
	return nil
}

// CloudType return the cloud type of Cluster
func (c *Cluster) CloudType() string {
	return "fake"
}

// Machine return the low level information of a node
func (c *Cluster) Resources() *cluster.Resources {
	return c.Res
}
