// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config for read the data form 'yaml' or 'json' file and copy the data to a struct.
package config

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"path/filepath"
)

func NewConfig(readDir string) *Config {

	return &Config {
		readDir: readDir,
		cfgT: make(map[string]interface{}, 10),
	}
}

type Config struct {

	cfgT map[string]interface{}
	readDir string
}

func (this *Config) SetDir(readDir string) {

	this.readDir = readDir
}

func (this *Config) ReloadAll() {

	for name, _ := range this.cfgT {
		this.ReloadFile(name)
	}
}

func (this *Config) ReloadFile(name string) {

	inst := this.cfgT[name]
	delete(this.cfgT, name)
	this.ReadStruct(name, inst)
}

// try to read the file data to struct, if file not exist, this goes no error,
// every struct should have default values
func (this *Config) ReadStruct(name string, struc interface{}) error {

	if _struc, ok := this.cfgT[name]; !ok {
		file := filepath.Join(this.readDir, name)
		data, err := ioutil.ReadFile(file)
		if err != nil {
			return err
		}

		switch filepath.Ext(name) {
		case ".yaml", ".yml", ".json":
			err = yaml.Unmarshal(data, struc)
			if err != nil {
				return err
			}
		}

		this.cfgT[name] = struc
	} else {
		struc = _struc
	}
	return nil
}

