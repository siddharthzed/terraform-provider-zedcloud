// Copyright (c) 2018-2021 Zededa, Inc.
// SPDX-License-Identifier: Apache-2.0

package provider

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	zschemas "github.com/zededa/terraform-provider-zedcloud/schemas"
	zedcloudapi "github.com/zededa/zedcloud-api"
	"github.com/zededa/zedcloud-api/swagger_models"
	"log"
)

var EdgeAppDataSourceSchema = &schema.Resource{
	ReadContext: readDataSourceEdgeApp,
	Schema:      zschemas.EdgeAppSchema,
	Description: "Schema for data source zedcloud_edgeapp. Must specify id or name",
}

// The schema for a resource group data source
func getEdgeAppDataSourceSchema() *schema.Resource {
	return EdgeAppDataSourceSchema
}

func getEdgeApp(client *zedcloudapi.Client,
	name, id string) (*swagger_models.App, error) {
	rspData := &swagger_models.App{}
	err := client.GetObj(edgeAppUrlExtension, name, id, false, rspData)
	if err != nil {
		return nil, fmt.Errorf("[ERROR] FAILED to get EdgeApp %s ( id: %s). Err: %s",
			name, id, err.Error())
	}
	return rspData, nil
}

func flattenAppResources(cfg []*swagger_models.Resource) []interface{} {
}

func flattenAppPermission(cfg []*swagger_models.Permission) []interface{} {
}

func flattenAppModule(cfg *swagger_models.ModuleDetail) []interface{} {
}

func flattenAppInterfaces(cfg *swagger_models.Interface) []interface{} {
}

func flattenAppImages(cfg *swagger_models.VMManifestImage) []interface{} {
}

func flattenDescDetails(cfg *swagger_models.Details) []interface{} {
}

func flattenContainerDetail(cfg *swagger_models.ContainerDetail) []interface{} {
}

func flattenUserDataTemplate(cfg *swagger_models.UserDataTemplate) []interface{} {
}

func flattenVmManifest(cfg *swagger_models.VMManifest) []interface{} {
    if cfg == nil {
        return []interface{}{}
    }
    return []interface{}{map[string]interface{}{
        "apptype" : ptrValStr(cfg.AppType),
        "configuration" : flattenUserDataTemplate(cfg.Configuration),
        "container_detail" : flattenContainerDetailTemplate(cfg.ContainerDetail).
        "deployment_type" : ptrValStr(cfg.DeploymentType),
        "desc_details" : flattenDescDetails(cfg.Desc),
        "description" : cfg.Description,
        "display_name" : cfg.DisplayName,
        "enablevnc" : cfg.Enablevnc,
        "images" : flattenAppImages(cfg.Images),
        "interfaces" : flattenAppInterfaces(cfg.Interfaces),
        "module" : flattenAppModule(cfg.Module),
        "name" : cfg.Name,
        "owner" : flattenAppAuthor(cfg.Owner),
        "permissions" : flattenAppPermission(cfg.Permissions),
        "resources" : flattenAppResources(cfg.Resources),
        "vmmode" : ptrValStr(cfg.Vmmode),
    }
}

func flattenEdgeAppConfig(cfg *swagger_models.App, computedOnly bool) map[string]interface{} {
	data := map[string]interface{}{
		"id":            cfg.ID,
		"origin_type":   ptrValStr(cfg.OriginType),
		"parent_detail": flattenObjectParentDetail(cfg.ParentDetail),
		"revision":      flattenObjectRevision(cfg.Revision),
	}
	if !computedOnly {
		data["name"] = ptrValStr(cfg.Name)
		data["title"] = ptrValStr(cfg.Title)
		data["description"] = cfg.Description
		data["cpus"] = int(cfg.Cpus)
		data["drives"] = int(cfg.Drives)
		data["manifest"] = flattenVmManifest(cfg.ManifestJSON)
		data["memory"] = int(cfg.Memory)
		data["networks"] = int(cfg.Networks)
		data["storage"] = int(cfg.Storage)
		data["user_defined_version"] = cfg.UserDefinedVersion
	}
	flattenedDataCheckKeys(zschemas.EdgeAppSchema, data, computedOnly)
	return data
}

// Read the Resource Group
func readEdgeApp(ctx context.Context, d *schema.ResourceData, meta interface{},
	resource bool) diag.Diagnostics {
	var diags diag.Diagnostics

	client := (meta.(Client)).Client
	id := rdEntryStr(d, "id")
	name := rdEntryStr(d, "name")

	if client == nil {
		return diag.Errorf("nil Client.")
	}
	log.Printf("PROVIDER:  readEdgeAppResource - id: %s, name: %s", id, name)
	if (id == "") && (name == "") {
		return diag.Errorf("The arguments \"id\" or \"name\" are required, but no definition was found.")
	}
	cfg, err := getEdgeApp(client, name, id)
	if err != nil {
		return diag.Errorf("[ERROR] Edge App %s (id: %s) not found. Err: %s",
			name, id, err.Error())
	}

	// Take the Config and convert it to terraform object
	marshalData(d, flattenEdgeAppConfig(cfg, resource))
	return diags
}

func readDataSourceEdgeApp(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	return readEdgeApp(ctx, d, meta, false)
}
