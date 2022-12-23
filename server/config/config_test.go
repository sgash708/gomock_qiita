package config_test

import (
	"os"
	"testing"

	"github.com/sgash708/gomock_qiita/server/config"
)

func clearEnv() {
	os.Clearenv()
}

func TestLoadEnvConfig(t *testing.T) {
	cases := []struct {
		port       string
		driverName string
		dataSource string
	}{
		{port: "0000", driverName: "mysql", dataSource: "mysql datasource name"},
		{port: "1234", driverName: "test", dataSource: "mysql datasource name"},
	}

	for _, cs := range cases {
		clearEnv()
		if err := os.Setenv("PORT", cs.port); err != nil {
			t.Error(err)
		}
		if err := os.Setenv("DRIVER", cs.driverName); err != nil {
			t.Error(err)
		}
		if err := os.Setenv("DATASOURCE", cs.dataSource); err != nil {
			t.Error(err)
		}
		cfg, err := config.LoadEnvConfig()
		if err != nil {
			t.Error(err)
		}

		if cfg.Port != cs.port {
			t.Error(cfg)
		}
		if cfg.DriverName != cs.driverName {
			t.Error(cfg)
		}
		if cfg.DataSource != cs.dataSource {
			t.Error(cfg)
		}
	}
}
