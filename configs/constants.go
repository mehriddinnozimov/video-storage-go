package configs

import "time"

const MAX_STRING int = 255
const MIN_STRING int = 0
const MIN_PASSWORD int = 6
const MAX_PASSWORD int = 50
const MAX_TEXT int = 2550
const PHONE_NUMBER_LENGTH int = 12
const CODE_LENGTH int = 4

const MINUTE int = 1000 * 60
const HOUR int = MINUTE * 60
const DAY int = HOUR * 24

const ACCESS_TOKEN_EXPIRE int = HOUR * 24
const REFRESH_TOKEN_EXPIRE int = ACCESS_TOKEN_EXPIRE * 2

const STATUS_EXPIRE int = MINUTE * 5
const CODE_EXPIRE int = HOUR / 2
const TAKE_EXPIRE int = DAY * 31 * 3

const ESKIZ_BASE_URL string = "http://notify.eskiz.uz/api/"
const ESKIZ_CALLBACK_URL string = `${PROTOCOL}://${HOST}:${PORT}/eskiz/callback`

var ALLOW_IMAGE_MIME_TYPES = [...]string{"image/png", "image/jpeg", "image/webp"}
var ALLOW_VIDEO_MIME_TYPES = [...]string{"video/mp4"}
var ALLOW_PDF_MIME_TYPES = [...]string{"application/pdf"}

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
