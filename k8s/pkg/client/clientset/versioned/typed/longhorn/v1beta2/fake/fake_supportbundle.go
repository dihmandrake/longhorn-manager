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

package fake

import (
	v1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/apis/longhorn/v1beta2"
	longhornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/client/applyconfiguration/longhorn/v1beta2"
	typedlonghornv1beta2 "github.com/longhorn/longhorn-manager/k8s/pkg/client/clientset/versioned/typed/longhorn/v1beta2"
	gentype "k8s.io/client-go/gentype"
)

// fakeSupportBundles implements SupportBundleInterface
type fakeSupportBundles struct {
	*gentype.FakeClientWithListAndApply[*v1beta2.SupportBundle, *v1beta2.SupportBundleList, *longhornv1beta2.SupportBundleApplyConfiguration]
	Fake *FakeLonghornV1beta2
}

func newFakeSupportBundles(fake *FakeLonghornV1beta2, namespace string) typedlonghornv1beta2.SupportBundleInterface {
	return &fakeSupportBundles{
		gentype.NewFakeClientWithListAndApply[*v1beta2.SupportBundle, *v1beta2.SupportBundleList, *longhornv1beta2.SupportBundleApplyConfiguration](
			fake.Fake,
			namespace,
			v1beta2.SchemeGroupVersion.WithResource("supportbundles"),
			v1beta2.SchemeGroupVersion.WithKind("SupportBundle"),
			func() *v1beta2.SupportBundle { return &v1beta2.SupportBundle{} },
			func() *v1beta2.SupportBundleList { return &v1beta2.SupportBundleList{} },
			func(dst, src *v1beta2.SupportBundleList) { dst.ListMeta = src.ListMeta },
			func(list *v1beta2.SupportBundleList) []*v1beta2.SupportBundle {
				return gentype.ToPointerSlice(list.Items)
			},
			func(list *v1beta2.SupportBundleList, items []*v1beta2.SupportBundle) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
