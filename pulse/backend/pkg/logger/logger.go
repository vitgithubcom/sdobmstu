package logger

import (
    "fmt"
    "io"
    "log"
    "os"
    "runtime"
    "strings"
    "sync"
    "time"
)

// LogLevel представляет уровень логирования
type LogLevel int

const (
    DEBUG LogLevel = iota
    INFO
    WARN
    ERROR
    FATAL
)

var levelNames = map[LogLevel]string{
    DEBUG: "DEBUG",
    INFO:  "INFO",
    WARN:  "WARN",
    ERROR: "ERROR",
    FATAL: "FATAL",
}

var levelColors = map[LogLevel]string{
    DEBUG: "\033[36m",  // Cyan
    INFO:  "\033[32m",  // Green
    WARN:  "\033[33m",  // Yellow
    ERROR: "\033[31m",  // Red
    FATAL: "\033[35m",  // Magenta
}

const resetColor = "\033[0m"

// Logger представляет логгер
type Logger struct {
    mu       sync.Mutex
    level    LogLevel
    output   io.Writer
    prefix   string
    useColor bool
}

var (
    defaultLogger *Logger
    once          sync.Once
)

// New создаёт новый логгер
func New(level LogLevel, output io.Writer, prefix string, useColor bool) *Logger {
    return &Logger{
        level:    level,
        output:   output,
        prefix:   prefix,
        useColor: useColor,
    }
}

// Default возвращает глобальный логгер
func Default() *Logger {
    once.Do(func() {
        defaultLogger = New(INFO, os.Stdout, "", true)
    })
    return defaultLogger
}

// SetLevel устанавливает уровень логирования
func (l *Logger) SetLevel(level LogLevel) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.level = level
}

// SetOutput устанавливает вывод логгера
func (l *Logger) SetOutput(output io.Writer) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.output = output
}

// SetPrefix устанавливает префикс
func (l *Logger) SetPrefix(prefix string) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.prefix = prefix
}

// SetUseColor включает/отключает цвет
func (l *Logger) SetUseColor(useColor bool) {
    l.mu.Lock()
    defer l.mu.Unlock()
    l.useColor = useColor
}

// log внутренняя функция логирования
func (l *Logger) log(level LogLevel, format string, args ...interface{}) {
    if level < l.level {
        return
    }

    l.mu.Lock()
    defer l.mu.Unlock()

    // Получаем информацию о вызывающем
    _, file, line, ok := runtime.Caller(2)
    if !ok {
        file = "???"
        line = 0
    }

    // Укорачиваем путь до файла
    parts := strings.Split(file, "/")
    if len(parts) > 3 {
        file = strings.Join(parts[len(parts)-3:], "/")
    }

    // Формируем сообщение
    timestamp := time.Now().Format("2006-01-02 15:04:05.000")
    levelName := levelNames[level]
    message := fmt.Sprintf(format, args...)

    // Цвет
    color := ""
    reset := ""
    if l.useColor {
        color = levelColors[level]
        reset = resetColor
    }

    // Сборка строки
    prefix := ""
    if l.prefix != "" {
        prefix = "[" + l.prefix + "] "
    }

    logLine := fmt.Sprintf("%s %s%s%s %s %s:%d %s\n",
        timestamp,
        color,
        levelName,
        reset,
        prefix,
        file,
        line,
        message,
    )

    l.output.Write([]byte(logLine))

    // Если FATAL, завершаем программу
    if level == FATAL {
        os.Exit(1)
    }
}

// ========== Удобные методы ==========

// Debug логирует сообщение уровня DEBUG
func (l *Logger) Debug(format string, args ...interface{}) {
    l.log(DEBUG, format, args...)
}

// Info логирует сообщение уровня INFO
func (l *Logger) Info(format string, args ...interface{}) {
    l.log(INFO, format, args...)
}

// Warn логирует сообщение уровня WARN
func (l *Logger) Warn(format string, args ...interface{}) {
    l.log(WARN, format, args...)
}

// Error логирует сообщение уровня ERROR
func (l *Logger) Error(format string, args ...interface{}) {
    l.log(ERROR, format, args...)
}

// Fatal логирует сообщение уровня FATAL и завершает программу
func (l *Logger) Fatal(format string, args ...interface{}) {
    l.log(FATAL, format, args...)
}

// ========== Глобальные функции (используют defaultLogger) ==========

// Debug глобальный Debug
func Debug(format string, args ...interface{}) {
    Default().Debug(format, args...)
}

// Info глобальный Info
func Info(format string, args ...interface{}) {
    Default().Info(format, args...)
}

// Warn глобальный Warn
func Warn(format string, args ...interface{}) {
    Default().Warn(format, args...)
}

// Error глобальный Error
func Error(format string, args ...interface{}) {
    Default().Error(format, args...)
}

// Fatal глобальный Fatal
func Fatal(format string, args ...interface{}) {
    Default().Fatal(format, args...)
}

// ========== Дополнительные утилиты ==========

// WithFields возвращает логгер с дополнительными полями (структурированный лог)
type LoggerWithFields struct {
    logger *Logger
    fields map[string]interface{}
}

// WithFields создаёт логгер с полями
func (l *Logger) WithFields(fields map[string]interface{}) *LoggerWithFields {
    return &LoggerWithFields{
        logger: l,
        fields: fields,
    }
}

// Debug логирует с полями
func (lwf *LoggerWithFields) Debug(format string, args ...interface{}) {
    lwf.log(DEBUG, format, args...)
}

// Info логирует с полями
func (lwf *LoggerWithFields) Info(format string, args ...interface{}) {
    lwf.log(INFO, format, args...)
}

// Warn логирует с полями
func (lwf *LoggerWithFields) Warn(format string, args ...interface{}) {
    lwf.log(WARN, format, args...)
}

// Error логирует с полями
func (lwf *LoggerWithFields) Error(format string, args ...interface{}) {
    lwf.log(ERROR, format, args...)
}

func (lwf *LoggerWithFields) log(level LogLevel, format string, args ...interface{}) {
    // Формируем строку с полями
    fieldStr := ""
    if len(lwf.fields) > 0 {
        fields := make([]string, 0, len(lwf.fields))
        for k, v := range lwf.fields {
            fields = append(fields, fmt.Sprintf("%s=%v", k, v))
        }
        fieldStr = " {" + strings.Join(fields, ", ") + "}"
    }

    lwf.logger.log(level, format+fieldStr, args...)
}