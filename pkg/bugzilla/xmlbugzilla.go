/* This file has been autogenerated by tools/gen_xml_code.sh, please
   add more bugzillas to the list if you find any missing field
*/

package bugzilla

import "encoding/xml"
type Cactual_time struct {
	XMLName xml.Name `xml:"actual_time,omitempty" json:"actual_time,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type Calias struct {
	XMLName xml.Name `xml:"alias,omitempty" json:"alias,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cassigned_to struct {
	XMLName xml.Name `xml:"assigned_to,omitempty" json:"assigned_to,omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cblocked struct {
	XMLName xml.Name `xml:"blocked,omitempty" json:"blocked,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type Cbug struct {
	XMLName xml.Name `xml:"bug,omitempty" json:"bug,omitempty"`
	Cactual_time *Cactual_time `xml:"actual_time,omitempty" json:"actual_time,omitempty"`
	Calias *Calias `xml:"alias,omitempty" json:"alias,omitempty"`
	Cassigned_to *Cassigned_to `xml:"assigned_to,omitempty" json:"assigned_to,omitempty"`
	Cblocked []*Cblocked `xml:"blocked,omitempty" json:"blocked,omitempty"`
	Cbug_file_loc *Cbug_file_loc `xml:"bug_file_loc,omitempty" json:"bug_file_loc,omitempty"`
	Cbug_id *Cbug_id `xml:"bug_id,omitempty" json:"bug_id,omitempty"`
	Cbug_severity *Cbug_severity `xml:"bug_severity,omitempty" json:"bug_severity,omitempty"`
	Cbug_status *Cbug_status `xml:"bug_status,omitempty" json:"bug_status,omitempty"`
	Ccc []*Ccc `xml:"cc,omitempty" json:"cc,omitempty"`
	Ccclist_accessible *Ccclist_accessible `xml:"cclist_accessible,omitempty" json:"cclist_accessible,omitempty"`
	Ccf_build_id *Ccf_build_id `xml:"cf_build_id,omitempty" json:"cf_build_id,omitempty"`
	Ccf_category *Ccf_category `xml:"cf_category,omitempty" json:"cf_category,omitempty"`
	Ccf_clone_of *Ccf_clone_of `xml:"cf_clone_of,omitempty" json:"cf_clone_of,omitempty"`
	Ccf_cloudforms_team *Ccf_cloudforms_team `xml:"cf_cloudforms_team,omitempty" json:"cf_cloudforms_team,omitempty"`
	Ccf_compliance_control_group *Ccf_compliance_control_group `xml:"cf_compliance_control_group,omitempty" json:"cf_compliance_control_group,omitempty"`
	Ccf_compliance_level *Ccf_compliance_level `xml:"cf_compliance_level,omitempty" json:"cf_compliance_level,omitempty"`
	Ccf_crm *Ccf_crm `xml:"cf_crm,omitempty" json:"cf_crm,omitempty"`
	Ccf_cust_facing *Ccf_cust_facing `xml:"cf_cust_facing,omitempty" json:"cf_cust_facing,omitempty"`
	Ccf_devel_whiteboard *Ccf_devel_whiteboard `xml:"cf_devel_whiteboard,omitempty" json:"cf_devel_whiteboard,omitempty"`
	Ccf_doc_type *Ccf_doc_type `xml:"cf_doc_type,omitempty" json:"cf_doc_type,omitempty"`
	Ccf_docs_score *Ccf_docs_score `xml:"cf_docs_score,omitempty" json:"cf_docs_score,omitempty"`
	Ccf_documentation_action *Ccf_documentation_action `xml:"cf_documentation_action,omitempty" json:"cf_documentation_action,omitempty"`
	Ccf_environment *Ccf_environment `xml:"cf_environment,omitempty" json:"cf_environment,omitempty"`
	Ccf_fixed_in *Ccf_fixed_in `xml:"cf_fixed_in,omitempty" json:"cf_fixed_in,omitempty"`
	Ccf_internal_whiteboard *Ccf_internal_whiteboard `xml:"cf_internal_whiteboard,omitempty" json:"cf_internal_whiteboard,omitempty"`
	Ccf_last_closed *Ccf_last_closed `xml:"cf_last_closed,omitempty" json:"cf_last_closed,omitempty"`
	Ccf_mount_type *Ccf_mount_type `xml:"cf_mount_type,omitempty" json:"cf_mount_type,omitempty"`
	Ccf_ovirt_team *Ccf_ovirt_team `xml:"cf_ovirt_team,omitempty" json:"cf_ovirt_team,omitempty"`
	Ccf_pgm_internal *Ccf_pgm_internal `xml:"cf_pgm_internal,omitempty" json:"cf_pgm_internal,omitempty"`
	Ccf_pm_score *Ccf_pm_score `xml:"cf_pm_score,omitempty" json:"cf_pm_score,omitempty"`
	Ccf_regression_status *Ccf_regression_status `xml:"cf_regression_status,omitempty" json:"cf_regression_status,omitempty"`
	Ccf_release_notes *Ccf_release_notes `xml:"cf_release_notes,omitempty" json:"cf_release_notes,omitempty"`
	Ccf_show_homepage *Ccf_show_homepage `xml:"cf_show_homepage,omitempty" json:"cf_show_homepage,omitempty"`
	Ccf_story_points *Ccf_story_points `xml:"cf_story_points,omitempty" json:"cf_story_points,omitempty"`
	Ccf_subsystem_team *Ccf_subsystem_team `xml:"cf_subsystem_team,omitempty" json:"cf_subsystem_team,omitempty"`
	Ccf_type *Ccf_type `xml:"cf_type,omitempty" json:"cf_type,omitempty"`
	Ccf_verified *Ccf_verified `xml:"cf_verified,omitempty" json:"cf_verified,omitempty"`
	Ccf_verified_branch *Ccf_verified_branch `xml:"cf_verified_branch,omitempty" json:"cf_verified_branch,omitempty"`
	Cclassification *Cclassification `xml:"classification,omitempty" json:"classification,omitempty"`
	Cclassification_id *Cclassification_id `xml:"classification_id,omitempty" json:"classification_id,omitempty"`
	Ccomment_sort_order *Ccomment_sort_order `xml:"comment_sort_order,omitempty" json:"comment_sort_order,omitempty"`
	Ccomponent *Ccomponent `xml:"component,omitempty" json:"component,omitempty"`
	Ccreation_ts *Ccreation_ts `xml:"creation_ts,omitempty" json:"creation_ts,omitempty"`
	Cdelta_ts *Cdelta_ts `xml:"delta_ts,omitempty" json:"delta_ts,omitempty"`
	Cdependson []*Cdependson `xml:"dependson,omitempty" json:"dependson,omitempty"`
	Cestimated_time *Cestimated_time `xml:"estimated_time,omitempty" json:"estimated_time,omitempty"`
	Ceverconfirmed *Ceverconfirmed `xml:"everconfirmed,omitempty" json:"everconfirmed,omitempty"`
	Cexternal_bugs []*Cexternal_bugs `xml:"external_bugs,omitempty" json:"external_bugs,omitempty"`
	Cflag []*Cflag `xml:"flag,omitempty" json:"flag,omitempty"`
	Ckeywords *Ckeywords `xml:"keywords,omitempty" json:"keywords,omitempty"`
	Clong_desc []*Clong_desc `xml:"long_desc,omitempty" json:"long_desc,omitempty"`
	Cop_sys *Cop_sys `xml:"op_sys,omitempty" json:"op_sys,omitempty"`
	Cpriority *Cpriority `xml:"priority,omitempty" json:"priority,omitempty"`
	Cproduct *Cproduct `xml:"product,omitempty" json:"product,omitempty"`
	Cqa_contact *Cqa_contact `xml:"qa_contact,omitempty" json:"qa_contact,omitempty"`
	Cremaining_time *Cremaining_time `xml:"remaining_time,omitempty" json:"remaining_time,omitempty"`
	Crep_platform *Crep_platform `xml:"rep_platform,omitempty" json:"rep_platform,omitempty"`
	Creporter *Creporter `xml:"reporter,omitempty" json:"reporter,omitempty"`
	Creporter_accessible *Creporter_accessible `xml:"reporter_accessible,omitempty" json:"reporter_accessible,omitempty"`
	Cresolution *Cresolution `xml:"resolution,omitempty" json:"resolution,omitempty"`
	Cshort_desc *Cshort_desc `xml:"short_desc,omitempty" json:"short_desc,omitempty"`
	Cstatus_whiteboard *Cstatus_whiteboard `xml:"status_whiteboard,omitempty" json:"status_whiteboard,omitempty"`
	Ctarget_milestone *Ctarget_milestone `xml:"target_milestone,omitempty" json:"target_milestone,omitempty"`
	Ctarget_release *Ctarget_release `xml:"target_release,omitempty" json:"target_release,omitempty"`
	Ctoken *Ctoken `xml:"token,omitempty" json:"token,omitempty"`
	Cversion *Cversion `xml:"version,omitempty" json:"version,omitempty"`
	Cvotes *Cvotes `xml:"votes,omitempty" json:"votes,omitempty"`
}

