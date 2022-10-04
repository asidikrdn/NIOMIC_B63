import React, { useState } from "react";

const TambahKaryawan = (props) => {
  const [formInput, setFormInput] = useState({
    id: "",
    nama_karyawan: "",
    jabatan: "",
    jenis_kelamin: "",
    tanggal_lahir: "",
  });

  const [errors, setErrors] = useState({
    id: "",
    nama_karyawan: "",
    jabatan: "",
    jenis_kelamin: "",
    tanggal_lahir: "",
  });

  const handleInputChange = (e) => {
    e.target.name === "id"
      ? e.target.value.length > 0
        ? setFormInput({
            ...formInput,
            [e.target.name]: parseInt(e.target.value),
          })
        : setFormInput({
            ...formInput,
            [e.target.name]: "",
          })
      : setFormInput({ ...formInput, [e.target.name]: e.target.value });
  };

  const handleFormSubmit = (e) => {
    e.preventDefault();
    let pesanErrors = {};

    // Validasi ID
    if (formInput.id.toString().trim() === "") {
      pesanErrors.id = "ID tidak boleh kosong";
    } else if (!/^[\d]{12}/.test(formInput.id)) {
      pesanErrors.id = "ID harus 13 digit angka";
    } else {
      pesanErrors.id = "";
    }

    // Validasi Nama
    formInput.nama_karyawan.trim() === ""
      ? (pesanErrors.nama_karyawan = "Nama Karyawan tidak boleh kosong")
      : (pesanErrors.nama_karyawan = "");

    // Validasi Jabatan
    formInput.jabatan.trim() === ""
      ? (pesanErrors.jabatan = "Jabatan tidak boleh kosong")
      : (pesanErrors.jabatan = "");

    // Validasi Jenis Kelamin
    formInput.jenis_kelamin.trim() === ""
      ? (pesanErrors.jenis_kelamin = "Jenis Kelamin tidak boleh kosong")
      : (pesanErrors.jenis_kelamin = "");

    // Validasi Tanggal Lahir
    formInput.tanggal_lahir.trim() === ""
      ? (pesanErrors.tanggal_lahir = "Tanggal Lahir tidak boleh kosong")
      : (pesanErrors.tanggal_lahir = "");

    // Update State Error
    setErrors(pesanErrors);

    // Pemeriksaan seluruh validasi
    let formValid = true;
    for (let inputName in pesanErrors) {
      pesanErrors[inputName] !== "" && (formValid = false);
    }
    // console.log(formValid);

    // Jika lolos validasi, jalankan perintah didalamnya
    if (formValid) {
      // console.log(formInput);

      // Menjalankan function handleTambahKaryawan yang ada di komponen App
      props.onTambahKaryawan(formInput);

      // Mengosongkan Form Input
      setFormInput({
        id: "",
        nama_karyawan: "",
        jabatan: "",
        jenis_kelamin: "",
        tanggal_lahir: "",
      });
    }
  };

  return (
    <tr>
      <td colSpan="6">
        <form onSubmit={handleFormSubmit}>
          <div className="row g-3">
            <div className="col">
              <input
                type="text"
                name="id"
                placeholder="Masukkan ID"
                className="form-control"
                value={formInput.id}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.id}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="nama_karyawan"
                placeholder="Masukkan Nama Karyawan"
                className="form-control"
                value={formInput.nama_karyawan}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.nama_karyawan}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="jabatan"
                placeholder="Masukkan Jabatan"
                className="form-control"
                value={formInput.jabatan}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.jabatan}</small>}
            </div>
            <div className="col">
              <input
                type="text"
                name="jenis_kelamin"
                placeholder="Masukkan Jenis Kelamin"
                className="form-control"
                value={formInput.jenis_kelamin}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.jenis_kelamin}</small>}
            </div>
            <div className="col">
              <input
                type="date"
                name="tanggal_lahir"
                placeholder="Masukkan Tanggal Lahir"
                className="form-control"
                value={formInput.tanggal_lahir}
                onChange={handleInputChange}
              />
              {errors && <small>{errors.tanggal_lahir}</small>}
            </div>
            <div className="col">
              <button type="submit" className="btn btn-primary">
                Tambah
              </button>
            </div>
          </div>
        </form>
      </td>
    </tr>
  );
};

export default TambahKaryawan;
