package wood

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

// Traceln logs a message at level Trace on the standard logger.
func Traceln(arguments ...interface{}) {
	if ignored(TraceLevel) {
		return
	}
	std.Traceln(decorate(arguments...)...)
}

// Debugln logs a message at level Debug on the standard logger.
func Debugln(arguments ...interface{}) {
	if ignored(DebugLevel) {
		return
	}
	std.Debugln(decorate(arguments...)...)
}

// Println logs a message at level Info on the standard logger.
func Println(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Println(decorate(arguments...)...)
}

// Infoln logs a message at level Info on the standard logger.
func Infoln(arguments ...interface{}) {
	if ignored(InfoLevel) {
		return
	}
	std.Infoln(decorate(arguments...)...)
}

// Warnln logs a message at level Warn on the standard logger.
func Warnln(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warnln(decorate(arguments...)...)
}

// Warningln logs a message at level Warn on the standard logger.
func Warningln(arguments ...interface{}) {
	if ignored(WarnLevel) {
		return
	}
	std.Warningln(decorate(arguments...)...)
}

// Errorln logs a message at level Error on the standard logger.
func Errorln(arguments ...interface{}) {
	std.Errorln(decorate(arguments...)...)
}

// Panicln logs a message at level Panic on the standard logger.
func Panicln(arguments ...interface{}) {
	std.Panicln(decorate(arguments...)...)
}

// Fatalln logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalln(arguments ...interface{}) {
	std.Fatalln(decorate(arguments...)...)
}
