package ucfgwrap_test

import (
	"testing"

	"github.com/elastic/go-ucfg"
	"github.com/sapcc/ucfgwrap"
)

func TestUnwrap(t *testing.T) {
	t.Setenv("WRAP_TEST", "stone")
	yaml := "key: ${WRAP_TEST}"
	config, err := ucfgwrap.FromYAML([]byte(yaml), ucfg.VarExp, ucfg.ResolveEnv)
	if err != nil {
		t.Fatalf("FromYAML: %s", err)
	}
	data := struct {
		Key string
	}{}
	err = config.Unpack(&data)
	if err != nil {
		t.Fatalf("Unpack: %s", err)
	}
	if data.Key != "stone" {
		t.Fatalf("Unexpected unpack value: %s", data.Key)
	}
}

func TestWrap(t *testing.T) {
	t.Setenv("WRAP_TEST", "jewel")
	base := "key:\n  sub: ${WRAP_TEST}"
	config, err := ucfgwrap.FromYAML([]byte(base), ucfg.VarExp, ucfg.ResolveEnv)
	if err != nil {
		t.Fatalf("FromYAML: %s", err)
	}
	data := struct {
		Key *ucfg.Config
	}{}
	err = config.Unpack(&data)
	if err != nil {
		t.Fatalf("Unpack: %s", err)
	}
	wrapped := config.Wrap(data.Key)
	sub := struct {
		Sub string
	}{}
	err = wrapped.Unpack(&sub)
	if err != nil {
		t.Fatalf("Unpack: %s", err)
	}
	if sub.Sub != "jewel" {
		t.Fatalf("Unexpected unpack value: %s", sub.Sub)
	}
}
