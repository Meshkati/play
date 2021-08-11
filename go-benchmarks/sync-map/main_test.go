package main

import "testing"

func BenchmarkSyncMap(b *testing.B) {
	SyncMap(b.N)
}

func BenchmarkRegularMap(b *testing.B) {
	RegularMap(b.N)
}

func BenchmarkSyncMapRw(b *testing.B) {
	SyncMapReadWrite(1, b.N)
}

func BenchmarkRegularMapRW(b *testing.B) {
	RegularMapReadWrite(1, b.N)
}

func BenchmarkSyncMapR5W1(b *testing.B) {
	SyncMapReadWrite(5, b.N)
}

func BenchmarkRegularMapR5W1(b *testing.B) {
	RegularMapReadWrite(5, b.N)
}

func BenchmarkSyncMapR100W1(b *testing.B) {
	SyncMapReadWrite(100, b.N)
}

func BenchmarkRegularMapR100W1(b *testing.B) {
	RegularMapReadWrite(100, b.N)
}

func BenchmarkSyncMapR999W1(b *testing.B) {
	SyncMapReadWrite(100, b.N)
}

func BenchmarkRegularMapR999W1(b *testing.B) {
	RegularMapReadWrite(100, b.N)
}
