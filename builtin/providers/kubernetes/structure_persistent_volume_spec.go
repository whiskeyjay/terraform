package kubernetes

import (
	"k8s.io/kubernetes/pkg/api/resource"
	"k8s.io/kubernetes/pkg/api/v1"
	"k8s.io/kubernetes/pkg/types"

	"github.com/hashicorp/terraform/helper/schema"
)

// Flatteners

func flattenAWSElasticBlockStoreVolumeSource(in *v1.AWSElasticBlockStoreVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["volume_id"] = in.VolumeID
	att["fs_type"] = in.FSType
	att["partition"] = in.Partition
	att["read_only"] = in.ReadOnly
	return att
}

func flattenAzureDiskVolumeSource(in *v1.AzureDiskVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["disk_name"] = in.DiskName
	att["data_disk_uri"] = in.DataDiskURI
	att["caching_mode"] = *in.CachingMode
	att["fs_type"] = *in.FSType
	att["read_only"] = *in.ReadOnly
	return att
}

func flattenAzureFileVolumeSource(in *v1.AzureFileVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["secret_name"] = in.SecretName
	att["share_name"] = in.ShareName
	att["read_only"] = in.ReadOnly
	return att
}

func flattenCephFSVolumeSource(in *v1.CephFSVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["monitors"] = in.Monitors
	att["path"] = in.Path
	att["user"] = in.User
	att["secret_file"] = in.SecretFile
	att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	att["read_only"] = in.ReadOnly
	return att
}

func flattenCinderVolumeSource(in *v1.CinderVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["volume_id"] = in.VolumeID
	att["fs_type"] = in.FSType
	att["read_only"] = in.ReadOnly
	return att
}

func flattenFCVolumeSource(in *v1.FCVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["target_ww_ns"] = in.TargetWWNs
	att["lun"] = *in.Lun
	att["fs_type"] = in.FSType
	att["read_only"] = in.ReadOnly
	return att
}

func flattenFlexVolumeSource(in *v1.FlexVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["driver"] = in.Driver
	att["fs_type"] = in.FSType
	att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	att["read_only"] = in.ReadOnly
	att["options"] = in.Options
	return att
}

func flattenFlockerVolumeSource(in *v1.FlockerVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["dataset_name"] = in.DatasetName
	att["dataset_uuid"] = in.DatasetUUID
	return att
}

func flattenGCEPersistentDiskVolumeSource(in *v1.GCEPersistentDiskVolumeSource) []interface{} {
	att := make(map[string]interface{})
	att["pd_name"] = in.PDName
	if in.FSType != "" {
		att["fs_type"] = in.FSType
	}
	if in.Partition != 0 {
		att["partition"] = in.Partition
	}
	att["read_only"] = in.ReadOnly
	return []interface{}{att}
}

func flattenGlusterfsVolumeSource(in *v1.GlusterfsVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["endpoints_name"] = in.EndpointsName
	att["path"] = in.Path
	att["read_only"] = in.ReadOnly
	return att
}

func flattenHostPathVolumeSource(in *v1.HostPathVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["path"] = in.Path
	return att
}

func flattenISCSIVolumeSource(in *v1.ISCSIVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["target_portal"] = in.TargetPortal
	att["iqn"] = in.IQN
	att["lun"] = in.Lun
	att["iscsi_interface"] = in.ISCSIInterface
	att["fs_type"] = in.FSType
	att["read_only"] = in.ReadOnly
	return att
}

func flattenLocalObjectReference(in *v1.LocalObjectReference) map[string]interface{} {
	att := make(map[string]interface{})
	att["name"] = in.Name
	return att
}

func flattenNFSVolumeSource(in *v1.NFSVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["server"] = in.Server
	att["path"] = in.Path
	att["read_only"] = in.ReadOnly
	return att
}

func flattenObjectReference(in *v1.ObjectReference) map[string]interface{} {
	att := make(map[string]interface{})
	att["kind"] = in.Kind
	att["namespace"] = in.Namespace
	att["name"] = in.Name
	att["uid"] = in.UID
	att["api_version"] = in.APIVersion
	att["resource_version"] = in.ResourceVersion
	att["field_path"] = in.FieldPath
	return att
}

