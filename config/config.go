package config

//Config stores bot ki'teer's config data
type Config struct {
	BotID         string
	CommandPrefix string
	BotAuthToken  string
}

func InitializeConfig(botID string) Config {

	configData := Config{
		BotID:         botID,
		CommandPrefix: "->b ",
		BotAuthToken:  "NjAwNjQxMDMxNzE2OTk1MDcy.XS7g5Q.VgOhZB2FlbzxDPfGR_D2YIphCqA",
	}

	return configData
}