type Cbug_file_loc struct {
	XMLName xml.Name `xml:"bug_file_loc,omitempty" json:"bug_file_loc,omitempty"`
}

type Cbug_id struct {
	XMLName xml.Name `xml:"bug_id,omitempty" json:"bug_id,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type Cbug_severity struct {
	XMLName xml.Name `xml:"bug_severity,omitempty" json:"bug_severity,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cbug_status struct {
	XMLName xml.Name `xml:"bug_status,omitempty" json:"bug_status,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cbug_when struct {
	XMLName xml.Name `xml:"bug_when,omitempty" json:"bug_when,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cbugzilla struct {
	XMLName xml.Name `xml:"bugzilla,omitempty" json:"bugzilla,omitempty"`
	Attrexporter string`xml:"exporter,attr"  json:",omitempty"`
	Attrmaintainer string`xml:"maintainer,attr"  json:",omitempty"`
	Attrurlbase string`xml:"urlbase,attr"  json:",omitempty"`
	Attrversion string`xml:"version,attr"  json:",omitempty"`
	Cbug *Cbug `xml:"bug,omitempty" json:"bug,omitempty"`
}

type Ccc struct {
	XMLName xml.Name `xml:"cc,omitempty" json:"cc,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccclist_accessible struct {
	XMLName xml.Name `xml:"cclist_accessible,omitempty" json:"cclist_accessible,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Ccf_build_id struct {
	XMLName xml.Name `xml:"cf_build_id,omitempty" json:"cf_build_id,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_category struct {
	XMLName xml.Name `xml:"cf_category,omitempty" json:"cf_category,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_clone_of struct {
	XMLName xml.Name `xml:"cf_clone_of,omitempty" json:"cf_clone_of,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type Ccf_cloudforms_team struct {
	XMLName xml.Name `xml:"cf_cloudforms_team,omitempty" json:"cf_cloudforms_team,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_compliance_control_group struct {
	XMLName xml.Name `xml:"cf_compliance_control_group,omitempty" json:"cf_compliance_control_group,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_compliance_level struct {
	XMLName xml.Name `xml:"cf_compliance_level,omitempty" json:"cf_compliance_level,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_crm struct {
	XMLName xml.Name `xml:"cf_crm,omitempty" json:"cf_crm,omitempty"`
}

