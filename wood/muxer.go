package wood

// Log logs a message the provided text using the provided log level
func Log(level Level, args ...interface{}) {
	switch level {
	case TraceLevel:
		Trace(args...)
	case DebugLevel:
		Debug(args...)
	case InfoLevel:
		Info(args...)
	case WarnLevel:
		Warn(args...)
	case ErrorLevel:
		Error(args...)
	case FatalLevel:
		Fatal(args...)
	case PanicLevel:
		Panic(args...)
	}
}

// Logln logs a message with newline the provided log level
func Logln(level Level, args ...interface{}) {
	switch level {
	case TraceLevel:
		Traceln(args...)
	case DebugLevel:
		Debugln(args...)
	case InfoLevel:
		Infoln(args...)
	case WarnLevel:
		Warnln(args...)
	case ErrorLevel:
		Errorln(args...)
	case FatalLevel:
		Fatalln(args...)
	case PanicLevel:
		Panicln(args...)
	}
}

// Logf logs a message with formatting using the provided log level
func Logf(level Level, args ...interface{}) {
	decorateF(level, args, func(format string, args []any) {
		std.Infof(format, args...)
	})
}
