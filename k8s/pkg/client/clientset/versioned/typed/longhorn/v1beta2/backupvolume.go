/*
Copyright The Longhorn Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package v1beta2

import (
	context "context"

	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	applyconfigurationlonghornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/client/applyconfiguration/longhorn/v1beta2"
	scheme "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/scheme"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	types "k8s.io/apimachinery/pkg/types"
	watch "k8s.io/apimachinery/pkg/watch"
	gentype "k8s.io/client-go/gentype"
)

// BackupVolumesGetter has a method to return a BackupVolumeInterface.
// A group's client should implement this interface.
type BackupVolumesGetter interface {
	BackupVolumes(namespace string) BackupVolumeInterface
}

// BackupVolumeInterface has methods to work with BackupVolume resources.
type BackupVolumeInterface interface {
	Create(ctx context.Context, backupVolume *longhornv1beta2.BackupVolume, opts v1.CreateOptions) (*longhornv1beta2.BackupVolume, error)
	Update(ctx context.Context, backupVolume *longhornv1beta2.BackupVolume, opts v1.UpdateOptions) (*longhornv1beta2.BackupVolume, error)
	// Add a +genclient:noStatus comment above the type to avoid generating UpdateStatus().
	UpdateStatus(ctx context.Context, backupVolume *longhornv1beta2.BackupVolume, opts v1.UpdateOptions) (*longhornv1beta2.BackupVolume, error)
	Delete(ctx context.Context, name string, opts v1.DeleteOptions) error
	DeleteCollection(ctx context.Context, opts v1.DeleteOptions, listOpts v1.ListOptions) error
	Get(ctx context.Context, name string, opts v1.GetOptions) (*longhornv1beta2.BackupVolume, error)
	List(ctx context.Context, opts v1.ListOptions) (*longhornv1beta2.BackupVolumeList, error)
	Watch(ctx context.Context, opts v1.ListOptions) (watch.Interface, error)
	Patch(ctx context.Context, name string, pt types.PatchType, data []byte, opts v1.PatchOptions, subresources ...string) (result *longhornv1beta2.BackupVolume, err error)
	Apply(ctx context.Context, backupVolume *applyconfigurationlonghornv1beta2.BackupVolumeApplyConfiguration, opts v1.ApplyOptions) (result *longhornv1beta2.BackupVolume, err error)
	// Add a +genclient:noStatus comment above the type to avoid generating ApplyStatus().
	ApplyStatus(ctx context.Context, backupVolume *applyconfigurationlonghornv1beta2.BackupVolumeApplyConfiguration, opts v1.ApplyOptions) (result *longhornv1beta2.BackupVolume, err error)
	BackupVolumeExpansion
}

// backupVolumes implements BackupVolumeInterface
type backupVolumes struct {
	*gentype.ClientWithListAndApply[*longhornv1beta2.BackupVolume, *longhornv1beta2.BackupVolumeList, *applyconfigurationlonghornv1beta2.BackupVolumeApplyConfiguration]
}

// newBackupVolumes returns a BackupVolumes
func newBackupVolumes(c *LonghornV1beta2Client, namespace string) *backupVolumes {
	return &backupVolumes{
		gentype.NewClientWithListAndApply[*longhornv1beta2.BackupVolume, *longhornv1beta2.BackupVolumeList, *applyconfigurationlonghornv1beta2.BackupVolumeApplyConfiguration](
			"backupvolumes",
			c.RESTClient(),
			scheme.ParameterCodec,
			namespace,
			func() *longhornv1beta2.BackupVolume { return &longhornv1beta2.BackupVolume{} },
			func() *longhornv1beta2.BackupVolumeList { return &longhornv1beta2.BackupVolumeList{} },
		),
	}
}
