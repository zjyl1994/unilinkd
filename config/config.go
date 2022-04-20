package config

import (
	"io/ioutil"
	"time"

	"gopkg.in/yaml.v2"
)

type configY struct {
	Listen   string
	Timezone string
	S3       *S3Config
	Links    []linkY
}

type linkY struct {
	Disable bool
	Mode    string
	Code    string
	Url     string
	Expire  string
}

type LinkS struct {
	Mode   string
	Url    string
	Expire time.Time
}

var (
	ListenAddr = "0.0.0.0:12315"
	Links      = make(map[string]LinkS)
	S3         *S3Config
)

func LoadConf(filename string) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return err
	}
	var conf configY
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return err
	}
	if len(conf.Listen) > 0 {
		ListenAddr = conf.Listen
	}
	S3 = conf.S3
	var tz = time.Local
	if len(conf.Timezone) > 0 {
		if t, err := time.LoadLocation(conf.Timezone); err != nil {
			return err
		} else {
			tz = t
		}
	}
	for _, v := range conf.Links {
		if v.Disable {
			continue
		}
		link := LinkS{
			Mode: v.Mode,
			Url:  v.Url,
		}
		if len(v.Expire) > 0 {
			if t, err := time.ParseInLocation("2006-01-02 15:04:05", v.Expire, tz); err != nil {
				return err
			} else {
				link.Expire = t
			}
		}
		Links[v.Code] = link
	}
	return nil
}
