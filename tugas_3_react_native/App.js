import React, {Component} from 'react';
import {View} from 'react-native';
import Barang from './Component/Barang';
import Header from './Component/Header';

class App extends Component {
  render() {
    return (
      <View>
        <Header data="SELAMAT DATANG" />
        <Barang />
      </View>
    );
  }
}

export default App;
