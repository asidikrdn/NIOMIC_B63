const buat_login = () => {
  // Menangkap kotak form login
  let kotakLogin = document.querySelector('.kotak_login');

  // Menghapus tombol berwarna merah
  kotakLogin.removeChild(document.getElementById('X'));

  // Membuat tulisan login
  let tulisanLogin = document.createElement('p');
  tulisanLogin.innerHTML = "Silahkan Mendaftar";
  tulisanLogin.className = 'tulisan_login';
  kotakLogin.appendChild(tulisanLogin);

  // Membuat element form
  let form = document.createElement('form');
  kotakLogin.appendChild(form);

  // Membuat form input nama
  let labelInputNama = document.createElement('label');
  labelInputNama.innerHTML = "Nama User";
  labelInputNama.setAttribute('for', 'nama')
  let inputNama = document.createElement('input');
  inputNama.setAttribute('type', 'text');
  inputNama.setAttribute('id', 'nama');
  inputNama.setAttribute('name', 'nama');
  inputNama.setAttribute('placeholder', 'Nama User..');
  inputNama.className = 'form_login';
  form.appendChild(labelInputNama);
  form.appendChild(inputNama);

  // Membuat form input nomor handphone
  let labelInputNoHP = document.createElement('label');
  labelInputNoHP.innerHTML = "Nomor Handphone";
  labelInputNoHP.setAttribute('for', 'NoHP')
  let inputNoHP = document.createElement('input');
  inputNoHP.setAttribute('type', 'text');
  inputNoHP.setAttribute('id', 'noHP');
  inputNoHP.setAttribute('name', 'noHP');
  inputNoHP.setAttribute('placeholder', 'Nomor Handphone.');
  inputNoHP.className = 'form_login';
  form.appendChild(labelInputNoHP);
  form.appendChild(inputNoHP);

  // Membuat form input Username
  let labelInputUsername = document.createElement('label');
  labelInputUsername.innerHTML = "Username";
  labelInputUsername.setAttribute('for', 'username')
  let inputUsername = document.createElement('input');
  inputUsername.setAttribute('type', 'text');
  inputUsername.setAttribute('id', 'username');
  inputUsername.setAttribute('name', 'username');
  inputUsername.setAttribute('placeholder', 'Username atau email ..');
  inputUsername.className = 'form_login';
  form.appendChild(labelInputUsername);
  form.appendChild(inputUsername);

  // Membuat form input Password
  let labelInputPassword = document.createElement('label');
  labelInputPassword.innerHTML = "Password";
  labelInputPassword.setAttribute('for', 'password')
  let inputPassword = document.createElement('input');
  inputPassword.setAttribute('type', 'password');
  inputPassword.setAttribute('id', 'password');
  inputPassword.setAttribute('name', 'password');
  inputPassword.setAttribute('placeholder', 'Password ..');
  inputPassword.className = 'form_login';
  form.appendChild(labelInputPassword);
  form.appendChild(inputPassword);

  // Membuat tombol submit
  let tombolLogin = document.createElement('input');
  tombolLogin.setAttribute('type', 'submit');
  tombolLogin.setAttribute('value', 'DAFTAR SEKARANG');
  tombolLogin.className = 'tombol_login';
  form.appendChild(tombolLogin);
}