package common

import (
	"reflect"
	"strings"
	"testing"
	"time"
)

func BenchmarkCopyWithOption(b *testing.B) {

	a := &RoleMartialartInfo{}
	a.status = 123
	b.Run("copy1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			to := &RoleMartialartInfo{}
			CopyWithOption(to, a, Option{DeepCopy: true, CaseSensitive: true})
		}
	})

	b.Run("copy2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			to := &RoleMartialartInfo{}
			DeepCopy(a, to, nil)
		}
	})

}

var defExcludeFields = map[string]bool{
	"uid":      true,
	"rid":      true,
	"rolename": true,
	"lid":      true,
}

type copyConfig struct {
	skip bool
}

func DeepCopy(src, dst interface{}, excludeFields map[string]bool) error {
	srcVal := reflect.ValueOf(src).Elem()
	dstVal := reflect.ValueOf(dst).Elem()
	return copyRecursive(srcVal, dstVal, make(map[uintptr]bool), excludeFields)
}

func copyRecursive(src, dst reflect.Value, visited map[uintptr]bool, excludeFields map[string]bool) error {
	if excludeFields == nil {
		excludeFields = defExcludeFields
	}

	// 处理指针类型
	if src.Kind() == reflect.Ptr {
		if src.IsNil() {
			return nil
		}

		// 检查循环引用
		addr := src.Pointer()
		if visited[addr] {
			return nil
		}
		visited[addr] = true

		src = src.Elem()
	}

	if dst.Kind() == reflect.Ptr {
		if dst.IsNil() {
			dst.Set(reflect.New(dst.Type().Elem()))
		}
		dst = dst.Elem()
	}

	// 根据类型处理
	switch src.Kind() {
	case reflect.Struct:
		for i := 0; i < src.NumField(); i++ {
			fieldType := src.Type().Field(i)
			fieldName := strings.ToLower(fieldType.Name)

			// 跳过未导出字段
			if fieldType.PkgPath != "" {
				continue
			}

			// 解析copy tag
			cfg := parseCopyTag(fieldType)
			if cfg.skip {
				continue
			}

			// 检查排除字段（仅当没有tag配置时）
			if !cfg.skip && excludeFields[fieldName] {
				continue
			}

			// 递归处理嵌套字段
			if err := copyRecursive(src.Field(i), dst.Field(i), visited, excludeFields); err != nil {
				return err
			}
		}

	case reflect.Slice:
		if src.IsNil() {
			return nil
		}

		dst.Set(reflect.MakeSlice(src.Type(), src.Len(), src.Cap()))

		for i := 0; i < src.Len(); i++ {
			if err := copyRecursive(src.Index(i), dst.Index(i), visited, excludeFields); err != nil {
				return err
			}
		}

	case reflect.Map:
		if src.IsNil() {
			return nil
		}
		dst.Set(reflect.MakeMap(src.Type()))
		for _, key := range src.MapKeys() {
			originalValue := src.MapIndex(key)
			newValue := reflect.New(originalValue.Type()).Elem()
			if err := copyRecursive(originalValue, newValue, visited, excludeFields); err != nil {
				return err
			}
			dst.SetMapIndex(key, newValue)
		}

	case reflect.Interface:
		if src.IsNil() {
			return nil
		}
		originalValue := src.Elem()
		newValue := reflect.New(originalValue.Type()).Elem()
		if err := copyRecursive(originalValue, newValue, visited, excludeFields); err != nil {
			return err
		}
		dst.Set(newValue)

	default:
		// 直接设置基础类型值
		if dst.CanSet() {
			dst.Set(src)
		}
	}
	return nil
}

func parseCopyTag(field reflect.StructField) copyConfig {
	cfg := copyConfig{}

	tag := field.Tag.Get("copy")
	if tag == "" {
		return cfg
	}

	// 支持多种指令，用逗号分隔
	options := strings.Split(tag, ",")
	for _, opt := range options {
		switch strings.TrimSpace(opt) {
		case "skip":
			cfg.skip = true
		}
	}
	return cfg
}

type RoleMartialartInfo struct {
	Rid int64

	PointPoints       int64
	PointBattleDamage int64

	GroupPoints       int64
	GroupBattleDamage int64

	SeasonPoints      int64
	PointWinCount     int64
	PointLoseCount    int64
	GroupWinCount     int64
	GroupLoseCount    int64
	KnockoutWinCount  int64
	KnockoutLoseCount int64
	KnockoutIndex     int32

	RoleInfo RoleSimpleInfo `handle:"json"`

	dbStatus
}

type RoleSimpleInfo struct {
	Rid       int64
	Rolename  string
	PlayerExp int64

	Lid            int64
	Leaguename     string
	LeagueFlag     int16
	LeaguePortrait int16

	ShowEmblemList []int32

	IllustrationsCount      int32
	SkillCount              int32
	CityStickerCollectCount int32
	RogueTowerClass         int32
	RogueTowerLevel         int32
	PveLevel                int32
	PvPScore                int32
	MainFP                  int32

	LastLogin time.Time
}

type dbStatus struct {
	status uint8
}
