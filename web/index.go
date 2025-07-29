package web

import (
	"math/rand"

	"github.com/a-h/templ"

	"github.com/TecharoHQ/anubis/lib/localization"
	"github.com/TecharoHQ/anubis/lib/policy/config"
)

// GetRandomCharacter returns a random character ('a' or 'b') for image selection
func GetRandomCharacter() string {
	characters := []string{"a", "b"}
	return characters[rand.Intn(len(characters))]
}

// getRandomCharacter returns a random character ('a' or 'b') for image selection
func getRandomCharacter() string {
	return GetRandomCharacter()
}

func Base(title string, body templ.Component, impressum *config.Impressum, localizer *localization.SimpleLocalizer) templ.Component {
	character := getRandomCharacter()
	return base(title, body, impressum, nil, nil, localizer, character)
}

func BaseWithChallengeAndOGTags(title string, body templ.Component, impressum *config.Impressum, challenge string, rules *config.ChallengeRules, ogTags map[string]string, localizer *localization.SimpleLocalizer, character string) (templ.Component, error) {
	return base(title, body, impressum, struct {
		Rules     *config.ChallengeRules `json:"rules"`
		Challenge string                 `json:"challenge"`
	}{
		Challenge: challenge,
		Rules:     rules,
	}, ogTags, localizer, character), nil
}

func Index(localizer *localization.SimpleLocalizer) templ.Component {
	character := getRandomCharacter()
	return index(localizer, character)
}

func IndexWithCharacter(localizer *localization.SimpleLocalizer, character string) templ.Component {
	return index(localizer, character)
}

func ErrorPage(msg, mail string, localizer *localization.SimpleLocalizer) templ.Component {
	character := getRandomCharacter()
	return errorPage(msg, mail, localizer, character)
}

func Bench(localizer *localization.SimpleLocalizer) templ.Component {
	character := getRandomCharacter()
	return bench(localizer, character)
}
