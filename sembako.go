package main

import (
	"fmt"
)

const NMAX int = 2023

type dataBarang struct {
	namaBarang  string
	harga, stok int
}

type dataTransaksi struct {
	namaBarang, jenis string
	waktu             date
	nilai             int
	jumlahWaktu       int
}

type date struct {
	tanggal, bulan, tahun int
}

type arrBarang [NMAX]dataBarang
type arrTransaksi [NMAX]dataTransaksi

func inputDataBarang(T *arrBarang, N *int) {
	var i, harga, stok int
	var nama string

	fmt.Println("Masukkan data barang:")
	fmt.Scan(&nama, &harga, &stok)
	for T[i].namaBarang != "x" && T[i].harga != -1 && T[i].stok != -1 {
		T[i].namaBarang = nama
		T[i].harga = harga
		T[i].stok = stok
		fmt.Scan(&nama, &harga, &stok)
		i++
	}
	*N = len(T)
}

func inputDataTransaksi(T *arrTransaksi, N *int) {
	var i, tanggal, bulan, tahun int
	var nama, jenis string

	fmt.Println("Masukkan data transaksi:")
	fmt.Scan(&nama, &tanggal, &bulan, &tahun, &jenis)

	for T[i].namaBarang != "x" && T[i].waktu.tanggal != -1 && T[i].waktu.bulan != -1 && T[i].waktu.tahun != -1 && T[i].jenis != "x" {
		T[i].namaBarang = nama
		T[i].waktu.tanggal = tanggal
		T[i].waktu.bulan = bulan
		T[i].waktu.tahun = tahun
		T[i].jenis = jenis
		T[i].jumlahWaktu = (365 * tahun) + (30 * bulan) + tanggal
		fmt.Scan(&nama, &tanggal, &bulan, &tahun, &jenis)
		i++
	}
	*N = len(T)
}

func idxData_nama(T arrBarang, N int, x string) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].namaBarang == x
		i++
	}
	return i
}

func idxData_harga(T arrBarang, N int, x int) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].harga == x
		i++
	}
	return i
}

func idxData_stok(T arrBarang, N int, x int) int {
	var found bool
	var i int

	for i < N && !found {
		found = T[i].stok == x
		i++
	}
	return i
}

