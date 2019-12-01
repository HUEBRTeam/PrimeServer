package main

type Config struct {
	APIKey        string
	Online        bool
	MD5Check      bool
	ServerAddress string
	ScoreType     string // "default", "exscore"
}
