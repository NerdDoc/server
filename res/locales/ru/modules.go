package ru

import (
	"github.com/NerdDoc/server/modules"
)

func init() {
	modules.RegisterModules("ru", []modules.Module{
		// AREA
		// For modules related to countries, please add the translations of the countries' names
		// or open an issue to ask for translations.

		{
			Tag: modules.AreaTag,
			Patterns: []string{
				"Какая площадь у",
				"Сколько занимает ",
			},
			Responses: []string{
				"Площадь %s составляет %g квадратных километров ",
			},
			Replacer: modules.AreaReplacer,
		},

		// CAPITAL
		{
			Tag: modules.CapitalTag,
			Patterns: []string{
				"Какая столица у ",
				"Столица ",
				"Назови мне столицу ",
			},
			Responses: []string{
				"Столицей %s является город %s",
			},
			Replacer: modules.CapitalReplacer,
		},

		// CURRENCY
		{
			Tag: modules.CurrencyTag,
			Patterns: []string{
				"Какая валюта используется в",
				"Какая валюта у ",
				"Валюта ",
			},
			Responses: []string{
				"Валютой %s является %s",
			},
			Replacer: modules.CurrencyReplacer,
		},

		// MATH
		// A regex translation is also required in `language/math.go`, please don't forget to translate it.
		// Otherwise, remove the registration of the Math module in this file.

		{
			Tag: modules.MathTag,
			Patterns: []string{
				"Скажи мне результат ",
				"Посчитай ",
			},
			Responses: []string{
				"Результат %s",
				"Это будет %s",
			},
			Replacer: modules.MathReplacer,
		},

		// MOVIES
		// A translation of movies genres is also required in `language/movies.go`, please don't forget
		// to translate it.
		// Otherwise, remove the registration of the Movies modules in this file.

		{
			Tag: modules.GenresTag,
			Patterns: []string{
				"Я люблю приключения и анимацию",
				"Мне нравятся мюзиклы и фантастика",
				"Я смотрю фильмы sci-fi",
			},
			Responses: []string{
				"Отличный выбор! А запомню это.",
				"Поняла! Я отправлю эту информацию твоему клиенту.",
			},
			Replacer: modules.GenresReplacer,
		},

		{
			Tag: modules.MoviesTag,
			Patterns: []string{
				"Можешь найти мне фильм о",
				"Найди мне фильм",
				"Я хотел бы посмотреть фильм о",
			},
			Responses: []string{
				"Вот что я нашла для тебя “%s” с рейтингом %.02f/5",
				"Конечно, я тут нашла фильм “%s” рейтинг у него %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesAlreadyTag,
			Patterns: []string{
				"Я уже видел этот фильм",
				"Я его уже посмотрел",
				"О, я уже смотрел этот фильм",
			},
			Responses: []string{
				"О, у меня есть еще один “%s” его рейтинг %.02f/5",
			},
			Replacer: modules.MovieSearchReplacer,
		},

		{
			Tag: modules.MoviesDataTag,
			Patterns: []string{
				"Мне скучно",
				"Я незнаю что мне делать",
			},
			Responses: []string{
				"Хочу предложить тебе посмотреть фильм %s “%s” рейтинг %.02f/5",
			},
			Replacer: modules.MovieSearchFromInformationReplacer,
		},

		// NAME
		{
			Tag: modules.NameGetterTag,
			Patterns: []string{
				"Ты знаешь как меня зовут?",
				"кто я",
				"как меня зовут",
				"как меня звать",
				"какое у меня имя",
			},
			Responses: []string{
				"Тебя зовут %s!",
				"Ты %s",
			},
			Replacer: modules.NameGetterReplacer,
		},

		{
			Tag: modules.NameSetterTag,
			Patterns: []string{
				"Меня зовут ",
				"Ты можешь звать меня ",
			},
			Responses: []string{
				"Отлично! Привет %s",
				"Как дела %s?",
			},
			Replacer: modules.NameSetterReplacer,
		},

		// RANDOM
		{
			Tag: modules.RandomTag,
			Patterns: []string{
				"Скажи мне случайное число",
				"Сгенерируй случайное число",
			},
			Responses: []string{
				"Число %s",
				"Твое случайное число %s",
			},
			Replacer: modules.RandomNumberReplacer,
		},

		// REMINDERS
		// Translations are required in `language/date/date`, `language/date/rules` and in `language/reason`,
		// please don't forget to translate it.
		// Otherwise, remove the registration of the Reminders modules in this file.

		{
			Tag: modules.ReminderSetterTag,
			Patterns: []string{
				"Напомни мне приготовить завтрак в 8 вечера",
				"Напомни мне позвонить маме во вторник",
				"Запиши, что у меня завтра экзамен",
				"Напомни мне, что у меня завтра телефонная конференция в 9 вечера",
				"Уведоми меня что 5 января у меня выходной",
			},
			Responses: []string{
				"Записала! Я напомню тебе: “%s” для %s",
			},
			Replacer: modules.ReminderSetterReplacer,
		},

		{
			Tag: modules.ReminderGetterTag,
			Patterns: []string{
				"О чем я просил тебя мне напомнить",
				"Дай мне список напоминаний",
				"Какие у меня напоминания на сегодня",
				"Список напоминаний",
				"Что ты должна напомнить",
			},
			Responses: []string{
				"Ты просил меня напомнить о следующих вещах: \n %s",
			},
			Replacer: modules.ReminderGetterReplacer,
		},

		// SPOTIFY
		// A translation is needed in `language/music`, please don't forget to translate it.
		// Otherwise, remove the registration of the Spotify modules in this file.

		{
			Tag: modules.SpotifySetterTag,
			Patterns: []string{
				"Вот мои токены spotify",
				"Мои spotify секреты",
			},
			Responses: []string{
				"Логинюсь, подождите",
			},
			Replacer: modules.SpotifySetterReplacer,
		},

		{
			Tag: modules.SpotifyPlayerTag,
			Patterns: []string{
				"Воспроизвести на Spotify",
			},
			Responses: []string{
				"Воспроизвожу %s от %s на Spotify.",
			},
			Replacer: modules.SpotifyPlayerReplacer,
		},

		{
			Tag: modules.JokesTag,
			Patterns: []string{
				"Расскажи мне шутку",
				"Заставь меня смеяться",
				"Пошути",
			},
			Responses: []string{
				"Ну вот, %s",
				"Вот одна, %s",
			},
			Replacer: modules.JokesReplacer,
		},
	})

	// COUNTRIES
	// Please translate this method for adding the correct article in front of countries names.
	// Otherwise, remove the countries modules from this file.

	modules.ArticleCountries["ru"] = ArticleCountries
}

// ArticleCountries returns the country with its article in front.
func ArticleCountries(name string) string {
	if name == "United States" {
		return "the " + name
	}

	return name
}
