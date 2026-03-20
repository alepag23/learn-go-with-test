package helloworld

const (
	franch              = "Franch"
	franchBonjourPrefix = "Bonjour "
	spanish             = "Spanish"
	spanishHolaPrefix   = "Hola "
	englishHelloPrefix  = "Hello "
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHolaPrefix
	case franch:
		prefix = franchBonjourPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}
