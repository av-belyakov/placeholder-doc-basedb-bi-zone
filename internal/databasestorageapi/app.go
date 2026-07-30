package databasestorageapi

import (
	"errors"

	"github.com/av-belyakov/placeholder_doc-basedb_bi.zone/interfaces"
)

// New настраивает новый модуль взаимодействия с API Database
func New(logger interfaces.Logger, counter interfaces.Counter, opts ...DatabaseStorageOptions) (*DatabaseStorage, error) {
	dbs := &DatabaseStorage{
		counter:  counter,
		logger:   logger,
		chInput:  make(chan SettingsChanInput),
		chOutput: make(chan SettingsChanOutput),
		settings: settingsDatabaseStorage{
			maxGetDocumentsSize: 10,
		},
	}

	for _, opt := range opts {
		if err := opt(dbs); err != nil {
			return dbs, err
		}
	}

	return dbs, nil
}

// WithHost имя или ip адрес хоста API
func WithHost(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'host' cannot be empty")
		}

		dbs.settings.host = v

		return nil
	}
}

// WithPort порт API
func WithPort(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v <= 0 || v > 65535 {
			return errors.New("an incorrect network port value was received")
		}

		dbs.settings.port = v

		return nil
	}
}

// WithUser имя пользователя для доступа к БД
func WithUser(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'user' cannot be empty")
		}

		dbs.settings.user = v

		return nil
	}
}

// WithPasswd пароль для доступа к БД
func WithPasswd(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		if v == "" {
			return errors.New("the value of 'passwd' cannot be empty")
		}

		dbs.settings.passwd = v

		return nil
	}
}

// WithNameDB наименование БД
func WithNameDB(v string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.namedb = v

		return nil
	}
}

// WithStorage наименование коллекции или индекса БД
func WithStorage(v map[string]string) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.storages = v

		return nil
	}
}

// WithMaxGetDocumentSize максимальное количество запрашиваемых документов (по умолчанию 10)
func WithMaxGetDocumentSize(v int) DatabaseStorageOptions {
	return func(dbs *DatabaseStorage) error {
		dbs.settings.maxGetDocumentsSize = v

		return nil
	}
}
