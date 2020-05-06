// Copyright 2013 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package barrage

import (
	"encoding/json"
	"sinblog.cn/FunAnime-Server/model"
	"strconv"
	"time"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client
}

type BarrageType struct {
	Key       string  `json:"key"`
	Time      float64 `json:"time"`
	Text      string  `json:"text"`
	FontSize  int     `json:"fontSize"`
	Color     string  `json:"color"`
	VideoID   string  `json:"videoId"`
	CreatorID int64   `json:"creatorId"`
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			go model.CreateBarrage(dealBarrageMessage(message))
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func dealBarrageMessage(message []byte) *model.FaBarrage {
	newBarrage := new(BarrageType)
	_ = json.Unmarshal(message, newBarrage)

	videoId, _ := strconv.ParseInt(newBarrage.VideoID, 10, 64)
	return &model.FaBarrage{
		VideoId:      videoId,
		Creator:      newBarrage.CreatorID,
		BarrageText:  string(message),
		Status:       1,
		CreateTime:   time.Now(),
		ModifyTime:   time.Now(),
	}
}
