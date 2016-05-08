// Copyright 2016 orivil Authors. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

// Package config for read the data form 'yaml' or 'json' file and copy the data to a struct.
package config

import (
	"gopkg.in/yaml.v2"
	"gopkg.in/orivil/helper.v0"
	"path/filepath"
	"os"
	"io/ioutil"
	"log"
)

func NewConfig(readDir string) *Config {

	if !helper.IsExist(readDir) {
		err := os.MkdirAll(readDir, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	return &Config {
		readDir: readDir,
		cfgT: make(map[string]interface{}, 10),
	}
}

type Config struct {

	cfgT map[string]interface{}
	readDir string
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

// try to read the data form config file to struct, if the file not exist,
// it will be auto generated
func (this *Config) ReadStruct(name string, struc interface{}) error {

	if _struc, ok := this.cfgT[name]; !ok {
		filename := filepath.Join(this.readDir, name)
		data, err := ioutil.ReadFile(filename)
		if os.IsNotExist(err) {
			data, err = yaml.Marshal(struc)
			if err != nil {
				return err
			}
			err = ioutil.WriteFile(filename, data, 0666)
			if err != nil {
				return err
			}
			log.Printf("generating config file:%s\n", filename)
		} else if err != nil {
			return err
		} else {
			switch filepath.Ext(name) {
			case ".yaml", ".yml", ".json":
				err = yaml.Unmarshal(data, struc)
				if err != nil {
					return err
				}
			}
			this.cfgT[name] = struc
		}
	} else {
		struc = _struc
	}
	return nil
}

