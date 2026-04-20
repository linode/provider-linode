package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/linode/provider-linode/apis/v1beta1"
)

func TestPrepareTerraformProviderConfiguration_URLFromCreds(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name  string
		creds map[string]string
		want  map[string]any
	}{
		{
			name:  "token and custom url",
			creds: map[string]string{"token": "mytoken", "url": "https://api.custom.linode.com"},
			want:  map[string]any{"token": "mytoken", "url": "https://api.custom.linode.com"},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			config := prepareTerraformProviderConfiguration(tc.creds, v1beta1.ProviderConfiguration{})
			for k, v := range tc.want {
				assert.Equal(t, v, config[k], "key %q", k)
			}
		})
	}
}

func TestPrepareTerraformProviderConfiguration_NoURL(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name  string
		creds map[string]string
	}{
		{
			name:  "no url",
			creds: map[string]string{"token": "mytoken"},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			config := prepareTerraformProviderConfiguration(tc.creds, v1beta1.ProviderConfiguration{})
			_, ok := config["url"]
			assert.False(t, ok, "expected url to be absent when not in creds, but it was present")
		})
	}
}

func TestPrepareTerraformProviderConfiguration_TwoConfigs_DifferentURLs(t *testing.T) {
	t.Parallel()
	tcs := []struct {
		name   string
		credsA map[string]string
		credsB map[string]string
	}{
		{
			name:   "different urls",
			credsA: map[string]string{"token": "tokenA", "url": "https://api-a.linode.com"},
			credsB: map[string]string{"token": "tokenB", "url": "https://api-b.linode.com"},
		},
	}
	for _, tc := range tcs {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			configA := prepareTerraformProviderConfiguration(tc.credsA, v1beta1.ProviderConfiguration{})
			configB := prepareTerraformProviderConfiguration(tc.credsB, v1beta1.ProviderConfiguration{})
			assert.NotEqual(t, configA["url"], configB["url"], "expected different URLs")
			assert.Equal(t, "https://api-a.linode.com", configA["url"], "configA url")
			assert.Equal(t, "https://api-b.linode.com", configB["url"], "configB url")
		})
	}
}

func TestPrepareTerraformProviderConfiguration_ProviderConfigFields(t *testing.T) {
	t.Parallel()
	t.Run("provider config fields", func(t *testing.T) {
		t.Parallel()
		creds := map[string]string{"token": "tok"}
		pc := v1beta1.ProviderConfiguration{
			UserAgentPrefix:        "my-prefix",
			SkipInstanceReadyPoll:  true,
			SkipInstanceDeletePoll: true,
			SkipImplicitReboots:    true,
			DisableInternalCache:   true,
			MinRetryDelayms:        100,
			MaxRetryDelayms:        200,
			LKEEventPollms:         300,
			LKENodeReadyPollms:     400,
			ObjAccessKey:           "accesskey",
			ObjSecretKey:           "secretkey",
			ObjUseTempKeys:         true,
			ObjForceDelete:         true,
		}
		config := prepareTerraformProviderConfiguration(creds, pc)
		checks := map[string]any{
			"ua_prefix":                 "my-prefix",
			"skip_instance_ready_poll":  true,
			"skip_instance_delete_poll": true,
			"skip_implicit_reboots":     true,
			"disable_internal_cache":    true,
			"min_retry_delay_ms":        100,
			"max_retry_delay_ms":        200,
			"lke_event_poll_ms":         300,
			"lke_node_ready_poll_ms":    400,
			"obj_access_key":            "accesskey",
			"obj_secret_key":            "secretkey",
			"obj_use_temp_keys":         true,
			"obj_bucket_force_delete":   true,
		}
		for k, want := range checks {
			assert.Equal(t, want, config[k], "key %q", k)
		}
	})
}

func TestPrepareTerraformProviderConfiguration_ZeroValuesOmitted(t *testing.T) {
	t.Parallel()
	t.Run("zero values omitted", func(t *testing.T) {
		t.Parallel()
		creds := map[string]string{"token": "tok"}
		config := prepareTerraformProviderConfiguration(creds, v1beta1.ProviderConfiguration{})
		omitted := []string{"ua_prefix", "min_retry_delay_ms", "max_retry_delay_ms", "lke_event_poll_ms", "lke_node_ready_poll_ms", "obj_access_key", "obj_secret_key"}
		for _, k := range omitted {
			_, ok := config[k]
			assert.False(t, ok, "expected key %q to be absent for zero value, but it was present", k)
		}
	})
}
