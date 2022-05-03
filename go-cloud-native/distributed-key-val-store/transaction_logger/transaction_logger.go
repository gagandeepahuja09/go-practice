package transaction_logger

type TransactionLogger interface {
	WritePut(key, value string)
	WriteDelete(key string)
}
