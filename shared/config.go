package shared

type Config struct {
	TimeoutSeconds uint16
	BearerToken    string
	BaseUserApiUrl string
}

func NewConfig() (result Config) {
	result = DefaultConfig

	//TODO: parse env tags
	//TODO: parse 'flag' tags
	return result
}

//goland:noinspection SpellCheckingInspection
var DefaultConfig = Config{
	TimeoutSeconds: 200,
	BearerToken:    "Bearer 8df384818f8c6c5f7abf35db59d497db56e839e808a3f6a2369bb4a921bcbde0", // test-hw3-go-njn
	BaseUserApiUrl: "https://gorest.co.in/public/v2/users",
}
