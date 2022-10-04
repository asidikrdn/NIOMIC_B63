import React, { useRef, useState, useEffect } from "react";

const RowKaryawan = (props) => {
  const [formInput, setFormInput] = useState({});
  const [editStatus, setEditStatus] = useState(false);
  const [errors, setErrors] = useState({
    id: "",
    nama_karyawan: "",
    jabatan: "",
    jenis_kelamin: "",
    tanggal_lahir: "",
  });
  const dataReset = useRef({});

  // Menggunakan useEffect agar data yang tampil selalu diupdate dari state karyawan pada App.js
  // Hal ini mencegah perbedaan data antara yang tampil dengan yang tersimpan
  useEffect(() => {
    setFormInput(props.karyawan);
  }, [props.karyawan]);

  const handleInputChange = (e) => {
    setFormInput({ ...formInput, [e.target.name]: e.target.value });
  };

  const handleEditCondition = () => {
    if (editStatus) {
      setFormInput(dataReset.current);
      setEditStatus(!editStatus);
    } else {
      dataReset.current = { ...formInput };
      setEditStatus(!editStatus);
    }
    // console.log(dataReset.current);
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
      ? (pesanErrors.nama_karyawan = "nama_karyawan tidak boleh kosong")
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
      props.onEditKaryawan(formInput);
      setEditStatus(false);
    }
  };

  return (
    <>
      {editStatus ? (
        <tr>
          <td colSpan="6">
            <form onSubmit={handleFormSubmit}>
              <div className="row">
                <div className="col">
                  <input
                    type="text"
                    name="id"
                    placeholder="Masukkan ID"
                    className="form-control"
                    value={formInput.id}
                    onChange={handleInputChange}
                    disabled
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
                  <button type="submit" className="btn btn-success me-2">
                    Simpan
                  </button>
                  <button
                    type="reset"
                    className="btn btn-warning"
                    onClick={handleEditCondition}
                  >
                    Batal
                  </button>
                </div>
              </div>
            </form>
          </td>
        </tr>
      ) : (
        <tr>
          <td>{formInput.id}</td>
          <td>{formInput.nama_karyawan}</td>
          <td>{formInput.jabatan}</td>
          <td>{formInput.jenis_kelamin}</td>
          <td>{formInput.tanggal_lahir}</td>
          <td>
            <button
              className="btn btn-secondary me-2"
              id={formInput.id}
              onClick={handleEditCondition}
            >
              Edit
            </button>
            <button
              className="btn btn-danger"
              id={formInput.id}
              onClick={props.onDeleteKaryawan}
            >
              Hapus
            </button>
          </td>
        </tr>
      )}
    </>
  );
};

export default RowKaryawan;
