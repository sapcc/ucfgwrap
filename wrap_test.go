package ucfgwrap_test

import (
	"testing"

	"github.com/elastic/go-ucfg"
	"github.com/sapcc/ucfgwrap"
)

func TestWrap(t *testing.T) {
	t.Setenv("WRAP_TEST", "stone")
	yaml := "key: ${WRAP_TEST}"
	config, err := ucfgwrap.FromYAML([]byte(yaml), ucfg.VarExp, ucfg.ResolveEnv)
	if err != nil {
		t.Fatalf("FromYAML: %s", err)
	}
	data := struct {
		Key string
	}{Key: ""}
	err = config.Unpack(&data)
	if err != nil {
		t.Fatalf("Unpack: %s", err)
	}
	if data.Key != "stone" {
		t.Fatalf("Unexpected unpack value: %s", data.Key)
	}
}
