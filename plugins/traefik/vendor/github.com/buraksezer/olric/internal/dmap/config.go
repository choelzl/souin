package dmap

import (
	"fmt"
	"time"

	"github.com/buraksezer/olric/config"
)

// dmapConfig keeps DMap config control parameters and access-log for keys in a dmap.
type dmapConfig struct {
	engine          *config.Engine
	maxIdleDuration time.Duration
	ttlDuration     time.Duration
	maxKeys         int
	maxInuse        int
	lruSamples      int
	evictionPolicy  config.EvictionPolicy
}

func (c *dmapConfig) load(dc *config.DMaps, name string) error {
	// Try to set config configuration for this dmap.
	c.maxIdleDuration = dc.MaxIdleDuration
	c.ttlDuration = dc.TTLDuration
	c.maxKeys = dc.MaxKeys
	c.maxInuse = dc.MaxInuse
	c.lruSamples = dc.LRUSamples
	c.evictionPolicy = dc.EvictionPolicy
	c.engine = dc.Engine

	if dc.Custom != nil {
		// config.DMap struct can be used for fine-grained control.
		cs, ok := dc.Custom[name]
		if ok {
			if c.maxIdleDuration != cs.MaxIdleDuration {
				c.maxIdleDuration = cs.MaxIdleDuration
			}
			if c.ttlDuration != cs.TTLDuration {
				c.ttlDuration = cs.TTLDuration
			}
			if c.evictionPolicy != cs.EvictionPolicy {
				c.evictionPolicy = cs.EvictionPolicy
			}
			if c.maxKeys != cs.MaxKeys {
				c.maxKeys = cs.MaxKeys
			}
			if c.maxInuse != cs.MaxInuse {
				c.maxInuse = cs.MaxInuse
			}
			if c.lruSamples != cs.LRUSamples {
				c.lruSamples = cs.LRUSamples
			}
			if c.evictionPolicy != cs.EvictionPolicy {
				c.evictionPolicy = cs.EvictionPolicy
			}
			if c.engine == nil {
				c.engine = cs.Engine
			}
		}
	}

	// TODO: Create a new function to verify config config.
	if c.evictionPolicy == config.LRUEviction {
		if c.maxInuse <= 0 && c.maxKeys <= 0 {
			return fmt.Errorf("maxInuse or maxKeys have to be greater than zero")
		}
		// set the default value.
		if c.lruSamples == 0 {
			c.lruSamples = config.DefaultLRUSamples
		}
	}
	return nil
}
