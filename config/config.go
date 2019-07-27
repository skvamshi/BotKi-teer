package config

//Config stores bot ki'teer's config data
type Config struct {
	BotID         string
	CommandPrefix string
	BotAuthHeader string
}

func InitializeConfig() *Config {

	return &Config{
		BotID:         "",
		CommandPrefix: "->b ",
		BotAuthHeader: "Bot NjAwNjQxMDMxNzE2OTk1MDcy.XS7g5Q.VgOhZB2FlbzxDPfGR_D2YIphCqA",
	}

}
