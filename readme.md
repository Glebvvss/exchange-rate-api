# Exchange rate api

API що дозволяє дізнатись поточній курс валют USD/UAH та підписатися на щоденне сповіщення поточного курсу.

## API
| Endpoint       | Description                              |
|----------------|------------------------------------------|
| /api/rate      | _дізнатись поточній курс валют_          |
| /api/subscribe | _підпісатися на щоденну email розсилку. Слід зауважити, що цей ендпоиінт приймає данні у двох форматах у форматі json та form-data в залежності від контенту body в запиті_  |

## CRON
| Crontab        | Description                              |
|----------------|------------------------------------------|
| 0 5 * * *      | _щоденна email розсилка поточного курсу валют (розсилка починається вранці тому що потенційно це час коли навантаження на сервер буде меншим у порівнянні з денним, проте не глибокої ночі бо курс може змінитись у проміжок часу між тим як надійшло посилання і користувач його прочитав)_|

## Деталі
- docker та docker compose (для старту додатку достатньо виконати команду docker compose up або docker-compose up)
- за замовчуванням додаток використовує наступний лінк https://localhost:8080
- конфігурація додатку знаходиться у файлі .env що підключаеться за допомогою змінних середовища
- розсилка email потребує заповнення змінних сереивища _EMAIL_FROM_ та _EMAIL_FROM_PASS_ у файлі .env. За замовчуванням використовується smtp сервер gmail, проте використовувати інші smtp сервери також можливо змінивши відповідні змінні у файлі .env
- логування помилок у файл (за замовчуванням app.log)
- міграції розташовані в директорії _migrations_ запускаються автоматично при старті додатку (слід зауважити, що після старту демона додатку підняття міграцій почнеться через 10 секунд щоб дати повністю ініціалізуватись конейнеру с базою данний)
- в разі якщо виникла проблема з міграціями при старті (таке може бути якщо ініціалізація контейнеру бази зайняла білше часу ніж передбачувалось) міграції можна запускати вручну за допомогою наступної команди _docker exec ex_rate_app -it "/migrate.bin"_
- тести слід запускати в контейнері тому що в них присутні залежності від змінних серидовища (можливо в майбутньому має сенс реалізувати інтеграційні тести які будут взаємодіяти з базою данних, але наразі вони відсутні тому що потребуют більш складної інфраструктури)