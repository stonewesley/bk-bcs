/*
 * Tencent is pleased to support the open source community by making Blueking Container Service available.
 * Copyright (C) 2019 THL A29 Limited, a Tencent company. All rights reserved.
 * Licensed under the MIT License (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 * http://opensource.org/licenses/MIT
 * Unless required by applicable law or agreed to in writing, software distributed under
 * the License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions and
 * limitations under the License.
 *
 */

package types

import (
	"fmt"

	commtypes "github.com/Tencent/bk-bcs/bcs-common/common/types"
)

//Application for container
type BcsDaemonset struct {
	commtypes.ObjectMeta `json:",inline"`
	Instances            uint64
	RunningInstances     uint64
	Status               string
	LastStatus           string
	Created              int64
	LastUpdateTime       int64
	//force delete daemonset
	ForceDeleting bool
	//taskgroup id list
	Pods []*commtypes.BcsPodIndex
	// Populated by the system.
	// Read-only.
	// Value must be treated as opaque by clients and .
	ResourceVersion string `json:"-"`
	// RC current original definition
	//RawJson []byte `json:"raw_json,omitempty"`
}

//get daemonset the unique uuid
func (d *BcsDaemonset) GetUuid() string {
	return fmt.Sprintf("%s.%s", d.NameSpace, d.Name)
}

type BcsDaemonsetDef struct {
	commtypes.ObjectMeta `json:",inline"`
	Version              *Version `json:"version"`
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *BcsDaemonset) DeepCopyInto(out *BcsDaemonset) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	if in.Pods != nil {
		in, out := &in.Pods, &out.Pods
		*out = make([]*commtypes.BcsPodIndex, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(commtypes.BcsPodIndex)
				**out = **in
			}
		}
	}

	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AgentSchedInfoSpec.
func (in *BcsDaemonset) DeepCopy() *BcsDaemonset {
	if in == nil {
		return nil
	}
	out := new(BcsDaemonset)
	in.DeepCopyInto(out)
	return out
}
