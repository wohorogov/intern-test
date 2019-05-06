Тестовое задание Сима-ленд мобайл.

**Каталог.**

1\. Сделать список категорий плитками картинка + название. Данные для отображения получить по запросу

```
GET https://www.sima-land.ru/api/v3/category/?is_not_empty=1&with_adult=1&page=1&level=1&sort=priority_home&expand-root=1&expand=name_alias
```

![catalog][catalog]

2\. При нажатии на категорию открыть список подкатегорий. *path* для запроса взять из родительской категории

```
 GET https://www.sima-land.ru/api/v3/category/?is_not_empty=1&with_adult=1&sort=priority&expand-root=1&expand=name_alias&level=2&path=3003
``` 

![subcategory][subcategory]


Готовое решение сделайте на Swift, если потребуется подключить сторонние библиотеки, подключайте с использованием CocoaPods

Проект залейте в открытый репозиторий GitHub/Gitlab/BitBucket/... или закрытый GitHub с открытым доступом юзеру `WildyDS`

Пример реализации можно увидеть в приложении [Сима-ленд](https://itunes.apple.com/ru/app/сима-ленд-интернет-магазин/id1057565689)

[catalog]: images/catalog.PNG "catalog"
[subcategory]: images/subcategory.PNG "subcategory"
