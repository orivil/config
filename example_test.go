package config_test
import (
	"fmt"
	"gopkg.in/orivil/config.v0"
)

var configI18n = &struct {
	Languages map[string]string
}{
	// set default value
	Languages: map[string]string{
		"简体中文": "zh-CN",
		"english": "en",
	},
}

func ExampleConfig() {
	// 1. new config instance
	cfg := config.NewConfig("./testdata")

	// 2. read file data to struct
	cfg.ReadStruct("i18n.yml", configI18n)

	// 3. get config data
	fmt.Println(configI18n.Languages)

	// Output:
	// map[简体中文:zh-CN english:en]
}