func idxMaksimum(T arrBarang, X arrTransaksi, N, N1, x int) int {
	var idx int
	var j int = 1

	if x == 1 {
		for j < N {
			if T[idx].harga < T[j].harga {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 2 {
		for j < N {
			if T[idx].stok < T[j].stok {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 3 {
		for j < N {
			if X[idx].jumlahWaktu < X[j].jumlahWaktu {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 4 {
		for j < N {
			if X[idx].nilai < X[j].nilai {
				idx = j
			}
			j++
		}
		return idx
	}
	return 0
}

func idxMinimum(T arrBarang, X arrTransaksi, N, N1, x int) int {
	var idx int
	var j int = 1

	if x == 1 {
		for j < N {
			if T[idx].harga > T[j].harga {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 2 {
		for j < N {
			if T[idx].stok > T[j].stok {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 3 {
		for j < N {
			if X[idx].jumlahWaktu > X[j].jumlahWaktu {
				idx = j
			}
			j++
		}
		return idx
	} else if x == 4 {
		for j < N {
			if X[idx].nilai > X[j].nilai {
				idx = j
			}
			j++
		}
		return idx
	}
	return 0
}

func menu(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("1. modifikasi data")
	fmt.Println("2. cari data")
	fmt.Println("3. tampilan data")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		modifikasi(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cari(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {

	}
}

func modifikasi(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, idx int
	var nama string

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("modifikasi:")
	fmt.Println("1. tambah data")
	fmt.Println("2. ubah data")
	fmt.Println("3. hapus data")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		modifikasiTambah(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nama barang: ")
		fmt.Scan(&nama)

		idx = idxData_nama(*T, *N, nama)
		modifikasiUbah(&*T, &*X, &*N, &*N1, idx)
	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nama barang: ")
		fmt.Scan(&nama)

		idx = idxData_nama(*T, *N, nama)
		modifikasiHapus(&*T, &*X, &*N, &*N1, idx)
	}
}

func modifikasiTambah(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var harga, stok, tanggal, bulan, tahun int

	fmt.Print("\nbuat data")
	fmt.Print("\nnama: ")
	fmt.Scan(&nama)
	fmt.Print("harga: ")
	fmt.Scan(&harga)
	fmt.Print("stok: ")
	fmt.Scan(&stok)
	fmt.Print("waktu (DD MM YYYY): ")
	fmt.Scan(&tanggal, &bulan, &tahun)

	T[*N+1].namaBarang = nama
	T[*N+1].harga = harga
	T[*N+1].stok = stok
	fmt.Println("\nsukses ditambah")

	X[*N1+1].namaBarang = nama
	X[*N1+1].waktu.tanggal = tanggal
	X[*N1+1].waktu.bulan = bulan
	X[*N1+1].waktu.tahun = tahun
	X[*N1+1].jenis = "tambah"
	X[*N1+1].nilai = stok
}

func modifikasiHapus(T *arrBarang, X *arrTransaksi, N, N1 *int, idx int) {
	var tanggal, bulan, tahun int

	fmt.Print("waktu (DD MM YYYY): ")
	fmt.Scan(&tanggal, &bulan, &tahun)

	X[*N1+1].namaBarang = T[idx].namaBarang
	X[*N1+1].waktu.tanggal = tanggal
	X[*N1+1].waktu.bulan = bulan
	X[*N1+1].waktu.tahun = tahun
	X[*N1+1].jenis = "hapus"
	X[*N1+1].nilai = 0 - T[idx].stok

	for i := idx; i < *N; i++ {
		T[i] = T[i+1]
	}
	fmt.Println("\nsukses dihapus")
}

func modifikasiUbah(T *arrBarang, X *arrTransaksi, N, N1 *int, idx int) {
	var pilihan, tanggal, bulan, tahun int
	var a string
	var b, c int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("modifikasi:")
	fmt.Println("1. nama barang")
	fmt.Println("2. harga barang")
	fmt.Println("3. stok")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		fmt.Print("ubah nama: ")
		fmt.Scan(&a)
		T[idx].namaBarang = a
		fmt.Println("sukses diubah")
	} else if pilihan == 2 {
		fmt.Print("ubah harga: ")
		fmt.Scan(&b)
		T[idx].harga = b
		fmt.Println("sukses diubah")

	} else if pilihan == 3 {
		fmt.Print("ubah stok: ")
		fmt.Scan(&c)
		fmt.Print("waktu (DD MM YYYY): ")
		fmt.Scan(&tanggal, &bulan, &tahun)
		T[idx].stok = c
		fmt.Println("\nsukses diubah")
		X[*N1+1].namaBarang = T[idx].namaBarang
		X[*N1+1].waktu.tanggal = tanggal
		X[*N1+1].waktu.bulan = bulan
		X[*N1+1].waktu.tahun = tahun
		X[*N1+1].jenis = "ubah"
		X[*N1+1].nilai = c - T[idx].stok
	}
}

func cari(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nama barang")
	fmt.Println("2. harga barang")
	fmt.Println("3. stok")
	fmt.Println("4. transaksi")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		cariNama(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cariHarga(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {
		cariStok(&*T, &*X, &*N, &*N1)
	} else if pilihan == 4 {
		cariTransaksi(&*T, &*X, &*N, &*N1)
	}
}

func cariNama(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var j int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("====================")
	fmt.Print("nama barang: ")
	fmt.Scan(&nama)

	for i := 0; i < *N; i++ {
		fmt.Println("nama barang", "\t", "harga", "\t", "stok")
		if T[i].namaBarang == nama {
			fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
			j++
		}
	}

	if j == 0 {
		fmt.Print("\ntidak ada barang")
	}
}

func cariHarga(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, j, harga, idx int

	cariEkstrim(&pilihan)

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 1)
		fmt.Println("nama barang", "\t", "harga", "\t", "stok")
		fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 1)
		fmt.Println("nama barang", "\t", "harga", "\t", "stok")
		fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("harga barang: ")
		fmt.Scan(&harga)

		for i := 0; i < *N; i++ {
			fmt.Println("nama barang", "\t", "harga", "\t", "stok")
			if T[i].harga == harga {
				fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
				j++
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	}
}

func cariStok(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, j, stok, idx int

	cariEkstrim(&pilihan)

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 2)
		fmt.Println("nama barang", "\t", "harga", "\t", "stok")
		fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
	} else if pilihan == 2 {
		idx = idxMinimum(*T, *X, *N, *N1, 2)
		fmt.Println("nama barang", "\t", "harga", "\t", "stok")
		fmt.Println(T[idx].namaBarang, "\t", T[idx].harga, "\t", T[idx].stok)
	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("stok: ")
		fmt.Scan(&stok)

		for i := 0; i < *N; i++ {
			fmt.Println("nama barang", "\t", "harga", "\t", "stok")
			if T[i].stok == stok {
				fmt.Println(T[i].namaBarang, "\t", T[i].harga, "\t", T[i].stok)
				j++
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	}
}

func cariTransaksi(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari data transaksi:")
	fmt.Println("1. nama barang")
	fmt.Println("2. waktu transaksi")
	fmt.Println("3. jenis")
	fmt.Println("4. nilai")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		cariTransaksiNama(&*T, &*X, &*N, &*N1)
	} else if pilihan == 2 {
		cariTransaksiWaktu(&*T, &*X, &*N, &*N1)
	} else if pilihan == 3 {
		cariTransaksiJenis(&*X, &*N1)
	} else if pilihan == 4 {
		cariTransaksiNilai(&*T, &*X, &*N, &*N1)
	}
}

func cariTransaksiNama(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var nama string
	var j int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("====================")
	fmt.Print("nama barang: ")
	fmt.Scan(&nama)

	for i := 0; i < *N; i++ {
		fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
		if X[i].jenis == "tambah" {
			fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
		}
	}

	if j == 0 {
		fmt.Print("\ntidak ada transaksi")
	}
}

func cariTransaksiWaktu(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, idx, tanggal, bulan, tahun, waktu, j int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nilai maksimum")
	fmt.Println("2. nilai maksimum")
	fmt.Println("3. lainnya")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 3)
		fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
		fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].nilai, "\t", X[idx].waktu.tanggal, "\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun)

	} else if pilihan == 2 {
		idx = idxMaksimum(*T, *X, *N, *N1, 3)
		fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
		fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].nilai, "\t", X[idx].waktu.tanggal, "\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun)

	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("waktu (DD MM YYYY): ")
		fmt.Scan(&tanggal, &bulan, &tahun)

		waktu = (365 * tahun) + (30 * bulan) + tanggal
		for i := 0; i < *N; i++ {
			fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
			if X[i].jumlahWaktu == waktu {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
				j++
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	}
}

func cariTransaksiNilai(T *arrBarang, X *arrTransaksi, N, N1 *int) {
	var pilihan, idx, nilai, j int

	cariEkstrim(&pilihan)

	if pilihan == 1 {
		idx = idxMaksimum(*T, *X, *N, *N1, 4)
		fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
		fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].nilai, "\t", X[idx].waktu.tanggal, "\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun)

	} else if pilihan == 2 {
		idx = idxMaksimum(*T, *X, *N, *N1, 4)
		fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
		fmt.Println(X[idx].namaBarang, "\t", X[idx].jenis, "\t", X[idx].nilai, "\t", X[idx].waktu.tanggal, "\t", X[idx].waktu.bulan, "\t", X[idx].waktu.tahun)

	} else if pilihan == 3 {
		fmt.Println("====================")
		fmt.Println("  INVENTORI BARANG  ")
		fmt.Println("====================")
		fmt.Print("nilai: ")
		fmt.Scan(&nilai)

		for i := 0; i < *N; i++ {
			fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
			if X[i].nilai == nilai {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
				j++
			}
		}

		if j == 0 {
			fmt.Print("\ntidak ada barang")
		}
	}
}

func cariTransaksiJenis(X *arrTransaksi, N1 *int) {
	var pilihan int

	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari jenis data transaksi:")
	fmt.Println("1. tambah data")
	fmt.Println("2. ubah data")
	fmt.Println("3. hapus data")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)

	if pilihan == 1 {
		for i := 0; i < *N1; i++ {
			fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
			if X[i].jenis == "tambah" {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
			}
		}
	} else if pilihan == 2 {
		for i := 0; i < *N1; i++ {
			fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
			if X[i].jenis == "ubah" {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
			}
		}
	} else if pilihan == 3 {
		for i := 0; i < *N1; i++ {
			fmt.Println("namaBarang", "\t", "jenis", "\t", "nilai", "\t", "tanggal", "\t", "bulan", "\t", "tahun")
			if X[i].jenis == "hapus" {
				fmt.Println(X[i].namaBarang, "\t", X[i].jenis, "\t", X[i].nilai, "\t", X[i].waktu.tanggal, "\t", X[i].waktu.bulan, "\t", X[i].waktu.tahun)
			}
		}
	}
}

func cariEkstrim(pilihan *int) {
	fmt.Println("====================")
	fmt.Println("  INVENTORI BARANG  ")
	fmt.Println("cari:")
	fmt.Println("1. nilai maksimum")
	fmt.Println("2. nilai maksimum")
	fmt.Println("3. lainnya")
	fmt.Println("====================")
	fmt.Print("pilih: ")
	fmt.Scan(&pilihan)
}

func main() {
	
	var dataTransaksi arrTransaksi
	var NBarang, NTransaksi int
	// inputDataBarang(&dataBarang, &NBarang)
	// inputDataTransaksi(&dataTransaksi, &NTransaksi)
	menu(&dataBarang, &dataTransaksi, &NBarang, &NTransaksi)
	fmt.Println(dataBarang[1].namaBarang, dataBarang[1].harga, dataBarang[1].stok)
	fmt.Println(dataTransaksi[1].namaBarang, dataTransaksi[1].jenis, dataTransaksi[1].waktu.tanggal, dataTransaksi[1].waktu.bulan, dataTransaksi[1].waktu.tahun, dataTransaksi[1].nilai)
}
