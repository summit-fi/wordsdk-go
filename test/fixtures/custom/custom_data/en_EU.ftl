cldr-era-narrow = { $index ->
    *[0] BC
     [1] AD
}

cldr-era = { $index ->
    *[0] Before Christ
     [1] Anno Domini
}

cldr-month-narrow = { $index ->
    *[0] J
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
    *[0] J
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
    *[0] January
     [1] February
     [2] March
     [3] April
     [4] May
     [5] June
     [6] July
     [7] August
     [8] September
     [9] October
    [10] November
    [11] December
}

cldr-month-standalone = { $index ->
    *[0] January
     [1] February
     [2] March
     [3] April
     [4] May
     [5] June
     [6] July
     [7] August
     [8] September
     [9] October
    [10] November
    [11] December
}

cldr-month-short = { $index ->
    *[0] Jan
     [1] Feb
     [2] Mar
     [3] Apr
     [4] May
     [5] Jun
     [6] Jul
     [7] Aug
     [8] Sep
     [9] Oct
    [10] Nov
    [11] Dec
}

cldr-month-standalone-short = { $index ->
    *[0] Jan
     [1] Feb
     [2] Mar
     [3] Apr
     [4] May
     [5] Jun
     [6] Jul
     [7] Aug
     [8] Sep
     [9] Oct
    [10] Nov
    [11] Dec
}

cldr-weekday-narrow = { $index ->
    *[0] S
     [1] M
     [2] T
     [3] W
     [4] T
     [5] F
     [6] S
}

cldr-weekday-standalone-narrow = { $index ->
    *[0] S
     [1] M
     [2] T
     [3] W
     [4] T
     [5] F
     [6] S
}

cldr-weekday = { $index ->
    *[0] Sunday
     [1] Monday
     [2] Tuesday
     [3] Wednesday
     [4] Thursday
     [5] Friday
     [6] Saturday
}

cldr-weekday-standalone = { $index ->
    *[0] Sunday
     [1] Monday
     [2] Tuesday
     [3] Wednesday
     [4] Thursday
     [5] Friday
     [6] Saturday
}

cldr-weekday-short = { $index ->
    *[0] Sun
     [1] Mon
     [2] Tue
     [3] Wed
     [4] Thu
     [5] Fri
     [6] Sat
}

cldr-weekday-standalone-short = { $index ->
    *[0] Sun
     [1] Mon
     [2] Tue
     [3] Wed
     [4] Thu
     [5] Fri
     [6] Sat
}

cldr-quater-short = { $index ->
    *[0] Q1
     [1] Q2
     [2] Q3
     [3] Q4
}

cldr-quater = { $index ->
    *[0] 1st quarter
     [1] 2nd quarter
     [2] 3rd quarter
     [3] 4th quarter
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
    .Md = M/d
    .MEd = EEE, M/d
    .MMM = LLL
    .MMMd = MMM d
    .MMMEd = EEE, MMM d
    .MMMM = LLLL
    .MMMMd = MMMM d
    .MMMMEEEEd = EEEE, MMMM d
    .QQQ = QQQ
    .QQQQ = QQQQ
    .y = y
    .yM = M/y
    .yMd = M/d/y
    .yMEd = EEE, M/d/y
    .yMMM = MMM y
    .yMMMd = MMM d, y
    .yMMMEd = EEE, MMM d, y
    .yMMMM = MMMM y
    .yMMMMd = MMMM d, y
    .yMMMMEEEEd = EEEE, MMMM d, y
    .yQQQ = QQQ y
    .yQQQQ = QQQQ y
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

cldr-decimal-sep = .

cldr-group-sep = ,

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

cldr-currency-pattern = ¤#,##0.00

cldr-def-currency-code = EUR