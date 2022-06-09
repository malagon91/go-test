package logger

import (
	"api-bootstrap-echo/constants"
	"fmt"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/diode"
	"github.com/rs/zerolog/log"
)

var (
	hostname string
)

const (
	filePathPlaceholder  string = "logs/%s_%s"
	logLevelEnvName      string = "LOG_LEVEL"          // Default: DEBUG
	enableLogFileEnvName string = "ENABLE_FILE_LOG"    // Default: false
	enableConsoleEnvName string = "ENABLE_CONSOLE_LOG" // Default: true

	logLevelTraceStr string = "TRACE"
	logLevelDebugStr string = "DEBUG"
	logLevelInfoStr  string = "INFO"
	logLevelWarnStr  string = "WARN"
	logLevelErrorStr string = "ERROR"

	logLevelTrace int = -1
	logLevelDebug int = 0
	logLevelInfo  int = 1
	logLevelWarn  int = 2
	logLevelError int = 3
)

func init() {
	filePath := fmt.Sprintf(filePathPlaceholder, constants.ServiceName, getHostname())
	logLevel, logLevelStr := getLogLevel()
	enableLogFile, _ := strconv.ParseBool(getEnvVar(enableLogFileEnvName, "false"))
	enableConsoleLogs, _ := strconv.ParseBool(getEnvVar(enableConsoleEnvName, "true"))
	var outputs []io.Writer

	zerolog.SetGlobalLevel(zerolog.Level(logLevel))

	if enableConsoleLogs {
		wr := diode.NewWriter(os.Stdout, 1000, 10*time.Millisecond, func(missed int) {
			fmt.Printf("Logger dropped %d messages", missed)
		})
		outputs = append(outputs, wr)
	}

	if enableLogFile {
		logf, _ := rotatelogs.New(filePath+".%Y%m%d%H.log",
			rotatelogs.WithClock(rotatelogs.Local),
			rotatelogs.WithRotationTime(time.Hour*1),
		)
		fileOutput := zerolog.ConsoleWriter{Out: logf, NoColor: true}
		fileOutput.FormatLevel = func(text interface{}) string {
			return ""
		}
		fileOutput.FormatMessage = func(text interface{}) string {
			return text.(string)
		}
		outputs = append(outputs, fileOutput)
	}

	multi := zerolog.MultiLevelWriter(outputs...)
	log.Logger = zerolog.New(multi).With().Timestamp().Logger()

	Debug("logger", "init", fmt.Sprintf("Log level set to %s", logLevelStr))
}

func getLogLevel() (int, string) {
	logLevelStr := getEnvVar(logLevelEnvName, logLevelDebugStr)
	switch logLevelStr {
	case logLevelTraceStr:
		return logLevelTrace, logLevelTraceStr
	case logLevelInfoStr:
		return logLevelInfo, logLevelInfoStr
	case logLevelWarnStr:
		return logLevelWarn, logLevelWarnStr
	case logLevelErrorStr:
		return logLevelError, logLevelErrorStr
	default:
		return logLevelDebug, logLevelDebugStr
	}
}

func getEnvVar(env string, defaultValue string) string {
	value, found := os.LookupEnv(env)
	if !found || value == "" {
		return defaultValue
	}
	return value
}

func getHostname() string {
	if hostname == "" {
		h, err0 := os.Hostname()
		if err0 != nil {
			hostname = "UNKNOWN"
		} else {
			hostname = h
		}
	}
	return hostname
}

// Trace :
func Trace(module string, function string, text string) {
	log.Trace().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function),
		).
		Msgf(text)
}

// Performance :
func Performance(module string, function string, start time.Time) {
	log.Debug().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function),
		).
		Dict("analysis", zerolog.Dict().
			Str("duration_ms", fmt.Sprintf("%.2f", float32(time.Since(start).Nanoseconds())/1000000.0)),
		).
		Msgf("")
}

// Debug :
func Debug(module string, function string, text string) {
	log.Debug().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function),
		).
		Msgf(text)
}

// Info :
func Info(module string, function string, target string, text string) {
	log.Info().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function).
			Str("target", target),
		).
		Msgf(text)
}

// Warn :
func Warn(module string, function string, target string, text string) {
	log.Warn().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function).
			Str("target", target),
		).
		Msgf(text)
}

// Error :
func Error(module string, function string, target string, text string) {
	log.Error().
		Str("id", uuid.New().String()).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function).
			Str("target", target),
		).
		Msgf(text)
}

// Fatal : It's a critical error with an exit statement.
func Fatal(module string, function string, target string, text string) {
	id := uuid.New().String()
	message := fmt.Sprintf(`{"level":"fatal","id":"%s","context":{"module":"%s","func":"%s","target":"%s"},"message":"%s"}`,
		id, module, function, target, text)
	fmt.Println(message)
	log.Fatal().
		Str("id", id).
		Dict("context", zerolog.Dict().
			Str("module", module).
			Str("func", function).
			Str("target", target),
		).
		Msgf(text)
}

// GetLoggerConfig :
func GetLoggerConfig() middleware.LoggerConfig {
	return middleware.LoggerConfig{
		Skipper: func(c echo.Context) bool {
			return c.Path() == constants.HealthResource ||
				c.Path() == constants.MetricsResource
		},
		Format: `{"time":"${time_rfc3339}","id":"${id}",method":"${method}","uri":"${uri}",` +
			`"status":${status},"duration":"${latency_human}"}` + "\n",
		CustomTimeFormat: "2006-01-02 15:04:05.00000",
	}
}