func flattenPersistentVolumeSource(in v1.PersistentVolumeSource) []interface{} {
	att := make(map[string]interface{})

	if in.GCEPersistentDisk != nil {
		att["gce_persistent_disk"] = flattenGCEPersistentDiskVolumeSource(in.GCEPersistentDisk)
	}
	if in.AWSElasticBlockStore != nil {
		att["aws_elastic_block_store"] = flattenAWSElasticBlockStoreVolumeSource(in.AWSElasticBlockStore)
	}
	if in.HostPath != nil {
		att["host_path"] = flattenHostPathVolumeSource(in.HostPath)
	}
	if in.Glusterfs != nil {
		att["glusterfs"] = flattenGlusterfsVolumeSource(in.Glusterfs)
	}
	if in.NFS != nil {
		att["nfs"] = flattenNFSVolumeSource(in.NFS)
	}
	if in.RBD != nil {
		att["rbd"] = flattenRBDVolumeSource(in.RBD)
	}
	if in.ISCSI != nil {
		att["iscsi"] = flattenISCSIVolumeSource(in.ISCSI)
	}
	if in.Cinder != nil {
		att["cinder"] = flattenCinderVolumeSource(in.Cinder)
	}
	if in.CephFS != nil {
		att["ceph_fs"] = flattenCephFSVolumeSource(in.CephFS)
	}
	if in.FC != nil {
		att["fc"] = flattenFCVolumeSource(in.FC)
	}
	if in.Flocker != nil {
		att["flocker"] = flattenFlockerVolumeSource(in.Flocker)
	}
	if in.FlexVolume != nil {
		att["flex_volume"] = flattenFlexVolumeSource(in.FlexVolume)
	}
	if in.AzureFile != nil {
		att["azure_file"] = flattenAzureFileVolumeSource(in.AzureFile)
	}
	if in.VsphereVolume != nil {
		att["vsphere_volume"] = flattenVsphereVirtualDiskVolumeSource(in.VsphereVolume)
	}
	if in.Quobyte != nil {
		att["quobyte"] = flattenQuobyteVolumeSource(in.Quobyte)
	}
	if in.AzureDisk != nil {
		att["azure_disk"] = flattenAzureDiskVolumeSource(in.AzureDisk)
	}
	if in.PhotonPersistentDisk != nil {
		att["photon_persistent_disk"] = flattenPhotonPersistentDiskVolumeSource(in.PhotonPersistentDisk)
	}
	return []interface{}{att}
}

func flattenPersistentVolumeSpec(in v1.PersistentVolumeSpec) []interface{} {
	att := make(map[string]interface{})
	att["capacity"] = flattenResourceList(in.Capacity)
	att["persistent_volume_source"] = flattenPersistentVolumeSource(in.PersistentVolumeSource)
	att["access_modes"] = flattenPersistentVolumeAccessModes(in.AccessModes)
	if in.ClaimRef != nil {
		att["claim_ref"] = flattenObjectReference(in.ClaimRef)
	}
	att["persistent_volume_reclaim_policy"] = in.PersistentVolumeReclaimPolicy
	return []interface{}{att}
}

func flattenPhotonPersistentDiskVolumeSource(in *v1.PhotonPersistentDiskVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["pd_id"] = in.PdID
	att["fs_type"] = in.FSType
	return att
}

func flattenQuobyteVolumeSource(in *v1.QuobyteVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["registry"] = in.Registry
	att["volume"] = in.Volume
	att["read_only"] = in.ReadOnly
	att["user"] = in.User
	att["group"] = in.Group
	return att
}

func flattenRBDVolumeSource(in *v1.RBDVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["ceph_monitors"] = in.CephMonitors
	att["rbd_image"] = in.RBDImage
	att["fs_type"] = in.FSType
	att["rbd_pool"] = in.RBDPool
	att["rados_user"] = in.RadosUser
	att["keyring"] = in.Keyring
	att["secret_ref"] = flattenLocalObjectReference(in.SecretRef)
	att["read_only"] = in.ReadOnly
	return att
}

func flattenVsphereVirtualDiskVolumeSource(in *v1.VsphereVirtualDiskVolumeSource) map[string]interface{} {
	att := make(map[string]interface{})
	att["volume_path"] = in.VolumePath
	att["fs_type"] = in.FSType
	return att
}

// Expanders

