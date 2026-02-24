cldr-era-narrow = { $index ->
    *[0] до н. е.
     [1] н. е.
}

cldr-era = { $index ->
    *[0] до нашої ери
     [1] нашої ери
}

cldr-month-narrow = { $index ->
    *[0] с
     [1] л
     [2] б
     [3] к
     [4] т
     [5] ч
     [6] л
     [7] с
     [8] в
     [9] ж
    [10] л
    [11] г
}

cldr-month-standalone-narrow = { $index ->
    *[0] С
     [1] Л
     [2] Б
     [3] К
     [4] Т
     [5] Ч
     [6] Л
     [7] С
     [8] В
     [9] Ж
    [10] Л
    [11] Г
}

cldr-month = { $index ->
    *[0] січня
     [1] лютого
     [2] березня
     [3] квітня
     [4] травня
     [5] червня
     [6] липня
     [7] серпня
     [8] вересня
     [9] жовтня
    [10] листопада
    [11] грудня
}

cldr-month-standalone = { $index ->
    *[0] січень
     [1] лютий
     [2] березень
     [3] квітень
     [4] травень
     [5] червень
     [6] липень
     [7] серпень
     [8] вересень
     [9] жовтень
    [10] листопад
    [11] грудень
}

cldr-month-short = { $index ->
    *[0] січ.
     [1] лют.
     [2] бер.
     [3] квіт.
     [4] трав.
     [5] черв.
     [6] лип.
     [7] серп.
     [8] вер.
     [9] жовт.
    [10] лист.
    [11] груд.
}

cldr-month-standalone-short = { $index ->
    *[0] січ.
     [1] лют.
     [2] бер.
     [3] квіт.
     [4] трав.
     [5] черв.
     [6] лип.
     [7] серп.
     [8] вер.
     [9] жовт.
    [10] лист.
    [11] груд.
}

cldr-weekday-narrow = { $index ->
    *[0] Н
     [1] П
     [2] В
     [3] С
     [4] Ч
     [5] П
     [6] С
}

cldr-weekday-standalone-narrow = { $index ->
    *[0] Н
     [1] П
     [2] В
     [3] С
     [4] Ч
     [5] П
     [6] С
}

cldr-weekday = { $index ->
    *[0] неділя
     [1] понеділок
     [2] вівторок
     [3] середа
     [4] четвер
     [5] пʼятниця
     [6] субота
}

cldr-weekday-standalone = { $index ->
    *[0] неділя
     [1] понеділок
     [2] вівторок
     [3] середа
     [4] четвер
     [5] пʼятниця
     [6] субота
}

cldr-weekday-short = { $index ->
    *[0] нд
     [1] пн
     [2] вт
     [3] ср
     [4] чт
     [5] пт
     [6] сб
}

cldr-weekday-standalone-short = { $index ->
    *[0] нд
     [1] пн
     [2] вт
     [3] ср
     [4] чт
     [5] пт
     [6] сб
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
    *[0] дп
     [1] пп
}

cldr-date-patterns =
    .d = d
    .E = ccc
    .EEEE = cccc
    .LLL = LLL
    .LLLL = LLLL
    .M = MM
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
    .yMEd = EEE, dd.MM.y
    .yMMM = LLL y р.
    .yMMMd = d MMM y р.
    .yMMMEd = EEE, d MMM y р.
    .yMMMM = LLLL y р.
    .yMMMMd = d MMMM y р.
    .yMMMMEEEEd = EEEE, d MMMM y р.
    .yQQQ = QQQ y
    .yQQQQ = QQQQ y р.
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

cldr-decimal-pattern = #,##0.###

cldr-scientific-pattern = #E0

cldr-percent-pattern = #,##0%

cldr-currency-pattern = #,##0.00 ¤

cldr-def-currency-code = UAH

#custom
date-format-datetime = { CLDRDATETIME($date, pattern: "yMdjm") }
