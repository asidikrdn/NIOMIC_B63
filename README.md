# Perkenalan React Native

React Native salah satu framework untuk membuat mobile app.
+ Terdapat fitur hot reload
+ Banyak perusahaan besar yang menggunakan framework react native


## SetUp React Native

Menggunakan React Native CLI
- Download & Install JDK
- Download & Install NodeJS
- Download & Install Android Studio
Android Studio > Configure > Android SDK > Android 12 & SDK Platform 31

- Environtment Variables > New 
	Variabel Name : `ANDROID_HOME`
	Variabel Value : <Android SDK Location>
- Environtment Variables > Path > Edit > New
	<SDK Location>\platform-tools

Cek versi react native 	=> `react-native --version`
Cek vesi node 		=> `node --version`
Cek versi java 		=> `java -version`

## Create New React Native Project

`npx react-native init ProjectName`

Buka VSCode > Buka folder ProjectName

Buka android studio > avd manager > create virtual device > pilih device yang ingin digunakan untuk simulasi > install os untuk virtual device > beri nama virtual device > finish > jalankan virtual device tersebut

Buka terminal di vscode > jalankan `npx react-native start` lalu `npx react-native run-android`

## Basic Component React Native

- Buka folder project di vscode
- jalankan virtual device
- jalankan project react native via terminal di vscode


AppRegistry.registerterComponent => fungsi untuk menampilkan halaman tertentu pada saat app dijalankan/dibuka pertama kali

- Buka App.js > Edit text tertentu
- aktifkan fitur hot reload > klik virtual device > CTRL+M > klik hot reloading

### Compoment dasar yang sering digunakan
- `<View>` : A container that supports layout with flexbox, style, some touch handling, and accessibility controls
- `<Text>` : Displays, styles, and nests strings of text and even handles touch events
- `<StatusBar>` : Configure and modify status bar on devices
- `<Image>` : Displays different types of images
- `<ScrollView>` : A generic scrolling container that can contain multiple components and views
- `<TextInput>` : Allows the user to enter text