type Ccf_cust_facing struct {
	XMLName xml.Name `xml:"cf_cust_facing,omitempty" json:"cf_cust_facing,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_devel_whiteboard struct {
	XMLName xml.Name `xml:"cf_devel_whiteboard,omitempty" json:"cf_devel_whiteboard,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_doc_type struct {
	XMLName xml.Name `xml:"cf_doc_type,omitempty" json:"cf_doc_type,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_docs_score struct {
	XMLName xml.Name `xml:"cf_docs_score,omitempty" json:"cf_docs_score,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Ccf_documentation_action struct {
	XMLName xml.Name `xml:"cf_documentation_action,omitempty" json:"cf_documentation_action,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_environment struct {
	XMLName xml.Name `xml:"cf_environment,omitempty" json:"cf_environment,omitempty"`
}

type Ccf_fixed_in struct {
	XMLName xml.Name `xml:"cf_fixed_in,omitempty" json:"cf_fixed_in,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_internal_whiteboard struct {
	XMLName xml.Name `xml:"cf_internal_whiteboard,omitempty" json:"cf_internal_whiteboard,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_last_closed struct {
	XMLName xml.Name `xml:"cf_last_closed,omitempty" json:"cf_last_closed,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_mount_type struct {
	XMLName xml.Name `xml:"cf_mount_type,omitempty" json:"cf_mount_type,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_ovirt_team struct {
	XMLName xml.Name `xml:"cf_ovirt_team,omitempty" json:"cf_ovirt_team,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_pgm_internal struct {
	XMLName xml.Name `xml:"cf_pgm_internal,omitempty" json:"cf_pgm_internal,omitempty"`
}

