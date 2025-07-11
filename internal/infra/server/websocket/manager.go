package websocket

import "github.com/gorilla/websocket"

var Upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var userConnections = make(map[string]*websocket.Conn)
var userSubscriptions = make(map[string]string)

func RegisterConnection(user string, team string, conn *websocket.Conn) {
	userConnections[user] = conn
	userSubscriptions[user] = team
}

func RemoveConnection(user string) {
	delete(userConnections, user)
	delete(userSubscriptions, user)
}

func GetConnectionsByTeam(team string) []*websocket.Conn {
	var connections []*websocket.Conn
	for userID, subscribedTeam := range userSubscriptions {
		if subscribedTeam == team {
			conn := userConnections[userID]
			connections = append(connections, conn)
		}
	}
	return connections
}
