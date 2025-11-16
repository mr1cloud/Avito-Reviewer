package logger

import (
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

type RotateFileHook struct {
	logWriter *lumberjack.Logger
	formatter logrus.Formatter
	level     logrus.Level
}

func (h *RotateFileHook) Levels() []logrus.Level {
	return logrus.AllLevels[:h.level+1]
}

func (h *RotateFileHook) Fire(entry *logrus.Entry) error {
	b, err := h.formatter.Format(entry)
	if err != nil {
		return err
	}
	_, err = h.logWriter.Write(b)
	return err
}

type RotateFileConfig struct {
	Filename   string `yaml:"file" env:"FILE" env-default:"service.log"`
	MaxSize    int    `yaml:"max_size" env:"MAX_SIZE" env-default:"100" env-description:"Maximum size in megabytes of the log file before it gets rotated"`
	MaxAge     int    `yaml:"max_age" env:"MAX_AGE" env-default:"30" env-description:"Maximum number of days to retain old log files based on the timestamp encoded in their filename"`
	MaxBackups int    `yaml:"max_backups" env:"MAX_BACKUPS" env-default:"0" env-description:"Maximum number of old log files to retain"`
	LocalTime  bool   `yaml:"local_time" env:"LOCAL_TIME" env-default:"false" env-description:"Use the host's local time to format timestamps in backup files"`
	Compress   bool   `yaml:"compress" env:"COMPRESS" env-default:"false" env-description:"Compress old log files to save space"`
}

func NewRotateFileHook(cfg RotateFileConfig, formatter logrus.Formatter, level logrus.Level) logrus.Hook {
	hook := &RotateFileHook{
		logWriter: &lumberjack.Logger{
			Filename:   cfg.Filename,
			MaxSize:    cfg.MaxSize,
			MaxAge:     cfg.MaxAge,
			MaxBackups: cfg.MaxBackups,
			LocalTime:  cfg.LocalTime,
			Compress:   cfg.Compress,
		},
		formatter: formatter,
		level:     level,
	}

	return hook
}

type Logger struct {
	base    *logrus.Logger
	service string
	entry   *logrus.Entry
}

func NewLogger(service string, level string, rotateFileConfig *RotateFileConfig) *Logger {
	var logger Logger

	logger.base = logrus.New()
	logger.base.SetFormatter(&logrus.TextFormatter{
		ForceColors:     true,
		DisableColors:   false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339,
	})

	parsedLevel, err := logrus.ParseLevel(level)
	if err != nil {
		parsedLevel = logrus.InfoLevel
	}
	logger.base.SetLevel(parsedLevel)

	if rotateFileConfig != nil {
		logger.base.AddHook(NewRotateFileHook(
			*rotateFileConfig,
			&logrus.TextFormatter{
				DisableColors:   true,
				FullTimestamp:   true,
				TimestampFormat: time.RFC3339,
			},
			logger.base.Level,
		))
	}

	if service != "" {
		logger.service = service

		logger.entry = logger.base.WithFields(logrus.Fields{
			"service": service,
		})
	} else {
		logger.entry = logrus.NewEntry(logger.base)
	}

	return &logger
}

func (l *Logger) Info(args ...interface{}) {
	l.entry.Info(args...)
}

func (l *Logger) Infof(template string, args ...interface{}) {
	l.entry.Infof(template, args...)
}

func (l *Logger) Warn(args ...interface{}) {
	l.entry.Warn(args...)
}

func (l *Logger) Warnf(template string, args ...interface{}) {
	l.entry.Warnf(template, args...)
}

func (l *Logger) Error(args ...interface{}) {
	l.entry.Error(args...)
}

func (l *Logger) Errorf(template string, args ...interface{}) {
	l.entry.Errorf(template, args...)
}

func (l *Logger) Debug(args ...interface{}) {
	l.entry.Debug(args...)
}

func (l *Logger) Debugf(template string, args ...interface{}) {
	l.entry.Debugf(template, args...)
}

func (l *Logger) Fatal(args ...interface{}) {
	l.entry.Fatal(args...)
}

func (l *Logger) Fatalf(template string, args ...interface{}) {
	l.entry.Fatalf(template, args...)
}

func (l *Logger) WithFields(args ...interface{}) *Logger {
	fields := logrus.Fields{}

	for i := 0; i < len(args)-1; i += 2 {
		key, ok := args[i].(string)
		if !ok {
			continue
		}
		fields[key] = args[i+1]
	}

	newEntry := l.entry.WithFields(fields)

	return &Logger{
		base:    l.base,
		service: l.service,
		entry:   newEntry,
	}
}
