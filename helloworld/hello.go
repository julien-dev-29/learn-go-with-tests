package helloworld

const (
	spanish = "Spanish"
	french  = "French"

	englishYoloPrefix = "Yolo "
	spanishYoloPrefix = "Hola "
	frenchYoloPrefix  = "Coucou "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "les kikis!"
	}
	return greetingPrefix(language) + name

}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishYoloPrefix
	case french:
		prefix = frenchYoloPrefix
	default:
		prefix = englishYoloPrefix
	}
	return
}