package log

import (
	"bitbucket.org/bridce/ms-clean-code/configs"
	"fmt"
	"github.com/olivere/elastic/v7"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"gopkg.in/sohlich/elogrus.v7"
	"sync"
)

var (
	instace *LogCustom
	once    sync.Once
)

type LogCustom struct {
	Logrus *logrus.Logger
}

type stackTracer interface {
	StackTrace() errors.StackTrace
}

func ProvideConnElk(config *configs.Config) *LogCustom {
	return &LogCustom{
		Logrus: InitConnElk(config),
	}
}

func InitConnElk(config *configs.Config) *logrus.Logger {
	var log *logrus.Logger

	configElstc := config.Elastic

	log = logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{PrettyPrint: true})

	client, err := elastic.NewClient(elastic.SetURL(
		fmt.Sprintf("http://%v:%v", configElstc.Host, configElstc.Port)),
		elastic.SetSniff(false),
		elastic.SetBasicAuth(configElstc.Username, configElstc.Password))
	fmt.Println("http://"+configElstc.Host+":"+configElstc.Port, "<<< ---- elasticsearch url")
	if err != nil {
		selfLogError(err, "config/log: elastic client", log)
	} else {
		hook, err := elogrus.NewAsyncElasticHook(
			client, configElstc.Host, logrus.DebugLevel, configElstc.Index)
		if err != nil {
			selfLogError(err, "config/log: elastic client sync with logrus", log)
		}
		log.Hooks.Add(hook)
	}

	return log
}

func (l *LogCustom) Error(err error, description, respTime string, traceHeader map[string]string, req, resp, reqBE, respBE interface{}) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	l.Logrus.WithFields(logrus.Fields{
		"trace_header":  traceHeader,
		"error_cause":   stFormat,
		"error_message": err.Error(),
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Error(description)
}

func (l *LogCustom) Success(req, resp, reqBE, respBE interface{}, description, respTime string, traceHeader map[string]string) {

	l.Logrus.WithFields(logrus.Fields{
		"trace_header":  traceHeader,
		"request":       req,
		"response":      resp,
		"request_be":    reqBE,
		"response_be":   respBE,
		"response_time": respTime,
	}).Info(description)
}

func selfLogError(err error, description string, log *logrus.Logger) {
	err = errors.WithStack(err)
	st := err.(stackTracer).StackTrace()
	stFormat := fmt.Sprintf("%+v", st[1:2])

	log.WithFields(logrus.Fields{
		"error_cause":   stFormat,
		"error_message": err.Error(),
	}).Error(description)
}