type Ccf_pm_score struct {
	XMLName xml.Name `xml:"cf_pm_score,omitempty" json:"cf_pm_score,omitempty"`
	Number int16 `xml:",chardata" json:",omitempty"`
}

type Ccf_regression_status struct {
	XMLName xml.Name `xml:"cf_regression_status,omitempty" json:"cf_regression_status,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_release_notes struct {
	XMLName xml.Name `xml:"cf_release_notes,omitempty" json:"cf_release_notes,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_show_homepage struct {
	XMLName xml.Name `xml:"cf_show_homepage,omitempty" json:"cf_show_homepage,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_story_points struct {
	XMLName xml.Name `xml:"cf_story_points,omitempty" json:"cf_story_points,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_subsystem_team struct {
	XMLName xml.Name `xml:"cf_subsystem_team,omitempty" json:"cf_subsystem_team,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_type struct {
	XMLName xml.Name `xml:"cf_type,omitempty" json:"cf_type,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_verified struct {
	XMLName xml.Name `xml:"cf_verified,omitempty" json:"cf_verified,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccf_verified_branch struct {
	XMLName xml.Name `xml:"cf_verified_branch,omitempty" json:"cf_verified_branch,omitempty"`
}

type Cclassification struct {
	XMLName xml.Name `xml:"classification,omitempty" json:"classification,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cclassification_id struct {
	XMLName xml.Name `xml:"classification_id,omitempty" json:"classification_id,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Ccomment_count struct {
	XMLName xml.Name `xml:"comment_count,omitempty" json:"comment_count,omitempty"`
	Number int8 `xml:",chardata" json:",omitempty"`
}

type Ccomment_sort_order struct {
	XMLName xml.Name `xml:"comment_sort_order,omitempty" json:"comment_sort_order,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccommentid struct {
	XMLName xml.Name `xml:"commentid,omitempty" json:"commentid,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type Ccomponent struct {
	XMLName xml.Name `xml:"component,omitempty" json:"component,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ccreation_ts struct {
	XMLName xml.Name `xml:"creation_ts,omitempty" json:"creation_ts,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cdelta_ts struct {
	XMLName xml.Name `xml:"delta_ts,omitempty" json:"delta_ts,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cdependson struct {
	XMLName xml.Name `xml:"dependson,omitempty" json:"dependson,omitempty"`
	Number int32 `xml:",chardata" json:",omitempty"`
}

