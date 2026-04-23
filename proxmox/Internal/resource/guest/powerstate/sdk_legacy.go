package powerstate

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"
	"github.com/Telmate/terraform-provider-proxmox/v2/proxmox/Internal/util"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func sdkLegacy(d *schema.ResourceData) *pveSDK.PowerState {
	v, ok := d.GetOk(LegacyRoot)
	if !ok {
		return util.Pointer(pveSDK.PowerStateRunning)
	}
	switch v.(string) {
	case enumRunning:
		return util.Pointer(pveSDK.PowerStateRunning)
	case enumStopped:
		return util.Pointer(pveSDK.PowerStateStopped)
	default:
		return nil
	}
}
