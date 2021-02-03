package news

import (
	"time"

	"github.com/go-kit/kit/log"
)

// LoggingMiddleware logs method calls, parameters, results, and elapsed time.
func LoggingMiddleware(logger log.Logger) Middleware {
	return func(next Service) Service {
		return loggingMiddleware{
			next:   next,
			logger: logger,
		}
	}
}

type loggingMiddleware struct {
	next   Service
	logger log.Logger
}

func (mw loggingMiddleware) List() (socks []NewsItem, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "List",
			"result", len(socks),
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.List()
}

func (mw loggingMiddleware) Health() (health []Health) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Health",
			"result", len(health),
			"took", time.Since(begin),
		)
	}(time.Now())
	return mw.next.Health()
}
