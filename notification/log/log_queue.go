package log

import (
	"github.com/golang/glog"
	"google.golang.org/protobuf/proto"

	"github.com/bookstairs/bookworm/notification"
	"github.com/bookstairs/bookworm/util"
)

func init() {
	notification.MessageQueues = append(notification.MessageQueues, &LogQueue{})
}

type LogQueue struct {
}

func (k *LogQueue) GetName() string {
	return "log"
}

func (k *LogQueue) Initialize(configuration util.Configuration, prefix string) (err error) {
	return nil
}

func (k *LogQueue) SendMessage(key string, message proto.Message) (err error) {

	glog.V(0).Infof("%v: %+v", key, message)
	return nil
}
