// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This code is generated by Magic Modules using the following:
//
//     Configuration: https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/products/alloydb/Backup.yaml
//     Template:      https://github.com/GoogleCloudPlatform/magic-modules/tree/main/mmv1/templates/terraform/resource.go.tmpl
//
//     DO NOT EDIT this file directly. Any changes made to this file will be
//     overwritten during the next generation cycle.
//
// ----------------------------------------------------------------------------

package alloydb

import (
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strings"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/customdiff"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"

	"github.com/hashicorp/terraform-provider-google/google/tpgresource"
	transport_tpg "github.com/hashicorp/terraform-provider-google/google/transport"
	"github.com/hashicorp/terraform-provider-google/google/verify"
)

func ResourceAlloydbBackup() *schema.Resource {
	return &schema.Resource{
		Create: resourceAlloydbBackupCreate,
		Read:   resourceAlloydbBackupRead,
		Update: resourceAlloydbBackupUpdate,
		Delete: resourceAlloydbBackupDelete,

		Importer: &schema.ResourceImporter{
			State: resourceAlloydbBackupImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(10 * time.Minute),
			Update: schema.DefaultTimeout(10 * time.Minute),
			Delete: schema.DefaultTimeout(10 * time.Minute),
		},

		CustomizeDiff: customdiff.All(
			tpgresource.SetLabelsDiff,
			tpgresource.SetAnnotationsDiff,
			tpgresource.DefaultProviderProject,
		),

		Schema: map[string]*schema.Schema{
			"backup_id": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The ID of the alloydb backup.`,
			},
			"cluster_name": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: tpgresource.ProjectNumberDiffSuppress,
				Description:      `The full resource name of the backup source cluster (e.g., projects/{project}/locations/{location}/clusters/{clusterId}).`,
			},
			"location": {
				Type:        schema.TypeString,
				Required:    true,
				ForceNew:    true,
				Description: `The location where the alloydb backup should reside.`,
			},
			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `Annotations to allow client tools to store small amount of arbitrary data. This is distinct from labels. https://google.aip.dev/128
An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.


**Note**: This field is non-authoritative, and will only manage the annotations present in your configuration.
Please refer to the field 'effective_annotations' for all of the annotations present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-provided description of the backup.`,
			},
			"display_name": {
				Type:        schema.TypeString,
				Optional:    true,
				Description: `User-settable and human-readable display name for the Backup.`,
			},
			"encryption_config": {
				Type:        schema.TypeList,
				Optional:    true,
				Description: `EncryptionConfig describes the encryption config of a cluster or a backup that is encrypted with a CMEK (customer-managed encryption key).`,
				MaxItems:    1,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kms_key_name": {
							Type:        schema.TypeString,
							Optional:    true,
							ForceNew:    true,
							Description: `The fully-qualified resource name of the KMS key. Each Cloud KMS key is regionalized and has the following format: projects/[PROJECT]/locations/[REGION]/keyRings/[RING]/cryptoKeys/[KEY_NAME].`,
						},
					},
				},
			},
			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
				Description: `User-defined labels for the alloydb backup. An object containing a list of "key": value pairs. Example: { "name": "wrench", "mass": "1.3kg", "count": "3" }.


**Note**: This field is non-authoritative, and will only manage the labels present in your configuration.
Please refer to the field 'effective_labels' for all of the labels present on the resource.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"type": {
				Type:         schema.TypeString,
				Computed:     true,
				Optional:     true,
				ValidateFunc: verify.ValidateEnum([]string{"TYPE_UNSPECIFIED", "ON_DEMAND", "AUTOMATED", "CONTINUOUS", ""}),
				Description:  `The backup type, which suggests the trigger for the backup. Possible values: ["TYPE_UNSPECIFIED", "ON_DEMAND", "AUTOMATED", "CONTINUOUS"]`,
			},
			"cluster_uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The system-generated UID of the cluster which was used to create this resource.`,
			},
			"create_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Create time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"delete_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Delete time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"effective_annotations": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of annotations (key/value pairs) present on the resource in GCP, including the annotations configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"effective_labels": {
				Type:        schema.TypeMap,
				Computed:    true,
				Description: `All of labels (key/value pairs) present on the resource in GCP, including the labels configured through Terraform, other clients and services.`,
				Elem:        &schema.Schema{Type: schema.TypeString},
			},
			"encryption_info": {
				Type:        schema.TypeList,
				Computed:    true,
				Description: `EncryptionInfo describes the encryption information of a cluster or a backup.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"encryption_type": {
							Type:        schema.TypeString,
							Computed:    true,
							Description: `Output only. Type of encryption.`,
						},
						"kms_key_versions": {
							Type:        schema.TypeList,
							Computed:    true,
							Description: `Output only. Cloud KMS key versions that are being used to protect the database or the backup.`,
							Elem: &schema.Schema{
								Type: schema.TypeString,
							},
						},
					},
				},
			},
			"etag": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `For Resource freshness validation (https://google.aip.dev/154)`,
			},
			"expiry_quantity": {
				Type:     schema.TypeList,
				Computed: true,
				Description: `Output only. The QuantityBasedExpiry of the backup, specified by the backup's retention policy.
Once the expiry quantity is over retention, the backup is eligible to be garbage collected.`,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"retention_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Output only. The backup's position among its backups with the same source cluster and type, by descending chronological order create time (i.e. newest first).`,
						},
						"total_retention_count": {
							Type:        schema.TypeInt,
							Computed:    true,
							Description: `Output only. The length of the quantity-based queue, specified by the backup's retention policy.`,
						},
					},
				},
			},
			"expiry_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. The time at which after the backup is eligible to be garbage collected.
