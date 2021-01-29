package prioritize

import (
	"math/rand"
	"time"

	"github.com/sirupsen/logrus"
	extenderv1 "k8s.io/kube-scheduler/extender/v1"
)

func init() {
	rand.Seed(time.Now().UnixNano())

}

func Roll() int64 {

	number := rand.Int63n(extenderv1.MaxExtenderPriority + 1)
	logrus.Info("Rolled the dice and it is ", number)

	return number

}
