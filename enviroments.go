package gemini

type Environment string

const (
	// Sandbox Hits sandbox servers. Does NOT pull live data.
	Sandbox Environment = "https://api.sandbox.gemini.com/v1"
	// Live Hits production servers. PULLS live data.
	Live Environment = "https://api.gemini.com/v1"
)

var validEnvironments = [...]Environment{
	Sandbox,
	Live,
}

// IsValid Checks if the enviroment supplied is valid.
func (e Environment) IsValid() bool {
	for _, v := range validEnvironments {
		if e == v {
			return true
		}
	}
	return false
}
