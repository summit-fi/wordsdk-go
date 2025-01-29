
shared-photos =
    {$userName} {$photoCount ->
    [one] added a new photo
    *[other] added {$photoCount} new photos
        }to {$userGender ->
    [male] his stream
    [female] her stream
    *[other] their stream
        }.

### A Resource comment
-brand-name = Mozilla
-other-brand = Chrome

use-case = { $Browser ->
[mozilla] {-brand-name }
*[other] { -other-brand }
}.


spot =
    { $center ->
    *[hotel] Welcome to your { room }
    [restaurant] Welcome to your { table }
    [tennis]  Welcome to your { court }
       }.

room = { $count ->
*[one] one room
[other] {$count} rooms
}

table = { $count ->
*[one] one table
[other] {$count} tables
}

court = { $count ->
*[one] one court
[other] {$count} courts
}

function-test = { TIME($date) }

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