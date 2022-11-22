# Tugas 16

Struktur Database :

```SQL
CREATE TABLE IF NOT EXISTS `tbl_mahasiswa` (
  `nim` varchar(12) NOT NULL,
  `nama` varchar(255) NOT NULL,
  `umur` int(11) NOT NULL,
  `asal_kota` varchar(255) NOT NULL,
  `asal_sekolah` varchar(255) NOT NULL,
  `nilai_ujian` float(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `tbl_mahasiswa` (`nim`, `nama`, `umur`, `asal_kota`, `asal_sekolah`, `nilai_ujian`) VALUES
("201106041165", 'Ahmad', 21, "Bogor", "SMK Negeri 3 Kota Bogor", 3.99),
("201106041166", 'Sidik', 22, "Brebes", "SMK Negeri 1 Brebes", 3.98),
("201106041167", 'Rudini', 23, "Depok", "SMK Negeri 2 Depok", 4.00);

ALTER TABLE `tbl_mahasiswa` ADD PRIMARY KEY (`nim`);
```

Cara menjalankan program :

1. Clone source code project ini
2. Buka folder project `Tugas_16_GO`
3. Jalankan database mysql
4. Buat database baru dengan nama `db_tugas16`, lalu import file `db_tugas16.sql` yang ada di folder `Tugas_16_GO`
5. Kembali ke folder `Tugas_16_GO`, kemudian jalankan program `tugas16.go`
