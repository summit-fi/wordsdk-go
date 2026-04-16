cldr-era-narrow = { $index ->
    *[0] до н. э.
     [1] н. э.
}

cldr-era = { $index ->
    *[0] до Рождества Христова
     [1] от Рождества Христова
}

cldr-month-narrow = { $index ->
    *[0] Я
     [1] Ф
     [2] М
     [3] А
     [4] М
     [5] И
     [6] И
     [7] А
     [8] С
     [9] О
    [10] Н
    [11] Д
}

cldr-month-standalone-narrow = { $index ->
    *[0] Я
     [1] Ф
     [2] М
     [3] А
     [4] М
     [5] И
     [6] И
     [7] А
     [8] С
     [9] О
    [10] Н
    [11] Д
}

cldr-month = { $index ->
    *[0] января
     [1] февраля
     [2] марта
     [3] апреля
     [4] мая
     [5] июня
     [6] июля
     [7] августа
     [8] сентября
     [9] октября
    [10] ноября
    [11] декабря
}

cldr-month-standalone = { $index ->
    *[0] январь
     [1] февраль
     [2] март
     [3] апрель
     [4] май
     [5] июнь
     [6] июль
     [7] август
     [8] сентябрь
     [9] октябрь
    [10] ноябрь
    [11] декабрь
}

cldr-month-short = { $index ->
    *[0] янв.
     [1] февр.
     [2] марта
     [3] апр.
     [4] мая
     [5] июня
     [6] июля
     [7] авг.
     [8] сент.
     [9] окт.
    [10] нояб.
    [11] дек.
}

cldr-month-standalone-short = { $index ->
    *[0] янв.
     [1] февр.
     [2] март
     [3] апр.
     [4] май
     [5] июнь
     [6] июль
     [7] авг.
     [8] сент.
     [9] окт.
    [10] нояб.
    [11] дек.
}

cldr-weekday-narrow = { $index ->
    *[0] В
     [1] П
     [2] В
     [3] С
     [4] Ч
     [5] П
     [6] С
}

cldr-weekday-standalone-narrow = { $index ->
    *[0] В
     [1] П
     [2] В
     [3] С
     [4] Ч
     [5] П
     [6] С
}

cldr-weekday = { $index ->
    *[0] воскресенье
     [1] понедельник
     [2] вторник
     [3] среда
     [4] четверг
     [5] пятница
     [6] суббота
}

cldr-weekday-standalone = { $index ->
    *[0] воскресенье
     [1] понедельник
     [2] вторник
     [3] среда
     [4] четверг
     [5] пятница
     [6] суббота
}

cldr-weekday-short = { $index ->
    *[0] Вс
     [1] Пн
     [2] Вт
     [3] Ср
     [4] Чт
     [5] Пт
     [6] Сб
}

cldr-weekday-standalone-short = { $index ->
    *[0] Вс
     [1] Пн
     [2] Вт
     [3] Ср
     [4] Чт
     [5] Пт
     [6] Сб
}

cldr-quater-short = { $index ->
    *[0] 1-й кв.
     [1] 2-й кв.
     [2] 3-й кв.
     [3] 4-й кв.
}

cldr-quater = { $index ->
    *[0] 1-й квартал
     [1] 2-й квартал
     [2] 3-й квартал
     [3] 4-й квартал
}

cldr-ampm = { $index ->
    *[0] AM
     [1] PM
}

cldr-date-patterns =
    .d = d
    .E = ccc
    .EEEE = cccc
    .LLL = LLL
    .LLLL = LLLL
    .M = L
    .Md = dd.MM
    .MEd = EEE, dd.MM
    .MMM = LLL
    .MMMd = d MMM
    .MMMEd = EEE, d MMM
    .MMMM = LLLL
    .MMMMd = d MMMM
    .MMMMEEEEd = EEEE, d MMMM
    .QQQ = QQQ
    .QQQQ = QQQQ
    .y = y
    .yM = MM.y
    .yMd = dd.MM.y
    .yMEd = EEE, d.MM.y
    .yMMM = LLL y г.
    .yMMMd = d MMM y г.
    .yMMMEd = EEE, d MMM y г.
    .yMMMM = LLLL y г.
    .yMMMMd = d MMMM y г.
    .yMMMMEEEEd = EEEE, d MMMM y г.
    .yQQQ = QQQ y г.
    .yQQQQ = QQQQ y г.
    .H = HH
    .Hm = HH:mm
    .Hms = HH:mm:ss
    .j = HH
    .jm = HH:mm
    .jms = HH:mm:ss
    .jmv = HH:mm v
    .jmz = HH:mm z
    .jz = HH z
    .m = m
    .ms = mm:ss
    .s = s
    .v = v
    .z = z
    .zzzz = zzzz
    .ZZZZ = ZZZZ


# cldr number

cldr-decimal-sep = ,

cldr-group-sep = {" "}

cldr-percent = %

cldr-zero-digit = 0

cldr-plus-sign = +

cldr-minus-sign = -

cldr-exp-symbol = E

cldr-permill = ‰

cldr-infinity = ∞

cldr-nan = NaN

cldr-decimal-pattern = ###,##0.###

cldr-scientific-pattern = #E0

cldr-percent-pattern = #,##0%

cldr-currency-pattern = ###,##0.00 ¤#

cldr-def-currency-code = UAH