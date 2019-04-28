/*
Copyright (c) 2019 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// IMPORTANT: This file has been generated automatically, refrain from modifying it manually as all
// your changes will be lost when the file is generated again.

package v1 // github.com/openshift-online/uhc-sdk-go/pkg/client/clustersmgmt/v1

import (
	"fmt"
	time "time"

	"github.com/openshift-online/uhc-sdk-go/pkg/client/helpers"
)

// clusterData is the data structure used internally to marshal and unmarshal
// objects of type 'cluster'.
type clusterData struct {
	Kind              *string                       "json:\"kind,omitempty\""
	ID                *string                       "json:\"id,omitempty\""
	HREF              *string                       "json:\"href,omitempty\""
	Name              *string                       "json:\"name,omitempty\""
	Flavour           *flavourData                  "json:\"flavour,omitempty\""
	Console           *clusterConsoleData           "json:\"console,omitempty\""
	MultiAZ           *bool                         "json:\"multi_az,omitempty\""
	Nodes             *clusterNodesData             "json:\"nodes,omitempty\""
	API               *clusterAPIData               "json:\"api,omitempty\""
	Region            *cloudRegionData              "json:\"region,omitempty\""
	DisplayName       *string                       "json:\"display_name,omitempty\""
	DNS               *dnsData                      "json:\"dns,omitempty\""
	Properties        map[string]string             "json:\"properties,omitempty\""
	State             *ClusterState                 "json:\"state,omitempty\""
	Managed           *bool                         "json:\"managed,omitempty\""
	Memory            *clusterMetricData            "json:\"memory,omitempty\""
	CPU               *clusterMetricData            "json:\"cpu,omitempty\""
	Storage           *clusterMetricData            "json:\"storage,omitempty\""
	ExternalID        *string                       "json:\"external_id,omitempty\""
	AWS               *awsData                      "json:\"aws,omitempty\""
	Network           *networkData                  "json:\"network,omitempty\""
	CreationTimestamp *time.Time                    "json:\"creation_timestamp,omitempty\""
	CloudProvider     *cloudProviderData            "json:\"cloud_provider,omitempty\""
	OpenshiftVersion  *string                       "json:\"openshift_version,omitempty\""
	Subscription      *subscriptionData             "json:\"subscription,omitempty\""
	Groups            *groupListLinkData            "json:\"groups,omitempty\""
	Creator           *string                       "json:\"creator,omitempty\""
	Version           *versionData                  "json:\"version,omitempty\""
	IdentityProviders *identityProviderListLinkData "json:\"identity_providers,omitempty\""
}

// MarshalCluster writes a value of the 'cluster' to the given target,
// which can be a writer or a JSON encoder.
func MarshalCluster(object *Cluster, target interface{}) error {
	encoder, err := helpers.NewEncoder(target)
	if err != nil {
		return err
	}
	data, err := object.wrap()
	if err != nil {
		return err
	}
	return encoder.Encode(data)
}

// wrap is the method used internally to convert a value of the 'cluster'
// value to a JSON document.
func (o *Cluster) wrap() (data *clusterData, err error) {
	if o == nil {
		return
	}
	data = new(clusterData)
	data.ID = o.id
	data.HREF = o.href
	data.Kind = new(string)
	if o.link {
		*data.Kind = ClusterLinkKind
	} else {
		*data.Kind = ClusterKind
	}
	data.Name = o.name
	data.Flavour, err = o.flavour.wrap()
	if err != nil {
		return
	}
	data.Console, err = o.console.wrap()
	if err != nil {
		return
	}
	data.MultiAZ = o.multiAZ
	data.Nodes, err = o.nodes.wrap()
	if err != nil {
		return
	}
	data.API, err = o.api.wrap()
	if err != nil {
		return
	}
	data.Region, err = o.region.wrap()
	if err != nil {
		return
	}
	data.DisplayName = o.displayName
	data.DNS, err = o.dns.wrap()
	if err != nil {
		return
	}
	data.Properties = o.properties
	data.State = o.state
	data.Managed = o.managed
	data.Memory, err = o.memory.wrap()
	if err != nil {
		return
	}
	data.CPU, err = o.cpu.wrap()
	if err != nil {
		return
	}
	data.Storage, err = o.storage.wrap()
	if err != nil {
		return
	}
	data.ExternalID = o.externalID
	data.AWS, err = o.aws.wrap()
	if err != nil {
		return
	}
	data.Network, err = o.network.wrap()
	if err != nil {
		return
	}
	data.CreationTimestamp = o.creationTimestamp
	data.CloudProvider, err = o.cloudProvider.wrap()
	if err != nil {
		return
	}
	data.OpenshiftVersion = o.openshiftVersion
	data.Subscription, err = o.subscription.wrap()
	if err != nil {
		return
	}
	data.Groups, err = o.groups.wrapLink()
	if err != nil {
		return
	}
	data.Creator = o.creator
	data.Version, err = o.version.wrap()
	if err != nil {
		return
	}
	data.IdentityProviders, err = o.identityProviders.wrapLink()
	if err != nil {
		return
	}
	return
}

// UnmarshalCluster reads a value of the 'cluster' type from the given
// source, which can be an slice of bytes, a string, a reader or a JSON decoder.
func UnmarshalCluster(source interface{}) (object *Cluster, err error) {
	decoder, err := helpers.NewDecoder(source)
	if err != nil {
		return
	}
	data := new(clusterData)
	err = decoder.Decode(data)
	if err != nil {
		return
	}
	object, err = data.unwrap()
	return
}

// unwrap is the function used internally to convert the JSON unmarshalled data to a
// value of the 'cluster' type.
func (d *clusterData) unwrap() (object *Cluster, err error) {
	if d == nil {
		return
	}
	object = new(Cluster)
	object.id = d.ID
	object.href = d.HREF
	if d.Kind != nil {
		switch *d.Kind {
		case ClusterKind:
			object.link = false
		case ClusterLinkKind:
			object.link = true
		default:
			err = fmt.Errorf(
				"expected kind '%s' or '%s' but got '%s'",
				ClusterKind,
				ClusterLinkKind,
				*d.Kind,
			)
			return
		}
	}
	object.name = d.Name
	object.flavour, err = d.Flavour.unwrap()
	if err != nil {
		return
	}
	object.console, err = d.Console.unwrap()
	if err != nil {
		return
	}
	object.multiAZ = d.MultiAZ
	object.nodes, err = d.Nodes.unwrap()
	if err != nil {
		return
	}
	object.api, err = d.API.unwrap()
	if err != nil {
		return
	}
	object.region, err = d.Region.unwrap()
	if err != nil {
		return
	}
	object.displayName = d.DisplayName
	object.dns, err = d.DNS.unwrap()
	if err != nil {
		return
	}
	object.properties = d.Properties
	object.state = d.State
	object.managed = d.Managed
	object.memory, err = d.Memory.unwrap()
	if err != nil {
		return
	}
	object.cpu, err = d.CPU.unwrap()
	if err != nil {
		return
	}
	object.storage, err = d.Storage.unwrap()
	if err != nil {
		return
	}
	object.externalID = d.ExternalID
	object.aws, err = d.AWS.unwrap()
	if err != nil {
		return
	}
	object.network, err = d.Network.unwrap()
	if err != nil {
		return
	}
	object.creationTimestamp = d.CreationTimestamp
	object.cloudProvider, err = d.CloudProvider.unwrap()
	if err != nil {
		return
	}
	object.openshiftVersion = d.OpenshiftVersion
	object.subscription, err = d.Subscription.unwrap()
	if err != nil {
		return
	}
	object.groups, err = d.Groups.unwrapLink()
	if err != nil {
		return
	}
	object.creator = d.Creator
	object.version, err = d.Version.unwrap()
	if err != nil {
		return
	}
	object.identityProviders, err = d.IdentityProviders.unwrapLink()
	if err != nil {
		return
	}
	return
}
