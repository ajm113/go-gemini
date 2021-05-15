package gemini

import "os"

var (
	testPublicKey   = os.Getenv("GEMINI_PUBLIC_KEY")
	testSecret      = os.Getenv("GEMINI_SECRET")
	testEnvironment = Sandbox
)

var testOptions = ClientOptions{
	PublicKey:   testPublicKey,
	Secret:      testSecret,
	Environment: testEnvironment,
}

var testClient, _ = NewClient(testOptions)
