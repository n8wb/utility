/**
 * Copyright 2019 Whiteblock Inc. All rights reserved.
 * Use of this source code is governed by a BSD-style
 * license that can be found in the LICENSE file.
 */
package rabbitmq

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/streadway/amqp"
	"github.com/whiteblock/utility/utils"
	"golang.org/x/sync/semaphore"
	"sync"
)

type Client struct {
	Queue         string
	QueueURL      string
	MaxConcurreny int64
	MaxRetries    int64
	callback      func(msg amqp.Delivery) error
	conn          *amqp.Connection
	once          *sync.Once
	sem           *semaphore.Weighted
}

func (c *Client) Init(callback func(msg amqp.Delivery) error) error {
	c.once = &sync.Once{}
	if c.MaxConcurreny < 1 {
		return fmt.Errorf("MaxConcurreny must be atleast 1")
	}
	c.sem = semaphore.NewWeighted(c.MaxConcurreny)
	if c.MaxRetries < 1 {
		return fmt.Errorf("MaxRetries must be atleast 1")
	}
	c.callback = callback
	return c.init()
}

// CreateQueue creates the coresponding queue with the given parameters
func (c *Client) CreateQueue(durable, autoDelete, exclusive, noWait bool, args amqp.Table) error {
	ch, err := c.conn.Channel()
	if err != nil {
		return utils.LogError(err)
	}
	defer ch.Close()
	_, err = ch.QueueDeclare(c.Queue, durable, autoDelete, exclusive, noWait, args)
	return utils.LogError(err)
}

// Close cleans up the connections and resources used by this client
func (c *Client) Close() {
	if c == nil {
		return
	}
	if c.conn != nil {
		c.conn.Close()
	}
}

// Run starts the client. This function should be called only once and does not return
func (c *Client) Run() {
	c.once.Do(func() { c.loop() })
}

func (c *Client) kickBackMessage(msg amqp.Delivery) {
	pub := amqp.Publishing{
		Headers: msg.Headers,
		// Properties
		ContentType:     msg.ContentType,
		ContentEncoding: msg.ContentEncoding,
		DeliveryMode:    msg.DeliveryMode,
		Priority:        msg.Priority,
		CorrelationId:   msg.CorrelationId,
		ReplyTo:         msg.ReplyTo,
		Expiration:      msg.Expiration,
		MessageId:       msg.MessageId,
		Timestamp:       msg.Timestamp,
		Type:            msg.Type,
		Body:            msg.Body,
	}
	_, exists := pub.Headers["retryCount"]
	if !exists {
		pub.Headers["retryCount"] = int64(0)
	}
	if pub.Headers["retryCount"].(int64) > c.MaxRetries {
		log.WithFields(log.Fields{"retries": c.MaxRetries}).Debug("discarded message after too many retries")
		return
	}
	pub.Headers["retryCount"] = pub.Headers["retryCount"].(int64) + 1
	ch, err := c.conn.Channel()
	if err != nil {
		utils.LogError(err)
		return
	}
	defer ch.Close()
	err = ch.Publish(msg.Exchange, msg.RoutingKey, false, false, pub)
	if err != nil {
		utils.LogError(err)
		return
	}

	utils.LogError(msg.Reject(false))
}

func (c *Client) handleMessage(msg amqp.Delivery) {
	defer c.sem.Release(1)
	err := c.callback(msg)
	if err != nil {
		utils.LogError(err)
		go c.kickBackMessage(msg)
		return
	}
	utils.LogError(msg.Ack(false))
}

func (c *Client) loop() {
	ch, err := c.conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(c.Queue, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}
	for msg := range msgs {
		c.sem.Acquire(context.Background(), 1)
		go c.handleMessage(msg)
	}
}

func (c *Client) init() (err error) {
	c.conn, err = amqp.Dial(c.QueueURL)
	return utils.LogError(err)
}
