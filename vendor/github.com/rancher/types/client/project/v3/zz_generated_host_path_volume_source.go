package client

const (
	HostPathVolumeSourceType      = "hostPathVolumeSource"
	HostPathVolumeSourceFieldPath = "path"
	HostPathVolumeSourceFieldType = "type"
)

type HostPathVolumeSource struct {
	Path string `json:"path,omitempty"`
	Type string `json:"type,omitempty"`
}
