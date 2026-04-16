cldr-era-narrow = { $index ->
    *[0] a. C.
     [1] d. C.
}

cldr-era = { $index ->
    *[0] antes de Cristo
     [1] después de Cristo
}

cldr-month-narrow = { $index ->
    *[0] E
     [1] F
     [2] M
     [3] A
     [4] M
     [5] J
     [6] J
     [7] A
     [8] S
     [9] O
    [10] N
    [11] D
}

cldr-month-standalone-narrow = { $index ->
    *[0] E
     [1] F
     [2] M
     [3] A
     [4] M
     [5] J
     [6] J
     [7] A
     [8] S
     [9] O
    [10] N
    [11] D
}

cldr-month = { $index ->
    *[0] enero
     [1] febrero
     [2] marzo
     [3] abril
     [4] mayo
     [5] junio
     [6] julio
     [7] agosto
     [8] septiembre
     [9] octubre
    [10] noviembre
    [11] diciembre
}

cldr-month-standalone = { $index ->
    *[0] enero
     [1] febrero
     [2] marzo
     [3] abril
     [4] mayo
     [5] junio
     [6] julio
     [7] agosto
     [8] septiembre
     [9] octubre
    [10] noviembre
    [11] diciembre
}

cldr-month-short = { $index ->
    *[0] ene
     [1] feb
     [2] mar
     [3] abr
     [4] may
     [5] jun
     [6] jul
     [7] ago
     [8] sep
     [9] oct
    [10] nov
    [11] dic
}

cldr-month-standalone-short = { $index ->
    *[0] ene
     [1] feb
     [2] mar
     [3] abr
     [4] may
     [5] jun
     [6] jul
     [7] ago
     [8] sept
     [9] oct
    [10] nov
    [11] dic
}

cldr-weekday-narrow = { $index ->
    *[0] D
     [1] L
     [2] M
     [3] X
     [4] J
     [5] V
     [6] S
}

cldr-weekday-standalone-narrow = { $index ->
    *[0] D
     [1] L
     [2] M
     [3] X
     [4] J
     [5] V
     [6] S
}

cldr-weekday = { $index ->
    *[0] domingo
     [1] lunes
     [2] martes
     [3] miércoles
     [4] jueves
     [5] viernes
     [6] sábado
}

cldr-weekday-standalone = { $index ->
    *[0] domingo
     [1] lunes
     [2] martes
     [3] miércoles
     [4] jueves
     [5] viernes
     [6] sábado
}

cldr-weekday-short = { $index ->
    *[0] dom
     [1] lun
     [2] mar
     [3] mié
     [4] jue
     [5] vie
     [6] sáb
}

cldr-weekday-standalone-short = { $index ->
    *[0] dom
     [1] lun
     [2] mar
     [3] mié
     [4] jue
     [5] vie
     [6] sáb
}

cldr-quater-short = { $index ->
    *[0] T1
     [1] T2
     [2] T3
     [3] T4
}

cldr-quater = { $index ->
    *[0] 1.º trimestre
     [1] 2.º trimestre
     [2] 3.º trimestre
     [3] 4.º trimestre
}

cldr-ampm = { $index ->
    *[0] a.m.
     [1] p.m.
}

cldr-date-patterns =
    .d = d
    .E = ccc
    .EEEE = cccc
    .LLL = LLL.
    .LLLL = LLLL
    .M = L
    .Md = d/M
    .MEd = EEE, d/M
    .MMM = LLL.
    .MMMd = d 'de' MMM.
    .MMMEd = EEE, d 'de' MMM.
    .MMMM = LLLL
    .MMMMd = d 'de' MMMM
    .MMMMEEEEd = EEEE, d 'de' MMMM
    .QQQ = QQQ
    .QQQQ = QQQQ
    .y = y
    .yM = M/y
    .yMd = d/M/y
    .yMEd = EEE, d/M/y
    .yMMM = MMM. 'de' y
    .yMMMd = d 'de' MMM. 'de' y
    .yMMMEd = EEE, d MMM. y
    .yMMMM = MMMM 'de' y
    .yMMMMd = d 'de' MMMM 'de' y
    .yMMMMEEEEd = EEEE, d 'de' MMMM 'de' y
    .yQQQ = QQQ 'de' y
    .yQQQQ = QQQQ 'de' y
    .H = H
    .Hm = H:mm
    .Hms = H:mm:ss
    .j = h a
    .jm = h:mm a
    .jms = h:mm:ss a
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

cldr-group-sep = .

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

cldr-currency-pattern = ¤ #,##0

cldr-def-currency-code = COP