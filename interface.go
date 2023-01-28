package iterago

type Iterator[T any] interface {
	Next() Option[T]
	HasNext() bool
}
