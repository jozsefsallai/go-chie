// Package chie is a powerful encryption/decryption tool for Cave Story's TSC
// script format
//
// Decryption
//
//	tsc := chie.NewTSCParser()
//
//	err := tsc.FromFile("Ballo1.tsc")
//	if err != nil {
//		panic(err)
//	}
//
//	err = tsc.Decrypt().ToFile("Ballo1.txt")
//	if err != nil {
//		panic(err)
//	}
//
// Encryption
//	tsc := chie.NewTSCParser()
//
//	err := tsc.FromFile("Ballo1.txt")
//	if err != nil {
//		panic(err)
//	}
//
//	err = tsc.Encrypt().ToFile("Ballo1.tsc")
//	if err != nil {
//		panic(err)
//	}
package chie
