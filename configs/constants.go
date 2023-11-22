package configs

import (
	"time"
)

const MAX_STRING int = 255
const MIN_STRING int = 0
const MIN_PASSWORD int = 6
const MAX_PASSWORD int = 50
const MAX_TEXT int = 2550
const PHONE_NUMBER_LENGTH int = 12
const CODE_LENGTH int = 4

const MINUTE = time.Minute
const HOUR = time.Hour
const DAY = HOUR * 24

const ACCESS_TOKEN_EXPIRE = DAY
const REFRESH_TOKEN_EXPIRE = ACCESS_TOKEN_EXPIRE * 2

const SECOND = 1
const MINUTE_BY_SECOND = 60 * SECOND
const HOUR_BY_SECOND = 60 * MINUTE_BY_SECOND
const DAY_BY_SECOUND = 24 * HOUR_BY_SECOND

const ACCESS_TOKEN_EXPIRE_BY_SECOND = DAY_BY_SECOUND
const REFRESH_TOKEN_EXPIRE_BY_SECOND = ACCESS_TOKEN_EXPIRE_BY_SECOND * 2

const STATUS_EXPIRE = MINUTE * 5
const CODE_EXPIRE = HOUR / 2
const TAKE_EXPIRE = DAY * 31 * 3

const ESKIZ_BASE_URL string = "http://notify.eskiz.uz/api/"
const ESKIZ_CALLBACK_URL string = `${PROTOCOL}://${HOST}:${PORT}/eskiz/callback`

var ALLOW_IMAGE_MIME_TYPES = []string{"image/png", "image/jpeg", "image/webp"}
var ALLOW_VIDEO_MIME_TYPES = []string{"video/mp4"}

const TEST_QUESTION_COUNT_FOR_CUSTOMER int = 10
const TEST_QUESTION_TOTAL_COUNT int = 40

const CONTEXT_TIMEOUT = 10 * time.Second

var ENV = GetEnv()

type collection struct {
	User  string
	Video string
}

var Collection = collection{
	User:  "user",
	Video: "video",
}