type Cestimated_time struct {
	XMLName xml.Name `xml:"estimated_time,omitempty" json:"estimated_time,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type Ceverconfirmed struct {
	XMLName xml.Name `xml:"everconfirmed,omitempty" json:"everconfirmed,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Cexternal_bugs struct {
	XMLName xml.Name `xml:"external_bugs,omitempty" json:"external_bugs,omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cflag struct {
	XMLName xml.Name `xml:"flag,omitempty" json:"flag,omitempty"`
	Attrid string`xml:"id,attr"  json:",omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Attrsetter string`xml:"setter,attr"  json:",omitempty"`
	Attrstatus string`xml:"status,attr"  json:",omitempty"`
	Attrtype_id string`xml:"type_id,attr"  json:",omitempty"`
}

type Ckeywords struct {
	XMLName xml.Name `xml:"keywords,omitempty" json:"keywords,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Clong_desc struct {
	XMLName xml.Name `xml:"long_desc,omitempty" json:"long_desc,omitempty"`
	Attrisprivate string`xml:"isprivate,attr"  json:",omitempty"`
	Cbug_when *Cbug_when `xml:"bug_when,omitempty" json:"bug_when,omitempty"`
	Ccomment_count *Ccomment_count `xml:"comment_count,omitempty" json:"comment_count,omitempty"`
	Ccommentid *Ccommentid `xml:"commentid,omitempty" json:"commentid,omitempty"`
	Cthetext *Cthetext `xml:"thetext,omitempty" json:"thetext,omitempty"`
	Cwho *Cwho `xml:"who,omitempty" json:"who,omitempty"`
}

type Cop_sys struct {
	XMLName xml.Name `xml:"op_sys,omitempty" json:"op_sys,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cpriority struct {
	XMLName xml.Name `xml:"priority,omitempty" json:"priority,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cproduct struct {
	XMLName xml.Name `xml:"product,omitempty" json:"product,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cqa_contact struct {
	XMLName xml.Name `xml:"qa_contact,omitempty" json:"qa_contact,omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cremaining_time struct {
	XMLName xml.Name `xml:"remaining_time,omitempty" json:"remaining_time,omitempty"`
	Number float32 `xml:",chardata" json:",omitempty"`
}

type Crep_platform struct {
	XMLName xml.Name `xml:"rep_platform,omitempty" json:"rep_platform,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Creporter struct {
	XMLName xml.Name `xml:"reporter,omitempty" json:"reporter,omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Creporter_accessible struct {
	XMLName xml.Name `xml:"reporter_accessible,omitempty" json:"reporter_accessible,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Cresolution struct {
	XMLName xml.Name `xml:"resolution,omitempty" json:"resolution,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cshort_desc struct {
	XMLName xml.Name `xml:"short_desc,omitempty" json:"short_desc,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cstatus_whiteboard struct {
	XMLName xml.Name `xml:"status_whiteboard,omitempty" json:"status_whiteboard,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ctarget_milestone struct {
	XMLName xml.Name `xml:"target_milestone,omitempty" json:"target_milestone,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ctarget_release struct {
	XMLName xml.Name `xml:"target_release,omitempty" json:"target_release,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cthetext struct {
	XMLName xml.Name `xml:"thetext,omitempty" json:"thetext,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Ctoken struct {
	XMLName xml.Name `xml:"token,omitempty" json:"token,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cversion struct {
	XMLName xml.Name `xml:"version,omitempty" json:"version,omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

type Cvotes struct {
	XMLName xml.Name `xml:"votes,omitempty" json:"votes,omitempty"`
	Flag bool `xml:",chardata" json:",omitempty"`
}

type Cwho struct {
	XMLName xml.Name `xml:"who,omitempty" json:"who,omitempty"`
	Attrname string`xml:"name,attr"  json:",omitempty"`
	Content string `xml:",chardata" json:",omitempty"`
}