func expandAWSElasticBlockStoreVolumeSource(l []interface{}) *v1.AWSElasticBlockStoreVolumeSource {
	if len(l) == 0 {
		return &v1.AWSElasticBlockStoreVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.AWSElasticBlockStoreVolumeSource{
		VolumeID:  in["volume_id"].(string),
		FSType:    in["fs_type"].(string),
		Partition: in["partition"].(int32),
		ReadOnly:  in["read_only"].(bool),
	}
	return &obj
}

func expandAzureDiskVolumeSource(l []interface{}) *v1.AzureDiskVolumeSource {
	if len(l) == 0 {
		return &v1.AzureDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.AzureDiskVolumeSource{
		DiskName:    in["disk_name"].(string),
		DataDiskURI: in["data_disk_uri"].(string),
		CachingMode: ptrToAzureDataDiskCachingMode(in["caching_mode"].(v1.AzureDataDiskCachingMode)),
		FSType:      ptrToString(in["fs_type"].(string)),
		ReadOnly:    ptrToBool(in["read_only"].(bool)),
	}
	return &obj
}

func expandAzureFileVolumeSource(l []interface{}) *v1.AzureFileVolumeSource {
	if len(l) == 0 {
		return &v1.AzureFileVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.AzureFileVolumeSource{
		SecretName: in["secret_name"].(string),
		ShareName:  in["share_name"].(string),
		ReadOnly:   in["read_only"].(bool),
	}
	return &obj
}

func expandCephFSVolumeSource(l []interface{}) *v1.CephFSVolumeSource {
	if len(l) == 0 {
		return &v1.CephFSVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.CephFSVolumeSource{
		Monitors:   sliceOfString(in["monitors"].([]interface{})),
		Path:       in["path"].(string),
		User:       in["user"].(string),
		SecretFile: in["secret_file"].(string),
		SecretRef:  expandLocalObjectReference(in["secret_ref"].([]interface{})),
		ReadOnly:   in["read_only"].(bool),
	}
	return &obj
}

func expandCinderVolumeSource(l []interface{}) *v1.CinderVolumeSource {
	if len(l) == 0 {
		return &v1.CinderVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.CinderVolumeSource{
		VolumeID: in["volume_id"].(string),
		FSType:   in["fs_type"].(string),
		ReadOnly: in["read_only"].(bool),
	}
	return &obj
}

func expandFCVolumeSource(l []interface{}) *v1.FCVolumeSource {
	if len(l) == 0 {
		return &v1.FCVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.FCVolumeSource{
		TargetWWNs: sliceOfString(in["target_ww_ns"].([]interface{})),
		Lun:        ptrToInt32(in["lun"].(int32)),
		FSType:     in["fs_type"].(string),
		ReadOnly:   in["read_only"].(bool),
	}
	return &obj
}

func expandFlexVolumeSource(l []interface{}) *v1.FlexVolumeSource {
	if len(l) == 0 {
		return &v1.FlexVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.FlexVolumeSource{
		Driver:    in["driver"].(string),
		FSType:    in["fs_type"].(string),
		SecretRef: expandLocalObjectReference(in["secret_ref"].([]interface{})),
		ReadOnly:  in["read_only"].(bool),
		Options:   expandStringMap(in["options"].(map[string]interface{})),
	}
	return &obj
}

func expandFlockerVolumeSource(l []interface{}) *v1.FlockerVolumeSource {
	if len(l) == 0 {
		return &v1.FlockerVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.FlockerVolumeSource{
		DatasetName: in["dataset_name"].(string),
		DatasetUUID: in["dataset_uuid"].(string),
	}
	return &obj
}

func expandGCEPersistentDiskVolumeSource(l []interface{}) *v1.GCEPersistentDiskVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return &v1.GCEPersistentDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.GCEPersistentDiskVolumeSource{
		PDName: in["pd_name"].(string),
	}
	if v, ok := in["fs_type"].(string); ok {
		obj.FSType = v
	}
	if v, ok := in["partition"].(int); ok {
		obj.Partition = int32(v)
	}
	if v, ok := in["read_only"].(bool); ok {
		obj.ReadOnly = v
	}
	return &obj
}

func expandGlusterfsVolumeSource(l []interface{}) *v1.GlusterfsVolumeSource {
	if len(l) == 0 {
		return &v1.GlusterfsVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.GlusterfsVolumeSource{
		EndpointsName: in["endpoints_name"].(string),
		Path:          in["path"].(string),
		ReadOnly:      in["read_only"].(bool),
	}
	return &obj
}

func expandHostPathVolumeSource(l []interface{}) *v1.HostPathVolumeSource {
	if len(l) == 0 {
		return &v1.HostPathVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.HostPathVolumeSource{
		Path: in["path"].(string),
	}
	return &obj
}

func expandISCSIVolumeSource(l []interface{}) *v1.ISCSIVolumeSource {
	if len(l) == 0 {
		return &v1.ISCSIVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.ISCSIVolumeSource{
		TargetPortal:   in["target_portal"].(string),
		IQN:            in["iqn"].(string),
		Lun:            in["lun"].(int32),
		ISCSIInterface: in["iscsi_interface"].(string),
		FSType:         in["fs_type"].(string),
		ReadOnly:       in["read_only"].(bool),
	}
	return &obj
}

func expandLocalObjectReference(l []interface{}) *v1.LocalObjectReference {
	if len(l) == 0 {
		return &v1.LocalObjectReference{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.LocalObjectReference{
		Name: in["name"].(string),
	}
	return &obj
}

func expandNFSVolumeSource(l []interface{}) *v1.NFSVolumeSource {
	if len(l) == 0 {
		return &v1.NFSVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.NFSVolumeSource{
		Server:   in["server"].(string),
		Path:     in["path"].(string),
		ReadOnly: in["read_only"].(bool),
	}
	return &obj
}

func expandObjectReference(l []interface{}) *v1.ObjectReference {
	if len(l) == 0 {
		return &v1.ObjectReference{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.ObjectReference{
		Kind:            in["kind"].(string),
		Namespace:       in["namespace"].(string),
		Name:            in["name"].(string),
		UID:             in["uid"].(types.UID),
		APIVersion:      in["api_version"].(string),
		ResourceVersion: in["resource_version"].(string),
		FieldPath:       in["field_path"].(string),
	}
	return &obj
}

func expandPersistentVolumeSource(l []interface{}) v1.PersistentVolumeSource {
	if len(l) == 0 || l[0] == nil {
		return v1.PersistentVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.PersistentVolumeSource{}

	if v, ok := in["gce_persistent_disk"].([]interface{}); ok && len(v) > 0 {
		obj.GCEPersistentDisk = expandGCEPersistentDiskVolumeSource(v)
	}
	if v, ok := in["aws_elastic_block_store"].([]interface{}); ok && len(v) > 0 {
		obj.AWSElasticBlockStore = expandAWSElasticBlockStoreVolumeSource(v)
	}
	if v, ok := in["host_path"].([]interface{}); ok && len(v) > 0 {
		obj.HostPath = expandHostPathVolumeSource(v)
	}
	if v, ok := in["glusterfs"].([]interface{}); ok && len(v) > 0 {
		obj.Glusterfs = expandGlusterfsVolumeSource(v)
	}
	if v, ok := in["nfs"].([]interface{}); ok && len(v) > 0 {
		obj.NFS = expandNFSVolumeSource(v)
	}
	if v, ok := in["rbd"].([]interface{}); ok && len(v) > 0 {
		obj.RBD = expandRBDVolumeSource(v)
	}
	if v, ok := in["iscsi"].([]interface{}); ok && len(v) > 0 {
		obj.ISCSI = expandISCSIVolumeSource(v)
	}
	if v, ok := in["cinder"].([]interface{}); ok && len(v) > 0 {
		obj.Cinder = expandCinderVolumeSource(v)
	}
	if v, ok := in["ceph_fs"].([]interface{}); ok && len(v) > 0 {
		obj.CephFS = expandCephFSVolumeSource(v)
	}
	if v, ok := in["fc"].([]interface{}); ok && len(v) > 0 {
		obj.FC = expandFCVolumeSource(v)
	}
	if v, ok := in["flocker"].([]interface{}); ok && len(v) > 0 {
		obj.Flocker = expandFlockerVolumeSource(v)
	}
	if v, ok := in["flex_volume"].([]interface{}); ok && len(v) > 0 {
		obj.FlexVolume = expandFlexVolumeSource(v)
	}
	if v, ok := in["azure_file"].([]interface{}); ok && len(v) > 0 {
		obj.AzureFile = expandAzureFileVolumeSource(v)
	}
	if v, ok := in["vsphere_volume"].([]interface{}); ok && len(v) > 0 {
		obj.VsphereVolume = expandVsphereVirtualDiskVolumeSource(v)
	}
	if v, ok := in["quobyte"].([]interface{}); ok && len(v) > 0 {
		obj.Quobyte = expandQuobyteVolumeSource(v)
	}
	if v, ok := in["azure_disk"].([]interface{}); ok && len(v) > 0 {
		obj.AzureDisk = expandAzureDiskVolumeSource(v)
	}
	if v, ok := in["photon_persistent_disk"].([]interface{}); ok && len(v) > 0 {
		obj.PhotonPersistentDisk = expandPhotonPersistentDiskVolumeSource(v)
	}

	return obj
}

func expandPersistentVolumeSpec(l []interface{}) v1.PersistentVolumeSpec {
	if len(l) == 0 || l[0] == nil {
		return v1.PersistentVolumeSpec{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.PersistentVolumeSpec{
		Capacity:    expandMapToResourceList(in["capacity"].(map[string]interface{})),
		AccessModes: expandPersistentVolumeAccessModes(in["access_modes"].(*schema.Set).List()),
	}

	if v, ok := in["persistent_volume_source"].([]interface{}); ok && len(v) > 0 {
		obj.PersistentVolumeSource = expandPersistentVolumeSource(in["persistent_volume_source"].([]interface{}))
	}
	if v, ok := in["claim_ref"].([]interface{}); ok && len(v) > 0 {
		obj.ClaimRef = expandObjectReference(in["claim_ref"].([]interface{}))
	}
	if v, ok := in["persistent_volume_reclaim_policy"].(v1.PersistentVolumeReclaimPolicy); ok && len(v) > 0 {
		obj.PersistentVolumeReclaimPolicy = in["persistent_volume_reclaim_policy"].(v1.PersistentVolumeReclaimPolicy)
	}

	return obj
}

func expandPhotonPersistentDiskVolumeSource(l []interface{}) *v1.PhotonPersistentDiskVolumeSource {
	if len(l) == 0 {
		return &v1.PhotonPersistentDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.PhotonPersistentDiskVolumeSource{
		PdID:   in["pd_id"].(string),
		FSType: in["fs_type"].(string),
	}
	return &obj
}

func expandQuobyteVolumeSource(l []interface{}) *v1.QuobyteVolumeSource {
	if len(l) == 0 {
		return &v1.QuobyteVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.QuobyteVolumeSource{
		Registry: in["registry"].(string),
		Volume:   in["volume"].(string),
		ReadOnly: in["read_only"].(bool),
		User:     in["user"].(string),
		Group:    in["group"].(string),
	}
	return &obj
}

func expandRBDVolumeSource(l []interface{}) *v1.RBDVolumeSource {
	if len(l) == 0 {
		return &v1.RBDVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.RBDVolumeSource{
		CephMonitors: sliceOfString(in["ceph_monitors"].([]interface{})),
		RBDImage:     in["rbd_image"].(string),
		FSType:       in["fs_type"].(string),
		RBDPool:      in["rbd_pool"].(string),
		RadosUser:    in["rados_user"].(string),
		Keyring:      in["keyring"].(string),
		SecretRef:    expandLocalObjectReference(in["secret_ref"].([]interface{})),
		ReadOnly:     in["read_only"].(bool),
	}
	return &obj
}

func expandVsphereVirtualDiskVolumeSource(l []interface{}) *v1.VsphereVirtualDiskVolumeSource {
	if len(l) == 0 {
		return &v1.VsphereVirtualDiskVolumeSource{}
	}
	in := l[0].(map[string]interface{})
	obj := v1.VsphereVirtualDiskVolumeSource{
		VolumePath: in["volume_path"].(string),
		FSType:     in["fs_type"].(string),
	}
	return &obj
}

func flattenResourceList(l v1.ResourceList) map[string]string {
	m := make(map[string]string)
	for k, v := range l {
		m[string(k)] = v.String()
	}
	return m
}

func expandMapToResourceList(m map[string]interface{}) v1.ResourceList {
	out := make(map[v1.ResourceName]resource.Quantity)
	for stringKey, v := range m {
		key := v1.ResourceName(stringKey)
		value, err := resource.ParseQuantity(v.(string))
		if err != nil {
			// TODO
		}

		out[key] = value
	}
	return out
}

func ifaceFromString(s string) interface{} {
	return s
}

func flattenPersistentVolumeAccessModes(in []v1.PersistentVolumeAccessMode) *schema.Set {
	var out = make([]interface{}, len(in), len(in))
	for i, v := range in {
		out[i] = string(v)
	}
	return schema.NewSet(schema.HashString, out)
}

func expandPersistentVolumeAccessModes(s []interface{}) []v1.PersistentVolumeAccessMode {
	out := make([]v1.PersistentVolumeAccessMode, len(s), len(s))
	for i, v := range s {
		out[i] = v1.PersistentVolumeAccessMode(v.(string))
	}
	return out
}

func ptrToAzureDataDiskCachingMode(in v1.AzureDataDiskCachingMode) *v1.AzureDataDiskCachingMode {
	return &in
}
