import React, { useEffect, useState } from "react";
import RowKaryawan from "./components/RowKaryawan";
import TambahKaryawan from "./components/TambahKaryawan";

// const dataKaryawan = [
//   {
//     id: 1557298991816,
//     nama_karyawan: "allen saputra",
//     jabatan: "Security",
//     jenis_kelamin: "laki-laki",
//     tanggal_lahir: "1993-12-08",
//   },
//   {
//     id: 1537298993816,
//     nama_karyawan: "satria",
//     jabatan: "Marketing",
//     jenis_kelamin: "laki-laki",
//     tanggal_lahir: "1990-12-12",
//   },
//   {
//     id: 1557298963816,
//     nama_karyawan: "putra agung",
//     jabatan: "Security",
//     jenis_kelamin: "laki-laki",
//     tanggal_lahir: "1991-12-12",
//   },
//   {
//     id: 1557298993816,
//     nama_karyawan: "gerysena",
//     jabatan: "Marketing",
//     jenis_kelamin: "laki-laki",
//     tanggal_lahir: "2011-01-31",
//   },
// ];

const App = () => {
  const [karyawan, setKaryawan] = useState([]);
  const [errorStatus, setErrorStatus] = useState(false);
  const [loadingStatus, setLoadingStatus] = useState(true);

  // =========================================================================
  // Function untuk mengambil data karyawan dari API
  // -------------------------------------------------------------------------
  const getKaryawan = async () => {
    try {
      let response = await fetch("http://localhost:3004/data-karyawan");
      if (!response.ok) {
        setErrorStatus(true);
        throw new Error(
          "Data Gagal Dimuat, silahkan periksa url api atau koneksi jaringan anda"
        );
      }
      let data = await response.json();
      // console.log(data);
      setKaryawan(data);
    } catch (e) {
      console.log(e);
    } finally {
      setLoadingStatus(false);
    }
  };
  // =========================================================================

  // =========================================================================
  // Pada render pertama kali, jalankan function getKaryawan untuk mengambil data karyawan dari database
  // -------------------------------------------------------------------------
  useEffect(() => {
    getKaryawan();
  }, []);
  // =========================================================================

  // =========================================================================
  // Function untuk memeriksa duplikasi data saat menambahkan karyawan
  // -------------------------------------------------------------------------
  const handleDuplicateData = (data) => {
    let duplicateData = false;

    karyawan.forEach((kar) => {
      // console.log(`${kar.id} === ${data.id} : ${kar.id === data.id}`);
      kar.id === data.id && (duplicateData = true);
    });

    // console.log(duplicateData);

    return duplicateData; // Mengembalikan nilai duplicateData
  };
  // -------------------------------------------------------------------------
  // Function untuk menambahkan karyawan baru ke data karyawan
  const handleTambahKaryawan = async (data) => {
    if (handleDuplicateData(data)) {
      alert("Data karyawan sudah ada");
    } else {
      // Menambah data karyawan ke state
      // setKaryawan([...karyawan, data]);

      console.log(data);
      console.log(JSON.stringify(data));

      // Menambah data karyawan ke database
      let url = `http://localhost:3004/data-karyawan/`;
      try {
        let response = await fetch(url, {
          method: "POST",
          body: JSON.stringify(data),
          headers: {
            "Content-type": "application/json; charset=UTF-8",
          },
        });
        if (!response.ok) throw new Error(response.status);
      } catch (e) {
        console.log(`Terjadi gangguan saat mengirim data dengan pesan: "${e}"`);
      }

      // Mengupdate kembalu data karyawan
      getKaryawan();
    }
  };
  // =========================================================================

  // =========================================================================
  // Function untuk mengupdate data karyawan
  // -------------------------------------------------------------------------
  const handleEditKaryawan = async (data) => {
    // Update data karyawan pada state
    // let indexDataEdited = karyawan.findIndex((kar) => kar.id === data.id);
    // console.log(indexDataEdited);
    // let newKaryawan = [...karyawan];
    // newKaryawan.splice(indexDataEdited, 1, data);
    // console.log(karyawan);
    // console.log(newKaryawan);
    // setKaryawan(newKaryawan);

    // Update data karyawan pada database menggunakan api method
    let url = `http://localhost:3004/data-karyawan/${data.id}`;
    try {
      let response = await fetch(url, {
        method: "PUT",
        body: JSON.stringify(data),
        headers: {
          "Content-type": "application/json; charset=UTF-8",
        },
      });
      if (!response.ok) throw new Error(response.status);
    } catch (e) {
      console.log(`Terjadi gangguan saat mengupdate data dengan pesan: "${e}"`);
    }

    // Mengupdate kembalu data karyawan
    getKaryawan();
  };
  // =========================================================================

  // =========================================================================
  // Function untuk menghapus data karyawan
  // -------------------------------------------------------------------------
  const handleDeleteKaryawan = async (e) => {
    // console.log(e.target.id);

    // Menghapus data karyawan dari state
    // let karyawanAfterDelete = karyawan.filter((kar) => {
    //   return kar.id !== parseInt(e.target.id);
    // });
    // setKaryawan(karyawanAfterDelete);
    // console.log(karyawanAfterDelete);

    // Menghapus data karyawan dari database berdasarkan id-nya
    let url = `http://localhost:3004/data-karyawan/${e.target.id}`;
    try {
      let response = await fetch(url, { method: "DELETE" });
      if (!response.ok) throw new Error(response.status);
    } catch (e) {
      console.log(`Terjadi gangguan saat menghapus data dengan pesan: "${e}"`);
    }

    // Mengupdate kembalu data karyawan
    getKaryawan();
  };
  // =========================================================================

  // console.log(karyawan);

  return (
    <div className="container">
      <div className="row mt-5">
        <div className="col">
          <h1 className="text-center">Data Karyawan</h1>
          <table className="table mt-3">
            <thead>
              <tr>
                <th>ID</th>
                <th>Nama</th>
                <th>Jabatan</th>
                <th>Jenis Kelamin</th>
                <th>Tanggal Lahir</th>
                <th></th>
              </tr>
            </thead>
            <tbody className="table-group-divider">
              {loadingStatus ? (
                <tr>
                  <td colSpan={6}>
                    <div class="d-flex justify-content-center">
                      <div class="spinner-border" role="status">
                        <span class="sr-only"></span>
                      </div>
                    </div>
                  </td>
                </tr>
              ) : errorStatus ? (
                <tr>
                  <td colSpan={6}>
                    <div class="d-flex justify-content-center">
                      <h1>Terjadi Error saat memuat data...</h1>
                    </div>
                  </td>
                </tr>
              ) : (
                karyawan.map((kar) => {
                  return (
                    <RowKaryawan
                      key={kar.id}
                      karyawan={kar}
                      onEditKaryawan={handleEditKaryawan}
                      onDeleteKaryawan={handleDeleteKaryawan}
                    />
                  );
                })
              )}
              <TambahKaryawan onTambahKaryawan={handleTambahKaryawan} />
            </tbody>
          </table>
        </div>
      </div>
    </div>
  );
};

export default App;
