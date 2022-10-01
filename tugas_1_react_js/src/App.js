import React, { Component } from "react";
import MenuUtama from "./Page/MenuUtama";
import MenuProduct from "./Page/MenuProduct";
import MenuKontak from "./Page/MenuKontak";
import MenuTentangKami from "./Page/MenuTentangKami";

const Header = () => {
  return <h1>Ini Halaman Untuk Header</h1>;
};

const Footer = () => {
  return <h1>Ini Halaman Untuk Footer</h1>;
};

class App extends Component {
  render() {
    return (
      <>
        <Header></Header>
        <MenuUtama></MenuUtama>
        <MenuProduct></MenuProduct>
        <MenuKontak></MenuKontak>
        <MenuTentangKami></MenuTentangKami>
        <Footer></Footer>
      </>
    );
  }
}

export default App;
