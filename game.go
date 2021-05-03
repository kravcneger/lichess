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

func (c *Client) ExportOngoingGameOfUser(user_id string) (*Game, error) {
	req, err := c.newRequest("GET", fmt.Sprintf("/api/user/%s/current-game", user_id), nil)

	game := new(Game)
	_, err = c.do(req, &game)
	if err != nil {
		return nil, err
	}
	return game, err
}

type OngoingGame struct {
	FullId   string `json:"fullId"`
	GameId   string `json:"gameId"`
	Fen      string `json:"fen"`
	Color    string `json:"color"`
	LastMove string `json:"lastMove"`
	Speed    string `json:"speed"`
	Perf     string `json:"perf"`
	Rated    bool   `json:"rated"`
	Opponent struct {
		Id       string `json:"id"`
		Username string `json:"username"`
		Rating   int    `json:"rating"`
	}
	IsMyTurn bool `json:"isMyTurn"`
	//	GameState *GameState
}

type OngoingGames struct {
	Games []OngoingGame `json:"NowPlaying"`
}

func (c *Client) OngoingGames() (*[]OngoingGame, error) {
	req, err := c.newRequest("GET", "/api/account/playing", nil)

	ongoingGames := new(OngoingGames)
	_, err = c.do(req, &ongoingGames)
	if err != nil {
		return nil, err
	}
	return &ongoingGames.Games, nil
}
