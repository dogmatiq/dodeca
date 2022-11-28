package logging

import "go.uber.org/zap"

// Zap returns a Logger that writes to a Zap logger.
func Zap(target *zap.Logger) Logger {
	return zapAdaptor{
		target.Sugar(),
		target.Core().Enabled(zap.DebugLevel),
	}
}

type zapAdaptor struct {
	target       *zap.SugaredLogger
	captureDebug bool
}

func (a zapAdaptor) Log(f string, v ...interface{}) {
	a.target.Infof(f, v...)
}

func (a zapAdaptor) LogString(s string) {
	a.target.Info(s)
}

func (a zapAdaptor) Debug(f string, v ...interface{}) {
	a.target.Debugf(f, v...)
}

func (a zapAdaptor) DebugString(s string) {
	a.target.Debug(s)
}

func (a zapAdaptor) IsDebug() bool {
	return a.captureDebug
}
