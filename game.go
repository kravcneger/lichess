package lichess

import "fmt"

type Game struct {
	Id         string `json:"id"`
	Rated      bool   `json:"rated"`
	Variant    string `json:"variant"`
	Speed      string `json:"speed"`
	Perf       string `json:"perf"`
	CreatedAt  uint   `json:"createdAt"`
	LastMoveAt uint   `json:"lastMoveAt"`
	Status     string `json:"status"`
	Players    struct {
		White struct {
			User       User
			Rating     int `json:"rating"`
			RatingDiff int `json:"ratingDiff"`
		}
		Black struct {
			User       User
			Rating     int `json:"rating"`
			RatingDiff int `json:"ratingDiff"`
		}
	}
	Moves string `json:"moves"`
	Clock struct {
		Initial   int `json:"initial"`
		Increment int `json:"increment"`
		TotalTime int `json:"totalTime"`
	}
}

func (c *Client) ExportGame(game_id string) (*Game, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/game/export/%s", game_id), nil)

	game := new(Game)
	_, err = c.do(req, &game)
	if err != nil {
		return nil, err
	}
	return game, err
}
