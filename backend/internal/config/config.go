package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
	"strings"

	"gopkg.in/yaml.v2"

	"github.com/hashfunc/project-milky-way/ext/kakao"
	"github.com/hashfunc/project-milky-way/internal"
	"github.com/hashfunc/project-milky-way/internal/database"
)

const (
	DefaultConfigPath = "server.config.yml"
)

type Config struct {
	Name     string          `yaml:"name" config:"required"`
	Bind     string          `yaml:"bind"`
	CORS     CORSConfig      `yaml:"cors"`
	Database database.Config `yaml:"database"`
	Secret   *Secret         `yaml:"secret"`
}

type CORSConfig struct {
	AllowOrigins     string `yaml:"allow-origins"`
	AllowMethods     string `yaml:"allow-methods"`
	AllowHeaders     string `yaml:"allow-headers"`
	AllowCredentials bool   `yaml:"allow-credentials"`
	ExposeHeaders    string `yaml:"expose-headers"`
	MaxAge           int    `yaml:"max-age"`
}

type Secret struct {
	Kakao *kakao.Config `yaml:"kakao"`
}

func LoadConfigFile() (*Config, error) {
	data, err := ioutil.ReadFile(DefaultConfigPath)

	if err != nil {
		return nil, internal.NewError("설정 파일 열기 오류", err)
	}

	config := new(Config)

	if err := yaml.UnmarshalStrict(data, config); err != nil {
		return nil, internal.NewError("설정 파일 오류", err)
	}

	if err := validate(config, nil); err != nil {
		return nil, internal.NewError("설정값 오류", err)
	}

	return config, nil
}

func validate(config interface{}, parents []string) error {
	element := reflect.ValueOf(config).Elem()

	for index := 0; index < element.NumField(); index += 1 {
		field := element.Type().Field(index)
		fieldValue := element.Field(index)
		value := fieldValue.Interface()

		if field.Tag.Get("config") == "required" {
			if (value == nil) || (fieldValue.Kind() == reflect.String && value == "") {
				fieldPath := strings.Join(append(parents, field.Name), ".")
				text := fmt.Sprintf("필수 항목 누락: %s", fieldPath)
				return errors.New(text)
			}
		}

		if fieldValue.Kind() == reflect.Struct {
			ptr := fieldValue.Addr().Interface()
			if err := validate(ptr, append(parents, field.Name)); err != nil {
				return err
			}
		}
	}

	return nil
}
