package logic

import "log"

// broadcaster 广播器
type broadcaster struct {
	// 所有聊天室用户
	users map[string]*User

	// 所有 channel 统一管理，可以避免外部乱用
	enteringChannel chan *User
	leavingChannel  chan *User
	messageChanel   chan *Message

	// 判断该昵称用户是否可进入聊天室（重复与否）：true 能，false 不能
	checkUserChannel      chan string
	checkUserCanInChannel chan bool

	// 获取用户列表
	requestUsersChannel chan struct{}
	usersChannel        chan []*User
}

var Broadcaster = &broadcaster{
	users: make(map[string]*User),

	enteringChannel: make(chan *User),
	leavingChannel:  make(chan *User),
	messageChanel:   make(chan *Message),

	checkUserChannel:      make(chan string),
	checkUserCanInChannel: make(chan bool),

	requestUsersChannel: make(chan struct{}),
	usersChannel:        make(chan []*User),
}

func (b *broadcaster) Start() {
	for {
		select {
		case user := <-b.enteringChannel:
			b.users[user.NickName] = user

			// b.SendUserList()
		case user := <-b.leavingChannel:
			delete(b.users, user.NickName)
			// 避免 goroutine 泄露
			user.CloseMessageChannel()

			// b.SendUserList()
		case msg := <-b.messageChanel:
			if msg.To == "" {
				for _, user := range b.users {
					if user.UID == msg.User.UID {
						continue
					}
					user.MessageChannel <- msg
				}
				OfflineProcessor.Save(msg)
			} else {
				if user, ok := b.users[msg.To]; ok {
					user.MessageChannel <- msg
				} else {
					// 对方不在线或用户不存在，直接忽略消息
					log.Println("user:", msg.To, "not exists!")
				}
			}
		case nickname := <-b.checkUserChannel:
			if _, ok := b.users[nickname]; ok {
				b.checkUserCanInChannel <- false
			} else {
				b.checkUserCanInChannel <- true
			}
		case <-b.requestUsersChannel:
			userList := make([]*User, 0, len(b.users))
			for _, user := range b.users {
				userList = append(userList, user)
			}

			b.usersChannel <- userList
		}
	}
}

func (b *broadcaster) CanEnterRoom(nickname string) bool {
	b.checkUserChannel <- nickname

	return <-b.checkUserCanInChannel
}

func (b *broadcaster) UserEntering(u *User) {
	b.enteringChannel <- u
}

func (b *broadcaster) UserLeaving(u *User) {
	b.leavingChannel <- u
}

func (b *broadcaster) Broadcast(msg *Message) {
	b.messageChanel <- msg
}

func (b *broadcaster) GetUserList() []*User {
	b.requestUsersChannel <- struct{}{}
	return <-b.usersChannel
}
