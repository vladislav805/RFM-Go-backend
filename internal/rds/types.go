package rds

// https://tract.media/rds/

// PI (Programme Identification) – уникальный код радиостанции, состоящий из четырёх символов в 16-ричной системе.
// Первый символ означает страну вещания.
// Второй символ несёт в себе информацию о типе покрытия сигналом (локальная станция, использующая только один передатчик / региональная станция / национальная станция - вещающая во всех регионах страны).
// Третий и четвертый символы указывают на номер радиопрограммы (позывного). Последняя группа цифр позывного обозначает номер программы.
// Кстати: при сохранении радиостанции в пресет автомагнитолы в городе «Н» в память приёмника вписывается PI-код этой радиостанции (например, на кнопку «1») . При нахождении в городе «М», где данная радиостанция вещает на другой частоте, можно нажать на ту же самую кнопку пресета, которая была назначена в городе «Н», и если у приёмника есть функция PI Search (PI SEEK), он просканирует PI-код и поймёт, что это та самая станция, которую вы сохранили, и с лёгкостью её включит – достаточно лишь нажать нужную кнопку. В данном случае – «1».
// 1 - Germany, Moldova
// 2 - Estonia
// 3 - Poland, Turkey
// 9 - Latvia
// 6 - Ukraine, Finland
// 7 - Russia, Georgia, Armenia
// C - Литва
// E - Bulgaria, Romania
// F - Belarus
type PI uint16

// PTY (Programme TYpe) – код, обозначающий жанр радиостанции (всего 30 вариантов).
type PTY uint8

// PS (Programme Service name) – строка, предназначенная для сокращённого написания названия радиостанции – максимум 8 символов.
// Но, к сожалению, вопреки национальному и европейскому стандартам, где указано, что функция PS должна быть статической и не должна использоваться для передачи текста, многие радиостанции запихивают в эту строку что ни попадя: название песни, слоган радиостанции, полное название радиостанции, к тому же количество символов часто превышает допустимый предел.
type PS string

// RT (RadioText) – поле, куда передаётся всё, что радиостанции хотят донести до слушателей. Это может быть и название композиции или исполнителя, и адрес сайта, и слоган. В распоряжении радиостанций целых 64 или 32 символа – всё зависит от группы, в которой передаётся RT: 2А (64 символа) или 2В (32 символа).
type RT string

// AF (Alternative Frequencies list) - список альтернативных частот. Эта функция применяется во время поездок за пределы одного города. Вот, как она работает: когда вы покидаете зону слабого сигнала от вышки «А» и находитесь в зоне сильного сигнала от вышки «Б», ваш приёмник автоматически переключается на более сильный сигнал. Но для того чтобы это работало, должны соблюдаться определённые условия: должны совпадать PI-коды, в RDS-кодере должны быть прописаны альтернативные частоты (от 87.6 МГц до 107.9 МГц), сигнал хотя бы одной из них должен стать сильнее, чем текущая частота. К тому же список должен содержать те частоты, на которых подаётся сигнал RDS от кодера и которые имеют достаточное пересечение по покрытию.
// Прописывается AF одним из двух методов. Первый основан на ретрансляции сигнала в других городах (можно прописать до 25 частот). Второй необходим для того, чтобы делать переключение между региональными передатчиками, у которых есть основной сигнал, но программы и музыкальные композиции в разных городах сформированы по-разному.
type AF []uint32
