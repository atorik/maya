/*
Copyright 2019 The OpenEBS Authors

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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	internalinterfaces "github.com/openebs/maya/pkg/client/generated/informer/externalversions/internalinterfaces"
)

// Interface provides access to all the informers in this group version.
type Interface interface {
	// BackupCStors returns a BackupCStorInformer.
	BackupCStors() BackupCStorInformer
	// BackupCStorLasts returns a BackupCStorLastInformer.
	BackupCStorLasts() BackupCStorLastInformer
	// CASTemplates returns a CASTemplateInformer.
	CASTemplates() CASTemplateInformer
	// CStorPools returns a CStorPoolInformer.
	CStorPools() CStorPoolInformer
	// CStorRestores returns a CStorRestoreInformer.
	CStorRestores() CStorRestoreInformer
	// CStorVolumes returns a CStorVolumeInformer.
	CStorVolumes() CStorVolumeInformer
	// CStorVolumeReplicas returns a CStorVolumeReplicaInformer.
	CStorVolumeReplicas() CStorVolumeReplicaInformer
	// Disks returns a DiskInformer.
	Disks() DiskInformer
	// RunTasks returns a RunTaskInformer.
	RunTasks() RunTaskInformer
	// StoragePools returns a StoragePoolInformer.
	StoragePools() StoragePoolInformer
	// StoragePoolClaims returns a StoragePoolClaimInformer.
	StoragePoolClaims() StoragePoolClaimInformer
}

type version struct {
	factory          internalinterfaces.SharedInformerFactory
	namespace        string
	tweakListOptions internalinterfaces.TweakListOptionsFunc
}

// New returns a new Interface.
func New(f internalinterfaces.SharedInformerFactory, namespace string, tweakListOptions internalinterfaces.TweakListOptionsFunc) Interface {
	return &version{factory: f, namespace: namespace, tweakListOptions: tweakListOptions}
}

// BackupCStors returns a BackupCStorInformer.
func (v *version) BackupCStors() BackupCStorInformer {
	return &backupCStorInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// BackupCStorLasts returns a BackupCStorLastInformer.
func (v *version) BackupCStorLasts() BackupCStorLastInformer {
	return &backupCStorLastInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CASTemplates returns a CASTemplateInformer.
func (v *version) CASTemplates() CASTemplateInformer {
	return &cASTemplateInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CStorPools returns a CStorPoolInformer.
func (v *version) CStorPools() CStorPoolInformer {
	return &cStorPoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// CStorRestores returns a CStorRestoreInformer.
func (v *version) CStorRestores() CStorRestoreInformer {
	return &cStorRestoreInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumes returns a CStorVolumeInformer.
func (v *version) CStorVolumes() CStorVolumeInformer {
	return &cStorVolumeInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// CStorVolumeReplicas returns a CStorVolumeReplicaInformer.
func (v *version) CStorVolumeReplicas() CStorVolumeReplicaInformer {
	return &cStorVolumeReplicaInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// Disks returns a DiskInformer.
func (v *version) Disks() DiskInformer {
	return &diskInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// RunTasks returns a RunTaskInformer.
func (v *version) RunTasks() RunTaskInformer {
	return &runTaskInformer{factory: v.factory, namespace: v.namespace, tweakListOptions: v.tweakListOptions}
}

// StoragePools returns a StoragePoolInformer.
func (v *version) StoragePools() StoragePoolInformer {
	return &storagePoolInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}

// StoragePoolClaims returns a StoragePoolClaimInformer.
func (v *version) StoragePoolClaims() StoragePoolClaimInformer {
	return &storagePoolClaimInformer{factory: v.factory, tweakListOptions: v.tweakListOptions}
}