# SRS "Ленивый поварёнок"

## Introduction
Серёжа
### Purpose
Серёжа
### Document conventions
Серёжа
### Intended Audience and Reading Suggestions
Серёжа
### Project scope
Серёжа
### References
Серёжа
## Overall Description
Ваня
### Product perspective
Ваня
### Product features
Ваня
### User classes and characteristics
Ваня
### Operating environment
Ваня
### Design and implementation constraints
Ваня
### User documentation
Ваня
### Assumptions and dependencies
Ваня
## System features

### System feature — `Interface.product_list.take_photo`, ввод данных в список с помощью камеры

#### Description and priority

Позволяет пользователю использовать камеру устройства и отправляет полученное изображение на back-end, чтобы выделить продукты.

#### Stimulus/Response sequences

* Стимул: Пользователь нажимает на соотвествующую кнопку.

* Реакция: Запуск камеры устройства для съёмки фотографий.

Опция использует камеру устройства для создания фотографии и передачи результата на дальнейшую обработку, добавляя полученный результат в список продуктов.

### System feature — `Interface.product_list.keyboard`, ручной ввод данных в список

#### Description and priority

Позволяет пользователю вручную редактировать список продуктов в текущем списке.

#### Stimulus/Response sequences

* Стимул: Пользователь выбирает элемент списка и/или нажимает соответствующую кнопку.

* Реакция: Изменение состава текущего списка продуктов.

Опция должна поддерживать добавление, удаление и изменение элементов списка продуктов, при этом не изменяя состояния других элементов.

### System feature — `Interface.gen_recepies`, создание списка рецептов по спику продуктов 

#### Description and priority

Передаёт back-end'у текущий список продуктов и ожидает список рецептов.

#### Stimulus/Response sequences

* Стимул: Пользователь нажимает соответствующую кнопку.

* Реакция: Приложение возвращает пользователю список рецептов.

Ручка, отправляющая текущий список продуктов на back-end.

## External interface requirements

### User interfaces

Пользовательский интерфейс должен:

* Демонстрировать пользователю текущий список продуктов

* Позволять редактировать список

    + С помощью камеры устройства, считывая продукты с фотографии

    + Вручную добавлять и удалять продукты из списка

* Обладать элементом взаимодействия, с помощью которого активный список продуктов конвертируется в список рецептов

### Software interfaces
Общее
### Hardware interfaces
Серёжа
### Communication interfaces
Ваня, Олег
## Non functional requirements
Олег
### Performance requirements
Олег
### Safety requirements
Олег
### Software quality attributes
Олег
### Security requirements
Олег
## Other requirements

## Appendix A: Glossary
?
## Appendix B: Analysis Models
?
## Appendix C: Issues list
?