package main

import (
	"hw1/api/http"
	"hw1/api/http/session"
	"hw1/repository/dbram"
	"hw1/repository/dbusers"
	"hw1/usecases/service"
	"log"

	_ "hw1/docs" // Загрузка сгенерированных файлов Swagger
)

// @title Task Manager API
// @version 1.0
// @description This is a simple API for managing tasks with status and result tracking.
// @host localhost:8080
// @BasePath /
func main() {
	port := ":8080"

	dbRam := dbram.NewObject()
	db := service.NewObject(dbRam)

	dbRamUser := dbusers.NewUserDB()
	dbUser := service.NewUser(dbRamUser)

	provider := session.NewMemoryProvider() //
	cookieName := "session_id"              // для записи
	maxLifetime := int64(3600)              // час

	dbManager := session.NewManager(provider, cookieName, maxLifetime)

	log.Println("server is running on port " + port)

	err := http.CreateAndRunServer(db, port, dbManager, dbUser)

	if err != nil {
		log.Fatalf("Fail to start %v\n", err)
	}
}

Я имею большой опыт, множество разных проектов и ролей в них, от разработки графового представления цифровых схем на с++ в университете, до сайта парсера резюме и умный поиск подходящих вакансий. Вследствие моего опыта я решил, что самым интересным и перспективным направлением для меня будет ds, меня на 2 курсе заинтересовала работа с нейронными сетями на одноименном майноре. Все указано в резюме. Самое интересное для меня оказалось работа в команде и сайт по hr резюме. Сейчас я прохожу также обучение в Школе 21, где активно изучаю Data Science. Мне нравится когда моя работа проявляется на проекте, как было  в поиске резюме и вакансий. Сейчас я стремлюсь углубить свои знания в области машинного обучения, улучшить навыки работы с пайплайнами и научиться применять модели в продакшене — от подготовки данных до развёртывания.

Я также хочу пройти Школу аналитиков-разработчиков, потому что давно интересуюсь анализом данных и хочу подтянуть как фундамент, так и практику — особенно в реальных проектах. Было интересно собирать полноценную систему с нуля, от данных до конечного продукта. Именно в таких задачах мне реально интересно — когда есть и инженерия, и интеллект, и польза. Хочу строить стабильные решения. После школы хочу развиваться как ds с хорошим бэкграундом в инженерии — строить пайплайны, работать с большими данными, подключать ML-модели. Интересуют сферы, где технологии реально могут менять жизни. Хочу быть тем, кто не просто делает модели, а помогает делать крутые и полезные продукты.
