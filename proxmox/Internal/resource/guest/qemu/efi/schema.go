package efi

import (
	pveSDK "github.com/Telmate/proxmox-api-go/proxmox"

	"github.com/hashicorp/go-cty/cty"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

const (
	Root = "efidisk"

	schemaEfiType         = "efitype"
	schemaFormat          = "format"
	schemaPreEnrolledKeys = "pre_enrolled_keys"
	schemaStorage         = "storage"

	defaultEfiType         = string(pveSDK.EfiDiskType4M)
	defaultFormat          = string(pveSDK.QemuDiskFormat_Raw)
	defaultPreEnrolledKeys = false
)

func Schema() *schema.Schema {
	return &schema.Schema{
		Type:     schema.TypeList,
		Optional: true,
		MaxItems: 1,
		Elem: &schema.Resource{
			Schema: map[string]*schema.Schema{
				schemaEfiType: {
					Type:     schema.TypeString,
					Optional: true,
					Default:  defaultEfiType,
					ForceNew: true,
					ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
						return diag.FromErr(pveSDK.EfiDiskType(i.(string)).Validate())
					}},
				schemaFormat: {
					Type:     schema.TypeString,
					Optional: true,
					Default:  defaultFormat,
					ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
						return diag.FromErr(pveSDK.QemuDiskFormat(i.(string)).Validate())
					}},
				schemaPreEnrolledKeys: {
					Type:     schema.TypeBool,
					Optional: true,
					Default:  defaultPreEnrolledKeys,
					ForceNew: true},
				schemaStorage: {
					Type:     schema.TypeString,
					Required: true,
					ValidateDiagFunc: func(i any, p cty.Path) diag.Diagnostics {
						return diag.FromErr(pveSDK.StorageName(i.(string)).Validate())
					}},
			}}}
}
