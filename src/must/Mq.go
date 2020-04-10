package must

import (
	"encoding/json"
	"fmt"
	"zhanyia/src/common"

	"github.com/streadway/amqp"
)

type Mq struct {
	MQConnection *amqp.Connection
}

// 创建Mq实例
func init() {
	common.AllGlobal["Mq"] = &Mq{}
	conn, err := amqp.Dial("amqp://admin:123456@127.0.0.1:5672/")
	if err != nil {
		panic(err)
	}
	common.AllGlobal["Mq"].(*Mq).MQConnection = conn
}

const (
	ExchangeName = "LttFanout"
	ReportName   = "LttReport"
	RoomName     = "RedMi"
)

// 绑定报表
func (obj *Mq) BindReportQueue() {
	// 获取携程
	ch, err := obj.MQConnection.Channel()
	if err != nil {
		fmt.Println("Mq SubHubBindQueueExchange Channel has err", err)
		return
	}
	// 交换机声明
	_, err = ch.QueueDeclare(ReportName, true, true, false, false, nil)
	if err != nil {
		fmt.Println("Mq SubHubBindQueueExchange QueueDeclare has err", err)
		return
	}
	err = ch.Qos(1, 0, false)
	if err != nil {
		fmt.Println("Mq SubHubBindQueueExchange qos has err ", err)
		return
	}

	go func() {
		// 创建消息者
		msgList, err := ch.Consume(ReportName, "ReportConsumer", false, false, false, false, nil)
		if err != nil {
			fmt.Println("Mq SubHubBindQueueExchange Consume has err", err)
			return
		}

		// 读取mq获取的数据
		for msg := range msgList {
			// 处理数据
			var msgStr string
			fmt.Println(msg)
			err = json.Unmarshal(msg.Body, &msgStr)
			if err != nil {
				fmt.Println("获取消息失败", err)
				// 读取失败-否认
				err = msg.Ack(false)
				if err != nil {
					fmt.Println("MQ Consume Ack false has err", err)
					continue
				}
			} else {
				// 读取成功-确认
				err = msg.Ack(true)
				if err != nil {
					fmt.Println("MQ Consume Ack true has err", err)
					continue
				}
				fmt.Println("获取消息成功！", msgStr)
			}
		}
	}()
}

// 设定交换机消费
func (obj *Mq) ExchangeRoomNameQueue() {
	// 获取携程
	ch, err := obj.MQConnection.Channel()
	if err != nil {
		fmt.Println("Mq ExchangeRoomNameQueue Channel has err", err)
		return
	}

	err = ch.ExchangeDeclare("RouteMe", "direct", false, true, false, false, nil)
	if err != nil {
		fmt.Println("Mq ExchangeRoomNameQueue ExchangeDeclare has err", err)
		return
	}
}

// SendReport 发送
func (obj *Mq) SendReport(msg []byte) error {

	// 获取携程
	ch, err := obj.MQConnection.Channel()
	if err != nil {
		fmt.Println("Mq SendReport Channel has err", err)
		return err
	}

	// 发送给报表
	err = ch.Publish("RouteMe", "", false, false, amqp.Publishing{
		ContentType: "text/plain",
		Body:        msg,
	})

	if err != nil {
		fmt.Println("Mq SendReport Publish has err", err)
		return err
	}

	return nil
}
