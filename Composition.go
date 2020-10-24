package main
import (
	"fmt"
)

type Game struct {
	*Player
	gameID int
}

type Player struct {
	name string
	age int
	location string
}

func (p *Player) showWelcomeMessage() string {
	displayMessage := fmt.Sprintf("Welcome Mr.%s",p.name)
	return displayMessage
}

func newGame(name string, age int,location string, gameid int) *Game  {
	
	return &Game{
		Player: &Player{name,age,location},
		gameID: gameid,
	}
}

func main() {

	// playerInfo := Game {
	// 				Player{"Kalidoss",22,"Chennai"},
	// 				123,
	// }

	playerInfo := newGame("Kalidoss",29,"Chennai",12345)
	fmt.Println(playerInfo.showWelcomeMessage())
	fmt.Println(playerInfo.name)

}