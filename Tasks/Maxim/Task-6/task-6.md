Task 5. Нужно описать модель библиотеки.
Есть 3 сущности: "Автор", "Книга", "Читатель"
Физически книга только одна и может быть только у одного читателя.
Нужно составить таблицы для библиотеки так, чтобы это учесть. 


Есть з сущности - пользователь, чат, сообщение
У пользователя есть имя и дата регистрации у чата есть название и дата создания
У сообщения есть текст, автор и дата создания
Пользователь может состоять в нескольких чатах одновременно
Сообщение обязательно принадлежит чату, сообщение не может принадлежать более чем 1 чату одновременно
Нужно описать предметную область в виде таблиц

```sql
CREATE TABLE user(
    id INT PRIMARY KEY AUTOINCREMENT, 
    name VARCHAR (100) NOT NULL, 
    created DATETIME
) ;

CREATE TABLE chat(
    id INT PRIMARY KEY AUTOINCREMENT, 
    name VARCHAR (100) NOT NULL, 
    created DATETIME
) ;

CREATE TABLE message (
    id INT PRIMARY KEY AUTOINCREMENT, 
    user_id INT NOT NULL, 
    chat_id INT NOT NULL, 
    created DATETIME,
    FOREIGN KEY chat_id REFERENCES chat (id)
    FOREIGN KEY user_id REFERENCES user(id)
)

CREATE TABLE user_chat(
    user_id INT
    chat_id INT
    FOREIGN KEY chat_id REFERENCES chat (id)
    FOREIGN KEY user_id REFERENCES user (id)
    PRIMARY KEY(user_id, chat_id)
)

// Выбрать все чаты пользователя Вася в формате (chat_id, chat_name)

SELECT c. chat_id, c. name
FROM chat c
JOIN user_chat uc
ON c. id=uc. chat_id
JOIN user u
ON uc.user_id= user (u.user_id)
WHERE u. name="Vasia"
```