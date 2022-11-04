package demo_strcuture

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_AdapterTestForLoop(t *testing.T) {

	adapter := NewKMSHandler(KMS_GCP)

	adapter.Register()

	adapter.Init()

	gw := sync.WaitGroup{}

	for i := 0; i < 6000; i++ {
		gw.Add(1)

		go func() {
			encrypt := adapter.Encrypt(&KMSValueT{plaintext: []byte("Test Text"), aead: []byte("Test EncryptKey")})

			decrypt := adapter.Decrypt(&KMSValueT{plaintext: encrypt.result, aead: []byte("Test EncryptKey")})

			fmt.Printf("Cipher Text : %s \n", encrypt.result)
			fmt.Printf("Decrypted Text :  %s \n", decrypt.result)

			gw.Done()
		}()
	}

	gw.Wait()

	assert.Equal(t, 1, 1)
}

// Test_AdapterTestForLoopStartFromInit 10000번 호출하는 코드 일때
// - gcp_key_projectsample/demo_strcuture.(*KMSGcp).Init - 40376612 byte 필요 / 40 Mb
//  - 1,000번 -> 4 mb
// 	- 10,000번 -> 40 mb
// 	- 100,000번 -> 400 mb
// 	- 1,000,000번 -> 4,000 mb -> 4 gb
// 	- 10,000,000번 -> 40,000 mb -> 40 gb
//	- 100,000,000번 -> 400,000 mb -> 400 gb
// - gcp_key_projectsample/demo_strcuture.(*KMSGcp).Encrypt - 26742180 byte 필요 / 26 mb
// - gcp_key_projectsample/demo_strcuture.(*KMSGcp).Encrypt - 15207436 byte 필요 / 15 mb
func Test_AdapterTestForLoopStartFromInit(t *testing.T) {

	adapter := NewKMSHandler(KMS_GCP)

	adapter.Register()

	for i := 0; i < 10000; i++ {
		adapter.Init()

		encrypt := adapter.Encrypt(&KMSValueT{plaintext: []byte("Test Text"), aead: []byte("Test EncryptKey")})

		decrypt := adapter.Decrypt(&KMSValueT{plaintext: encrypt.result, aead: []byte("Test EncryptKey")})

		fmt.Printf("Cipher Text : %s \n", encrypt.result)
		fmt.Printf("Decrypted Text :  %s \n", decrypt.result)
	}
}

func Test_AdapterTestForSingle(t *testing.T) {

	adapter := NewKMSHandler(KMS_GCP)

	adapter.Register()

	adapter.Init()

	encrypt := adapter.Encrypt(&KMSValueT{plaintext: []byte("Test Text"), aead: []byte("Test EncryptKey")})

	decrypt := adapter.Decrypt(&KMSValueT{plaintext: encrypt.result, aead: []byte("Test EncryptKey")})

	fmt.Printf("Cipher Text : %s \n", encrypt.result)
	fmt.Printf("Decrypted Text :  %s \n", decrypt.result)

	assert.Equal(t, 1, 1)
}
