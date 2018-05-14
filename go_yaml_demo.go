package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Global struct {
		LogPath string `yaml:"logPath"`
		LogName string `yaml:"logName"`
	} `yaml:"global"`

	KafkaConfig struct {
		Version    string   `yaml:"version"`
		BrokerList []string `yaml:"brokerList""`
	} `yaml:"kafkaConfig"`

	ClusterName string `yaml:"clusterName"`

	LogOps []struct {
		LogType    string `yaml:"logType"`
		Expression string `yaml:"expression"`
		Topic      string `yaml:"topic"`
	} `yaml:"logOps"`
}

func main() {
	var cfg Config
	var mCfg map[string]interface{}
	cfgFile, err := ioutil.ReadFile("./config.yaml")
	if err != nil {
		fmt.Printf("read config file failed: %s\n", err.Error())
		os.Exit(-1)
	}

	err = yaml.Unmarshal(cfgFile, &cfg)

	if err != nil {
		fmt.Printf("parse config failed: %s\n", err.Error())
		os.Exit(-1)
	}

	fmt.Printf("%+v\n", cfg)

	cfgJStr, err := json.Marshal(cfg)
	if err != nil {
		fmt.Printf("convert config to json string failed: %s\n", err.Error())
		os.Exit(-1)
	}

	fmt.Printf("%s\n", cfgJStr)

	err = yaml.Unmarshal(cfgFile, &mCfg)

	if err != nil {
		fmt.Printf("parse config faile: %s\n", err.Error())
		os.Exit(-1)
	}

	for k, v := range mCfg {
		fmt.Printf("%s: %+v\n", k, v)
	}
}
