package imitative_typing

// 提供加载proto文件
import (
	"github.com/golang/protobuf/proto"
	"io/ioutil"
)

func LoadProto(path string, msg proto.Message) proto.Message {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}

	err = proto.UnmarshalText(string(data), msg)
	if err != nil {
		panic(err)
	}
	return msg
}

func TryLoadProto(path string, msg proto.Message) (proto.Message, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	err = proto.UnmarshalText(string(data), msg)
	if err != nil {
		return nil, err
	}
	return msg, nil
}
