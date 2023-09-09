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

// Trace logs a message at level Trace on the standard logger.
func Trace(arguments ...interface{}) {
	if ignored(TraceLevel) {
		return
	}
	std.Trace(decorate(arguments...)...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(arguments ...interface{}) {
	if ignored(DebugLevel) {
		return
	}
	std.Debug(decorate(arguments...)...)
}

// Print logs a message at level Info on the standard logger.
func Print(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Print(decorate(arguments...)...)
}

// Info logs a message at level Info on the standard logger.
func Info(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Info(decorate(arguments...)...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warn(decorate(arguments...)...)
}

// Warning logs a message at level Warn on the standard logger.
func Warning(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warning(decorate(arguments...)...)
}

// Error logs a message at level Error on the standard logger.
func Error(arguments ...interface{}) {
	std.Error(decorate(arguments...)...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(arguments ...interface{}) {
	std.Panic(decorate(arguments...)...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(arguments ...interface{}) {
	std.Fatal(decorate(arguments...)...)
}
