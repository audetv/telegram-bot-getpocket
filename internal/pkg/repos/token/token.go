package token

type Bucket string

const (
	AccessTokens  Bucket = "access_tokens"
	RequestTokens Bucket = "request_tokens"
)

// RepositoryInterface TODO Зачем тут этот интерфейс, переделать.
type RepositoryInterface interface {
	Create(chatID int64, token string, bucket Bucket) error
	Get(chatID int64, bucket Bucket) (string, error)
}
