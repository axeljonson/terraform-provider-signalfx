/*
 * Metrics Metadata API
 *
 * API for creating, retrieving, updating, and deleting metric names and MTS metadata.<br> **NOTE:*() Although you can't set custom properties or tags for a metric, you *can* retrieve them for metrics and metric time series (**MTS**).
 *
 * API version: 3.2.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package metrics_metadata

// The metadata for a single tag:   * For a GET request, this is the metadata for an existing tag     retrieved from the `/tag` or `/tag/{name}` endpoint   * For a PUT request, this is the metadata for the tag specified by     the `{name}` path parameter of the `/tag/{name}` endpoint, after     the tag has been created or updated.
type Tag struct {
	// The tag. Tags are in UTF-8 format. The section [Tags Criteria](https://developers.signalfx.com/metrics/metrics_metadata_overview.html#_tags_criteria) lists the requirements for tags.
	Name string `json:"name,omitempty"`
	// A description of the tag. You can use up to 1024 UTF-8 characters.
	Description string `json:"description,omitempty"`
	// The custom properties for the tag, in the form of a JSON object (dictionary). Each property is a custom property name and value. The section [Custom Properties Criteria](https://developers.signalfx.com/metrics/metrics_metadata_overview.html#_custom_properties_criteria) lists the requirements for custom property names and values.
	CustomProperties map[string]string `json:"customProperties,omitempty"`
	// The time that the tag was created, in Unix time UTC-relative.
	Created int64 `json:"created,omitempty"`
	// SignalFx ID of the user who created the tag. If the value is \"AAAAAAAAAAA\", SignalFx created the tag.
	Creator string `json:"creator,omitempty"`
	// The time that the tag was last updated, in Unix time UTC-relative
	LastUpdated int64 `json:"lastUpdated,omitempty"`
	// SignalFx ID of the user who last updated the tag. If the value is \"AAAAAAAAAAA\", SignalFx last updated the metric.
	LastUpdatedBy string `json:"lastUpdatedBy,omitempty"`
}
