# httpreq
Приложение, которое принимает HTTP-запросы для создания пользователя, добавления друзей пользователю, изменения возраста пользователя, удаления пользователя,
удаления списка друзей пользователя. Есть proxy и подняты две реплики данного приложения. Приложение взаимодействует с базой данных MongoDB.

В папке cmd находятся файлы с кодом для proxy сервера, а также для двух реплик приложения.
В папке pkg находятся:
- в папке init_DB - код для поднятия MongoDB.
- в папке user - все хендлеры с тестами на них, а также структуры и вспомогательные функции.
