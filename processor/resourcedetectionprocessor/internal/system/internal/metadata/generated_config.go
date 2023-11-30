// Code generated by mdatagen. DO NOT EDIT.

package metadata

import "go.opentelemetry.io/collector/confmap"

// ResourceAttributeConfig provides common config for a particular resource attribute.
type ResourceAttributeConfig struct {
	Enabled bool `mapstructure:"enabled"`

	enabledSetByUser bool
}

func (rac *ResourceAttributeConfig) Unmarshal(parser *confmap.Conf) error {
	if parser == nil {
		return nil
	}
	err := parser.Unmarshal(rac, confmap.WithErrorUnused())
	if err != nil {
		return err
	}
	rac.enabledSetByUser = parser.IsSet("enabled")
	return nil
}

// ResourceAttributesConfig provides config for resourcedetectionprocessor/system resource attributes.
type ResourceAttributesConfig struct {
	HostArch           ResourceAttributeConfig `mapstructure:"host.arch"`
	HostCPUCacheL2Size ResourceAttributeConfig `mapstructure:"host.cpu.cache.l2.size"`
	HostCPUFamily      ResourceAttributeConfig `mapstructure:"host.cpu.family"`
	HostCPUModelID     ResourceAttributeConfig `mapstructure:"host.cpu.model.id"`
	HostCPUModelName   ResourceAttributeConfig `mapstructure:"host.cpu.model.name"`
	HostCPUStepping    ResourceAttributeConfig `mapstructure:"host.cpu.stepping"`
	HostCPUVendorID    ResourceAttributeConfig `mapstructure:"host.cpu.vendor.id"`
	HostID             ResourceAttributeConfig `mapstructure:"host.id"`
	HostIP             ResourceAttributeConfig `mapstructure:"host.ip"`
	HostName           ResourceAttributeConfig `mapstructure:"host.name"`
	OsDescription      ResourceAttributeConfig `mapstructure:"os.description"`
	OsType             ResourceAttributeConfig `mapstructure:"os.type"`
}

func DefaultResourceAttributesConfig() ResourceAttributesConfig {
	return ResourceAttributesConfig{
		HostArch: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUCacheL2Size: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUFamily: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUModelID: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUModelName: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUStepping: ResourceAttributeConfig{
			Enabled: false,
		},
		HostCPUVendorID: ResourceAttributeConfig{
			Enabled: false,
		},
		HostID: ResourceAttributeConfig{
			Enabled: false,
		},
		HostIP: ResourceAttributeConfig{
			Enabled: false,
		},
		HostName: ResourceAttributeConfig{
			Enabled: true,
		},
		OsDescription: ResourceAttributeConfig{
			Enabled: false,
		},
		OsType: ResourceAttributeConfig{
			Enabled: true,
		},
	}
}
