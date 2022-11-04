package demo_strcuture

type KMSHandler interface {
	Register()
	Encrypt(*KMSValueT) *KMSValueR
	Decrypt(*KMSValueT) *KMSValueR
}
