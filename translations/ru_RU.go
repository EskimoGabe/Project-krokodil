package translations

func initRussianTranslation() {
	translation := createTranslation()

	translation.put("requires-js", "Этот веб-сайт требует JavaScript для правильной работы.")

	translation.put("start-the-game", "Начало игры")
	translation.put("start", "Начинать")
	translation.put("game-not-started-title", "Ожидание начала игры")
	translation.put("waiting-for-host-to-start", "Подождите, пока владелец лобби не начнет игру.")

	translation.put("last-turn", "(последний ход: %s)")

	translation.put("round", "раунд")
	translation.put("toggle-soundeffects", "Звук включен- / выключить")
	translation.put("change-your-name", "Изменить имя пользователя")
	translation.put("randomize", "случайное имя")
	translation.put("apply", "Использовать")
	translation.put("save", "Сохранить")
	translation.put("toggle-fullscreen", "Активировать полноэкранный режим / деактивировать")
	translation.put("show-help", "показать помощь")
	translation.put("votekick-a-player", "Проголосовал за кик игрока")

	translation.put("change-lobby-settings-tooltip", "Изменить настройки лобби")
	translation.put("change-lobby-settings-title", "Настройки лобби")
	translation.put("lobby-settings-changed", "Настройки лобби изменены")
	translation.put("advanced-settings", "Расширенные настройки")
	translation.put("word-language", "Язык слов")
	translation.put("drawing-time-setting", "время рисования")
	translation.put("rounds-setting","Раунды")
	translation.put("max-players-setting", "Максимальное количество игроков")
	translation.put("public-lobby-setting", "Общественное лобби")
	translation.put("custom-words", "лишние слова")
	translation.put("custom-words-info", "Введите здесь дополнительные слова, разделяя отдельные слова запятой.")
	translation.put("custom-words-chance-setting", "Шанс на лишнее слово")
	translation.put("players-per-ip-limit-setting", "Максимальное количество игроков на IP")
	translation.put("enable-votekick-setting", "Разрешить голосование за удаление")
	translation.put("save-settings", "Сохранить настройки")
	translation.put("input-contains-invalid-data", "Ваши записи содержат неверные данные:")
	translation.put("please-fix-invalid-input", "Пожалуйста, исправьте свои записи и повторите попытку.")
	translation.put("create-lobby", "Создать комнату")

	translation.put("players", "Игрок")
	translation.put("refresh", "Обновить")
	translation.put("join-lobby", "Присоединиться к лобби")

	translation.put("message-input-placeholder", "Введите ответы и сообщения здесь")

	translation.put("choose-a-word", "выбрать слово")
	translation.put("waiting-for-word-selection", "Ожидание выбора слова")
	//This one doesn't use %s, since we want to make one part bold.
	translation.put("is-choosing-word", "просто выберите слово.")

	translation.put("close-guess", "'%s' почти.")
	translation.put("correct-guess", "Вы правильно угадали слово.")
	translation.put("correct-guess-other-player", "'%s' правильно угадал слово.")
	translation.put("round-over", "твоя очередь окончен, слово не выбрано.")
	translation.put("round-over-no-word", "твоя очередь, выбранное слово было '%s'.")
	translation.put("game-over-win", "Поздравляем, вы выиграли!")
	translation.put("game-over-tie", "Ничья!")
	translation.put("game-over", "Ты %s. c %s Баллы")

	translation.put("change-active-color", "Изменить активный цвет")
	translation.put("use-pencil", "использовать ручку")
	translation.put("use-eraser", "использовать ластик")
	translation.put("use-fill-bucket", "Используйте ведро для наполнения (Заполняет целевую область выбранным цветом)")
	translation.put("change-pencil-size-to", "Сменить ручку / ластик увеличить размер %s")
	translation.put("clear-canvas", "Очистить монтажную область")
	translation.put("undo", "Отменить последнее изменение (Не работает после \""+translation.Get("clear-canvas")+"\")")

	translation.put("connection-lost", "Соединение потеряно!")
	translation.put("connection-lost-text", "Попробуйте переподключиться"+
		" ...\n\nубеждаться, что ваше интернет-соединение работает.\nЕсли это "+
		"проблема сохраняется, связаться с веб-мастером.")
	translation.put("error-connecting", "Ошибка установления соединения")
	translation.put("error-connecting-text",
		"Scribble.rs не удалось установить соединение через сокет.\n\nПравда твой кажется "+
			"Интернет работал, но либо были сервера, либо \nваш брандмауэр настроен неправильно.\n\n"+
			"Попробуйте перезагрузить страницу.")

	translation.put("message-too-long", "Ваше сообщение слишком длинное.")

	//Help dialog
	translation.put("controls", "Элементы управления")
	translation.put("pencil", "Ручка")
	translation.put("eraser", "ластик")
	translation.put("fill-bucket", "Заливка")
	translation.put("switch-tools-intro", "Вы можете переключаться между инструментами с помощью сочетаний клавиш")
	translation.put("switch-pencil-sizes", "Вы можете изменить размер пера с помощью клавиш от %s до %s.")

	//Generic words
	//"close" as in "closing the window"
	translation.put("close", "Закрыть")
	translation.put("no", "нет")
	translation.put("yes", "Да")
	translation.put("system", "система")

	//Footer
	translation.put("source-code", "исходный код")
	translation.put("help", "Помощь")
	translation.put("contact", "Контакт")
	translation.put("submit-feedback", "Обратная связь")
	translation.put("stats", "статус")

	RegisterTranslation("ru", translation)
	RegisterTranslation("ru-ru", translation)
}
