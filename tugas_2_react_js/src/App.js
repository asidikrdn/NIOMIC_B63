import React, { Component } from "react";
import Header from "./Page/Header";
import Footer from "./Page/Footer";
import MenuUtama from "./Page/MenuUtama";
import MenuMakanan from "./Page/MenuMakanan";
import MenuKontak from "./Page/MenuKontak";
import MenuTentangKami from "./Page/MenuTentangKami";

class App extends Component {
  render() {
    return (
      <>
        <Header></Header>
        <MenuUtama></MenuUtama>
        <MenuMakanan></MenuMakanan>
        <MenuTentangKami></MenuTentangKami>
        <MenuKontak></MenuKontak>
        <Footer></Footer>
      </>
    );
  }
}

export default App;
