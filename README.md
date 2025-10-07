# в процессе 
Седова М. А., ЭФМО-01-25
# Практическое задание 3 Реализация простого HTTP-сервера на стандартной библиотеке net/http. Обработка запросов GET/POST

### Обязательная часть
- **GET /health** - возвращает `{"status":"ok"}`
<img width="1086" height="495" alt="image" src="https://github.com/user-attachments/assets/cc6dd79c-ff68-47f4-80a3-4287c87a6531" />

- **GET /tasks** - список задач с поддержкой фильтра `?q=текст`
<img width="834" height="600" alt="image" src="https://github.com/user-attachments/assets/73ceec42-6c18-4230-99f3-6f663eff29dd" />
<img width="1280" height="558" alt="image" src="https://github.com/user-attachments/assets/d686622d-2a42-4c05-8911-49dbb0fd1fae" />

- **POST /tasks** - создание задачи по `{"title":"..."}`
  <img width="1280" height="499" alt="image" src="https://github.com/user-attachments/assets/27e7229e-5571-44f3-9bc5-952472cf1666" />
  
- **GET /tasks/{id}** - получение задачи по ID
  <img width="1280" height="503" alt="image" src="https://github.com/user-attachments/assets/37cfd1d7-404a-41d9-bf45-9dd3db55175b" />

- **Создание пустой задачи**
<img width="1280" height="459" alt="image" src="https://github.com/user-attachments/assets/e815815b-f301-4a93-8149-6cf3c2e2d409" />

- **Проверка несуществующих задач**
<img width="1280" height="443" alt="image" src="https://github.com/user-attachments/assets/cef92877-49d6-4246-ba55-209fe0419055" />
<img width="1280" height="494" alt="image" src="https://github.com/user-attachments/assets/d57cb86c-aa10-4f15-8d7a-2fe76e31bec5" />

### Дополнительная часть

- **PATCH /tasks/{id}** - обновление задачи
<img width="1280" height="541" alt="image" src="https://github.com/user-attachments/assets/1464aece-d538-449b-9e9f-b851338040c5" />

- **DELETE /tasks/{id}** - удаление задачи
<img width="1280" height="428" alt="image" src="https://github.com/user-attachments/assets/bcb7f684-5274-49fd-ab3f-2b1752ef18c8" />

- **Валидация title** - 3-140 символов, иначе 422 ошибка
<img width="1280" height="578" alt="image" src="https://github.com/user-attachments/assets/5a78930a-9f3d-467f-b743-0d0617337803" />
<img width="1280" height="598" alt="image" src="https://github.com/user-attachments/assets/ddf31eb6-d5c5-40da-9a2a-661c213fb667" />

- **CORS-middleware** - заголовки для кросс-доменных запросов
<img width="1280" height="626" alt="image" src="https://github.com/user-attachments/assets/de3cd648-1fc6-4d21-bd7e-409e12fff38c" />
<img width="1280" height="668" alt="image" src="https://github.com/user-attachments/assets/8645b735-e911-4114-842e-871e9e0fa4d1" />

- **Graceful shutdown**
<img width="570" height="57" alt="image" src="https://github.com/user-attachments/assets/df754541-7960-4719-a3c9-e1eb061f4b2c" />


Файл REQUESTS.md содержит Postman коллекцию

- **Переменная PORT** - порт настраивается через $env:PORT
<img width="628" height="82" alt="image" src="https://github.com/user-attachments/assets/6aacb055-f50d-45b2-afcc-1b44bb18827f" />

- **PowerShell-скрипт** - run.ps1 с командами run, build, test
<img width="622" height="124" alt="image" src="https://github.com/user-attachments/assets/6b90c41b-ccdc-42ee-810b-c6e774008427" />
<img width="826" height="314" alt="image" src="https://github.com/user-attachments/assets/42adc422-e3a8-4fad-966f-14f4c7ceeea8" />

- ***Юнит-тесты** - 5 тестов в handlers_test.go с httptest
<img width="702" height="306" alt="image" src="https://github.com/user-attachments/assets/ddddd282-b905-4952-a06f-bd329acd718c" />

- **Дерево проекта**
<img width="465" height="654" alt="image" src="https://github.com/user-attachments/assets/dd3d0330-2b9f-486b-abf9-78cacfdb7f9b" />

