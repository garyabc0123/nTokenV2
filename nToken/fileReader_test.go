package nToken

import "testing"

func BenchmarkReadData(b *testing.B) {
	_, err := ReadData("/mnt/data/go/solr-ckip.dat")
	if err != nil {
		b.Error(err)
	}
}
