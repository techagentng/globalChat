package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
)

//keep info of users
type  client struct{
	conn net.Conn
	nick string
	room *room
	commands chan <-command
}

func (c *client) readInput() { //Function reads the input of the user
	for {
		msg, err := bufio.NewReader(c.conn).ReadString('\n') //Reading client input
		if err != nil{
			return
		}
		msg = strings.Trim(msg, "\r\n") //Parse the input
		args := strings.Split(msg, " ")
		cmd := strings.TrimSpace(args[0]) //Gets the command name

		switch cmd {
		case "/nick":
			c.commands <- command{ //Sends message to channel if line 26 command name is found
				id: CMD_NICK,
				client: c,
				args: args,
			}
		case "/join":
			c.commands <- command{
				id: CMD_JOIN,
				client: c,
				args: args,
			}
		case "/rooms":
			c.commands <- command{
				id: CMD_ROOMS,
				client: c,
				args: args,
			}
		case "/msg":
			c.commands <- command{
				id: CMD_MSG,
				client: c,
				args: args,
			}
		case "/quit":
			c.commands <- command{
				id: CMD_QUIT,
				client: c,
				args: args,
			}
		default:  //If command not found, message is sent to the user
			c.err(fmt.Errorf("unknown command: %s", cmd))
		}
	}
}

func (c *client) err(err error) {
 c.conn.Write([]byte("ERR: " + err.Error() + "\n"))
}

func (c *client) msg(msg string) {
	c.conn.Write([]byte("> " + msg + "\n"))
}