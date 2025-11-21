
# ⚽ ftbl-api

Простой REST API на **Go + MySQL** для управления футбольными командами и игроками.  
Проект демонстрирует реализацию **CRUD**, роутинг с **Chi** и работу с базой через **GORM**.

---

## 🚀 Технологии
-  **Go** (Chi, GORM)  
-  **MySQL** (через Docker)  
-  **Docker Compose**  
-  **godotenv** для конфигурации окружения  


---

## 🧩 Установка и запуск

```bash
# 1. Клонируем репозиторий
git clone https://github.com/daniyardautbaev/ftbl-api.git

# 2. Переходим в папку проекта


cd ftbl-api

# 3. Запускаем MySQL через Docker
docker compose up -d

# 4. Запускаем сервер
go run cmd/main.go
