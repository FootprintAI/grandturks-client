package cmd

import (
	"context"

	"github.com/sdinsure/agent/pkg/logger"
	k8slog "sigs.k8s.io/kind/pkg/log"
)

var (
	_ logger.Logger = &pkgLoggerAdaptor{}
)

type pkgLoggerAdaptor struct {
	k8sLogger k8slog.Logger
}

func newPkgLoggerAdaptor(k8sLogger k8slog.Logger) *pkgLoggerAdaptor {
	return &pkgLoggerAdaptor{
		k8sLogger: k8sLogger,
	}
}

func (p *pkgLoggerAdaptor) Info(fmtStr string, values ...interface{}) {
	p.k8sLogger.V(1).Infof(fmtStr, values...)
}

func (p *pkgLoggerAdaptor) Warn(fmtStr string, values ...interface{}) {
	p.k8sLogger.Warnf(fmtStr, values...)
}

func (p *pkgLoggerAdaptor) Error(fmtStr string, values ...interface{}) {
	p.k8sLogger.Errorf(fmtStr, values...)
}

func (p *pkgLoggerAdaptor) Fatal(fmtStr string, values ...interface{}) {
}

func (p *pkgLoggerAdaptor) Infox(_ context.Context, fmtStr string, values ...interface{}) {
	p.k8sLogger.V(1).Infof(fmtStr, values...)
}

func (p *pkgLoggerAdaptor) Warnx(_ context.Context, fmtStr string, values ...interface{}) {
	p.k8sLogger.Warnf(fmtStr, values...)
}

func (p *pkgLoggerAdaptor) Errorx(_ context.Context, fmtStr string, values ...interface{}) {
	p.k8sLogger.Errorf(fmtStr, values...)
}
