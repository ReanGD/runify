package config

import (
	"reflect"

	"github.com/mitchellh/mapstructure"
	"go.uber.org/zap/zapcore"
)

func stringToZapLevel() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() == reflect.String && t == reflect.TypeOf(zapcore.DebugLevel) {
			var result zapcore.Level
			err := result.UnmarshalText([]byte(data.(string)))
			return result, err
		}

		return data, nil
	}
}

func stringToLoggerFormatLevel() mapstructure.DecodeHookFunc {
	return func(f reflect.Type, t reflect.Type, data interface{}) (interface{}, error) {
		if f.Kind() == reflect.String && t == reflect.TypeOf(PlainFormat) {
			var result LoggerFormat
			err := result.UnmarshalText([]byte(data.(string)))
			return result, err
		}

		return data, nil
	}
}

func zapLevelDecoder(c *mapstructure.DecoderConfig) {
	c.DecodeHook = mapstructure.ComposeDecodeHookFunc(
		mapstructure.StringToTimeDurationHookFunc(),
		mapstructure.StringToSliceHookFunc(","),
		stringToZapLevel(),
		stringToLoggerFormatLevel())
}
