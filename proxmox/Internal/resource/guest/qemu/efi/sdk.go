package efi

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/Telmate/terraform-provider-proxmox/v2/proxmox/Internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func SDK(d *schema.ResourceData) *pveSDK.EfiDisk {
	v, ok := d.GetOk(Root)
	if !ok { // delete
		return &pveSDK.EfiDisk{Delete: true}
	}
	vv, ok := v.([]any)
	if ok && len(vv) != 1 {
		return nil
	}
	if settings, ok := vv[0].(map[string]any); ok {
		return &pveSDK.EfiDisk{
			Type:            util.Pointer(pveSDK.EfiDiskType(settings[schemaEfiType].(string))),
			Format:          util.Pointer(pveSDK.QemuDiskFormat(settings[schemaFormat].(string))),
			PreEnrolledKeys: util.Pointer(settings[schemaPreEnrolledKeys].(bool)),
			Storage:         util.Pointer(pveSDK.StorageName(settings[schemaStorage].(string))),
		}
	}
	return &pveSDK.EfiDisk{Delete: true}
}
