package config

import (
	"github.com/juanwolf/gomodoro/pkg/stores"
)

type StoreConfig interface {
	IsActivated() bool
	Instantiate() *stores.Store
}

type StoresConfig struct {
	SQLite *SQLiteConfig `mapstructure:"sqlite"`
}

func DefaultStoresConfig() StoresConfig {
	return StoresConfig{
		SQLite: &SQLiteConfig{
			Path: "./gomodoro.db",
		},
	}
}

func (s StoresConfig) GetStoresConfig() []StoreConfig {
	return []StoreConfig{s.SQLite}
}

type SQLiteConfig struct {
	Path string `mapstructure:"path"`
}

func (s *SQLiteConfig) IsActivated() bool {
	return s != nil
}

func (s SQLiteConfig) Instantiate() *stores.Store {
	sqlite, _ := stores.NewSQLite(s.Path)
	store := stores.Store(sqlite)
	return &store
}