It is the duration specified by the backup's retention policy, added to the backup's createTime.`,
			},
			"name": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The name of the backup resource with the format: * projects/{project}/locations/{region}/backups/{backupId}`,
			},
			"reconciling": {
				Type:     schema.TypeBool,
				Computed: true,
				Description: `Output only. Reconciling (https://google.aip.dev/128#reconciliation), if true, indicates that the service is actively updating the resource.
This can happen due to user-triggered updates or system actions like failover or maintenance.`,
			},
			"size_bytes": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The size of the backup in bytes.`,
			},
			"state": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The current state of the backup.`,
			},
			"terraform_labels": {
				Type:     schema.TypeMap,
				Computed: true,
				Description: `The combination of labels configured directly on the resource
 and default labels configured on the provider.`,
				Elem: &schema.Schema{Type: schema.TypeString},
			},
			"uid": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Output only. The system-generated UID of the resource. The UID is assigned when the resource is created, and it is retained until it is deleted.`,
			},
			"update_time": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Output only. Update time stamp. A timestamp in RFC3339 UTC "Zulu" format, with nanosecond resolution and up to nine fractional digits.
Examples: "2014-10-02T15:01:23Z" and "2014-10-02T15:01:23.045123456Z".`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceAlloydbBackupCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	displayNameProp, err := expandAlloydbBackupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeProp, err := expandAlloydbBackupType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(typeProp)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	descriptionProp, err := expandAlloydbBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	clusterNameProp, err := expandAlloydbBackupClusterName(d.Get("cluster_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("cluster_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(clusterNameProp)) && (ok || !reflect.DeepEqual(v, clusterNameProp)) {
		obj["clusterName"] = clusterNameProp
	}
	encryptionConfigProp, err := expandAlloydbBackupEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(encryptionConfigProp)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	labelsProp, err := expandAlloydbBackupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(labelsProp)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbBackupEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(annotationsProp)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	obj, err = resourceAlloydbBackupEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups?backupId={{backup_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new Backup: %#v", obj)
	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "POST",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutCreate),
		Headers:   headers,
	})
	if err != nil {
		return fmt.Errorf("Error creating Backup: %s", err)
	}

	// Store the ID now
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = AlloydbOperationWaitTime(
		config, res, project, "Creating Backup", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create Backup: %s", err)
	}

	log.Printf("[DEBUG] Finished creating Backup %q: %#v", d.Id(), res)

	return resourceAlloydbBackupRead(d, meta)
}

func resourceAlloydbBackupRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "GET",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, fmt.Sprintf("AlloydbBackup %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	if err := d.Set("name", flattenAlloydbBackupName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("display_name", flattenAlloydbBackupDisplayName(res["displayName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("uid", flattenAlloydbBackupUid(res["uid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("create_time", flattenAlloydbBackupCreateTime(res["createTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("update_time", flattenAlloydbBackupUpdateTime(res["updateTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("delete_time", flattenAlloydbBackupDeleteTime(res["deleteTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("labels", flattenAlloydbBackupLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("state", flattenAlloydbBackupState(res["state"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("type", flattenAlloydbBackupType(res["type"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("description", flattenAlloydbBackupDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("cluster_uid", flattenAlloydbBackupClusterUid(res["clusterUid"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("cluster_name", flattenAlloydbBackupClusterName(res["clusterName"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("reconciling", flattenAlloydbBackupReconciling(res["reconciling"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("encryption_config", flattenAlloydbBackupEncryptionConfig(res["encryptionConfig"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("encryption_info", flattenAlloydbBackupEncryptionInfo(res["encryptionInfo"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("etag", flattenAlloydbBackupEtag(res["etag"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("annotations", flattenAlloydbBackupAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("size_bytes", flattenAlloydbBackupSizeBytes(res["sizeBytes"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("expiry_time", flattenAlloydbBackupExpiryTime(res["expiryTime"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("expiry_quantity", flattenAlloydbBackupExpiryQuantity(res["expiryQuantity"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("terraform_labels", flattenAlloydbBackupTerraformLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("effective_labels", flattenAlloydbBackupEffectiveLabels(res["labels"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}
	if err := d.Set("effective_annotations", flattenAlloydbBackupEffectiveAnnotations(res["annotations"], d, config)); err != nil {
		return fmt.Errorf("Error reading Backup: %s", err)
	}

	return nil
}

func resourceAlloydbBackupUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	displayNameProp, err := expandAlloydbBackupDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("display_name"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	typeProp, err := expandAlloydbBackupType(d.Get("type"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("type"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, typeProp)) {
		obj["type"] = typeProp
	}
	descriptionProp, err := expandAlloydbBackupDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	encryptionConfigProp, err := expandAlloydbBackupEncryptionConfig(d.Get("encryption_config"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("encryption_config"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, encryptionConfigProp)) {
		obj["encryptionConfig"] = encryptionConfigProp
	}
	labelsProp, err := expandAlloydbBackupEffectiveLabels(d.Get("effective_labels"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_labels"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, labelsProp)) {
		obj["labels"] = labelsProp
	}
	annotationsProp, err := expandAlloydbBackupEffectiveAnnotations(d.Get("effective_annotations"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("effective_annotations"); !tpgresource.IsEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, annotationsProp)) {
		obj["annotations"] = annotationsProp
	}

	obj, err = resourceAlloydbBackupEncoder(d, meta, obj)
	if err != nil {
		return err
	}

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating Backup %q: %#v", d.Id(), obj)
	headers := make(http.Header)
	updateMask := []string{}

	if d.HasChange("display_name") {
		updateMask = append(updateMask, "displayName")
	}

	if d.HasChange("type") {
		updateMask = append(updateMask, "type")
	}

	if d.HasChange("description") {
		updateMask = append(updateMask, "description")
	}

	if d.HasChange("encryption_config") {
		updateMask = append(updateMask, "encryptionConfig")
	}

	if d.HasChange("effective_labels") {
		updateMask = append(updateMask, "labels")
	}

	if d.HasChange("effective_annotations") {
		updateMask = append(updateMask, "annotations")
	}
	// updateMask is a URL parameter but not present in the schema, so ReplaceVars
	// won't set it
	url, err = transport_tpg.AddQueryParams(url, map[string]string{"updateMask": strings.Join(updateMask, ",")})
	if err != nil {
		return err
	}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	// if updateMask is empty we are not updating anything so skip the post
	if len(updateMask) > 0 {
		res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
			Config:    config,
			Method:    "PATCH",
			Project:   billingProject,
			RawURL:    url,
			UserAgent: userAgent,
			Body:      obj,
			Timeout:   d.Timeout(schema.TimeoutUpdate),
			Headers:   headers,
		})

		if err != nil {
			return fmt.Errorf("Error updating Backup %q: %s", d.Id(), err)
		} else {
			log.Printf("[DEBUG] Finished updating Backup %q: %#v", d.Id(), res)
		}

		err = AlloydbOperationWaitTime(
			config, res, project, "Updating Backup", userAgent,
			d.Timeout(schema.TimeoutUpdate))

		if err != nil {
			return err
		}
	}

	return resourceAlloydbBackupRead(d, meta)
}

func resourceAlloydbBackupDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*transport_tpg.Config)
	userAgent, err := tpgresource.GenerateUserAgentString(d, config.UserAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := tpgresource.GetProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for Backup: %s", err)
	}
	billingProject = project

	url, err := tpgresource.ReplaceVars(d, config, "{{AlloydbBasePath}}projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}

	// err == nil indicates that the billing_project value was found
	if bp, err := tpgresource.GetBillingProject(d, config); err == nil {
		billingProject = bp
	}

	headers := make(http.Header)

	log.Printf("[DEBUG] Deleting Backup %q", d.Id())
	res, err := transport_tpg.SendRequest(transport_tpg.SendRequestOptions{
		Config:    config,
		Method:    "DELETE",
		Project:   billingProject,
		RawURL:    url,
		UserAgent: userAgent,
		Body:      obj,
		Timeout:   d.Timeout(schema.TimeoutDelete),
		Headers:   headers,
	})
	if err != nil {
		return transport_tpg.HandleNotFoundError(err, d, "Backup")
	}

	err = AlloydbOperationWaitTime(
		config, res, project, "Deleting Backup", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting Backup %q: %#v", d.Id(), res)
	return nil
}

func resourceAlloydbBackupImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*transport_tpg.Config)
	if err := tpgresource.ParseImportId([]string{
		"^projects/(?P<project>[^/]+)/locations/(?P<location>[^/]+)/backups/(?P<backup_id>[^/]+)$",
		"^(?P<project>[^/]+)/(?P<location>[^/]+)/(?P<backup_id>[^/]+)$",
		"^(?P<location>[^/]+)/(?P<backup_id>[^/]+)$",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := tpgresource.ReplaceVars(d, config, "projects/{{project}}/locations/{{location}}/backups/{{backup_id}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenAlloydbBackupName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDisplayName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupCreateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupUpdateTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDeleteTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenAlloydbBackupState(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupDescription(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupClusterUid(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupClusterName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupReconciling(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionConfig(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["kms_key_name"] =
		flattenAlloydbBackupEncryptionConfigKmsKeyName(original["kmsKeyName"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupEncryptionConfigKmsKeyName(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionInfo(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["encryption_type"] =
		flattenAlloydbBackupEncryptionInfoEncryptionType(original["encryptionType"], d, config)
	transformed["kms_key_versions"] =
		flattenAlloydbBackupEncryptionInfoKmsKeyVersions(original["kmsKeyVersions"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupEncryptionInfoEncryptionType(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEncryptionInfoKmsKeyVersions(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEtag(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("annotations"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenAlloydbBackupSizeBytes(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupExpiryTime(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupExpiryQuantity(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return nil
	}
	original := v.(map[string]interface{})
	if len(original) == 0 {
		return nil
	}
	transformed := make(map[string]interface{})
	transformed["retention_count"] =
		flattenAlloydbBackupExpiryQuantityRetentionCount(original["retentionCount"], d, config)
	transformed["total_retention_count"] =
		flattenAlloydbBackupExpiryQuantityTotalRetentionCount(original["totalRetentionCount"], d, config)
	return []interface{}{transformed}
}
func flattenAlloydbBackupExpiryQuantityRetentionCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbBackupExpiryQuantityTotalRetentionCount(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	// Handles the string fixed64 format
	if strVal, ok := v.(string); ok {
		if intVal, err := tpgresource.StringToFixed64(strVal); err == nil {
			return intVal
		}
	}

	// number values are represented as float64
	if floatVal, ok := v.(float64); ok {
		intVal := int(floatVal)
		return intVal
	}

	return v // let terraform core handle it otherwise
}

func flattenAlloydbBackupTerraformLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	if v == nil {
		return v
	}

	transformed := make(map[string]interface{})
	if l, ok := d.GetOkExists("terraform_labels"); ok {
		for k := range l.(map[string]interface{}) {
			transformed[k] = v.(map[string]interface{})[k]
		}
	}

	return transformed
}

func flattenAlloydbBackupEffectiveLabels(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func flattenAlloydbBackupEffectiveAnnotations(v interface{}, d *schema.ResourceData, config *transport_tpg.Config) interface{} {
	return v
}

func expandAlloydbBackupDisplayName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupType(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupDescription(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupClusterName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupEncryptionConfig(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedKmsKeyName, err := expandAlloydbBackupEncryptionConfigKmsKeyName(original["kms_key_name"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedKmsKeyName); val.IsValid() && !tpgresource.IsEmptyValue(val) {
		transformed["kmsKeyName"] = transformedKmsKeyName
	}

	return transformed, nil
}

func expandAlloydbBackupEncryptionConfigKmsKeyName(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (interface{}, error) {
	return v, nil
}

func expandAlloydbBackupEffectiveLabels(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func expandAlloydbBackupEffectiveAnnotations(v interface{}, d tpgresource.TerraformResourceData, config *transport_tpg.Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}

func resourceAlloydbBackupEncoder(d *schema.ResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	// The only other available type is AUTOMATED which cannot be set manually
	obj["type"] = "ON_DEMAND"
	return obj, nil
}
