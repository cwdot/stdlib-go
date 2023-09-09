package wood

// Logf logs a message with formatting using the provided log level
func Logf(level Level, args ...interface{}) {
	switch level {
	case TraceLevel:
		Tracef(args...)
	case DebugLevel:
		Debugf(args...)
	case InfoLevel:
		Infof(args...)
	case WarnLevel:
		Warnf(args...)
	case ErrorLevel:
		Errorf(args...)
	case FatalLevel:
		Fatalf(args...)
	case PanicLevel:
		Panicf(args...)
	}
}

// Tracef logs a message with formatting at level Trace on the standard logger.
func Tracef(arguments ...interface{}) {
	decorateF(TraceLevel, arguments, func(format string, args []any) {
		std.Tracef(format, args...)
	})
}

// Debugf logs a message with formatting at level Debug on the standard logger.
func Debugf(arguments ...interface{}) {
	decorateF(DebugLevel, arguments, func(format string, args []any) {
		std.Debugf(format, args...)
	})
}

// Printf logs a message with formatting at level Info on the standard logger.
func Printf(arguments ...interface{}) {
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		std.Printf(format, args...)
	})
}

// Infof logs a message with formatting at level Info on the standard logger.
func Infof(arguments ...interface{}) {
	decorateF(InfoLevel, arguments, func(format string, args []any) {
		std.Infof(format, args...)
	})
}

// Warnf logs a message at level Info on the standard logger.
func Warnf(arguments ...interface{}) {
	decorateF(WarnLevel, arguments, func(format string, args []any) {
		std.Warnf(format, args...)
	})
}

// Errorf logs a message with formatting at level Error on the standard logger.
func Errorf(arguments ...interface{}) {
	decorateF(ErrorLevel, arguments, func(format string, args []any) {
		std.Errorf(format, args...)
	})
}

// Panicf logs a message with formatting at level Panic on the standard logger.
func Panicf(arguments ...interface{}) {
	decorateF(PanicLevel, arguments, func(format string, args []any) {
		std.Panicf(format, args...)
	})
}

// Fatalf logs a message with formatting at level Fatal on the standard logger.
func Fatalf(arguments ...interface{}) {
	decorateF(FatalLevel, arguments, func(format string, args []any) {
		std.Fatalf(format, args...)
	})
}
