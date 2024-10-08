package main

import (
	"bytes"
	"fmt"
	"log"
	"reflect"
	"testing"

	swtpm_test "github.com/foxboron/swtpm_test"
	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

var (
	// Default SRK handle
	srkHandle tpmutil.Handle = 0x81000001

	// Default SRK handle
	localHandle tpmutil.Handle = 0x81010004

	// oaepLabel = []byte("age-encryption.org/v1/ssh-rsa")

	srkTemplate = tpm2.Public{
		Type:       tpm2.AlgRSA,
		NameAlg:    tpm2.AlgSHA256,
		Attributes: tpm2.FlagFixedTPM | tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth | tpm2.FlagRestricted | tpm2.FlagDecrypt | tpm2.FlagNoDA,
		RSAParameters: &tpm2.RSAParams{
			Symmetric: &tpm2.SymScheme{
				Alg:     tpm2.AlgAES,
				KeyBits: 128,
				Mode:    tpm2.AlgCFB,
			},
			KeyBits:    2048,
			ModulusRaw: make([]byte, 256),
		},
	}

	// This uses RSA/ES
	// TODO: Add test with RSA/AOEP stuff
	rsaKeyParamsDecrypt = tpm2.Public{
		Type:       tpm2.AlgRSA,
		NameAlg:    tpm2.AlgSHA256,
		Attributes: tpm2.FlagStorageDefault & ^tpm2.FlagRestricted,
		AuthPolicy: []byte{},
		RSAParameters: &tpm2.RSAParams{
			Sign: &tpm2.SigScheme{
				Alg:  tpm2.AlgOAEP,
				Hash: tpm2.AlgSHA256,
			},
			KeyBits:    2048,
			ModulusRaw: make([]byte, 256),
		},
	}
)

func TestCreateEncryptionAgeKey(t *testing.T) {
	var sealedHandle tpmutil.Handle
	var tpmPublicKeyDigest tpmutil.U16Bytes
	tpm := swtpm_test.NewSwtpm(t.TempDir())
	socket, err := tpm.Socket()
	if err != nil {
		t.Fatal(err)
	}

	rwc, err := tpm2.OpenTPM(socket)
	if err != nil {
		t.Fatal(err)
	}

	handle, _, err := tpm2.CreatePrimary(rwc, tpm2.HandleOwner, tpm2.PCRSelection{}, "", "", srkTemplate)
	if err != nil {
		t.Fatalf("failed CreatePrimary")
	}

	if err = tpm2.EvictControl(rwc, "", tpm2.HandleOwner, handle, srkHandle); err != nil {
		t.Fatalf("failed EvictControl")
	}

	t.Run("create key, persistent", func(t *testing.T) {
		priv, pub, _, _, _, err := tpm2.CreateKey(rwc, handle, tpm2.PCRSelection{}, "", "", rsaKeyParamsDecrypt)
		if err != nil {
			t.Fatalf("CreateKey error message: %v", err)
		}

		sealedHandle, _, err = tpm2.Load(rwc, srkHandle, "", pub, priv)
		if err != nil {
			t.Fatalf("Load error message: %v", err)
		}

		defer func() {
			err := tpm2.FlushContext(rwc, sealedHandle)
			if err != nil {
				log.Fatal(err)
			}
		}()

		if err = tpm2.EvictControl(rwc, "", tpm2.HandleOwner, sealedHandle, localHandle); err != nil {
			t.Fatalf("EvictControl error message: %v", err)
		}
	})

	t.Run("read persistent key", func(t *testing.T) {
		pub, _, _, err := tpm2.ReadPublic(rwc, localHandle)
		if err != nil {
			t.Fatalf("failed to Read Public: %v", err)
		}
		name, err := pub.Name()
		if err != nil {
			t.Fatalf("can't read public name: %v", err)
		}
		if !reflect.DeepEqual(name.Digest.Value, tpmPublicKeyDigest) {
			t.Fatalf("did not get the same key")
		}
	})

	t.Run("read persistent key", func(t *testing.T) {
		pub, _, _, err := tpm2.ReadPublic(rwc, localHandle)
		if err != nil {
			t.Fatalf("failed to Read Public: %v", err)
		}
		name, err := pub.Name()
		if err != nil {
			t.Fatalf("can't read public name: %v", err)
		}
		if !reflect.DeepEqual(name.Digest.Value, tpmPublicKeyDigest) {
			t.Fatalf("did not get the same key")
		}
	})

	t.Run("Encrypt/Decrypt", func(t *testing.T) {
		msg := []byte("test")
		scheme := &tpm2.AsymScheme{Alg: tpm2.AlgOAEP, Hash: tpm2.AlgSHA256}
		b, err := tpm2.RSAEncrypt(rwc, localHandle, msg, scheme, "label")
		if err != nil {
			t.Fatalf("failed RSAEncrypt: %v", err)
		}
		b, err = tpm2.RSADecrypt(rwc, localHandle, "", b, scheme, "label")
		if err != nil {
			t.Fatalf("failed RSADecrypt: %v", err)
		}
		fmt.Println(string(b))
		if !bytes.Equal(msg, b) {
			t.Fatalf("didn't match encrypted and decrypted things")
		}
	})
}
