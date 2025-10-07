# в процессе 
Седова М. А., ЭФМО-01-25
# Практическое задание 3 Реализация простого HTTP-сервера на стандартной библиотеке net/http. Обработка запросов GET/POST

### Обязательная часть
#### 1. Реализованные маршруты
- **GET /health** - возвращает `{"status":"ok"}`
<img width="1086" height="495" alt="image" src="https://github.com/user-attachments/assets/cc6dd79c-ff68-47f4-80a3-4287c87a6531" />

- **GET /tasks** - список задач с поддержкой фильтра `?q=текст`
<img width="1280" height="499" alt="image" src="https://github.com/user-attachments/assets/27e7229e-5571-44f3-9bc5-952472cf1666" />

- **POST /tasks** - создание задачи по `{"title":"..."}`
- **GET /tasks/{id}** - получение задачи по ID
