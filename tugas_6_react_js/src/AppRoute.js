import React from "react";
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";
import App from "./App";
import MenuMakanan from "./Page/MenuMakanan";
import MenuMinuman from "./Page/MenuMinuman";
import Kontak from "./Page/Kontak";

const AppRoute = () => {
  return (
    <Router>
      <div>
        <Routes>
          <Route path="/" element={<App />} />
          <Route path="/menu_makanan" element={<MenuMakanan />} />
          <Route path="/menu_minuman" element={<MenuMinuman />} />
          <Route path="/kontak" element={<Kontak />} />
        </Routes>
      </div>
    </Router>
  );
};

export default AppRoute;
